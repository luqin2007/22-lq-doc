# Coremod

对其它类进行修改的 mod

- 可对 Minecraft、其他 mod 的类进行修改
- 可在 Minecraft 启动前执行代码，用于环境预处理，释放前置 mod 等操作

## 注意事项

1. 不能调用 Minecraft，其他普通 Mod 甚至自己的普通 Mod 部分，防止类在 Coremod 完全修改完成前过早加载进 classLoader 导致崩溃

   **CoreMod 的 ModContainer 也属于运行时的类，不可加载**

   可使用 gradle 的 sourceSet 进行隔离

   ```groovy
   sourceSets {
       coremod {
           compileClasspath += main.compileClasspath
       }
       main {
           runtimeClasspath += coremod.output
       }
   }
   ```

   

2. Coremod 实现的类都需要一个无参构造

## Vanilla Coremod

### 直接修改 Class 文件

版本：1.5.2 and before

过程：

1. 下载/解压 MCP

2. 使用 decompile 反编译 Minecraft 源码

3. 修改 src 目录下的源文件

4. 使用 recompile 编译，并使用 reobfuscate 混淆

5. 使用 reobf 文件夹下修改后的 class 文件替换 minecraft.jar 中对应的文件

6. （1.6 and later）在 JVM 参数中关闭 FML 对 Minecraft 核心文件完整性及签名检查

   >  -Dfml.ignoreInvalidMinecraftCertificates=true -Dfml.ignorePatchDiscrepancies=true 

7. （1.6 and later）修改 json 文件绕过启动器对黑犀牛文件完整性的检查

**极不推荐**

### JavaAgent

由 JVM 提供支持，在 class 文件被加载时对其进行动态修改

虽 FML 未使用，但这也是一种合理的方法

#### premain 方法

在 main 前执行的静态方法

接收一个 Instrumentation 对象，可用于注册 ClassFileTransformer 

```java
package com.example;

import java.lang.instrument.Instrumentation;

public class ExampleAgent {
  public static void premain(String args, Instrumentation instrumentation){
      // 注册 ClassFileTransformer
      ExampleTransformer transformer = new ExampleTransformer();
      instrumentation.addTransformer(transformer);
  }
}
```



需要在 Manifest 中标记，可用 gradle.build 的 jar.manifest 写入

```groovy
jar {
    manifest {
        attributes([
                "Premain-Class": "com.example.ExampleAgent"
        ])
    }
}
```



#### ClassFileTransformer 接口

存在一个  transform  方法，返回修改后的 class 文件的 byte 数组，在 ClassLoader#defineClass 中调用

```java
package com.example;

import java.lang.instrument.ClassFileTransformer;
import java.lang.instrument.IllegalClassFormatException;
import java.security.ProtectionDomain;

public class ExampleTransformer implements ClassFileTransformer {
    @Override
    public byte[] transform(ClassLoader loader,
                            String className,
                            Class<?> classBeingRedefined,
                            ProtectionDomain protectionDomain,
                            byte[] classfileBuffer) throws IllegalClassFormatException {
        //TODO: 可以在此写ASM代码
        return classfileBuffer;
    }
}
```

- loader：类加载器
- className：类名
- classBeingRedefined：为 returnsform 设计，原 class 对象
- protectionDomain：保护域
-  classfileBuffer：原 class 文件 byte 数组

#### 安装方法

在 JVM 参数中设定：

*-javaagent:filename=args*

-  filename：jar 包名
-  premain：传入 premain 的参数

#### 弊端

1. 需要修改 JVM 参数以启用
2. 优先度低，一般是最后一个对类进行修改

### LaunchWrapper

为修改老版本游戏对新版本游戏（1.6 and later）文件夹变化的兼容，产生了 LaunchWrapper

LaunchWrapper 主要由 Forge 团队制作，大量参考 FML Relauncher，被用于 1.6-1.12.2 的加载

不支持 Java 9 及以上版本

#### ITweaker

```java
package com.example;

import java.io.File;
import java.util.Arrays;
import java.util.ArrayList;
import java.util.List;

import net.minecraft.launchwrapper.ITweaker;
import net.minecraft.launchwrapper.LaunchClassLoader;

public class ExampleTweaker implements ITweaker {

    private String[] args;

    public void acceptOptions(List<String> args, File gameDir, File assetsDir, String profile) {
        String[] additionArgs = {"--gameDir", gameDir.getAbsolutePath(), "--assetsDir", assetsDir.getAbsolutePath(), "--version", profile};
        List<String> fullArgs =  new ArrayList<String>();
        fullArgs.addAll(args);
        fullArgs.addAll(Arrays.asList(additionArgs));
        this.args = fullArgs.toArray(new String[fullArgs.size()]);
    }

    public void injectIntoClassLoader(LaunchClassLoader classLoader) {
        classLoader.registerTransformer("com.example.ClassTransformer");
    }

    public String getLaunchTarget() {
        return "net.minecraft.client.main.Main";
    }

    public String[] getLaunchArguments() {
        return args;
    }
}
```

