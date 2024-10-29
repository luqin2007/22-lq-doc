>[!error] Neo.ClientError.Statement.ExternalResourceFailed
>Cannot load the external resource at: file:/... .csv

`LOAD CSV` 从文件加载在 3.0+ 需要修改安全配置，之后重启数据库

```properties title:$NEO4J_HOME/conf/neo4j.conf
dbms.security.allow_csv_import_from_file_urls=true
```
