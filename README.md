# MySQL HTTP Agent

![license](https://img.shields.io/github/license/alibaba/dubbo.svg)

参照 [`chop-dbhi/sql-agent`](https://github.com/chop-dbhi/sql-agent) 实现一个 `MySQL` 的 `HTTP` 代理，用于实现通过 `HTTP API` 查询 `MySQL`

示例请求
```
$ curl -s --request POST \
--header 'Content-Type: application/json' \
--url 'http://localhost:5000/' \
--data '
{
    "connection": {
        "host": "192.168.10.108",
        "port": 3306,
        "username": "root",
        "password": "123456",
        "database": "mysql"
    },
    "sql": "SELECT 120 AS cnd FROM DUAL WHERE 1 = :id",
    "params": {
        "id": 1
    }
}'
```