# SVN分支

------

Branch 选项会给开发者创建出另外一条线路。当有人希望开发进程分开成两条不同的线路时，这个选项会非常有用。

比如项目 demo 下有两个小组，svn 下有一个 trunk 版。

由于客户需求突然变化，导致项目需要做较大改动，此时项目组决定由小组 1 继续完成原来正进行到一半的工作（某个模块），小组 2 进行新需求的开发。

那么此时，我们就可以为小组2建立一个分支，分支其实就是 trunk 版（主干线）的一个copy版，不过分支也是具有版本控制功能的，而且是和主干线相互独立的，当然，到最后我们可以通过（合并）功能，将分支合并到 trunk 上来，从而最后合并为一个项目。

我们在本地副本中创建一个 **my_branch** 分支。

```
root@runoob:~/svn/runoob01# ls
branches  tags  trunk
root@runoob:~/svn/runoob01# svn copy trunk/ branches/my_branch
A         branches/my_branch
root@runoob:~/svn/runoob01# 
```

查看状态：

```
root@runoob:~/svn/runoob01# svn status
A  +    branches/my_branch
A  +    branches/my_branch/HelloWorld.html
A  +    branches/my_branch/readme
```

提交新增的分支到版本库。

```
root@runoob:~/svn/runoob01# svn commit -m "add my_branch" 
Adding         branches/my_branch
Replacing      branches/my_branch/HelloWorld.html
Adding         branches/my_branch/readme

Committed revision 9.
```

接着我们就到 my_branch 分支进行开发，切换到分支路径并创建 index.html 文件。

```
root@runoob:~/svn/runoob01# cd branches/my_branch/
root@runoob:~/svn/runoob01/branches/my_branch# ls
HelloWorld.html  index.html  readme
```

将 index.html 加入版本控制，并提交到版本库中。

```
root@runoob:~/svn/runoob01/branches/my_branch# svn status
?       index.html
root@runoob:~/svn/runoob01/branches/my_branch# svn add index.html 
A         index.html
root@runoob:~/svn/runoob01/branches/my_branch# svn commit -m "add index.html"
Adding         index.html
Transmitting file data .
Committed revision 10.
```



切换到 trunk，执行 svn update，然后将 my_branch 分支合并到 trunk 中。



```
root@runoob:~/svn/runoob01/trunk# svn merge ../branches/my_branch/
--- Merging r10 into '.':
A    index.html
--- Recording mergeinfo for merge of r10 into '.':
 G   .
```

此时查看目录，可以看到 trunk 中已经多了 my_branch 分支创建的 index.html 文件。

```
root@runoob:~/svn/runoob01/trunk# ll
total 16
drwxr-xr-x 2 root root 4096 Nov  7 03:52 ./
drwxr-xr-x 6 root root 4096 Jul 21 19:19 ../
-rw-r--r-- 1 root root   36 Nov  7 02:23 HelloWorld.html
-rw-r--r-- 1 root root    0 Nov  7 03:52 index.html
-rw-r--r-- 1 root root   22 Nov  7 03:06 readme
```

将合并好的 trunk 提交到版本库中。

```
root@runoob:~/svn/runoob01/trunk# svn commit -m "add index.html"
Adding         index.html
Transmitting file data .
Committed revision 11.
```