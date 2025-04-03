> [!note] Vulkan 实例：包含一系列 Vulkan 运行必须的状态和信息的变量

Vulkan 对象创建通常分为两步：
- 创建对应的 `VkXxxCreateInfo` 结构体，提供对应的信息
- 调用 `vkXxxCreate` 方法，通常需要一个 info 结构体，一个内存分配器（可以是 `nullptr` 表示默认），一个接收的指针

不同 `VkInfo` 结构体有一些公共的字段

| 成员      | 类型                | 说明                                  |
| ------- | ----------------- | ----------------------------------- |
| `sType` | `VkStructureType` | 结构体类型，不同结构体值不同                      |
| `pNext` | `void*`           | 如有必要，指向一个用于扩展该结构体的结构体，通常为 `nullptr` |
| `flags` | 各类 `Flags`        | 默认为 0                               |

因此封装一个 Vulkan 对象的基类
- Vulkan 对象（`THandle`）实质是指是一个指针，可以直接存储
- 接收一个用于记录对象类型的字符串，用于输出错误信息
- 显式声明一个析构函数，用于删除对象
- 重写移动构造和移动运算，移动 `handle` 指针
- 删除复制构造和复制运算

```cpp title:vkutils.h
template<class THandle, const char *TYPENAME>
class VkHandle {

protected:
    THandle handle = VK_NULL_HANDLE;
    const char *VKType = TYPENAME;

    // 直接封装 Vulkan 对象
    explicit VkHandle(THandle handle) : handle(handle) {};

public:

    VkHandle() = default;

    virtual ~VkHandle() = 0;
    
    // 禁用复制构造
    VkHandle(const VkHandle&) = delete;

    // 移动构造
    VkHandle(VkHandle&& other) noexcept : handle(other.handle) {
        other.handle = VK_NULL_HANDLE;
    }  

    THandle* getAddress() {  
        return &handle;  
    }

    // 移动运算符
    VkHandle& operator=(VkHandle&& other) noexcept {
        if (this != other) {
            handle = other.handle;
            other.handle = VK_NULL_HANDLE;
        }
        return *this;
    }
    
    // 删除复制运算符
    VkHandle& operator=(const VkHandle&) = delete;
    
    // Vulkan 对象与封装对象可隐式转换
    operator THandle() const {
        return handle;
    }
};
```
# 版本选择

使用 `vkEnumerateInstanceVersion` 可用于查询 Vulkan SDK 支持的最新 Vulkan 版本，但该方法仅在 1.0 以上的版本存在

```cpp
// 不存在 vkEnumerateInstanceVersion 函数，说明为 1.0
uint32_t apiVersion = VK_VERSION_1_0;

if (vkGetInstanceProcAddr(VK_NULL_HANDLE, "vkEnumerateInstanceVersion")) {
    // 返回值表示是否成功获取
    return vkEnumerateInstanceVersion(&apiVersion);
}

return VK_SUCCESS;
```
# 层与扩展

声明 Vulkan 实例需要确认层与扩展，提供对特定需求、特定平台的支持，还包括特定厂商提供的功能。
- 层：有显著的作用范围，如验证层等。只有实例级层，没有设备（独占）级层
- 扩展：分为实例级和设备级扩展。某些扩展也需要特定层

GLFW 提供 `glfwGetRequiredInstanceExtensions` 方法获取运行平台所需要的扩展。Windows 平台获取的扩展为：
- `VK_KHR_surface`
- `VK_KHR_win32_surface`

```cpp
uint32_t extensionCount;
auto extensionNames = glfwGetRequiredInstanceExtensions(&extensionCount);
if (!extensionNames) {
    // 系统不支持 Vulkan
    std::cout << "[ InitializeWindow ] ERROR\n"
                 "Vulkan is not available on this machine!\n";
    return VK_OTHER_ERROR;
}
auto instanceExtensions = std::vector<const char *>{extensionNames, extensionNames + extensionCount};
```

设备级扩展只有 `VK_KHR_SWAPCHAIN_EXTENSION_NAME` 是必须的

```cpp
const char *deviceExtensions[] = {
        VK_KHR_SWAPCHAIN_EXTENSION_NAME,
};
uint32_t deviceExtensionsCount = 1;
```

调试需要的 `Debug Messenger` 需要 `VK_LAYER_KHRONOS_validation` 层和 `VK_EXT_DEBUG_UTILS_EXTENSION_NAME` 扩展，详见[[Debug Messenger]]

```cpp
// 层
std::vector<const char *> instanceLayers = {
#ifndef NDEBUG
        "VK_LAYER_KHRONOS_validation",
#endif
};

#ifndef NDEBUG  
instanceExtensions.push_back(VK_EXT_DEBUG_UTILS_EXTENSION_NAME);  
#endif
```
## 可用性检查

