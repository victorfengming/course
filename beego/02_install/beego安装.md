# beego安装

## 安装

### 在安装前需要准备:

- 连接网络,如果无法连接互联网目前beego是不可以安装的

- git,如果没有安装可以参考[这里](https://victorfengming.gitee.io/2019/08/20/progit-min/)

- 配置git的https验证

  ```shell
  git config --global http.sslVerify false
  ```
  
  

### beego 的安装是典型的 Go 安装包的形式：

```shell
go get github.com/astaxie/beego
```



![1598102211707](beego%E5%AE%89%E8%A3%85.assets/1598102211707.png)

## 升级

beego 升级分为 go 方式升级和源码下载升级：

- Go 升级,通过该方式用户可以升级 beego 框架，强烈推荐该方式：

  ```
  go get -u github.com/astaxie/beego
  ```

- 源码下载升级，用户访问 `https://github.com/astaxie/beego` ,下载源码，然后覆盖到 `$GOPATH/src/github.com/astaxie/beego` 目录，然后通过本地执行安装就可以升级了：

  ```
  go install  github.com/astaxie/beego
  ```

## beego 的 git 分支

 beego 的 master 分支为相对稳定版本，dev 分支为开发者版本。大致流程如下： 



![1598103015506](beego%E5%AE%89%E8%A3%85.assets/1598103015506.png)

