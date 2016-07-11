package main

import "fmt"

const (
  a1 = iota
  b1
  c1
)

const (
  a2 = iota
  b2
  c2
)

type Circle struct{
  radius float64
}

func main(){
  var s string = "Another World!"
  var b int32
  var ptr *int32 = &b
  b, _ = 5, 7
  const c = "I love BJ"
  
  fmt.Println("Hello, World!")
  fmt.Println(s)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(a1,b1,c1)
  fmt.Println(a2,b2,c2)
  fmt.Println(ptr)
  
  var x interface{}
     
   switch i := x.(type) {
      case nil:	  
         fmt.Printf(" x  :%T",i)                
      case int:	  
         fmt.Printf("x  int ")                       
      case float64:
         fmt.Printf("x  float64 ")           
      case func(int) float64:
         fmt.Printf("x  func(int)")                      
      case bool, string:
         fmt.Printf("x  bool  string" )       
      default:
         fmt.Printf("未知型")     
    }
  
  var a3 int
	
  for a3 := 0; a3<10; a3++{
    fmt.Println("a3 = ", a3)
  }
  
  for a3<15  {
    a3++
    fmt.Println("a3 == ", a3)
	
  }
  
  numbers := [6]int{1,3,5,7}
  for i, j:=range numbers {
    fmt.Printf(" %d th = %d\n", i, j )
  }
  
  fmt.Println(numbers)
  
  var a4, b4 = 3, 5
  
  fmt.Println(a4, b4)
  
  a4, b4 = b4, a4

  fmt.Println(a4, b4)  
   
  var c3 Circle
  c3.radius = 3.0
  fmt.Println(c3.getArea())
  
  var slice1 []int
  slice1 =[]int {1,2,3,4,4,4,4,4,4,4}
  
  fmt.Println(slice1)
  
  var arr1 [5]int
  arr2:=[5]int{1,2,3}
  
  slice2 := arr2[:]
  
  fmt.Println(slice2)
  
  slice2[3] = 100
  fmt.Println(slice2[1:3])  
  var i int
  for i<3 {
    fmt.Println(slice1[i])
	fmt.Println(arr1[i])
	fmt.Println(arr2[i])
	i++
  }
  
  slice3 := arr2[1:3]
  fmt.Println(slice3)
  slice3[0] = 200
  fmt.Println(arr2)
  
  kvs := map[string]string{"a": "apple", "b": "banana"}
  for k, v := range kvs {
	fmt.Printf("%s -> %s\n", k, v)
}
}

func (c Circle) getArea() float64 {
  return 3.14*c.radius*c.radius
}




