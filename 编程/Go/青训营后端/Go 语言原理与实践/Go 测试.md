测试是避免事故的最后一道屏障，几种测试成本依次降低，但覆盖率依次增加
- 回归测试：QA 通过终端模拟日常使用场景
- 集成测试：对系统功能（接口等）做测试验证
- 单元测试：开发阶段开发者对单独模块做验证
# 单元测试

单元测试可以保证代码的质量和效率
- 代码功能上的正确性，新代码本身正确性及未破坏原代码正确性
- 提高效率，代码出现错误时可在较短的周期内定位和修改

![[../../../../_resources/images/Pasted image 20241113000927.png]]

1. 测试代码应以 `_test.go` 结尾
2. 测试函数命名为 `func TestXxx(*testing.T)`
	- 测试失败时，使用 `testing.T::Errorf` 等函数抛出一个异常即可
3. 测试入口为  `TestMain(*testing.M)`
	- 执行所有测试函数：`testing.M::Run()` 函数
4. 使用 `go test 项目文件名 [<flags>] [<packages>]` 运行测试

`TestMain` 函数中可以进行初始化、释放资源等操作

```go
func TestMain(m *testing.M) {
     // do something before all testing
     code :=  m.Run()
     // do something after all testings
     os.Exit(code)
}
```

很多开源 Assert 包都可以提供 `Equal` 等函数简化比较操作，如 `github.com/stretchr/testify/assert` 等
## 覆盖率

代码覆盖率是评估单元测试对代码的覆盖度，在运行测试时添加 `--cover` 参数即可
- 一般覆盖项目主流程测试，覆盖率在 50%-60%，对于资金交易等覆盖率要求 80%+
- 测试分支应互相独立，全面覆盖，不重不漏
- 测试单元粒度应足够小，函数粒度足够小，函数应保证单一职责
## 组件依赖

![[../../../../_resources/images/Pasted image 20241113002512.png]]

组件依赖应尽量满足幂等和稳定性
- 幂等性：重复运行测试时，输出的内容应保持一致
- 稳定性：单元测试相互隔离，可以在任何时间对任何函数进行测试
# Mock 测试

常用 Mock 包如 `github.com/bouk/monkey`，可以为函数或方法打桩（替换）
- 打桩：`monkey.Patch(<target>, <replace>)`
- 还原：`monkey.Unpatch(<target>)`，通常使用 `defer`

> [!note] 原理：在运行时使用 `unsafe` 将原本函数地址替换成其他打桩函数地址

通过 Mock 测试可脱离测试程序对上下文环境的强依赖
# 基准测试

测试项目对 CPU、内存等资源的消耗，用于代码的性能分析，与单元测试相比测试函数命名为 `func BenchmarkXxx(*testing.B)`
- `B::ResetTimer()` 可以重置计时器，用于消除初始化环境的误差
- 直接调用函数进行串行测试，需要执行 `B::N` 次待测函数
- 通过 `B::RunParallel(func(*testing.PB))` 进行并行测试，传入的函数中检查 `PB::Next()` 是否测试完成

![[../../../../_resources/images/Pasted image 20241113003609.png]]

测试结果为 `ns/op`

>  使用 FastRand 替代 Rand 可大幅度提高效率

# 项目实践思路 

## 需求分析

- 需求用例、ER 图
- 分层结构 ：
	- Model：数据层，封装外部数据的增删改查（Repository，Model），负责数据
	- Entity：逻辑层，处理核心逻辑输出（Service，Entity）
	- View：视图层，处理外部交互逻辑（Controller，View），负责交互
- Web 框架：`gopkg.in/gin-gonic/gin`：`r := gin.Default()`
	- 构建路由：`r.Get/Post/...(<path>, func(*gin.Context))`
	- 运行：`r.Run()`

---

- 使用 `map` 初始化索引
- `sync.Once` 用于仅执行一次的代码，可用于初始化单例模式

![[../../../../_resources/images/Pasted image 20241113004420.png]]

