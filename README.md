# TrainingProgram

## 描述
字节后端训练营——校园排课系统

## 目录结构
- api：就是controller层
- cache：redis相关
- conf：配置包
- middleware：存放一些基础通用功能的包，如认证、session等
- model：存放一些数据实体的定义
- resource：存放项目的资源文件，如建库建表的sql语句
- server：存放服务器配置、路由等
- service：存放业务逻辑
- util：存放工具函数
- .env文件：存放项目的配置，项目启动时Godotenv会自动读取该文件，并把值保存为项目的环境变量
- main.go文件：项目入口
