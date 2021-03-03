## 4.4 Java批量操作文档

### 4.4.1 批量添加

```java
@Test
public void bulkCreateDoc(){
    // 1. 准备多个json数据
    Person p1 = new Person(1,"张三",23,new Date());
    Person p1 = new Person(2,"李四",26,new Date());
    Person p1 = new Person(3,"王五",27,new Date());
    
    String json1 = mapper.writeValueAsString(p1);
    String json2 = mapper.writeValueAsString(p2);
    String json3 = mapper.writeValueAsString(p3);
    // 2.创建Request,将准备好的数据封装进去
    BulkRequest request = new BulkRequest();
    request.add(
        new IndexRequest(
            index,type,p1.getId().toString()
        ).source(json1,XContentType.JSON)
    );
    request.add(
        new IndexRequest(
            index,type,p2.getId().toString()
        ).source(json2,XContentType.JSON)
    );
    request.add(
        new IndexRequest(
            index,type,p3.getId().toString()
        ).source(json3,XContentType.JSON)
    );
    // 3. 用client执行
    BulkResponse resp = client.bulk(request, RequestOption.DEFAULT);
    // 4. 输出结果
    System.out.println(resp.toString());
}
```

刷新,查看结果

![1598440519480](17_Java%E6%93%8D%E4%BD%9C%E6%96%87%E6%A1%A3_%E6%89%B9%E9%87%8F%E6%93%8D%E4%BD%9C.assets/1598440519480.png)

### 4.4.2 批量删除

```java
@Test
public void bulkDeleteDoc throws IOException(){
    // 1.封装Request对象
    BulkRequest request = new BulkRequest();
    request.add(
        new DeleteRequest(index,type,"1")
    );
    request.add(
        new DeleteRequest(index,type,"2")
    );
    request.add(
        new DeleteRequest(index,type,"3")
    );
    // 2.client执行
    BulkRequest resp = client.bulk(request, RequestOptions.DEFAULT)
    // 3.输出
    System.out.println(resp.toString());
}
```

查看结果,果然没了

![1598440821475](17_Java%E6%93%8D%E4%BD%9C%E6%96%87%E6%A1%A3_%E6%89%B9%E9%87%8F%E6%93%8D%E4%BD%9C.assets/1598440821475.png)