- acceptOptions：接收 Minecraft 启动参数
  - args：启动参数
  - gameDir：游戏路径
  - assetsDir：assets 文件夹
  - profile：版本名称
- injectIntoClassLoader：可使用 registerTransformer 进行注册  IClassTransformer  等操作  LaunchClassLoader  的行为
  - classLoader：将要加载 Minecraft 的 LaunchClassLoader 对象
- getLaunchTarget：返回需要启动游戏的主类，以第一个 Tweaker 为准
  - 一般为 net.minecraft.client.main.Main
- getLaunchArguments：返回 Minecraft 的启动参数，需要返回 MC 启动所需的所有参数

#### IClassTransformer

LauncherWrapper 提供的接口，用于修改 class 文件

```java
package com.example;

import net.minecraft.launchwrapper.IClassTransformer;

public class ClassTransformer implements IClassTransformer {
    public byte[] transform(String name, String transformedName, byte[] basicClass) {
        return basicClass;//特别注意需要返回basicClass
    }
}
```

- transform：返回修改后的 class 文件 byte 数组
  - name 为原类名。name 和 basicClass 使用 notch 名
  - transformedName 只有  IClassNameTransformer 存在时才会与 name 不同，为反混淆后的名称
    - 具体反混淆方式取决于 IClassNameTransformer 实现，如 FML 会将其反混淆成 MCP 名
    - basicClass 可能已被其他 IClassNameTransformer 修改过
  - **切记无论如何都要返回一个有效的 byte 数组，否则可能会引发 ClassNotFoundException，NoClassDefFoundError 等错误**

#### 安装

- 使用库文件挂载

  - 在 libraries 中加入自己编写的 Tweaker 和 LaunchWrapper

    ```json
    {
      "name": "com.example:ExampleTweaker:1.0"
    },
    {
      "name":"net.minecraft:launchwrapper:1.12"
    }
    ```

  - 修改 mainClass 为  net.minecraft.launchwrapper.Launch

  - arguments 中添加： -tweakClass com.example.ExampleTweaker

- 使用 ModLoader

  - 在 Manifest 中添加：

     TweakClass： com.example.AnotherExampleTweaker

  - 将 jar 包放入 mod 文件夹即可

  - FML 与 LiteLoader 的 1.6-1.12.2 版本都支持

#### 弊端

- 一般需要自带映射表或动态映射来进行版本兼容，不能保证运行时一定存在  IClassNameTransformer
- 使用 ModLoader 和 库文件加载 难以使用同一个 Tweaker，处理参数方式不同
- 参数中多个 tweakClass 可能出错，Tweaker 重复处理或不处理启动参数均会异常
- 库文件挂载麻烦，最终仍依赖 ModLoader

### ModLauncher

cpw 制作，用于 1.13FML 加载的库

会记录每个类被  ITransformationService 和  ILaunchPluginService 请求修改的历史

#### ITransformationService

用于预处理环境并提供 ITransformer 实例的接口

```java
package com.example;

import cpw.mods.modlauncher.api.IEnvironment;
import cpw.mods.modlauncher.api.ITransformationService;
import cpw.mods.modlauncher.api.ITransformer;

import joptsimple.OptionResult;
import joptsimple.OptionSpecBuilder;

import java.util.Arrays;
import java.util.Set;
import java.util.function.BiFunction;

public class ExampleService implements ITransformationService {
    String name() {
        return "ExampleService";
    }
    void initialize(IEnvironment environment) {}
    void onLoad(IEnvironment env, Set<String> otherServices) throws IncompatibleEnvironmentException {}
    List<ITransformer> transformers() {
        return Arrays.asList(new ExampleTransformer());
    }
}
```

- name：Server 名称

- arguments：传入双参数 Function 指定要读取的游戏参数，第一个代表读取地参数名，第二个为参数描述

  - 最终参数为 Server名.参数名

- argumentValues：读取参数的结果

- initialize：初始化，传入环境。

  **不要调用 Minecraft 和任意 mod 的任何代码，否则造成类被提前加载**

- onLoad：加载 Service 时使用，传入环境及其他 Service 列表

- transformers：返回 ITransformer 实例

#### ITransformer

修改类

