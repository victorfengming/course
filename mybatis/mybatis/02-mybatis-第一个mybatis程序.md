
# 2.第一个mybatis程序
思路:搭建环境-->导入mybatis-->编写代码->>测试

## 2.1 搭建环境

```shell script
==================================================================
Congratulations! Installed successfully!
==================================================================
Bt-Panel: http://39.106.139.40:8888/c1838dd2
username: i8mcadci
password: 4382cfc5

```

搭建数据库
```sql

CREATE DATABASE `mybatis`;

use `mybatis`;


create table `user`(
    `id` int(20) not null primary key,
    `name` varchar(30) default null,
    `pwd` varchar(30) default null
)engine=innodb default charset=utf8;

insert into `user`(`id`,`name`,`pwd`) VALUES
(1,'狂神','2135345'),
(2,'张三','4357756'),
(3,'李四','6724524')
```

新建项目

1. 新建一个普通的maven项目
2. 删除src目录
3. 导入maven依赖
```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.kuang</groupId>
    <artifactId>mybatis-study</artifactId>
    <version>1.0-SNAPSHOT</version>

    <dependencies>
<!--        mysql驱动-->
        <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
            <version>5.1.46</version>
        </dependency>
<!--        mybatis-->
        <dependency>
            <groupId>org.mybatis</groupId>
            <artifactId>mybatis</artifactId>
            <version>3.5.2</version>
        </dependency>
<!--        junit-->
        <dependency>
            <groupId>junit</groupId>
            <artifactId>junit</artifactId>
            <version>4.12</version>
        </dependency>
    </dependencies>

</project>
```

## 2.2 创建一个模块
- 编写mybatis的核心配置文件
- 编写mybatis工具类

![1597294870488](02-mybatis-%E7%AC%AC%E4%B8%80%E4%B8%AAmybatis%E7%A8%8B%E5%BA%8F.assets/1597294870488.png)

