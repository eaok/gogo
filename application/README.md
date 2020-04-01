# 解决的问题

|文件名|功能|
|------|---|
|01_copy.go|复制文件|
|02_dirs.go|打印所有文件夹和文件|
|03_githubName.go|查找github上还未注册的4位用户名|
|04_githubNameProducerConsumer.go|用生产者消费者模式实现查找github上还未注册的4位用户名|

------
05_goroutinePoll.go

如果无休止的开辟Goroutine依然会出现高频率的调度Goroutine，那么依然会浪费很多上下文切换的资源，导致做无用功。所以设计一个Goroutine池限制Goroutine的开辟个数在大型并发场景还是必要的。
![](http://blog.maser.top/gogo/goroutinePool.jpg)
