# 简介
gin demo project

## 配置
.env

## start
> go run main.go


## 项目结构

```
├─.env
├─config  配置
│  ├─locales 翻译相关的配置文件
│  ├─config
│  ├─log
├─handler 路由处理　controller
├─middleware 中间件（权限，日志，跨域）
├─model　　存储数据库模型和数据库操作
│  ├─user 用户
│  ├─db   数据库
│  ├─casbin　权限
├─routers 路由
├─serializer model struct 序列化
├─service 业务处理
├─test 测试
```

