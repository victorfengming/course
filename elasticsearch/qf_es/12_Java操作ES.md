# 四.Java操作elasticsearch

## 4.1 Java连接ES

> 创建一个maven工程
>
> 导入依赖

```shell
elasticsearch官方jar包
elasticsearch的高级API
junit
lombok
```

`pom.xml`

```xml
<dependencies>
<!--    1.elasticsearch-->
<dependency>
    <groupId>org.elasticsearch</groupId>
    <artifactId>elasticsearch</artifactId>
    <version>6.5.4</version>
</dependency>
<!--    2. elasticsearch的高级API-->
<dependency>
    <groupId>org.elasticsearch.client</groupId>
    <artifactId>elasticsearch-rest-high-level-client</artifactId>
    <version>6.5.4</version>
</dependency>
<!--    3.junit-->
<dependency>
    <groupId>junit</groupId>
    <artifactId>junit</artifactId>
    <version>4.12</version>
</dependency>
<!--    4. lombok-->
<dependency>
    <groupId>org.projectlombok</groupId>
    <artifactId>lombok</artifactId>
    <version>1.16.22</version>
</dependency>
</dependencies>
```

创建类`ESClient.java`

![1598430705488](12_Java%E6%93%8D%E4%BD%9CES.assets/1598430705488.png)

测试连接

![1598430738726](12_Java%E6%93%8D%E4%BD%9CES.assets/1598430738726.png)

