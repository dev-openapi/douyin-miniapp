# dyma

微信小程序接口，使用protobuf生成的

生成工具是[protoc-gen-go_api](https://github.com/dev-openapi/protoc-gen-go_api)


## 如何使用

在自己的项目中添加此包 github.com/dev-openapi/douyin-miniapp

初始化包 `dyLoginService := miniapp.NewLoginService()`

然后根据自己的需要，用`dyLoginService`去调用想用的接口

## 如何添加新接口

首先要安装protoc,protoc-gen-go和protoc-gen-go_api。前面两个如何安装，请参考网上教程搜索

```shell
go install github.com/dev-openapi/protoc-gen-go_api@latest
```

下载[protocol](https://github.com/dev-openapi/protocol)仓库

下载本项目，与protocol是同一级目录

然后根据文档在protocol项目中写好proto文件，再执行make gen-dyma会自动生成到本项目目录，最后提交两个项目即可