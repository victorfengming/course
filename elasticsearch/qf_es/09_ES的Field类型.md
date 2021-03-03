![1598361274078](09_ES%E7%9A%84Field%E7%B1%BB%E5%9E%8B.assets/1598361274078.png)

## 3.4 ES中field可以指定类型


String

[`text`](https://www.elastic.co/guide/en/elasticsearch/reference/6.5/text.html) and [`keyword`](https://www.elastic.co/guide/en/elasticsearch/reference/6.5/keyword.html)

**[Numeric datatypes](https://www.elastic.co/guide/en/elasticsearch/reference/6.5/number.html)**

```
long`, `integer`, `short`, `byte`, `double`, `float`, `half_float`, `scaled_float
```

**[Date datatype](https://www.elastic.co/guide/en/elasticsearch/reference/6.5/date.html)**

```
date
```

**[Boolean datatype](https://www.elastic.co/guide/en/elasticsearch/reference/6.5/boolean.html)**

```
boolean
```

**[Binary datatype](https://www.elastic.co/guide/en/elasticsearch/reference/6.5/binary.html)**

```
binary
```

**[Range datatypes](https://www.elastic.co/guide/en/elasticsearch/reference/6.5/range.html)**

```
integer_range`, `float_range`, `long_range`, `double_range`, `date_range
```



字符串:
>
>​	text: 一把被用于全文检索.将当前Field进行分词.
>
>​	keyword: 当前Field不会被分词
>
>数值类型:
>
>​	long:
>
>​	Integer:
>
>​	short:
>
>​	byte:
>
>​	double:
>
>​	float:
>
>​	half_float: 精度比float小一半.
>
>​	scaled_float: 根据一个long和scaled来表达一个浮点型,long-345,scaled-100 -> 3.45
>
>时间类型:
>
>​	date类型,针对时间类型指定具体格式
>
>布尔类型:
>
> 	boolean类型,表达true和false
>
>二进制类型:
>
>​	binary类型暂时支持Base64 encode string
>
>范围类型:
>
>​	long_range: 赋值时,无序指定具体内容,只需要存储一个范围即可,指定gt,lt,gte,lte
>
>​	integer_range:
>
>​	double_range:
>
>​	date_range:
>
>​	ip_range:
>
>​	float_range:
>
>经纬度类型:
>
>​	geo_point: 用来存储经纬度的
>
>ip类型:
>
>​	ip: 可以存储IPV4或者IPV6

![1598426267491](09_ES%E7%9A%84Field%E7%B1%BB%E5%9E%8B.assets/1598426267491.png)