层与扩展可以在创建实例前检查其可用性，也可以不检查。当创建 Vulkan 实例时，若存在无法满足的层和扩展，会产生失败结果
- 不满足层：`VK_ERROR_LAYER_NOT_PRESENT`
- 不满足扩展：`VK_ERROR_EXTENSION_NOT_PRESENT`
因此可以在创建实例失败时再去处理不支持的层和扩展。

检查层、扩展的方法如下：
- `vkEnumerateInstanceLayerProperties`
- `vkEnumerateInstanceExtensionProperties`
	- 需要额外一个 `layerName` 参数表示扩展用于哪个层
	- `layerName=nullptr` 表示检查适用于所有层
- `vkEnumerateDeviceExtensionProperties`

```cpp
VkResult checkInstanceLayers(std::span<const char *> layers) {
    uint32_t layerCount;
    std::vector<VkLayerProperties> availableLayers;
    // 获取可用层数量
    if (auto result = vkEnumerateInstanceLayerProperties(&layerCount, nullptr)) {
        std::cout << std::format("[ VkInstance ] ERROR\n"
                                 "Failed to get the count of instance layers!\n"
                                 "Error code: {}\n", int32_t(result));
        return result;
    }
    // 获取可用层
    if (layerCount) {
        availableLayers.resize(layerCount);
        if (auto result = vkEnumerateInstanceLayerProperties(&layerCount, availableLayers.data())) {
            std::cout << std::format("[ VkInstance ] ERROR\n"
                                     "Failed to enumerate instance layer properties!\n"
                                     "Error code: {}\n", int32_t(result));
            return result;
        }
        for (auto &layerName : layers) {
            bool found = false;
            for (const auto &al : availableLayers) {
                if (!strcmp(layerName, al.layerName)) {
                    found = true;
                    break;
                }
            }
            if (!found) {
                // 层扩展不存在，这里直接删除该层
                layerName = nullptr;
            }
        }
    }
    
    return VK_SUCCESS;
}
```
# 创建实例

创建实例所需的信息对象为 `VkInstanceCreateInfo`

| 成员                        | 类型                   | 说明                                       |
| ------------------------- | -------------------- | ---------------------------------------- |
| `sType`                   | `VkStructureType`    | `VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO` |
| `pApplicationInfo`        | `VkApplicationInfo*` | 指向描述本程序相关信息的结构体                          |
| `enabledLayerCount`       | `uint32_t`           | 所需额外开启的实例级别层数                            |
| `ppEnabledLayerNames`     | `const char* const*` | 指向由所需开启的层的名称构成的数组（同一名称可以重复出现）            |
| `enabledExtensionCount`   | `uint32_t`           | 所需额外开启的实例级别扩展数                           |
| `ppEnabledExtensionNames` | `const char* const*` | 指向由所需开启的扩展的名称构成的数组（同一名称可以重复出现）           |

`VkApplicationInfo` 结构体描述程序具体信息，`apiVersion` 字段填写使用的 VulkanAPI 版本，其他名称版本按需填

| 成员                   | 类型                | 说明                                   |
| -------------------- | ----------------- | ------------------------------------ |
| `sType`              | `VkStructureType` | `VK_STRUCTURE_TYPE_APPLICATION_INFO` |
| `pApplicationName`   | `const void*`     | 应用程序的名称                              |
| `applicationVersion` | `uint32_t`        | 应用程序的版本号                             |
| `pEngineName`        | `const void*`     | 引擎的名称                                |
| `engineVersion`      | `uint32_t`        | 引擎的版本号                               |
| `apiVersion`         | `uint32_t`        | VulkanAPI的版本号，必填                     |

```cpp
VkApplicationInfo applicationInfo = {
        .sType = VK_STRUCTURE_TYPE_APPLICATION_INFO,
        .apiVersion = apiVersion,
};

VkInstanceCreateInfo instanceCreateInfo = {
        .sType = VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO,
        .flags = flags,
        .pApplicationInfo = &applicationInfo,
        .enabledLayerCount = uint32_t(instanceLayers.size()),
        .ppEnabledLayerNames = instanceLayers.data(),
        .enabledExtensionCount = uint32_t(instanceExtensions.size()),
        .ppEnabledExtensionNames = instanceExtensions.data(),
};

if (auto result = vkCreateInstance(&instanceCreateInfo, nullptr, &instance)) {
    std::cout << std::format("[ VkInstance ] ERROR\n"
                             "Failed to create a vulkan instance!\n"
                             "Error code: {}\n", int32_t(result));
    return result;
}

// 创建成功，输出 Vulkan 版本
std::cout << std::format("Vulkan API Version: {}.{}.{}\n",
                         VK_VERSION_MAJOR(apiVersion),
                         VK_VERSION_MINOR(apiVersion),
                         VK_VERSION_PATCH(apiVersion));
```
# Window Surface

