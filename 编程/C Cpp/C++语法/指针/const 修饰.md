**判断方法：靠近谁，`const` 修饰谁**
# 指向常量的指针

对象本身不可修改，指针地址可以修改

```cpp
int value {5};
const int *pvalue { &value }; // const 修饰 int 类型
*pvalue = 6; // 错误： pvalue 指向 const int 类型
pvalue = nullptr; // 正确
```
# 指向对象的常量指针

对象本身可以修改，指针地址不可修改

```cpp
int value {5};
int *const pvalue { &value }; // const 修饰 pvalue 变量本身
*pvalue = 6; // 正确
pvalue = nullptr; // 错误
```
# 指向常量的常量指针

对象本身和指针地址都不可修改

```cpp
int value {5};
const int *const pvalue { &value };
*pvalue = 6; // 错误
pvalue = nullptr; // 错误
```
