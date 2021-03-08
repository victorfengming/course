### 4.3.2 修改文档

```java
@Test
public void updateDoc() throws IOException{
    // 1. 创建一个Map, 指定需要修改的内容
    Map<String, Object> doc = new HashMap<>();
    doc.put("name","长大三");
    String docId = "1";
    // 2. 创建request对象,封装数据
    UpdateRequest request = new UpdateRequest(index，type,docId);
    request.doc(doc);
    // 3. 通过client对象执行
    UpdateResponse update = client.update(request, RequestOptions.DEFAULT);
    // 4. 输出返回结果
    System.out.println(update.getResult().toString())
}
```

### 4.3.3 删除文档

```java
@Test
public void deleteDoc(){
    // 1. 封装Request对象
    DeleteRequest request = new DeleteRequest(index,type,"1");
    // 2. client执行
    DeleteResponse resp = client.delete(request, ResquestOptions.DEFAULT);
    // 3. 输出结果
    System.out.println(resp.getResult().toString());
}
```

