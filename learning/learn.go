package main

import (
  "fmt"
  "time"
)

// 结构体
type Books struct {
   title string
   author string
   subject string
   book_id int
}

// interface
type Phone interface {
    call()
}
type NokiaPhone struct {
}
func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}
type IPhone struct {
}
func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}


func main() {
/*  
  fmt.Println("Hello, World!")
  fmt.Println("Conne"+"ction.")
*/

// 第一种，指定变量类型，如果没有初始化，则变量默认为零值。
  // var a string = "Run go"
  // fmt.Println(a)

// 第二种，根据值自行判定变量类型。
  // var b = 2
  // fmt.Println(b)

// 第三种，省略 var, 注意若 := 左侧的不是声明新的变量，就产生编译错误，格式：
  // d := 10
  // fmt.Println(d)


// 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
/*
  const LENGTH int = 10
  const WIDTH int = 5  
  fmt.Println(LENGTH,WIDTH) 
*/

// 常量还可以用作枚举
// iota，特殊常量，可以认为是一个可以被编译器修改的常量。
// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
/*
  const (
    a = iota   //0
    b          //1
    c          //2
    d = "ha"   //独立值，iota += 1
    e          //"ha"   iota += 1
    f = 100    //iota +=1
    g          //100  iota +=1
    h = iota   //7,恢复计数
    i          //8
  )
  fmt.Println(a,b,c,d,e,f,g,h,i)
*/

// 运算符实例
/*
  var a int = 4
  var b int32
  var c float32  
  fmt.Printf("a 变量类型为 = %T\n", a );
  fmt.Printf("b 变量类型为 = %T\n", b );
  fmt.Printf("c 变量类型为 = %T\n", c );
  
  var ptr *int
  ptr = &a     // 'ptr' 包含了 'a' 变量的地址 
  fmt.Printf("a 的值为  %d\n", a);
  fmt.Printf("*ptr 为 %d\n", *ptr);
*/

// 函数用法  
/*
  s1,s2 := swap("Google","Huawei")
  fmt.Println(s1,s2)
*/

// array  
/*
  var array = [10] int {10,23,4,32,5544,34,54,64,43,23}
  var n [10] int // n 是一个长度为 10 的数组 
  
  var i,j int
  // 为数组 n 初始化元素      
  for i = 0; i < 10; i++ {
    n[i] = array[i] + 100  // 设置元素为 i + 100
  }
  // 输出每个数组元素的值
  for j = 0; j < 10; j++ {
    fmt.Printf("Element[%d] = %d\n", j, n[j] )
  }
*/

// pointer
/*
  var ip *int     

  ip = &a
  fmt.Printf("变量a的地址是: %x\n", &a)
  //指针变量的存储地址 
  fmt.Printf("变量ip储存的指针地址: %x\n", ip )
  // 使用指针访问值 
  fmt.Printf("变量ip的值: %d\n", *ip) 

  var  ptr1 *int
  fmt.Printf("ptr的值为: %x\n", ptr1)
*/


// 结构体
  var Book1 Books        // Declare Book1 of type Book 
  var Book2 Books        // Declare Book2 of type Book

  // book 1 描述
  Book1.title = "Go 语言"
  Book1.author = "www.runoob.com"
  Book1.subject = "Go 语言教程"
  Book1.book_id = 6495407
  // book 2 描述
  Book2.title = "Python 教程"
  Book2.author = "www.runoob.com"
  Book2.subject = "Python 语言教程"
  Book2.book_id = 6495700

  printBook(&Book1)  
  printBook(&Book2)

// 切片
/*
  可以声明一个未指定大小的数组来定义切片
  var identifier [] type
  切片不需要说明长度。
  或使用make()函数来创建切片:
  var slice1 [] type = make([]type, len)
  slice1 := make([]type, len)
  也可以指定容量，其中capacity为可选参数。
  make([]type, length, capacity)
  这里 len 是数组的长度并且也是切片的初始长度。
*/  
  // 创建切片
  numbers := []int{0,1,2,3,4,5,6,7,8}  
  printSlice(numbers)
  // 打印原始切片 
  fmt.Println("numbers ==", numbers)
  // 打印子切片从索引1(包含) 到索引4(不包含)
  fmt.Println("numbers[1:4] ==", numbers[1:4])
  // 默认下限为 0
  fmt.Println("numbers[:3] ==", numbers[:3])
  // 默认上限为 len(s)
  fmt.Println("numbers[4:] ==", numbers[4:])
  
  numbers1 := make([]int,0,5)
  printSlice(numbers1)

  // 打印子切片从索引  0(包含) 到索引 2(不包含)
  number2 := numbers[:2]
  printSlice(number2)
  
  // 打印子切片从索引 2(包含) 到索引 5(不包含) 
  number3 := numbers[2:5]
  printSlice(number3)

  // append
  // 允许追加空切片 
  numbers = append(numbers, 0)
  printSlice(numbers)
  // 向切片添加一个元素 
  numbers = append(numbers, 1)
  printSlice(numbers)
  // 同时添加多个元素 
  numbers = append(numbers, 2,3,4)
  printSlice(numbers)

  // 创建切片 numbers1 是之前切片的两倍容量
  numbers4 := make([]int, len(numbers), (cap(numbers))*2)
  // 拷贝 numbers 的内容到 numbers4
  copy(numbers4,numbers)
  printSlice(numbers4)   

// 范围
// Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
// 在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对的 key 值。 
// 使用range去求一个slice的和。使用数组跟这个很类似
  nums := []int{2, 3, 4}
  sum := 0
  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)  
  // 在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。
  for i, num := range nums {    
    fmt.Println("index: ", i)
    fmt.Println("num:   ", num)
  }

