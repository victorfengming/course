# Git Gitee

大家都知道国内访问 Github 速度比较慢，很影响我们的使用。

如果你希望体验到 Git 飞一般的速度，可以使用国内的 Git 托管服务——[Gitee（gitee.com）](https://gitee.com/?utm_source=remote_blog_cnjc)。

Gitee 提供免费的 Git 仓库，还集成了代码质量检测、项目演示等功能。对于团队协作开发，Gitee 还提供了项目管理、代码托管、文档管理的服务，5 人以下小团队免费。

接下来我们学习一下如何使用 Gitee。

由于我们的本地 Git 仓库和 Gitee 仓库之间的传输是通过SSH加密的，所以我们需要配置验证信息。

**1、我们先在 [Gitee](https://gitee.com/?utm_source=remote_blog_cnjc) 上注册账号并登录后，然后上传自己的 SSH 公钥。**

我们在 Git Github 章节已经生成了自己的 SSH 公钥，所以我们只需要将用户主目录下的 ~/.ssh/id_rsa.pub 文件的内容粘贴 Gitee 上。

选择右上角用户头像 -> 设置，然后选择 "SSH公钥"，填写一个便于识别的标题，然后把用户主目录下的 .ssh/id_rsa.pub 文件的内容粘贴进去：

![img](https://www.runoob.com/wp-content/uploads/2020/03/gitee1.png)

![img](https://www.runoob.com/wp-content/uploads/2020/03/gitee2.png)

成功添加后如下图所示：

![img](https://www.runoob.com/wp-content/uploads/2020/03/gitee3.png)

接下来我们创建一个项目。

点击右上角的 **+** 号，新建仓库：

![img](https://www.runoob.com/wp-content/uploads/2020/03/gitee4.png)

然后添加仓库信息：

![img](https://www.runoob.com/wp-content/uploads/2020/03/gitee5.png)

创建成功后看到如下信息：

![img](https://www.runoob.com/wp-content/uploads/2020/03/gitee6.png)

接下来我们看下连接信息：

![img](https://www.runoob.com/wp-content/uploads/2020/03/gitee7.png)

项目名称最好与本地库保持一致。

然后，我们在本地库上使用命令 **git remote add** 把它和 Gitee 的远程库关联：

```
git remote add origin git@gitee.com:imnoob/runoob-test.git
```

之后，就可以正常地用 git push 和 git pull 推送了！

如果在使用命令 git remote add 时报错：

```
git remote add origin git@gitee.com:imnoob/runoob-test.git
fatal: remote origin already exists.
```

这说明本地库已经关联了一个名叫 origin 的远程库，此时，可以先用 **git remote -v** 查看远程库信息：

```
git remote -v
origin    git@github.com:tianqixin/runoob.git (fetch)
origin    git@github.com:tianqixin/runoob.git (push)
```

可以看到，本地库已经关联了 origin 的远程库，并且，该远程库指向 GitHub。

我们可以删除已有的 GitHub 远程库：

```
git remote rm origin
```

再关联 Gitee 的远程库（注意路径中需要填写正确的用户名）：

```
git remote add origin git@gitee.com:imnoob/runoob-test.git
```

此时，我们再查看远程库信息：

```
git remote -v
origin    git@gitee.com:imnoob/runoob-test.git (fetch)
origin    git@gitee.com:imnoob/runoob-test.git (push)
```

现在可以看到，origin 已经被关联到 Gitee 的远程库了。

通过 git push 命令就可以把本地库推送到 Gitee 上。

有的小伙伴又要问了，一个本地库能不能既关联 GitHub，又关联 Gitee 呢？

答案是肯定的，因为 git 本身是分布式版本控制系统，可以同步到另外一个远程库，当然也可以同步到另外两个远程库。

使用多个远程库时，我们要注意，git 给远程库起的默认名称是 origin，如果有多个远程库，我们需要用不同的名称来标识不同的远程库。

仍然以 runoob-test 本地库为例，我们先删除已关联的名为 origin 的远程库：

```
git remote rm origin
```

然后，先关联 GitHub 的远程库：

```
git remote add github git@github.com:tianqixin/runoob-git-test.git
```

注意，远程库的名称叫 github，不叫 origin 了。

接着，再关联 Gitee 的远程库：

```
git remote add gitee git@gitee.com:imnoob/runoob-test.git
```

同样注意，远程库的名称叫 gitee，不叫 origin。

现在，我们用 git remote -v 查看远程库信息，可以看到两个远程库：

```
git remote -v
gitee    git@gitee.com:imnoob/runoob-test.git (fetch)
gitee    git@gitee.com:imnoob/runoob-test.git (push)
github    git@github.com:tianqixin/runoob.git (fetch)
github    git@github.com:tianqixin/runoob.git (push)
```

如果要推送到 GitHub，使用命令：

```
git push github master
```

如果要推送到 Gitee，使用命令：

```
git push gitee master
```

这样一来，我们的本地库就可以同时与多个远程库互相同步：

![img](https://www.runoob.com/wp-content/uploads/2020/03/gitee8.png)