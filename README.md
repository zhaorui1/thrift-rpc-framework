Thrift Rpc Framework
===================


此项目应用Thrift搭建的rpc服务框架，server端使用go语言开发，client提供go、php、java三种客户端调用sdk。



----------


安装
-------------

1. 下载项目至$GOPATH/src路径下.
2. git clone https://github.com/apache/thrift.git

----------
运行
-------------

1. 启动server。项目路径下运行：go run main.go.
2. go启动测试。路径：sdk/go/driver/。 运行：go test
3. php启动测试。路径：sdk/php/driver。运行：php client.go
