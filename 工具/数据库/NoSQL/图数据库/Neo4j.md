# Cypher

> [!note] CQL：Cypher Query Language，Neo4j 的查询语言，一种声明式模式匹配语言

- 语句以 `;` 结尾
- 使用 `$` 引用参数/变量
- 使用 `//` 表示注释
## 数据元素

节点使用 `()` 表示，括号内为变量、属性、标签等，空表示一个匿名节点
- 直接一个名字表示将节点赋值给同名变量
- 节点属性使用 `{key1:value1,key2:value2}` 的形式声明
- 节点标签使用 `:label` 的形式声明

> [!note] 每个节点有一个唯一 ID 属性，创建节点由数据库时自动添加
> - `id(n)` 返回唯一整数 id，已弃用
> - `elementId(n)` 返回唯一字符串 id

使用 `-[]->`、`<-[]-` 表示有向关系（路径），`-[]-` 表示无向关系（路径）
- 关系两端点是一个节点
- `[]` 之间可以添加变量、属性、标签等，规则与[[#节点]]相同
- 关系之间 `[]` 为空时可以省略
### 关键字

> [!hint] 类似 SQL，关键字大小写不敏感，习惯性大写

| 语句             | 含义   | 说明                      |
| -------------- | ---- | ----------------------- |
| create         | 创建   | 创建节点、创建关系               |
| match          | 匹配   | 用于匹配节点、关系等              |
| optional match | 可选匹配 | 可选匹配节点、关系等              |
| return         | 返回   | 返回结果                    |
| where          | 条件   | 条件语句                    |
| delete         | 删除   | 删除节点和关联关系               |
| remove         | 移除   | 删除标签和属性                 |
| set            | 修改   | 添加或更新属性                 |
| order by       | 排序   | RETURN子句，按升序或降序对行进行排序数据 |
| skip           | 跳过   | RETURN子句，跳过n条结果         |
| limit          | 限制   | RETURN子句，限制返回n个结果       |
| unique         | 唯一   | 唯一性约束                   |
| foreach        | 循环   | 循环处理每一个结果               |
| merge          | 合并   | 合并的方式创建节点、关系            |
| with           | 具有   | 使用聚合函数过滤结果时可用于连接不同的子句   |
| case           | 选择   | 可作为RETURN子句，用于多路选择      |
| unwind         | 展开   | 将一个列表展开为一个序列            |
| union          | 组合   | 将多个查询结果合并为一个结果          |
| call           | 调用   | 调用存储过程                  |
### 函数

```tabs
tab: 字符串

| 函数名称      | 功能描述                |
| --------- | ------------------- |
| toUpper   | 用于将字符串转换成大写         |
| toLower   | 用于将字符串转换成小写         |
| substring | 结束的整个子串             |
| replace   | 用于将某个字符串中的子串替换成目标子串 |
| left      | 返回原字符串左边指定长度的子串     |
| right     | 返回原字符串右边指定长度的子串     |
| ITrim     | 返回去掉左侧空格后的字符串       |
| rTrim     | 返回去掉右侧空格后的字符串       |
| trim      | 返回去掉两侧空格后的字符串       |
| split     | 返回以指定模式分隔后的字符串序列    |
| reverse   | 返回原字符串的倒序字符串        |
| toString  | 将参数转换成字符串返回         |

tab: 标量

| 函数名称       | 功能描述                     |
| ---------- | ------------------------ |
| size       | 返回列表元素个数，或者结果集元素个数       |
| length     | 返回路径长度，或者字符串长度           |
| id         | 返回关系或节点的id               |
| timestamp  | 返回当前时间                   |
| properties | 返回关系或节点的 Map 类型属性的KV列表   |
| toInteger  | 将浮点数或字符串转换成整数，如果失败返回null |
| toFloat    | 将整数或字符串转换成浮点数，如果失败返回null |

tab: 数学

| 函数名称  | 功能描述                      |
| ----- | ------------------------- |
| abs   | 返回参数的绝对值                  |
| ceil  | 返回大于等于参数的最小整数             |
| floor | 返回小于等于参数的最大整数             |
| round | 返回四舍五入整数                  |
| sign  | 返回参数的符号，正数返回1，0返回0，负数返回一1 |
| rand  | 返回[0，1)之间的一个随机数           |
| sgrt  | 返回参数平方根                   |

tab: 列表

| 函数名称      | 功能描述                                                                                        | 
| ------------- | ----------------------------------------------------------------------------------------------- |
| nodes         | 返回路径的所有节点                                                                              |
| labels        | 返回节点的所有标签                                                                              |
| head          | 返回列表的第一个元素                                                                            |
| tail          | 返回列表中除了首元素之外的所有元素                                                              |
| range         | 返回某个范围内的数值列表                                                                        |
| keys          | 以字符串列表形式返回节点、关系的所有 key值，即属性名称列表                                      |
| reduce        | 对列表中元素执行一个表达式，并将表达式结果存入一个累加器 变量，参数中可以设置累加器变量的初始值 |
| relationships | 返回路径参数中的所有关系                                                                        |

tab: 聚合

| 函数名称  | 功能描述    |
| ----- | ------- |
| count | 用于返回计数值 |
| max   | 用于返回最大值 |
| min   | 用于返回最小值 |
| sum   | 用于返回和   |
| avg   | 用于返回平均值 |

tab: 节点关系

| 函数名称   | 功能描述           |
| --------- | ----------- |
| startNode | 用于返回关系的开始节点 |
| endNode   | 用于返回关系的结束节点 |
| type      | 用于返回关系的类型标签 |

tab: 断言

| 函数名称   | 功能描述           |
| ------ | -------------- |
| all    | 是否都满足断言条件      |
| any    | 是否至少一个满足断言条件   |
| none   | 是否都不满足断言条件     |
| single | 是否有且只有一个满足断言条件 |
| exists | 参数内容是否存在       |

```

>[!note] 除自带的函数和存储过程外，Neo4j 也提供 [APOC](https://neo4j.com/docs/apoc/current/installation/) 用于复杂图数据处理与分析
>- 下载与 Neo4j 版本相同的 APOC 版本
>- 将下载的 `jar` 放入 `plugins` 目录中
>- 运行 `return apoc.version()` 查看插件版本号
>
>![[../../../../_resources/images/Pasted image 20241021235713.png]]
### 数据类型

基本数据类型包括 `boolean`，`byte` - `long`，`float`，`double`，`char`，`string`，在此之上还支持 `Map` 和 `List` 两种容器。

Neo4j 不支持 `datetime` 等表示时间日期的类型，可以通过系统提供的函数创建
- `date()`：创建 `yyyy-MM-dd` 格式的时间字符串
- `timestamp()`：获取当前时间的毫秒值（`System.currentTimeMillis()`）
- `apoc.data.format()`：APOC 库提供的日期格式化工具
## 元素操作

### 创建

使用 `CREATE` 创建节点和关系

```cypher
create(节点名)
create(节点名:标签{属性列表})
create(起点) -[变量:标签{属性列表}]-> (终点) return 变量;
create(终点) <-[变量:标签{属性列表}]- (起点) return 变量;
create(端点1) -[变量:标签{属性列表}]- (端点2) return 变量;
```

> [!note] 创建关系时，节点不存在的情况下会自动创建节点

---

```cypher
create(dept:Dept{deptno:10,dname:"Accounting",location:"Beijing"});
create(andy:Person:Student:Writer{name:'andy',age:23});
```

`````col
````col-md
flexGrow=1
===
带有 `RETURN 节点名` 可以将节点返回，以图形形式展示

```cypher
create(n:Person{name:"WuJing",born:1974}) return n;
```
````
````col-md
flexGrow=1
===
![[../../../../_resources/images/Pasted image 20241021231938.png]]
````
`````
使用 `UNWIND` 指令可以批量添加多个图对象

`````col
````col-md
flexGrow=2
===
```cypher
unwind[{name: "Alice",age:32}, {name:"Bob",age:42}] as row  
create(n:Person)  
set n.name=row.name, n.age=row.age  
return n;
```
````
````col-md
flexGrow=1
===
```cypher
UNWIND[对象列表] as row
create(n:标签)
set 属性
return n;
```
````
`````
也可以通过 `WITH` 配合 `FOREACH` 流式的处理数据，添加节点

```cypher
with ["a", "b", "c"] as coll  
foreach (value in coll | create(:person{name:value}));
```

> [!example] 创建关系和两端节点
> ```cypher
> create(a:Book{name:"程序设计基础"})-[r:base]->(b:Book{name:"数据结构"});
> ```

为已创建的节点添加关系，需要先[[#元素查找|查询]]出相关节点

> [!example] 查询两组节点，在他们之间一一创建关系
> ```cypher
> match(a:Person),(b:Movie)
> where a.name='Robert Zemeckis' and b.title='Forrest Gump'
> create (a)-[r:DIRECTED]->(b)
> return r;
> ```

### 查询

```cypher
METCH(节点名:标签)
WHERE 条件列表
RETURN 节点或属性;
```

使用 `match` 匹配关系
- 泛指一切关系：`()--()`
- 泛指单项关系：`()-->()` / `()<--()`
- 泛指可有约束的一切关系：`()-[]-()`
- 泛指可有约束的一切单向关系：`()-[]->()` / `()<-[]-()`

---

> [!example] 查询所有节点
> 
> ```cypher
> match(n) return n;
> ```
> ![[../../../../_resources/images/Pasted image 20241022015106.png]]

> [!example] 按标签或属性查询
> ```cypher
> match(n:Person{name:"WuJing"})
> return n;
> ```

> [!example] 按属性条件查询节点，支持 `and` 和 `or`
> ```cypher
> match(n)
> where n.born < 1955
> return n;
> ```

> [!example] 查询节点，同时返回节点，节点 id，节点属性，节点标签等信息
> 
> 获取节点标签需要 APOC 库支持
> 
> ```cypher
> match(n:Person)
> return n, elementId(n), n.name, apoc.node.labels(n);
> ```
> ![[../../../../_resources/images/Pasted image 20241022013620.png]]

> [!example] 模糊查询：查询 `name` 属性以 `wang` 开头且不区分大小写的节点
> ```cypher
> match(n)
> where n.name=~'(?i)wang.*'
> return n;
> ```
> ![[../../../../_resources/images/Pasted image 20241022014926.png]]

> [!example] 使用 `ORDER BY` 排序，`SKIP` 跳过一定数量的结果，`LIMIT` 限定返回结果数量
> 默认升序排序，后可接 `DESC` 关键字表示降序排序
> ```cypher
> match(n)
> return n
> order by n.age
> skip 3
> limit 5;
> ```

> [!example] 使用 `UNION` 合并多个查询结果
> ```cypher
> match(n:Person)
> return n.name, n.age limit 2
> union
> match(n:Costomer)
> return n.name, n.age limit 2;
> ```
>  ![[../../../../_resources/images/Pasted image 20241022020342.png]]

> [!example] 使用 `keys(p)` 函数查询节点或边的所有属性
> ```cypher
> match(a)
> where a.age=10
> return a, keys(a);
> ```
>  ![[../../../../_resources/images/Pasted image 20241022020519.png]]

> [!example] 查询所有关系，并返回关系及其类型
> ```cypher
> match()-[r]-() return r,type(r);
> ```
### 修改

使用 `MATCH` 查询出节点后，可接多条 `SET` 添加或修改节点或关系
- `n:标签`：添加标签
- `n.property=value`：添加或修改属性
- `a=b`：复制两个标签除 ID 外所有属性

---

> [!example] 为所有 age>=18 的 `Person` 节点添加 `Adult` 标签
> ```cypher
> match(n:Person)
> where n.age>=18
> set n:Adult
> return n;
> ```
> ![[../../../../_resources/images/Pasted image 20241022021645.png]]

> [!example] 将 `tiger` 节点的属性复制给 `andy`
> ```cypher
> create(andy:Person{name:'andy',age:23});
> create(tiger:Person{name:'tiger',age:20});
> match(a{name:'andy'}),(p{name:'tiger'})
> set a=p
> return a,p;
> ```
### 删除

使用 `MATCH` 查询出后，可接多条 `DELETE` 删除节点、关系或其标签或属性

使用 `OPTION MATCH` 子句可以表示不确定节点间是否有关系

---

> [!example] 删除所有 `:Employee` 标签节点
> ```cypher
> match(e:Employee) delete e;
> ```

> [!example] 删除节点属性和标签
> ```cypher
> match(n{name:'andy'})
> remove n.age
> remove n:Person:Student
> return n;
> ```

> [!example] 删除学生 `Tom Hanks` 指向课程 `NoSQL`，关系为 `Study` 的关系
> ```cypher
> match(:Student{name:'Tom Hanks'})-[s:Study]-(:Course{name:'NoSQL'})
> delete s;
> ```

> [!example] 删除所有节点，若节点间有关系，同时删除对应关系
> ```cypher
> match(n) option match(n)-[r]-() delete n,r;
> ```
### MERGE 子句

类似于 `getOrCreate` 的形式
- 当模式存在时，匹配模式
- 当模式不存在时，创建模式

---

> [!example] 匹配搜索模式：查找一个带有 `name="Michael Douglas"` 属性的 `:Person` 节点，若不存在时创建一个节点
> ```cypher
> merge (michael:Person{name:"Michael Douglas"})
> return michael;
> ```
> ![[../../../../_resources/images/Pasted image 20241022023505.png]]

> [!example] 匹配模式，当创建时设置 `created` 属性，当访问时修改 `lastAccessed` 属性
> - `ON CREATE`：仅当匹配失败，创建节点时执行
> - `ON MATCH`：仅当匹配成功，获取节点时执行
> 
> ```cypher
> merge(keanu:Person{name:"Keanu Reeves"})
> on create set keanu.created=timestamp()
> on match set keanu.lastAccessed=date()
> return keanu.name, keanu.created, keanu.lastAccessed;
> ```
> - 第一次执行：
> ![[../../../../_resources/images/Pasted image 20241022024340.png]]
> - 第二次执行：
> ![[../../../../_resources/images/Pasted image 20241022024358.png]]
## 聚合函数
### 数值运算

```cypher
create (a:Student{id:1, score:97, classNo:3});
create (a:Student{id:2, score:65, classNo:7});
create (a:Student{id:3, score:92, classNo:7});
create (a:Student{id:4, score:78, classNo:8});
create (a:Student{id:5, score:83, classNo:8});
match(n:Student{classNo:7}) return count(n);
```

>[!example] 计数 `COUNT`：计算 7 班学生数
> ```cypher
> match(n:Student{classNo:7}) return count(n);
> ```
> ![[../../../../_resources/images/Pasted image 20241023023019.png]]

> [!example] 求和 `SUM`，平均值 `AVG`：计算学生成绩总和和平均值
> ```cypher
> match(n:Student) return sum(n.score), avg(n.score);
> ```
> ![[../../../../_resources/images/Pasted image 20241023024241.png]]

> [!example] 最大值 `MAX`，最小值 `MIN`：计算学生成绩最大值、最小值
> ```cypher
> match(n:Student) return max(n.score),  min(n.score);
> ```
> ![[../../../../_resources/images/Pasted image 20241023024332.png]]
### 去重

结果前加 `DISTINCT` 可以对结果去重

>[!example] 查询某一类产品供应商厂家名
> ```cypher
> match(c:Category{name:"Produce"})<--(:Product)<--(s:Supplier)
> return distinct s.companyName;
> ```
### 取值集合

使用 `COLLECT` 可以将结果组成集合

> [!example] 查询供应商的所有产品类型
> ```cypher
> match (s:Supplier)-->(:Product)-->(c:Category)
> return s.companyName as Company, collect(distinct c.categoryName) as Categories
> ```
## 路径操作
## 索引操作
## 约束
## 存储过程
# Neo4j 集群
# 管理与监控