```java
package com.example;

import cpw.mods.modlauncher.api.ITransformer;
import cpw.mods.modlauncher.api.ITransformerVotingContext;
import cpw.mods.modlauncher.api.TransformerVoteResult;

import org.objectweb.asm.tree.ClassNode;

import java.util.HashSet;
import java.util.Set;

public class ExampleTransformer implements ITransformer<ClassNode> {
    ClassNode transform(ClassNode input, ITransformerVotingContext context) {
        return input;
    }
    TransformerVoteResult castVote(ITransformerVotingContext context) {
        return TransformerVoteResult.YES;
    }
    Set<Target> targets() {
        return new HashSet<Target>(Arrays.asList(Target.targetClass("abc")));
    }
}
```

- transform：接收 asm 的 node 和相关信息，这个 node 由泛型 T 确定，一般为 ClassNode 或 MethodNode
- castVote：决定是否修改
  - YES：调用 transform 方法
  - NO：跳过这个类修改器
  - RECJECT：引发  VoteRejectedException 错误
  - DEFER：引发  VoteDeadlockException 错误
- targets：被修改目标的 target 列表，需要统一的 Target 类型，只有在这里返回的目标才会调用其他两个方法

#### ServiceLoader

创建  META-INF/services/cpw.mods.modlauncher.api.ITransformationService，写入 ITransformationService 实现类的全类名

#### 安装

- 库文件挂载
  - 在 libraries 中加入 ModLauncher 以及依赖的 Coremod
  - 修改 mainClass 为  cpw.mods.modlauncher.Launcher 
- Forge 加载
  - 从 1.13.2-25.0.216 开始，只要放入 mods 文件夹中即可

## FML Coremod

### 1.3.2-1.5.2

#### Manifest

使用 FMLCorePlugin 指定 IFMLLoadingPlugin 全类名

> FMLCorePlugin: com.example.ExamplePlugin

#### IFMLLoadingPlugin

```java
package com.example;

import java.util.Map;

import cpw.mods.fml.relauncher.IFMLLoadingPlugin;

public class ExamplePlugin implements IFMLLoadingPlugin {

    @Override
    public String[] getLibraryRequestClass() {
        return null;
    }

    @Override
    public String[] getASMTransformerClass() {
        return new String[]{"com.example.ClassTransformer"};
    }

    @Override
    public String getModContainerClass() {
        return null;
    }

    @Override
    public String getSetupClass() {
        return null;
    }

    @Override
    public void injectData(Map<String, Object> data) {
    }

}
```

- getLibraryRequestClass：依赖类
- getASMTransformerClass：IClassTransformer 全类名数组
- getSetupClass：IFMLCallhook 实现类全类名
-  injectData：在 minecraft 启动后调用，可操作 Minecraft 的 class
  - mcLocation：Minecraft 文件夹 File
  - coremodList：Coremod 列表 List
  - coremodLocation：当前 coremod 文件 File

- 辅助注解
  - TransformerExclusions：指定不被 coremod 修改的包名前缀数组
  - MCVersion：指定 coremod 适用 Minecraft 版本

#### IClassTransformer

用于修改 class 文件

```java
package com.example;

import cpw.mods.fml.relauncher.IClassTransformer;

import org.objectweb.asm.ClassReader;
import org.objectweb.asm.ClassWriter;
import org.objectweb.asm.tree.ClassNode;
import org.objectweb.asm.tree.MethodNode;

public class ClassTransformer implements IClassTransformer {
    public byte[] transform(String name, String transformedName, byte[] bytes) {
        if (!"net.minecraft.src.GuiPlayerTabOverlay".equals(transformedName))
            return bytes;

        //使用ASM读入bytes
        ClassReader cr = new ClassReader(bytes);
        ClassNode cn = new ClassNode();
        cr.accept(cn, 0);

        //遍历methods
        for (MethodNode mn : cn.methods) {
            if(!"a".equals(mn.name)) 
                continue;

            //TODO: 在这里进行ASM操作
        }

        //返回修改后的bytes
        ClassWriter cw = new ClassWriter(ClassWriter.COMPUTE_FRAMES | ClassWriter.COMPUTE_MAXS);
        cn.accept(cw);
        return cw.toByteArray();
    }
}
```

- name：原类名
- transformedName：MCP 无混淆类名
- bytes：class 文件 byte 数组
  - 测试时为 MCP 混淆，运行时为 notch 混淆

### 1.6.1-1.12.2

#### 大事记

