1. 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

   - fix length

     发送方和接收方规定固定大小的缓冲区，也就是发送和接收都使用固定大小的 byte[] 数组长度，当字符长度不够时使用空字符弥补, 但是没办法解决拆包问题

   - Delimiter based

     以特殊的字符结尾，比如以“\n”结尾，这样我们就知道结束字符，从而避免了半包和粘包问题.

   - Length field base frame decoder

     在 TCP 协议的基础上封装一层数据请求协议，既将数据包封装成数据头（存储数据正文大小）+ 数据正文的形式，这样在服务端就可以知道每个数据包的具体长度了，知道了发送数据的具体边界之后，就可以解决半包和粘包的问题了