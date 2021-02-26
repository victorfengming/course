

# Exact VM

- 为了解决上一个虚拟机问题, jdk1.2时,sun 提供了此虚拟机.

- Exact Memory Management: 准确式内存管理
  - 也可以叫Non-Conservative/Accurate Memory Management
  - 虚拟机可以知道内存中某个位置的数据具体是什么类型
- 具备现代高性能虚拟机的雏形
  - 热点探测
  - 编译器与解释器混合工作模式
- 只在Solaris平台短暂使用,其他平台上还是classic vm
  - 英雄气短,终被Hotspot 虚拟机替换


# HotSpot vm

![image_15](img/image_15.png)

Hotspot 占有绝对的市场地位,称霸武林

# Jrockit

![image_16](img/image_16.png)

专注于服务器端应用

它可以不太关注程序启动速度,因此,他内部不包含解析器的实现,全部代码都靠即时编译器编译后运行

# IBM J9

![image_17](img/image_17.png)


# KVM
![image_18](img/image_18.png)

JAVA ME 产品线 - 现在不成了 诺基亚用的是这个

# Azul Vm

![image_19](img/image_19.png)

和硬件绑定,这样性能更好,可以针对这个硬件资源记性优化

Zing虚拟机在低延迟和快速预热的方面比较好

# Liquid Vm

![image_20](img/image_20.png)

# Microsoft vm
![Q`X@MPCQ7{R)5O6}D$0272C](img/Q%60X%40MPCQ7%7BR%295O6%7DD%240272C.png)
# Taobao Jvm

![5SAKI2[ZD(O14){V9FNB({G](img/5SAKI2%5BZD%28O14%29%7BV9FNB%28%7BG.png)

在多个虚拟机中共享对象

有点是速度高

缺点是和硬件intel的cpu没啥耦合

# Dalvik Vm

安卓系统,在5.0 之前 使用的就是Dalvik虚拟机


![image_22](img/image_22.png)

# others

![image_23](img/image_23.png)


![image_24](img/image_24.png)

# Graal Vm

![image_25](img/image_25.png)

执行程序更快,

跨语言的全站虚拟机






