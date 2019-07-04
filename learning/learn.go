package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
  fmt.Println("Conne"+"ction.")

  // var a string = "Run go"
  // fmt.Println(a)

  // var b, c int = 1, 2
  // fmt.Println(b, c)

  // d := 10
  // fmt.Println(d)

  const LENGTH int = 10
  const WIDTH int = 5  
  fmt.Println(LENGTH,WIDTH) 

  // const (
  //   a = iota   //0
  //   b          //1
  //   c          //2
  //   d = "ha"   //独立值，iota += 1
  //   e          //"ha"   iota += 1
  //   f = 100    //iota +=1
  //   g          //100  iota +=1
  //   h = iota   //7,恢复计数
  //   i          //8
  // )
  // fmt.Println(a,b,c,d,e,f,g,h,i)


	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );

	/*  & 和 * 运算符实例 */
	ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a);
	fmt.Printf("*ptr 为 %d\n", *ptr);

  s1,s2 := swap("Google","Huawei")
  fmt.Println(s1,s2)
  
  // array
  var array = [10] int {10,23,4,32,5544,34,54,64,43,23}

  var n [10] int /* n 是一个长度为 10 的数组 */
  var i,j int
  /* 为数组 n 初始化元素 */         
  for i = 0; i < 10; i++ {
    n[i] = array[i] + 100 /* 设置元素为 i + 100 */
  }

  /* 输出每个数组元素的值 */
  for j = 0; j < 10; j++ {
    fmt.Printf("Element[%d] = %d\n", j, n[j] )
  }
 
  // pointer
  var ip *int        /* 声明指针变量 */

  ip = &a  /* 指针变量的存储地址 */
  fmt.Printf("a 变量的地址是: %x\n", &a)
  /* 指针变量的存储地址 */
  fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
  /* 使用指针访问值 */
  fmt.Printf("*ip 变量的值: %d\n", *ip) 

  var  ptr1 *int
  fmt.Printf("ptr 的值为 : %x\n", ptr1)

}

func swap(x, y string) (string, string) {
   return y, x
}
