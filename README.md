## resource-access-control

## Requirements

* Go version >= 1.12.0



```
运行前环境准备：
1 启动fabric网络
2 在configs/dao.toml配置mysql数据源
3 建立数据库 resource_access_control
3 在数据库resource_access_control执行internal/dao/db.sql创建数据表
```

### docker部署

####镜像依赖
harbor.tusdao.com/library/alpine:3.10.2
####镜像构建
######参见docker.sh

