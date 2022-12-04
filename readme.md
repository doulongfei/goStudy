# go begin
### 1.ch1包下，分别练习了使用go操作linux  echo1.go echo2.go echo3.go 和cd1执行shell命令
### 2.使用原生http包发送get请求的fetch两个相关练习
### 3.使用gif相关库，生成gif图片 lissajous.go
最后使用 原生http实现了三个简单的http服务 server1-3.监听9001端口 ，在linux上运行lissajous  访问localhost:9001/gif
即可看到lissajous生成gif图片

## go 指针
> 	x := 1
> 	p := &x //代表某个变量的地址
> 	fmt.Println("值", *p)
> 	fmt.Println("地址&x=p=", p)
> 	*p = 2 //*xxx 是某个变量，可以直接对其进行操作（赋值）相当于变量
> 	 问：*xxx 和变量 x的区别是什么？？？
> 	每次我们对一个变量取地址，或者复制指针，我们都是为原变量创建了新的别名。例如，*p就是变量v的别名。
> 	指针特别有价值的地方在于我们可以不用名字而访问一个变量，

# go ch2
Go语言将数据类型分为四类：
**基础类型、复合类型、引用类型和接口类型。**
> 基础类型，包括：**数字、字符串和布尔型**。
> 复合数据类型——**数组**和**结构体**——是通过组合简单类型，来表达**更加复杂的数据结构**。
> 引用类型包括**指针、切片、字典、函数、通道**

## int int(float) :只取整数部分
> int8,int16,int32,int64:分别对应8、16、32、64bit
> 大小的有符号整数。与之对应的4个无符号整数
> > 这里还有两种一般对应特定CPU平台机器字大小的有符号和无符号整数int和uint；其中int是应用最广泛的数值类型。这两种类型都有同样的大小，32或64bit，但是我们不能对此做任何的假设；因为不同的编译器即使在相同的硬件平台上可能产生不同的大小。

## float32 float64

## 复数  complex64 complex126分别对应float32和float64
> 内建的real和imag函数分别返回复数的实部和虚部
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y)                 // "(-5+10i)"
fmt.Println(real(x*y))           // "-5"
fmt.Println(imag(x*y))   