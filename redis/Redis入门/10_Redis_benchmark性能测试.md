### #性能测试

redis-benchmark 是一个压力测试工具

官方自带的性能测试工具

redis-benchmark命令采纳数

### 菜鸟教程

![1597193793816](10_Redis_benchmark%E6%80%A7%E8%83%BD%E6%B5%8B%E8%AF%95.assets/1597193793816.png)

![1597193809673](10_Redis_benchmark%E6%80%A7%E8%83%BD%E6%B5%8B%E8%AF%95.assets/1597193809673.png)

我们来简单测试一下

> 测试:100个并发连接 	100000个请求
>
> redis-benchmark -h localhost -p 6379 -c 100 -n 100000

![1597194193885](10_Redis_benchmark%E6%80%A7%E8%83%BD%E6%B5%8B%E8%AF%95.assets/1597194193885.png)

分析结果

![1597194143929](10_Redis_benchmark%E6%80%A7%E8%83%BD%E6%B5%8B%E8%AF%95.assets/1597194143929.png)
