---
语言: cpp
语法类型: 标准库
---
智能指针
* `unique_ptr<T>`：唯一指针，不允许任何复制构造调用；允许使用 `std::move` 语义移动指针，但移动后原指针失效
* `shared_ptr<T>`：记录指向同一对象的所有指针，当引用计数归零时对象自动删除
* `weak_ptr<T>`：引用 `shared_ptr<T>` 的引用，但不影响其引用计数，可避免循环引用
    * 若 `shared_ptr<T>` 被释放，则对应 `weak_ptr<T>` 全部失效
* #cpp20 `atomic` 指针：原子智能指针：`atomic<shared_ptr>`，`atomic<weak_ptr>` 等

对于智能指针，可使用带有 `_pointer` 后缀的类型转换符，转换失败则返回一个指向 `nullptr` 的智能指针
* `static_cast` -> `static_pointer_cast`
* `dynamic_cast` -> `dynamic_pointer_cast`
* `const_cast` -> `const_pointer_cast`
