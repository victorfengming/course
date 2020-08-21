# Rest风格说明

一种软件架构风格,而不是标准,只是提供了一组设计原则和约束条件.它主要用于客户端和服务器交互类的软件.基于这个风格设计的软件可以更简洁,更有层次,更易于实现缓存等机制.

![1596675703304](09-elasticsearch-Rest%E9%A3%8E%E6%A0%BC.assets/1596675703304.png)

> 基础测试

## 关于索引的基本操作

1. 创建一个索引!

   ```
   PUT /索引名/~类型名~/文档id
   {请求体}
   ```

   ![1596675966702](09-elasticsearch-Rest%E9%A3%8E%E6%A0%BC.assets/1596675966702.png)

   完成了自动增加了索引! 数据也成功的添加了,这就是我说大家在初期可以把它当做数据库学习的原因

   ![1596676030312](09-elasticsearch-Rest%E9%A3%8E%E6%A0%BC.assets/1596676030312.png)

以后查询可以直接查,不走数据库

3. 那么name这个字段用不用指定类型呢,毕竟我们关系型数据库 是需要指定类型的啊

   ![1596676128844](09-elasticsearch-Rest%E9%A3%8E%E6%A0%BC.assets/1596676128844.png)

   4. 指定字段的类型

      ![1596676273693](09-elasticsearch-Rest%E9%A3%8E%E6%A0%BC.assets/1596676273693.png)

      获得这个规则! 可通过GET请求获取具体的信息

      ![1596676322260](09-elasticsearch-Rest%E9%A3%8E%E6%A0%BC.assets/1596676322260.png)

   5. 查看默认的信息

      ![1596676470954](09-elasticsearch-Rest%E9%A3%8E%E6%A0%BC.assets/1596676470954.png)

      就是你不写的话,有一个默认的不可分割的类型

      ![1596676513685](09-elasticsearch-Rest%E9%A3%8E%E6%A0%BC.assets/1596676513685.png)

      如果自己的文档字段没有指定,那么ES就会给我们配置字段类型

      ```
      虚心学习,这个世界上大佬很多!
      ```

      

      扩展:通过命令elasticsearch索引情况! 通过get_cat/可以获得es当前的很多信息

      ![1596676751370](09-elasticsearch-Rest%E9%A3%8E%E6%A0%BC.assets/1596676751370.png)

      ```
      修改索引 提交哈哈四使用PUT即可! 然后覆盖! 最新办法
      ```

      曾经! 是这样的

      ![1596677004607](09-elasticsearch-Rest%E9%A3%8E%E6%A0%BC.assets/1596677004607.png)

      现在的方法

      ![1596677026606](09-elasticsearch-Rest%E9%A3%8E%E6%A0%BC.assets/1596677026606.png)

      

```
删除索引
```

通过DELETE 命令实现删除,根据你的请求是删除索引还是删除文档记录

使用ERSTFUL风格是我们推荐大家使用的

![1597974169988](09-elasticsearch-Rest%E9%A3%8E%E6%A0%BC.assets/1597974169988.png)