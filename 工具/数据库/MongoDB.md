# 存储结构
## 数据库

数据库名称使用小写字符命名，名称长度不超过 64 字符，不得包含 ` `，`.`，`$`，`/`，`\` 和 `\0`

MongoDB 有几个默认数据库：

| 数据库名称    | 作用                               |
| -------- | -------------------------------- |
| `admin`  | 在此数据库中操作默认具有管理员权限，特定操作也只能在此数据库运行 |
| `local`  | 只存储在各自服务器实例上，存储本地单台服务器数据，不会被复制   |
| `config` | 分布式分片配置                          |
| `test`   | 默认连接的数据库                         |
## 集合

Collection，一组文档的集合，相当于关系型数据库的表 Table，但没有固定模式，不受关系范式约束。

>[!info] 集合名不能以 `system.` 为前缀，也不能是保留字符 `$`
## 文档

Document，相当于关系型数据库的一行数据。文档以嵌套键值树形式 BSON 数据的格式存储，文档除数据值外还存储了元数据。

> [!note] BSON：Binary Serialized Document Format，MongoDB 的数据存储和网络传输格式。
> 二进制存储的 JSON，在 JSON 的基础上将元素长度保存于头部信息，提高遍历速度
> - 键：字符串，不能是 `$`，不能以 `.` 开头，不能是 `null`
> - 值：JSON 基本类型，文档对象，数组，**其他类型对象**
> ```json
> {
>     "_id": ObjectId("xxxxxxx"),
>     "name": zhangsan,
>     "address": {
>         "city": "Beijing",
>         "code": 100876
>     },
>     "scores": [
>         { "name": "English", "grade": 3.0 },
>         { "name": "Chinese", "grade": 4.0 }
>     ]
> }
> ```

MongoDB 会为每个文档自动添加一个键 `_id` 作为主键，类型为 `ObjectId`，值唯一且不可变。

> [!note] ObjectId：12 字节文档唯一标识，通常以 16 字节表示
> 前 9 字节保证同一秒不同机器、不同进程产生的记录标识唯一，后 3 字节保证同一进程产生的标识唯一
> - 4 字节：时间戳，记录 `_id` 产生时间
> - 3 字节：机器识别码，主机的唯一标识码，通常为主机名的 Hash 值
> - 2 字节：进程标识符 PID
> - 3 字节：计数器，同一进程不断自增

> [!attention] `_id` 可以设置为其他类型，但需要保证其唯一性

文档之间的关系包括嵌套和引用链接两种
- 嵌套：`Embed`，文档内包含其他文档。嵌套数量和深度没有限制，但文档最大限制为 16MB
- 引用链接：`Link`，一个 `DBRef` 对象，更接近外键的概念，通过解引用可以链接到另一个文档
# 数据库工具与操作

MongoDB 提供 Server，Shell，Compass，Atlas CLI 等工具
- Server：MongoDB 数据库服务端
- Shell：开源命令行操作界面
- Compass：GUI 工具
- Atlas CLI：统一命令行界面，管理 MongoDB Atlas

| 操作命令                                    | 功能说明                                                                                |
| --------------------------------------- | ----------------------------------------------------------------------------------- |
| show dbs                                | 查询显示所有数据库                                                                           |
| show users                              | 显示当前DB所有用户                                                                          |
| show collections                        | 显示当前DB所有集合                                                                          |
| db.getName()，db                         | 可以显示当前数据库对象或集合                                                                      |
| use <dbname>                            | 切换/创建数据库                                                                            |
| db.help()                               | 显示可用数据库操作命令帮助信息                                                                     |
| db.dropDatabase()                       | 删除当前使用的数据库                                                                          |
| db.runCommand(cmdObj)                   | 执行数据库中的命令，cmdObj 为命令字符串                                                             |
| db.cloneDatabase(fromhost)              | 将指定机器上数据库的数据克隆到当前数据库                                                                |
| db.copyDatabase(fromdb, todb, fromhost) | 数据库复制，如 `db.copyDatabase("mydb", "temp", "127.0.0.1");` 表示将本机 mydb 的数据复制到 temp 数据库中 |
| db.repairDatabase()                     | 修复当前数据库                                                                             |
| db.stats()                              | 显示当前 db 状态，包含集合数、空间占用情况等                                                            |
| db.version()                            | 显示当前 db 版本                                                                          |
| db.getMongo()                           | 查看当前 db 的链接机器地址                                                                     |
| db.getCollectionNames()                 | 查看当前 db 所有集合名称                                                                      |
| db.getCollectionInfos([filter])         | 查看满足过滤条件集合的详细信息                                                                     |
| db.createUser(userDocument)             | 添加用户                                                                                |
| db.auth(username, password)             | 数据库认证                                                                               |
| db.dropUser(username)                   | 删除用户                                                                                |
> [!attention] 使用 `use <数据库名>` 创建数据库时，必须对数据库进行操作才能真正创建。仅 `use` 不会创建
> 
# 集合操作

| 操作命令                                | 功能说明                            |
| ----------------------------------- | ------------------------------- |
| db.createCollection (name, options) | 创建一个集合，集合名称需要 `""`，`options` 可选 |
| db.getCollection(name)              | 获取指定名称的 Collection，结果会加上数据库的名称  |
| db.getCollectionNames()             | 获取当前 db 所有 Collection 名称数组      |
| db.printCollectionStats()           | 显示当前 db 所有 Collection 索引等状态信息   |
| db.<集合名>.help()                     | 显示集合操作的相关命令提示                   |
| db.<集合名>.drop()                     | 删除集合                            |
| db.<集合名>.dataSize()                 | 查看集合数据占用空间大小                    |
>[!note] 创建文档时，若集合不存在，MongoDB 会自动创建集合
## 定长集合

定长集合多用于重点存储近期业务的场景，只关心最近数据

```js
db.createCollection("集合名", { capped: true, size: 6142800, max: 10000 })
```

- `capped`：是否为定长集合
- `size`：单个文档占用的字节数
- `max`：包含文档的最大数量
# 文档操作
## 文档参考

```cardlink
url: https://mongodb.net.cn/manual/
title: "参考_MonogDB 中文网"
host: mongodb.net.cn
```

## 数据类型

除 JSON 的基本数据类型外，MongoDB 还额外提供以下类型：

| 类型名称       | 说明                                               |
| ---------- | ------------------------------------------------ |
| ObjectId   | 12 字节的字符串，用于唯一标识文档，如 ` { "x": objectld() }`      |
| Date       | 日期毫秒数，不含时区，如 `{ "x": new Date() }`               |
| Timestamp  | 时间戳                                              |
| Regex      | JavaScript 正则表达式，如 `{ "x": /[abc]/ }`，可用于查询条件    |
| BinaryData | 二进制数据，任意字节的字符串，可用于存储非 UTF-8 字符串                  |
| Min key    | BSON 中的最小值                                       |
| Max key    | BSON 中的最大值                                       |
| JavaScript | 任何 JavaScript 代码，如 `{ "x": function() {/*…*/} }` |
| Array      | 数组，如 `{ "x": ["a", "b", "c"] )`                  |
| Object     | 嵌套文档，被嵌套的文档作为值来处理，如 `{ "x": { "y": 3 } }`        |
每个数据类型都包含一个数字和别名，可通过 `$type` 查询对应文档的 BSON 类型
## 操作符
### 常用操作符

| 类型   | 操作符                            | 说明                    | 操作示例                                                                           |
| ---- | ------------------------------ | --------------------- | ------------------------------------------------------------------------------ |
| 元素比较 | $lt, $lte, $gt, $gte, $eq, $ne | 大小，相等性判断              | db.c.find({ "a": { $gt: 100 } })                                               |
|      | $in, $nin                      | 属于，不属于                | db.c.find({ j: { $in: [2.4.6] } })                                             |
| 元素查询 | $exists                        | true / false          | db.c.find({ a: { $exists : true } })                                           |
|      | $type                          | 按数据类型查询               | db.c.find({ a: { $type : 2 } })                                                |
| 模式评估 | $mod                           | 取模                    | db.c.find({ a: { $mod: [10, 1] } }) 表示 a%10==1                                 |
|      | $regex                         | 正则匹配                  | db.c.find({ sku: { $regex: /^ABC/i } })                                        |
|      | $where                         | 弥补其他方式无法满足的查询条件，但效率较低 | db.c.find({ "$where": "this.a>3" })                                            |
|      | $text                          | 文本搜索，要先建索引            | db.c.find({ $text: { $search: "coffee" } })                                    |
| 逻辑运算 | $and, $or, $nor                | 与，或，异或                | db.c.find({ name: "bob", $nor: [ {a:1}, {b:2} ] })                             |
|      | $not                           | 非运算                   | db.c.find({ name: "bob", $not: [ {a:1}, {b:2} ] })                             |
| 数据查询 | $size                          | 匹配数组长度                | db.c.find({ a: { $size: 1 } })                                                 |
|      | $all                           | 数组中的元素是否完全匹配          | db.c.find({ a: { $all: [2, 3] } })                                             |
|      | $elemMatch                     | 数组中的元素是否都满足条件         | db.c.find({ product: { "$elemMatch": { shape: "square", color: "purple" } } }) |
| 位查询  | $bitsAllClear                  | 所有位的值为0               | db.c.find({ a: { $bitsAllClear: [1,5] } })                                     |
|      | $bitsAIlSet                    | 所有位的值为1               | db.c.find({ a: { $bitsAllSet: [1,5] } })                                       |
|      | $bitsAnyClear                  | 部分位的值为0               | db.c.find({ a: { $bitsAnyClear: [1,5] } })                                     |
|      | $bitsAnySet                    | 部分位的值为1               | db.c.find({ a: { $bitsAnySet: [1,5] } })                                       |
### 其他操作符

- 地理空间操作符：`$minDistance`，`$maxDistance`，`$box`，`$center`，`$geoWithin`，`$near` 等
- 类型转换符
	- `double`，`string`，`objectId`，`date`，`integer`，`long`，`decimal` 类型转换
	- `$convert` 聚合类型转换符
## 插入

> [!note] 插入数据时，若 collection 不存在，Mongodb 会自动创建对应 Collection
`````col
````col-md
flexGrow=1
===
```json
db.<Collection 名>.insertOne("Json 数据对象")
db.<Collection 名>.insertMany(["Json 数据对象"])
```
````
````col-md
flexGrow=1
===
![[../../_resources/images/Pasted image 20241101120939.png]]
````
`````
## 查询

- 获取所有文档

```js
db.<collection-name>.find()
db.<collection-name>.find({})
```

- 查询所有文档，并以树形结构显示

```js
db.<collection-name>.find().pretty()
db.<collection-name>.find({}).pretty()
```

- 按条件查询：`find`，`findOne`
`````col
````col-md
flexGrow=1
===
```js
db.<collection-name>.find({ condition })
```
````
````col-md
flexGrow=1
===
![[../../_resources/images/Pasted image 20241101124559.png]]
````
`````
- 筛选
	- `limit(n)`，`skip(n)`
	- `sort({key:1/-1})`：按某键排序，1 为升序，-1 为降序，可以有多个关键字
## 删除

- 删除所有文档：`db.<collection-name>.drop()`
- 删除特定文档：`deleteOne`，`deleteMany`，`findOneAndDelete`

`````col
````col-md
flexGrow=1
===
```js
db.ct.deleteOne({ "Name": "Tian" })
```
````
````col-md
flexGrow=1
===
![[../../_resources/images/Pasted image 20241101123036.png]]
````
`````

`````col
````col-md
flexGrow=1
===
```js
db.ct.deleteOne({ "Age": { $gt: 25 } })
```
````
````col-md
flexGrow=1
===
![[../../_resources/images/Pasted image 20241101123228.png]]
````
`````
## 修改

`````col
````col-md
flexGrow=1
===
```js
db.<collection-name>.updateOne(<query>, <update>, {
	upsert: <bool>,
	writeConcern: <document>,
	collection: <document>,
	arrayFilters: [条件],
	hint: doc|str
})
```
````
````col-md
flexGrow=1
===
```js
db.<collection-name>.updateMany(<query>, <update>, {
	upsert: <bool>,
	writeConcern: <document>,
	collection: <document>,
	arrayFilters: [条件],
	hint: doc|str
})
```
````
`````
- `<query>`：查询条件，一个对象，同 `find`
- [ ] 111
## 游标
## 语句块
## 链接引用
## 管道聚合
# 索引
# 数据库架构
# 管理与监控