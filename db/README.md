# db 数据库上下文

用于持久化的数据库，只需要选择一种，默认使用内存。

处理数据库连接和数据库关系上下文,不包含任何业务处理逻辑。

一般情况，一个微服务最多连接 1-2 个数据源（往往是 DB 和 Cache)

当前 case 里面持久化的的连接数据库是：

- sqlite
- mysql
- mongodb
- dgraph
- memory

运行代码的时候都要选择一种持久化类型

缓存：

- redis

## docker compose for dgraph
