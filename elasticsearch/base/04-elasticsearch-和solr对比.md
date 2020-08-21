# Elasticsearch和Solr差别

Elasticsearch是一个实时分布式搜索和分析引擎.它让你以前所未有的速度处理大数据成为可能.
维基百科使用它提供全文搜索并高亮关键字,以及输入实时搜索

Solr简介
Solr是Apache下的一个顶级开源项目,采用Java开发,它是基于Lucene的全文搜索服务器.solr提供了比Lucene更为丰富的查询语言,同时实现了可配置,可扩展,并对索引,搜索性能进行优化
他可以独立运行,运行在tomcat ,jety等这些Servlet容器中
solr对外提供类似于Web-server的API接口

![1596613190316](04-elasticsearch-%E5%92%8Csolr%E5%AF%B9%E6%AF%94.assets/1596613190316.png)

![1596613206055](04-elasticsearch-%E5%92%8Csolr%E5%AF%B9%E6%AF%94.assets/1596613206055.png)

随着数据量的增加,solr的搜索

![1596613233807](04-elasticsearch-%E5%92%8Csolr%E5%AF%B9%E6%AF%94.assets/1596613233807.png)

50倍的效率

![1596613255957](04-elasticsearch-%E5%92%8Csolr%E5%AF%B9%E6%AF%94.assets/1596613255957.png)

![1597974203192](04-elasticsearch-%E5%92%8Csolr%E5%AF%B9%E6%AF%94.assets/1597974203192.png)

## ElasticSearch vs Solr 总结

1. es基本是开箱即用(解压就可以用了!),非常简单.solr安装略微复杂一丢丢!
2. Solr利用Zookeeper进行分布式管理,而Elasticsearch自身带有分布式协调管理功能.
3. solr支持更多格式的数据,比如JSON,XML,CSV, 而elasticsearch仅仅支持json文件格式
4. Solr官网提供的功能很多,儿elasticsearch本身更注重核心功能,高级功能多有第三方插件提供,例如图形化界面需要kibana友好质层支撑
5. Solr查询块,但更新索引时慢(即插入删除慢),用于电商等查询多的应用;
   - ES建立索引块(即查询慢),即实时性查询快,用于facebook新浪等搜索.
   - Solr是传统搜索应用的有力解决方案,但Elasticsearch更适用于新兴的实时搜索应用.
6. Solr比较成熟,有一个更大,更成熟的用户,开发好贡献者社区,而Elasticsearch相对开发维护者较少,更新太快,学习使用成本较高.

