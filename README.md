# databuffer (using sync.Pool)
 客户端发过来的消息长度 基本上是稳定的某几个数值之间，
 使用sync.Pool 对make([]byte,n) 进行内存重用
 尽量减少gc时间
 如果没有 databuffer
 则使用方式是
  ```
     data:=make([]byte,dataSize)
  ```
  有了databuffer 以后的使用是
  ```
     data:=databuffer.GetBuffer(dataSize)
     // do something
     databuffer.PutBuffer(data) // 把data 这块内存放到缓存中
     ```
     
# link
  http://blog.csdn.net/tiaotiaoyly/article/details/38388081
  http://my.oschina.net/lubia/blog/175154