- 1.6
  - 修改游戏文件结构，单一 .minecraft 可同时使用多个游戏版本
  - 引入 LaunchWrapper
    - FML 通过 libraries 和 -tweakClass 参数安装，使用 LaunchWrapper 加载
    - 通过 LaunchWrapper 提供的 IClassNameTransformer 进行运行时反混淆，将运行时混淆方式从 notch 反混淆成 srg
    - Coremod 不通过 RelaunchClassLoader 而是 LaunchClassLoader 对 class 进行动态修改
    - 取消 coremod 文件夹，CoreMod 可直接放入 mods 文件夹，CoreMod 可直接包含普通 mod
- 1.7
  - 引入 ForgeGradle，Forge 开发脱离 MCP 工具
  - MCP 对 Minecraft 类进行分包，不在存放于 net.minecraft.src
- 1.8：FML 包名从 cpw.mods.fml 变更为 net.minecraftforge.fml

#### 制作

##### manifest

- FMLAT ：自定义  access transformer 修改原版代码的方法和字段的可见性而不需从 coremod 中定义

- FMLCorePlugin：IFMLLoadingPlugin 实现全类名**必备**，使用 . 分包

- FMLCorePluginContainsFMLMod：true/false，标记 jar 包是否还存在普通 mod。若为 false 将不会尝试寻找 @Mod 注解

- ForceLoadAsMod：与 FMLCorePluginContainsFMLMod：true 配合，保证普通 mod 正常加载

- 调试

  ```groovy
  minecraft {
      ...
      clientJvmArgs += "-Dfml.coreMods.load=com.example.ExamplePlugin"
      serverJvmArgs += "-Dfml.coreMods.load=com.example.ExamplePlugin"
  }
  ```

  

##### IFMLLoadingPlugin

```java
package com.example;

import java.util.Map;

import net.minecraftforge.fml.relauncher.IFMLLoadingPlugin;
import net.minecraftforge.fml.relauncher.IFMLLoadingPlugin.Name;

@Name("ExampleCoreMod")
public class ExamplePlugin implements IFMLLoadingPlugin {

    @Override
    public String[] getASMTransformerClass() {
        return new String[]{"com.example.ClassTransformer"};
    }

    @Override
    public String getModContainerClass() {
        return null;
    }

    @Override
    public String getSetupClass() {
        return null;
    }

    @Override
    public void injectData(Map<String, Object> data) {
    }

    @Override
    public String getAccessTransformerClass() {
        return null;
    }

}
```

- getASMTransformerClass：IClassTransformer 实现类全类名数组

- getModContainerClass：ModContainer 全类名，可空

- getSetupClass：IFMLCallhook 实现类全类名，可空

- injectData：与 1.3.2-1.5.2 的 IFMLLoadingPlugin 对应方法相同

  - runtimeDeobfuscationEnabled：运行时反混淆是否被开启

- getAccessTransformerClass：AccessTransformer 实现类全类名，可空

  一般为 null，直接用 FML 提供的访问级转换器（net.minecraftforge.fml.common.asm.transformers.AccessTransformer）

- 注解

  - TransformerExclusions：不被 CoreMod 修改的包名前缀数组
  - MCVersion：适用 Minecraft 版本
  - Name：CoreMod 名称，不指定则用类名。崩溃报告中实现
  - DependsOn：依赖。理论上是让 coremod 在前置存在的时候才能加载，实际上从未实现
  - SortingIndex：调用顺序，越小优先级越高，默认 0

##### IClassTransformer

字节码修改器

```java
package com.example;

import net.minecraft.launchwrapper.IClassTransformer;

import net.minecraftforge.fml.common.asm.transformers.deobf.FMLDeobfuscatingRemapper;

import org.objectweb.asm.ClassReader;
import org.objectweb.asm.ClassWriter;
import org.objectweb.asm.tree.ClassNode;
import org.objectweb.asm.tree.MethodNode;

// 修改 net.minecraft.client.gui.GuiPlayerTabOverlay的func_175249_a(srg) renderPlayerlist(mcp)方法
public class ClassTransformer implements IClassTransformer {
    public byte[] transform(String name, String transformedName, byte[] basicClass) {
        if (!"net.minecraft.client.gui.GuiPlayerTabOverlay".equals(transformedName))
            return basicClass;

        //使用ASM读入basicClass
        ClassReader cr = new ClassReader(basicClass);
        ClassNode cn = new ClassNode();
        cr.accept(cn, 0);

        //遍历methods
        for (MethodNode mn : cn.methods) {
            //调用FML接口获得方法名，运行时获得的是srg，测试时获得的是mcp
            String methodName = FMLDeobfuscatingRemapper.INSTANCE.mapMethodName(name, mn.name, mn.desc);
            if(!"func_175249_a".equals(methodName) && !"renderPlayerlist".equals(methodName)) 
                continue;

            //TODO: 在这里进行ASM操作
        }

        //返回修改后的bytes
        ClassWriter cw = new ClassWriter(ClassWriter.COMPUTE_FRAMES | ClassWriter.COMPUTE_MAXS);
        cn.accept(cw);
        return cw.toByteArray();
    }
}
```

