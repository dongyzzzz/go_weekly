1.总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

- fix length：发送方和接受方约定每次发送数据的长度是一定的,每次发送固定缓冲区大小
- delimiter based:在包尾增加回车换行符进行分割，例如FTP协议
- length field based frame decoder：将消息分为消息头和消息体，消息头中包含表示消息的总长度（或者消息体长度）的字段

2.实现一个从 socket connection 中解码出 goim 协议的解码器


TODO

