### Linux下安装

全网最新版本

1. 下载安装包! redis-5.0.8.tar.gz

2. 解压Redis的安装包! 程序/opt

   ![1597152081477](08_Linux%E4%B8%8B%E5%AE%89%E8%A3%85Redis.assets/1597152081477.png)

3. 进入解压后的文件,

   ![1597152243206](08_Linux%E4%B8%8B%E5%AE%89%E8%A3%85Redis.assets/1597152243206.png)

4. 基本的环境安装

   ```shell
   yum install gcc-c++
   make
   make install
   ```

   ![1597152258681](08_Linux%E4%B8%8B%E5%AE%89%E8%A3%85Redis.assets/1597152258681.png)

   ![1597152313871](08_Linux%E4%B8%8B%E5%AE%89%E8%A3%85Redis.assets/1597152313871.png)

5. redis默认的路径`usr/local/bin`

   ![1597152402392](08_Linux%E4%B8%8B%E5%AE%89%E8%A3%85Redis.assets/1597152402392.png)

6. 将redis配置文件.复制到我们的目录下

   ![1597152480502](08_Linux%E4%B8%8B%E5%AE%89%E8%A3%85Redis.assets/1597152480502.png)

7. redis默认不是后台启动的,修改配置文件

   ![1597152552417](08_Linux%E4%B8%8B%E5%AE%89%E8%A3%85Redis.assets/1597152552417.png)

8. 启动Redis服务

   ![1597152707942](08_Linux%E4%B8%8B%E5%AE%89%E8%A3%85Redis.assets/1597152707942.png)

9. 用一下试试

   使用redis-cli 进行连接测试

   ![1597193393331](09_Linux%E4%B8%8B%E5%AE%89%E8%A3%85Redis.assets/1597193393331.png)

   ![1597193450033](09_Linux%E4%B8%8B%E5%AE%89%E8%A3%85Redis.assets/1597193450033.png)

10. 查看redis的进程是否开启

    ![1597193508202](09_Linux%E4%B8%8B%E5%AE%89%E8%A3%85Redis.assets/1597193508202.png)

11. 如何关闭Redis服务

    ![1597193576677](09_Linux%E4%B8%8B%E5%AE%89%E8%A3%85Redis.assets/1597193576677.png)

12. 再次查看进程是否存在

    ![1597193602121](09_Linux%E4%B8%8B%E5%AE%89%E8%A3%85Redis.assets/1597193602121.png)

13. 后面我们会使用单机多redis启动测试集群

    







