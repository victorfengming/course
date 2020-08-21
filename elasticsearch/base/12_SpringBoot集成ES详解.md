集成SpringBoot

> 找文档!

 https://proxies.app.aidoru.net/-----https://www.elastic.co/guide/index.html 

![1597906806327](12_SpringBoot%E9%9B%86%E6%88%90ES%E8%AF%A6%E8%A7%A3.assets/1597906806327.png)

![1597906105998](12_SpringBoot%E9%9B%86%E6%88%90ES%E8%AF%A6%E8%A7%A3.assets/1597906105998.png)

![](12_SpringBoot%E9%9B%86%E6%88%90ES%E8%AF%A6%E8%A7%A3.assets/1597906924635.png)

1. 找到原生的依赖

   ![1597907074366](12_SpringBoot%E9%9B%86%E6%88%90ES%E8%AF%A6%E8%A7%A3.assets/1597907074366.png)

2. 找对象

   ![1597907115554](12_SpringBoot%E9%9B%86%E6%88%90ES%E8%AF%A6%E8%A7%A3.assets/1597907115554.png)

3. 分析这个类中的方法即可

   > 配置基本的项目

   ![1597907258818](12_SpringBoot%E9%9B%86%E6%88%90ES%E8%AF%A6%E8%A7%A3.assets/1597907258818.png)

   > 问题:一定要保证我们导入的依赖和我们的ES版本一致

   ![1597907493787](12_SpringBoot%E9%9B%86%E6%88%90ES%E8%AF%A6%E8%A7%A3.assets/1597907493787.png)

按照官网的操作我们要构建一个对象

![1597907737591](12_SpringBoot%E9%9B%86%E6%88%90ES%E8%AF%A6%E8%A7%A3.assets/1597907737591.png)

分析源码

狂神的Spring步骤:

1. 找对象

2. 放到spring中待用

3. 如果是springboot,那就先分析源码

   xxxAutoConfiguration,xxxProperties

   

![1597908504743](12_SpringBoot%E9%9B%86%E6%88%90ES%E8%AF%A6%E8%A7%A3.assets/1597908504743.png)

源码中提供的对象

虽然这里导入了3个类,静态内部类,核心类就一个

> 具体的api测试



![1597974149380](12_SpringBoot%E9%9B%86%E6%88%90ES%E8%AF%A6%E8%A7%A3.assets/1597974149380.png)