VkSurfaceKHR，用于 VulkanAPI 与平台功能对接
- GLFW：使用 `glfwCreateWindowSurface` 创建
- 纯 VulkanSDK：使用 `vkCreateWin32SurfaceKHR` 创建

> [!note] 若应用有多个窗口，则对应有多个 Window Surface 和多个交换链

```cpp title:VulkanInstance::setSurfaceFromGLFW
auto result = glfwCreateWindowSurface(handle, window, nullptr, &surface);
// if (result) { cout; return result; }
returnVkError("Failed to allocate a vulkan surface!");
return VK_SUCCESS;
```

> [!failure] `VK_RESULT_MAX_ENUM` 错误代码
> `VK_RESULT_MAX_ENUM` 本身不是错误代码。但 VkResult 中没有找不到函数时的枚举值，且 `VK_RESULT_MAX_ENUM` 不是任何已用的错误代码，因此使用该值。
# 封装实例对象

首先封装一些工具类（工具宏）

```cpp title:vkutils.h
#define _instance (vk::VulkanInstance::base())
#define _device   (vk::Device::base())

#define destroyHandle(func)         if (handle) { func(handle, nullptr);            handle = VK_NULL_HANDLE; }
#define destroyDeviceHandle(func)   if (handle) { func(_device, handle, nullptr);   handle = VK_NULL_HANDLE; }
#define destroyInstanceHandle(func) if (handle) { func(_instance, handle, nullptr); handle = VK_NULL_HANDLE; }
#define destroyDeviceHandles(func, handles) { for (auto _handle : (handles))   { func(_device, _handle, nullptr);    } (handles).resize(0); }
#define destroyDeviceHandleArray(func)      { for (auto i = 0; i < count; i++) { func(_device, handles[i], nullptr); } delete[]handles; handles = nullptr; }

#define getProcAddr(name) (reinterpret_cast<PFN_##name>(vkGetInstanceProcAddr(_instance, #name)))

#define printVkWarning(message)     { std::cout << "[ " << VKType << " ] WARNING\n" << (message) << "\nError code: " << int32_t(result) << "\n"; }
#define printVkWarningCode(code)    { std::cout << "[ " << VKType << " ] WARNING\nError code: " << int32_t(code) << "\n"; }
#define printVkMessage(message)     { std::cout << "[ " << VKType << " ] ERROR\n" << (message) << "\n"; }
#define returnVkMessage(message)    { std::cout << "[ " << VKType << " ] ERROR\n" << (message) << "\n"; return vk::VK_OTHER_ERROR; }

#define printVkCode(code)      if (result) { std::cout << "[ " << VKType << " ] ERROR\nError code: " << int32_t(code) << "\n"; }
#define printVkError(message)  if (result) { std::cout << "[ " << VKType << " ] ERROR\n" << (message) << "\nError code: " << int32_t(result) << "\n"; }
#define returnVkError(message) if (result) { std::cout << "[ " << VKType << " ] ERROR\n" << (message) << "\nError code: " << int32_t(result) << "\n"; return result; }

#ifndef NDEBUG
#define ENABLE_DEBUG true
#else
#define ENABLE_DEBUG false
#endif

// 构造单例类
// 保留默认构造，删除移动构造
// 提供 base 方法获取单例
#define SingletonClass(Type) \
private:  static Type singleton;  Type() = default;  Type(Type &&) = delete;  \
public:   static Type &base() { return singleton; }
```

封装 VulkanInstance 类

```cpp title:VulkanInstance.h
namespace vk {

    constexpr const char TYPENAME_INSTANCE[] = "VkInstance";

    class VulkanInstance : public VkHandle<VkInstance, TYPENAME_INSTANCE> {
    SingletonClass(VulkanInstance);

    private:
        uint32_t apiVersion = VK_VERSION_1_0;
        VkInstance handle = VK_NULL_HANDLE;
        VkSurfaceKHR surface = VK_NULL_HANDLE;
    public:

        virtual ~VulkanInstance();

        VkResult initialize(VkInstanceCreateFlags flags = 0);

        uint32_t getVersion() const;

        VkResult setSurfaceFromGLFW(GLFWwindow *window);

        VkSurfaceKHR getSurface() const;
    };
} // vk
```