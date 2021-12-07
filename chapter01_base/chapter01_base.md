[toc]
# chapter1 base
## 变量
### var
使用`var`定义变量，和c不同，类型放在变量名后面。
```go
var x int                    // 自动初始化为0
var a, b *int               // 相同类型的多个变量
var a, s = 100, "abc"   // 不同类类型初始化值
var (
x,y int
var a, s = 100, "abc"
)
```
变量的命名规则遵循骆驼命名法，首字母小写。
但全局变量希望能够被外部包所使用，需要将首个单词的首字母大写
### 赋值
声明与赋值（初始化）语句也可以组合起来。
```go
var identifier [type] = value
var a int = 15
var i = 5
var b bool = false
var str string = "Go says hello to the world!"
```
Go编译器可以根据变量值推算变量类型，有点像Python和Ruby这类动态语言。
```go
var a = 15
var b = false
var str = "Go says hello to the world!"
```
自动推断类型并不都适用
```go
var n int64 = 2
```
然而，`var a`这种语法是不正确的，因为编译器没有自动推算的依据。变量的类型可以在运行时自动推断。
```go
var (
HOME = os.Getenv("HOME")
USER = os.Getenv("USER")
GOROOT = os.Getenv("GOROOT")
)
```
### 简短形式
```go
a, b, c := 5, 7, "abc" // 注意等号前的冒号
```
## 常量
### const
用于定义不会改变的数据
常量的定义格式：`const identifier [type] = value`
```go
const b string = "abc"    // 显示定义
const b = "abc"            // 隐式定义
```
### iota
iota是go语言的常量计数器，只能在常量的表达式中使用。
```go
const (   
a = iota  // a = 0   
b          // b = 1   
c           // c = 2   
d = 5     // d = 5   
e          // e = 5)

```
Go语言将数据类型分为四类：基础类型、复合类型、引用类型和接口类型
## 基础数据类型（underlying type）
### boolean
布尔值true或false，
### numeric
数值型包括了整型和浮点型
```
有符号整形数类型:
int8,长度:1字节, 取值范围:(-128 ~ 127)
int16,长度:2字节,取值范围:(-32768 ~ 32767）
int32,长度:4字节,取值范围:(-2,147,483,648 ~ 2,147,483,647）
int64.长度:8字节,取值范围:(-9,223,372,036,854,775,808 ~ 9,223,372,036,854,775,807)
无符号整形数类型:
uint8,长度:1字节, 取值范围:(0 ~ 255)
uint16,长度:2字节,取值范围:(0 ~ 65535)
uint32,长度:4字节,取值范围:(0 ~ 4,294,967,295)
uint64.长度:8字节,取值范围:(0 ~ 18,446,744,073,709,551,615)

float32:（+- 1e-45 -> +- 3.4 * 1e38）32位浮点类型
float64:（+- 5 1e-324 -> 107 1e308）64位浮点类型

complex64: 由两个float32类型的值分别表示复数的实数部分和虚数部分
complex128: 由两个float64类型的值表示复数的实数部分和虚数部分

byte        alias for uint8
rune        alias for int32
```
### [string](./string.go)
string类型不可变，可以为空。如需修改string内容需先转换为[]byte或[]rune，修改后的string重新分配内存
标准库中有四个包对字符串处理尤为重要：
*strings*包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能。
*bytes*包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。
*strconv*包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。
*unicode*包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类

```go
s := "hello, world"
fmt.Println(len(s))     // "12"
fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
```

## 引用数据类型（Reference types）

### array
#### 概况
数组是由相同元素的集合组成的数据结构，计算机会为数组分配一段连续的内存来保存其中的元素。
GO数组在初始化之后大小无法改变。存储类型相同并且大小相同的数组才是同一类型。

#### 初始化
两种创建方式
```go
var arr0 = new([3]int)
arr1 := [3]int{1, 2, 3}		// 指定数组大小
arr2 := [...]int{1, 2, 3}	// 编译器对数组大小进行推导
```
#### 访问
Go在编译期间会对数组进行静态类型检查判断数组是否越界。在运行时发现越界操作会由运行时的**runtime.panicIndex**和**runtime.goPanincIndex**触发程序运行错误并退出。

### [slice](slice.go)
#### 概括
切片是对数组一个连续片段的的引用，所以不需要使用额外的内存并且比数组更高效，类似于python中的list。切片的长度是动态的，声明时只需指定元素类型。

#### 初始化
初始切片的方式
```go
arr[0:3] or slice[0:3]		// 通过下标的方式获得数组或切片的一部分
slice := []int{1, 2, 3}		// 使字面量初始化新的切片
slice := make([]int, 10)	// 使用关键字make城建切片
```
#### 访问、赋值、append
切片的赋值是一次浅拷贝的操作，对一个切片修改会影响另外一个切片。
每次append都会形成新的切片变量，如果底层数组没有扩容，那么就共享数据，如果扩容，则数据不共享。
扩容策略：
1. 如果期望容量大于当前容量的两倍就会使用期望容器
2. 如果当前切片的长度小于1024就会将容量翻倍
3. 如果当前切片的长度大于1024就会每次增加25%的容量，直到大于期望容量
#### copy
copy将类型为T的切边从源地址src拷贝到目标地址dst，覆盖dst的相关元素，并返回拷贝的元素个数。拷贝个数是src和dst的长度最小值。如果src是字符串那么元素类型就是byte。如还想继续使用src，在copy后执行sec=dst。

### map
#### 概括
map 是引用类型，数组切片可以操作一块连续内存的能力，对同质元素的统一管理。而字典则赋予了不连续不同类的内存变量的关联性。
#### 创建

```go
var m map[int]string = make(map[int]string) // 使用make初始化
var m4 = make(map[int]string, 16)           //make指定长度，避免多次扩容
var m1 map[int]string = map[int]string{ //指定类型
    90: "excellent",
    80: "good",
    60: "pass",
}
var m2 = map[int]string{} //类型推导
```

#### 读写

```go
var score = fruits["apple"]
```

### struct
#### 概括
#### 初始化

```go
var c1 Circle = Circle{}
var c2 Circle
var c3 *Circle = new(Circle)
var c *Circle = nil
```

### pointer
## 自定义类型

## interface
### 概括
Go通过接口实现面向对象的特性。
#### 初始化
```go
type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
}
```
#### 接口嵌套
一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。
#### 类型断言
一个接口类型的变量 varI 中可以包含任何类型的值，必须有一种方式来检测它的动态类型，即运行时在变量中存储的值的实际类型。
type-switch

```go
switch t := areaIntf.(type){
case*Square:
    fmt.Printf("Type Square %T with value %v\n", t, t)
case*Circle:
    fmt.Printf("Type Circle %T with value %v\n", t, t)
casenil:
    fmt.Printf("nil value: nothing to check?\n")
default:
    fmt.Printf("Unexpected type %T\n", t)
}
```
## channel
### 概括
channel是进程内的通信方式



