## GRPC-GO

参考 https://juejin.cn/post/7068215203686514724

**1、下载protoc工具**

[protoc](https://github.com/protocolbuffers/protobuf/releases)

下载完成后解压后记得将路径添加到环境变量中


**2、下载go的依赖包**
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

**3、编写proto文件**

**4、生成go文件**
```shell
protoc -I . helloworld.proto --go_out=plugins=grpc:.
```
