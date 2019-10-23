/* variable store the memory
static type declaration,dynamic type,mixed type 
short hand type
varible type */
 package main
 import "fmt"

 func main() {
 //stastic type
 var x float64
 x=20.15
 fmt.Printf("Static type x %T\n ",x)
 fmt.Println("x value",x)
 //dynamic type or short hand type
 y:=25
 fmt.Println("value y",y)
 fmt.Printf("dynamic type y %T\n ",y)
 fmt.Println("**************")
 //mixed type
 var a,b,c=1,2,"sree"
 fmt.Println("value a",a,"value b",b,"value c",c)
 fmt.Printf("value a %T\n a",a)
 fmt.Printf("value b %T\n b",b)
 fmt.Printf("value c %T\n c",c)
 }


 //   o/p 
//  Hello Go
// Hello Golang