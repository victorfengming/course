mybatis

环境回顾

# mybatis
环境
- jdk1.8
- mysql5.7
- maven3.6.1
- idea

回顾
- jdbc
- mysql
- javase
- maven
- Junit

框架:配置文件的.最好的方式:看官方文档

# 简介
## 1.1. 什么是mybatis

mybatis本是apache的以一个开源项目,ibatis

mybatis是一款优秀的持久层框架,他支持定制化的sql,存储过程以及高级映射.mybatis避免了几乎所有的jdbc代码和手动设置参数以及获取结果集.mybatis可以使用简单的xml或注解来配置和映射原生类型,接口和java的pojo为数据库中的记录.

如何获得mybatis?

1. maven仓库

2. github

## 1.2 持久化
数据持久化,
持久化就是讲程序的数据在持久状态和瞬时状态转化的过程

内存:断电即失

数据库(jdbc),io文件持久化.

生活: 冷藏.肉丢到冰箱里面,罐头

为什么需要持久化?

有一些对象不能让他丢掉.

内存太贵了

## 1.3 持久层
Dao层,Service层,Controller层...
- 完成持久化工作的代码块儿
- 层界限十分明显.

## 1.4 为什么需要mybatis?
- 帮助程序员将数据存入到数据库中
- 方便
- 传统的jdbc代码太复杂了.简化.框架.
- 不用mybatis也可以.更容易上手.技术没有高低之分

## mybatis特点

- 简单易学
- 灵活:mybatis不会对应用程序或数据库加任何影响
- 解除sql与程序代码的耦合:通过提供DAO层,将业务逻辑

最终要的一点:使用的人多

spring spring boot spirng mvc

# 2.第一个mybatis程序
思路:搭建环境-->导入mybatis-->编写配置