// range也可以用在map的键值对上。
  kvs := map[string]string{"a": "apple", "b": "banana"}
  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }

// range也可以用来枚举 Unicode 字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
  for i, c := range "go" {
    fmt.Println(i, c)
  }

// map
// 可以使用内建函数 make 也可以使用 map 关键字来定义 Map:
// 声明变量，默认 map 是 nil
// var map_variable map[key_data_type]value_data_type
// 使用 make 函数 
// map_variable := make(map[key_data_type]value_data_type)
// 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
  // 创建集合 
  var countryCapitalMap map[string]string 
  countryCapitalMap = make(map[string]string)

  // map插入key - value对,各个国家对应的首都
  countryCapitalMap [ "France" ] = "巴黎"
  countryCapitalMap [ "Italy" ] = "罗马"
  countryCapitalMap [ "Japan" ] = "东京"
  countryCapitalMap [ "India " ] = "新德里"

  // 使用键输出地图值
  for country := range countryCapitalMap {
    fmt.Println(country, " 的首都是", countryCapitalMap [country])
  }

  // 查看元素在集合中是否存在
  capital, ok := countryCapitalMap [ "American" ] // 如果确定是真实的,则存在,否则不存在 
  if (ok) {
    fmt.Println("American 的首都是", capital)
  } else {
    fmt.Println("American 的首都不存在")
  }
  // 删除元素
  delete(countryCapitalMap, "France")
  fmt.Println(countryCapitalMap)

// interface
  var phone Phone
  phone = new(NokiaPhone)
  phone.call()
  phone = new(IPhone)
  phone.call()

// 并发
  go say("world")
  say("hello")

// channel  
  s := []int{7, 2, 8, -9, 4, 0}
  c := make(chan int)
  go go_sum(s[:len(s)/2], c)
  go go_sum(s[len(s)/2:], c)
  x, y := <-c, <-c // 从通道 c 中接收
  fmt.Println(x, y, x+y)

// 通道缓冲
  // 定义了一个可以存储整数类型的带缓冲通道
  // 缓冲区大小为2
  ch := make(chan int, 2)
  // 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
  // 而不用立刻需要去同步读取数据
  ch <- 1
  ch <- 2
  // 获取这两个数据
  fmt.Println(<-ch)
  fmt.Println(<-ch)

// 遍历通道  
  ch1 := make(chan int, 10)
  go fibonacci(cap(ch1), ch1)
  // range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
  // 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
  // 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
  // 会结束，从而在接收第 11 个数据的时候就阻塞了。
  for i := range ch1 {
    fmt.Println(i)
  }

}

func fibonacci(n int, c chan int) {
  x, y := 0, 1
  for i := 0; i < n; i++ {
    c <- x
    x, y = y, x+y
  }
  close(c)
}

func go_sum(s []int, c chan int) {
  sum0 := 0
  for _, v := range s {
    sum0 += v
  }
  c <- sum0 // 把 sum 发送到通道 c
}


func say(s string) {
  for i := 0; i < 5; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
  }
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func printBook( book *Books ) {
  fmt.Printf( "Book title : %s\n", book.title);
  fmt.Printf( "Book author : %s\n", book.author);
  fmt.Printf( "Book subject : %s\n", book.subject);
  fmt.Printf( "Book book_id : %d\n", book.book_id);
}
  

/*
func swap(x, y string) (string, string) {
   return y, x
}
*/
