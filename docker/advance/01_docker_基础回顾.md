# Docker安装





# Docker 安装

## 环境准备

1. 需要会一点点儿的Linux基础
2. Centos7
3. 我们使用Xshell连接远程服务器进行操作

## 环境查看

```
centos版本
3.10.0-514.26.2.el7.x86_64

```

```shell
## 系统版本
NAME="CentOS Linux"
VERSION="7 (Core)"
ID="centos"
ID_LIKE="rhel fedora"
VERSION_ID="7"
PRETTY_NAME="CentOS Linux 7 (Core)"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:centos:centos:7"
HOME_URL="https://www.centos.org/"
BUG_REPORT_URL="https://bugs.centos.org/"

CENTOS_MANTISBT_PROJECT="CentOS-7"
CENTOS_MANTISBT_PROJECT_VERSION="7"
REDHAT_SUPPORT_PRODUCT="centos"
REDHAT_SUPPORT_PRODUCT_VERSION="7"

```

## 安装

帮助文档

```
# 1. 卸载旧的版本

$ sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine

```

2.# 需要的安装包

```shell
sudo yum install -y yum-utils
```

# 3.设置镜像的仓库

```shell
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo

```

默认是从国外下载的,我们是使用国内的

docker阿里云国内镜像地址

```shell
sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
# 推荐使用阿里云的
```

在安装之前建议更新yum更新索引

```
sudo yum makecache fast
```

# 4. 现在就可以安装最新的docker的引擎

```shell
sudo yum install docker-ce docker-ce-cli containerd.io
```

docker-ce 是社区版本的
docker-ee 是企业版本的

# 5. 如果想要安装不是最新版本的,也可以指定安装的版本

```shell
sudo yum install docker-ce-<VERSION_STRING> docker-ce-cli-<VERSION_STRING> containerd.io
```

安装完成之后就可以启动docker,最后helloworld

```shell
sudo systemctl start docker
```


6. 使用docker version 看看安装成功没


```shell
[root@iz8g9301trfnpxz /]# docker version
Client: Docker Engine - Community
 Version:           19.03.12
 API version:       1.40
 Go version:        go1.13.10
 Git commit:        48a66213fe
 Built:             Mon Jun 22 15:46:54 2020
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          19.03.12
  API version:      1.40 (minimum version 1.12)
  Go version:       go1.13.10
  Git commit:       48a66213fe
  Built:            Mon Jun 22 15:45:28 2020
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.2.13
  GitCommit:        7ad184331fa3e55e52b890ea95e65ba581ae3429
 runc:
  Version:          1.0.0-rc10
  GitCommit:        dc9208a3303feef5b3839f4323d9beb36df0a9dd
 docker-init:
  Version:          0.18.0
  GitCommit:        fec3683

```





7. docker helloworld

```shell
[root@iz8g9301trfnpxz /]# docker run hello-world

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/

```

8. 查看一下下载的helloworld的镜像
   在不在

```shell
[root@iz8g9301trfnpxz /]# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
hello-world         latest              bf756fb1ae65        6 months ago        13.3kB
[root@iz8g9301trfnpxz /]# 
```

![1597802996844](01_docker_%E5%9F%BA%E7%A1%80%E5%9B%9E%E9%A1%BE.assets/1597802996844.png)

### docker Compose

### docker Swarm

### docker Stack

### docker Secret

### docker Config

容器单独它是没有意义的

有意义的是容器编排

### k8s