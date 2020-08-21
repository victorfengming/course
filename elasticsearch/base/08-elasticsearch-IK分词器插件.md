# IK分词器插件

狂神说

```
什么是分词器
```

![1596637415868](08-elasticsearch-IK%E5%88%86%E8%AF%8D%E5%99%A8%E6%8F%92%E4%BB%B6.assets/1596637415868.png)

如果使用中文,建议使用ik分词器

IK体用了两个分词算法:ik_smart和ik_max_word,其中ik_smart为最少(ˉ▽￣～) 切~~分,ik_max_word为最细粒度划分! 一会我们测试

```
安装
```

1. https://github.com/medcl/elasticsearch-analysis-ik

2. 下载完毕之后,放入到我们的elasticsearch插件即可

   ![1596637767841](08-elasticsearch-IK%E5%88%86%E8%AF%8D%E5%99%A8%E6%8F%92%E4%BB%B6.assets/1596637767841.png)

   3. 重启观察ES
   
      ![1596674648793](08-elasticsearch-IK%E5%88%86%E8%AF%8D%E5%99%A8%E6%8F%92%E4%BB%B6.assets/1596674648793.png)
   
   4. elasticsearch-plugin 可以通过这个命令来查看加载的插件
   
      ![1596674703775](08-elasticsearch-IK%E5%88%86%E8%AF%8D%E5%99%A8%E6%8F%92%E4%BB%B6.assets/1596674703775.png)
   
   5. 使用kibana测试!
   
      ```
      查看不同的分词器效果
      ```
   
      ![1596674979545](08-elasticsearch-IK%E5%88%86%E8%AF%8D%E5%99%A8%E6%8F%92%E4%BB%B6.assets/1596674979545.png)
   
      
   
      其中ik_smart为最细粒度划分! 穷尽词库的可能! 字典 ! 
   
      ![1596675058784](08-elasticsearch-IK%E5%88%86%E8%AF%8D%E5%99%A8%E6%8F%92%E4%BB%B6.assets/1596675058784.png)
   
      ```
      我们输入超级喜欢狂神说Java
      ```
   
      ![1596675178617](08-elasticsearch-IK%E5%88%86%E8%AF%8D%E5%99%A8%E6%8F%92%E4%BB%B6.assets/1596675178617.png)
   
      

发现问题: 狂神说被拆开了!

这种自己需要的词,需要自己加到我们的分词器字典中!

```
ik分词器增加我们自己的配置
```

![1596675349989](08-elasticsearch-IK%E5%88%86%E8%AF%8D%E5%99%A8%E6%8F%92%E4%BB%B6.assets/1596675349989.png)

重启ES,看细节

![1596675426305](08-elasticsearch-IK%E5%88%86%E8%AF%8D%E5%99%A8%E6%8F%92%E4%BB%B6.assets/1596675426305.png)

再次测试一下狂神说,

![1596675486318](08-elasticsearch-IK%E5%88%86%E8%AF%8D%E5%99%A8%E6%8F%92%E4%BB%B6.assets/1596675486318.png)



以后的话我们需要自己配置自己的词,只需要在自定义的dic文件中进行配置即可! 

![1597974175451](08-elasticsearch-IK%E5%88%86%E8%AF%8D%E5%99%A8%E6%8F%92%E4%BB%B6.assets/1597974175451.png)