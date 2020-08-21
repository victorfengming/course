# Elasticsearch安装

声明:JDK1.8, 最低要求 , Elasticsearch客户端,界面工具!

Java开发,elasticsearch的版本和我们之后对应的Java的核心jar包! 版本对应! JDK环境是正常的

**这里一定要保证**



![1596614089583](05-elasticsearch-%E5%AE%89%E8%A3%85%E5%92%8Chead%E6%8F%92%E4%BB%B6.assets/1596614089583.png)

下载

![1596614104028](05-elasticsearch-%E5%AE%89%E8%A3%85%E5%92%8Chead%E6%8F%92%E4%BB%B6.assets/1596614104028.png)

![1596614119890](05-elasticsearch-%E5%AE%89%E8%A3%85%E5%92%8Chead%E6%8F%92%E4%BB%B6.assets/1596614119890.png)

一定要在服务器上面搭建

下载地址:https://www.elastic.co/cn/downloads/elasticsearch

官网下载巨慢,翻墙,网盘中下载即可

华为云: https://mirrors.huaweicloud.com/elasticsearch/7.6.2/ 

我们学习的话Window和Linux都可以学习
==我们这里现在window下学习==

ELK三剑客,解压即用!(web项目! 前端环境! npm 下载依赖)

Node.js python2

```
window下安装!
```

1. 解压

    

   解压就可以使用

熟悉目录

![1596615600130](05-elasticsearch-%E5%AE%89%E8%A3%85%E5%92%8Chead%E6%8F%92%E4%BB%B6.assets/1596615600130.png)

```
bin 启动文件
config 配置文件
	log4j2 日志配置文件
	jvm.options java 虚拟机相关的配置
	elasticsearch.yml elasticsearch的配置文件! 默认 9200 端口! 跨域!
lib 相关jar包
modules 功能模块
plugins 插件! ik
logs 日志
```

启动

![1596615647818](05-elasticsearch-%E5%AE%89%E8%A3%85%E5%92%8Chead%E6%8F%92%E4%BB%B6.assets/1596615647818.png)

![1596615856687](05-elasticsearch-%E5%AE%89%E8%A3%85%E5%92%8Chead%E6%8F%92%E4%BB%B6.assets/1596615856687.png)

访问测试

![1596615846096](05-elasticsearch-%E5%AE%89%E8%A3%85%E5%92%8Chead%E6%8F%92%E4%BB%B6.assets/1596615846096.png)

老师,能不能给我个可视化界面啊

```
安装可视化界面 es head的插件
```



1. 下载地址:  https://github.com/mobz/elasticsearch-head 

2. 启动

   ```shell
   npm install
   npm run start
   ```

3. 连接测试发现,存在跨域问题:配置es

   ```shell
   http.cors.enabled: true
   http.cors.allow-origin: "*"
   ```

4. 重启es服务器,然后再次连接

   ![1596616716125](05-elasticsearch-%E5%AE%89%E8%A3%85%E5%92%8Chead%E6%8F%92%E4%BB%B6.assets/1596616716125.png)

   你们初学,你就把es的当做一个数据库!(可以建立索引(库),文档(库中的数据))

   你就把索引当做数据库(表,文档,类型)

   ```
   这个head我们就把他当做数据展示工具,我们后面所有的查询,kibana
   ```

   

   ![1597974207666](05-elasticsearch-%E5%AE%89%E8%A3%85%E5%92%8Chead%E6%8F%92%E4%BB%B6.assets/1597974207666.png)

   

