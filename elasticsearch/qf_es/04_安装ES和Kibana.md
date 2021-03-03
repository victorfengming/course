https://www.bilibili.com/video/BV1Qz411e7yx?p=5

# 二.Elasticsearch安装

## 2.1 `ES` 安装&`Kibana`

### 准备部分



可以在http://hub.docker.com/里面搜索

`elasticsearch`和`kibana`

咱们统一一下版本

用`6.5.4`的版本

![1598358080059](04_%E5%AE%89%E8%A3%85ES%E5%92%8CKibana.assets/1598358080059.png)

![1598358221749](04_%E5%AE%89%E8%A3%85ES%E5%92%8CKibana.assets/1598358221749.png)

### 一起来操作

在服务器的`/opt/`路径下面创建`docker_es`目录

创建文件`docker-compose.yml`,内容如下:

```yml
version: "3.1"
services: 
  elasticsearch: 
    image: daocloud.io/library/elasticsearch:6.5.4
    restart: always
    container_name: elasticsearch
    ports: 
      - 9200:9200
  kibana: 
    image: daocloud.io/library/kibana:6.5.4
    restart: always
    container_name: kibana
    ports: 
      - 5601:5601
    environment: 
      - elasticsearch_url=http://39.106.139.40:9200
      # 上面要填上你自己的虚拟机的ip和端口
    depends_on: 
      - elasticsearch
```

> 上面要注意yml文件的冒号后面的空格要有,不能有tab缩进,都是踩过的坑,汗水经验

![1598359315223](04_%E5%AE%89%E8%A3%85ES%E5%92%8CKibana.assets/1598359315223.png)

然后在目录中运行命令

```shell
docker-compose up -d
```

试一下

```shell
[root@iz8g9301trfnpxz docker_es]# docker-compose up -d
Creating network "docker_es_default" with the default driver
Pulling elasticsearch (daocloud.io/library/elasticsearch:6.5.4)...
6.5.4: Pulling from library/elasticsearch
a02a4930cb5d: Already exists
dd8a94cca3f9: Pull complete
bd73f551dee4: Pull complete
70de352c4efc: Pull complete
0b5ae4c7310f: Pull complete
489d9f8b18f1: Pull complete
8ba96caf5951: Pull complete
f1df04f27c5f: Pull complete
Digest: sha256:5ca85697b6273f63196b44c32311c5a2d1135af9cfd919e5922e49c5045d04b8
Status: Downloaded newer image for daocloud.io/library/elasticsearch:6.5.4
Pulling kibana (daocloud.io/library/kibana:6.5.4)...
6.5.4: Pulling from library/kibana
a02a4930cb5d: Already exists
39999a66d2b5: Pull complete
4b2aeb7c30c9: Pull complete
00616cdd3c09: Pull complete
e98987d4a5ba: Pull complete
5dfef65737c6: Pull complete
a8d538f84d1a: Pull complete
961861c10f7a: Pull complete
5ccab4bdbc02: Pull complete
Digest: sha256:632ecebdf89a36052e3eba281fdfa621a2afe5cd6b8061ad380ba3b3f0b25c01
Status: Downloaded newer image for daocloud.io/library/kibana:6.5.4
Creating elasticsearch ... done
Creating kibana        ... done

```

如果是上面这样就成了

如果你不放心,你可以执行

```shell
docker-compose logs -f
```

查看一下

。。。

卡了

```shell

[root@iz8g9301trfnpxz docker_es]# 
docker
dockre
s





[root@iz8g9301trfnpxz docker_es]# 



[root@iz8g9301trfnpxz docker_es]# 
[root@iz8g9301trfnpxz docker_es]# docker
dokxer
dokxer

```

然后咱们访问一下`ES`

![1598360012487](04_%E5%AE%89%E8%A3%85ES%E5%92%8CKibana.assets/1598360012487.png)

出现返回的`json`字符串,说明`ES`启动成功

---

然后我们访问一下`kibana`

![1598360150480](04_%E5%AE%89%E8%A3%85ES%E5%92%8CKibana.assets/1598360150480.png)

进来之后就是我们的kibana的面板

![1598360181039](04_%E5%AE%89%E8%A3%85ES%E5%92%8CKibana.assets/1598360181039.png)

我们主要关注`dev Tools`

咱们可以在这样的`Dev Tools`中,去编写基于`RESful`风格的`WEB`接口,去访问`ES`

![1598360322609](04_%E5%AE%89%E8%A3%85ES%E5%92%8CKibana.assets/1598360322609.png)

这样就可以进行测试了

---

还有一个需要关注的就是`Management`

![1598360372124](04_%E5%AE%89%E8%A3%85ES%E5%92%8CKibana.assets/1598360372124.png)

这里面能看到关于ES的信息