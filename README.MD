# 介绍
网络调试工具，支持以客户端或服务端模式启动，监听或连接制定的地址。接收到数据后输出到命令行,接收命令行输入并向目标写数据。
# 参数
#### -d
传输数据类型，默认为 text，可以修改为 hex。
#### -h
host，主机名，默认为 0.0.0.0。
#### -p
port，端口号，默认为 8080。
#### -pr
protocol，协议。tcp 或 udp， 默认为 tcp。