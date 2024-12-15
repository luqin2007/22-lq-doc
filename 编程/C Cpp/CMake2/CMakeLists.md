CMake 脚本位于 `CMakeLists.txt` 中，基本结构为：

```cmake
# 设置 CMake 最低版本
cmake_minimum_required(VERSION <minimum-cmake-version>)
# 设置项目名
project(<project-name>)

# 添加可执行文件
add_executable(<target> <dependencies>)
```
