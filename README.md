# postgres 相关

查找外键
```
SELECT
    tc.table_schema,
    tc.constraint_name,
    tc.table_name,
    kcu.column_name,
    ccu.table_schema AS foreign_table_schema,
    ccu.table_name AS foreign_table_name,
    ccu.column_name AS foreign_column_name
FROM
    information_schema.table_constraints AS tc
    JOIN information_schema.key_column_usage AS kcu
      ON tc.constraint_name = kcu.constraint_name
      AND tc.table_schema = kcu.table_schema
    JOIN information_schema.constraint_column_usage AS ccu
      ON ccu.constraint_name = tc.constraint_name
      AND ccu.table_schema = tc.table_schema
WHERE tc.constraint_type = 'FOREIGN KEY' AND tc.table_name='my_tablename';
```

删除外键
```
ALTER TABLE my_tablename DROP CONSTRAINT fk_to_drop;
```

# 使用GoAdmin框架

- [github](https://github.com/GoAdminGroup/go-admin)
- [论坛](http://discuss.go-admin.com)
- [文档](https://book.go-admin.cn)

```
.
├── Dockerfile          Dockerfile
├── Makefile            Makefile
├── build               二进制构建目标文件夹
├── config.json         配置文件
├── go.mod              go.mod
├── go.sum              go.sum
├── html                前端html文件
├── logs                日志
├── main.go             main.go
├── main_test.go        CI测试
├── pages               页面控制器
├── tables              数据模型
└── uploads             上传文件夹
```


### 使用命令行工具

```
adm generate -l cn -c adm.ini
```

