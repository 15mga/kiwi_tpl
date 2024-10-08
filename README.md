## 安装Protocol Buffers编译器

访问 [Protocol Buffers Releases](https://github.com/protocolbuffers/protobuf/releases) 页面，下载适合你操作系统的预编译二进制文件并解压。可以直接将`protoc`执行文件放到go/bin目录下。

## 安装Go Protocol Buffers插件

```bash
go intall google.golang.org/protobuf/cmd/protoc-gen-go
```

## 安装kiwi代码生成工具

```bash
# kiwi代码生成工具
go install github.com/15mga/kiwi_tool/cmd/proto_kiwi/protoc-gen-kiwi.go
# 将MongoDB的Schema结构中Id字段的bson tag修改为"_id"
go install github.com/15mga/kiwi_tool/cmd/protoc_mgo_bson/protoc-mgo-bson.go
```