与 1.3.2-1.5.2 的 IClassTransformer 相似

#### 安装

直接丢 mods 文件夹

### 1.13.2+

#### 引入 ModLauncher

- LaunchWrapper 被舍弃
- Mod 加载机制被重写
- Forge 实现部分模块化，由 CoreMods 模块调度运行 CoreMod
- 修改 class 行为需要在 JavaScript 脚本中受限进行，几乎无法使用 ASM Core API，转而使用 ASM Tree API
- 大量使用 Java8 新特性
- 不再进行运行时反混淆，而是在安装过程中调用 SpecialSource 生成 srg 混淆的 Minecraft 核心文件
- binpatch 同样使用 srg 混淆，运行时对 Minecraft 进行修改
- 1.13.2-25.0.199 更新 ModLauncher 1.1，不再使用 URLClassLauncher 进行 Mod 加载，崩溃报告中会显示类修改来源
- 1.13.2-25.0.216 增加 ModDirTransformerDiscoverer，ModLauncher 的 ITransformationService 可被加载

#### Java or JavaScript

从 1.13.2-25.0.216 开始，Forge 开始尝试加载仅实现 ITransformationService 的 Mod，自此可以用 Java 进行 CoreMod 编写

在 1.14.4-28.0.17 中，Forge 实现了一个 ITransformationService，使用 JavaScript 进行重写

| 项目     | Java                                                         | JavaScript                                              |
| -------- | ------------------------------------------------------------ | ------------------------------------------------------- |
| 可追踪性 | 在崩溃报告和调试信息中显示类修改来源的转换器名，需要规范命名 | 在崩溃报告和调试信息中显示类修改来源的 modid 和转换器名 |
| 编写难度 | 手动实现 ITransformationService，ITransformer，声明 ServiceLoader | 只需要按照格式声明要修改的内容并完成修改的函数          |
| 代码提示 | 完整                                                         | 仅基本语法提示                                          |
| 操作风险 | 不能在 CoreMod 阶段调用任何原版和其他 mod，需要开发者注意    | 不会导致类提前加载                                      |
| 自由程度 | 除 Java 安全级别，完全自由                                   | 受限的类访问，只能使用 Java 基本类，ASM 与 ASMAPI       |
| ASM      | Tree API，Core API                                           | Tree API                                                |

#### Java

参考 ModLauncher 部分

#### JavaScript

##### coremods.json

位于 META-INF/coremods.json

```javascript
{
    "脚本名称":"脚本路径"
    ...
}
```

*脚本路径的根目录为 jar 包根目录*

```javascript
{
  "example": "example.js"
}
```



##### JavaScript

在 js 中，只能完成修改类这一个操作，避免了 CoreMod 调用 Minecraft 和普通 Mod 等导致错误

*Nashorn 引擎在 Java11 已弃用*

```javascript
// 修改net/minecraft/client/gui/GuiPlayerTabOverlay中func_175249_a(srg)方法
var ASMAPI = Java.type('net.minecraftforge.coremod.api.ASMAPI');
var Opcodes = Java.type('org.objectweb.asm.Opcodes');

function initializeCoreMod() {
    return {
        'PlayerTabTransformer': {
            'target': {
                'type': 'CLASS',
                'name': 'net/minecraft/client/gui/GuiPlayerTabOverlay'
            },
            'transformer': function (cn) {
                //遍历ClassNode下的methods
                cn.methods.forEach(function (mn) {
                    if (mn.name === 'func_175249_a') {
                        //请在这里对ClassNode和MethodNode进行ASM操作
                    }
                });

                //返回修改后的ClassNode对象
                return cn;
            }
        }
    };
}
```

- type：目标类型
  - 1.14.3-27.0.16 增加 FIELD 和 METHOD
  - 之前版本仅有 CLASS
- transformer：修改类
  - 参数：与 type 对应的 Node 对象（ClassNode，FieldNode，MethodNode）
  - 返回值：修改完后的 Node 对象
  - 函数会在目标类被加载时调用
  - 若不需要 ClassNode，尽量使用 METHOD 作为 type，减少遍历代码
- Java.type
  - 相当于 import
  - 仅能载入 ASM 库中的类和 Forge 提供的 ASMAPI