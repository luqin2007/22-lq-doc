> [!error] Neo.ClientError.Procedure.ProcedureRegistrationFailed
> apoc.node.labels is unavailable because it is sandboxed and has dependencies outside of the sandbox. Sandboxing is controlled by the dbms.security.procedures.unrestricted setting. Only unrestrict procedures you can trust with access to database internals.

`apoc.node.labels` 等指令需要配置安全配置，`neo4j restart` 重启数据库

```properties title:/conf/neo4j.conf
dbms.security.procedures.unrestricted=apoc.*
dbms.security.procedures.whitelist=apoc.*
```
