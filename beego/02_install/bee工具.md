# bee工具的使用

## 简介

bee工具是用于管理我们的beego项目的工具,你可以用bee来创建你的beego项目

并且,很容易的进行 beego 项目的创建、热编译、开发、测试、和部署。 



![1598103178724](bee%E5%B7%A5%E5%85%B7.assets/1598103178724.png)

## 安装

与beego的安装方式一样,我们可以通过如下的方式安装 bee 工具：

```shell
go get github.com/beego/bee
```

安装完之后，`bee` 可执行文件默认存放在 `$GOPATH/bin` 里面，所以您需要把 `$GOPATH/bin` 添加到您的环境变量中，才可以进行下一步。

如何添加环境变量，请参考:[go环境变量配置 (GOROOT和GOPATH)](./goroot_gopath.html)
如果你本机设置了 `GOBIN`，那么上面的命令就会安装到 `GOBIN` 下，请添加 GOBIN 到你的环境变量中



