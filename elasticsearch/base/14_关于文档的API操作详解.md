> 文档操作Api

```java
void testAddDocument(){
    // 创建对象
    User user = new User("狂神说",3);
    //创建请求
    IndexRequest kuang_index = new IndexRequest("kuang_index");
    // 规则 put /kuang_index/_doc/1
    request.id("1")
    request.timeout(TimeValue.timeValueSeconds(1));
    request.timeout("1s");
    // 将我们的数据放入请求 json
    IndexRequest source = request.source(JSON.toJSONString(user),XContentType.JSON);
    // 客户端发送请求
    IndexResponse indexResponse = client.index(request,RequestOptions.DEFAULT);
    System.out.println(indexResponse.toString());
    System.out.println(indexResponse.status());
}
```



![1597910328977](14_%E5%85%B3%E4%BA%8E%E6%96%87%E6%A1%A3%E7%9A%84API%E6%93%8D%E4%BD%9C%E8%AF%A6%E8%A7%A3.assets/1597910328977.png)

结果

这里的返回的全部内容和我们的命令是一样的

![1597910356381](14_%E5%85%B3%E4%BA%8E%E6%96%87%E6%A1%A3%E7%9A%84API%E6%93%8D%E4%BD%9C%E8%AF%A6%E8%A7%A3.assets/1597910356381.png)

更新文档信息

![1597910576621](14_%E5%85%B3%E4%BA%8E%E6%96%87%E6%A1%A3%E7%9A%84API%E6%93%8D%E4%BD%9C%E8%AF%A6%E8%A7%A3.assets/1597910576621.png)

删除文档记录

![1597910686502](14_%E5%85%B3%E4%BA%8E%E6%96%87%E6%A1%A3%E7%9A%84API%E6%93%8D%E4%BD%9C%E8%AF%A6%E8%A7%A3.assets/1597910686502.png)

特殊的,真的项目一般都会批量插入数据

![1597911025645](14_%E5%85%B3%E4%BA%8E%E6%96%87%E6%A1%A3%E7%9A%84API%E6%93%8D%E4%BD%9C%E8%AF%A6%E8%A7%A3.assets/1597911025645.png)

![1597911057770](14_%E5%85%B3%E4%BA%8E%E6%96%87%E6%A1%A3%E7%9A%84API%E6%93%8D%E4%BD%9C%E8%AF%A6%E8%A7%A3.assets/1597911057770.png)

![1597911133927](14_%E5%85%B3%E4%BA%8E%E6%96%87%E6%A1%A3%E7%9A%84API%E6%93%8D%E4%BD%9C%E8%AF%A6%E8%A7%A3.assets/1597911133927.png)

是否失败,最后返回false 代表成功

如果你不设置id,默认id随机

> // 查询

![1597911490131](14_%E5%85%B3%E4%BA%8E%E6%96%87%E6%A1%A3%E7%9A%84API%E6%93%8D%E4%BD%9C%E8%AF%A6%E8%A7%A3.assets/1597911490131.png)

> 注意点

![1597911595769](14_%E5%85%B3%E4%BA%8E%E6%96%87%E6%A1%A3%E7%9A%84API%E6%93%8D%E4%BD%9C%E8%AF%A6%E8%A7%A3.assets/1597911595769.png)

![1597974144176](14_%E5%85%B3%E4%BA%8E%E6%96%87%E6%A1%A3%E7%9A%84API%E6%93%8D%E4%BD%9C%E8%AF%A6%E8%A7%A3.assets/1597974144176.png)