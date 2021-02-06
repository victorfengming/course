![image_10](img/image_10.png)


这里你如果想要设计一门语言,可以去了解

java虚拟机的规范的字节码文件



# jvm架构模型

![image_11](img/image_11.png)

### Java的架构模型主要有两种
- 基于栈式架构的特点:
- 
- 基于寄存器的架构特点:


```java
package com.atguigu.java;


public class StackStruTest {
    public static void main(String[] args) {
        int i = 2 + 3;

    }
}

```



```cmd
D:\IdeaProjects\study_note_01\out\production\code\com\atguigu\java>javap -v StackStruTest.class
Classfile /D:/IdeaProjects/study_note_01/out/production/code/com/atguigu/java/StackStruTest.class
  Last modified 2020-10-22; size 440 bytes
  MD5 checksum 574568c45d72c6731f6ae90223888b6f
  Compiled from "StackStruTest.java"
public class com.atguigu.java.StackStruTest
  minor version: 0
  major version: 52
  flags: ACC_PUBLIC, ACC_SUPER
Constant pool:
   #1 = Methodref          #3.#19         // java/lang/Object."<init>":()V
   #2 = Class              #20            // com/atguigu/java/StackStruTest
   #3 = Class              #21            // java/lang/Object
   #4 = Utf8               <init>
   #5 = Utf8               ()V
   #6 = Utf8               Code
   #7 = Utf8               LineNumberTable
   #8 = Utf8               LocalVariableTable
   #9 = Utf8               this
  #10 = Utf8               Lcom/atguigu/java/StackStruTest;
  #11 = Utf8               main
  #12 = Utf8               ([Ljava/lang/String;)V
  #13 = Utf8               args
  #14 = Utf8               [Ljava/lang/String;
  #15 = Utf8               i
  #16 = Utf8               I
  #17 = Utf8               SourceFile
  #18 = Utf8               StackStruTest.java
  #19 = NameAndType        #4:#5          // "<init>":()V
  #20 = Utf8               com/atguigu/java/StackStruTest
  #21 = Utf8               java/lang/Object
{
  public com.atguigu.java.StackStruTest();
    descriptor: ()V
    flags: ACC_PUBLIC
    Code:
      stack=1, locals=1, args_size=1
         0: aload_0
         1: invokespecial #1                  // Method java/lang/Object."<init>":()V
         4: return
      LineNumberTable:
        line 17: 0
      LocalVariableTable:
        Start  Length  Slot  Name   Signature
            0       5     0  this   Lcom/atguigu/java/StackStruTest;

  public static void main(java.lang.String[]);
    descriptor: ([Ljava/lang/String;)V
    flags: ACC_PUBLIC, ACC_STATIC
    Code:
      stack=1, locals=2, args_size=1
         0: iconst_5
         1: istore_1
         2: return
      LineNumberTable:
        line 19: 0
        line 21: 2
      LocalVariableTable:
        Start  Length  Slot  Name   Signature
            0       3     0  args   [Ljava/lang/String;
            2       1     1     i   I
}
SourceFile: "StackStruTest.java"


```


![image_12](img/image_12.png)