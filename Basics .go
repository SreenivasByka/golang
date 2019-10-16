				              Go
Go language is a programming language initially developed at Google in the year 2007 by Robert Griesemer, Rob Pike, and Ken Thompson, down load go software and install after check the install area can find the lot folders like api,bin,pkg,doc,src,test,lib,misc and files like AUTHORS,LICENSE,README,VERSION,CONTRIBUTORS (authors mail id info). bin- all .exe files available here ,src- files store here
Write go program any where follow    --Comments, Package Declaration, Import Packages, Functions ,Variables, Statements and Expressions
//first program

package main    //main pkg
import "fmt"   //import fmt pkg
func main() {

    fmt.Println("Hello Go Programming")
}
//after finish the program go run first.go
 E:\go\prat>go run first.go
O/P -  Hello Go Programming

fmt
. Println("Hello go!")
*Line Separator- nothing but statement terminator other programming use like ';'
*Comments -  //  and /*  and */  
*Identifiers -identify a variable or function or any other user defind item. Go case=sensitive like Man and man  both are different identifiers. examples of the identifiers-
byka sree b_sr b_123    name20 _tmp   z34er 
*Keyword- Go reserved keywords not use as variable or constant or any identifier.
break,default,func,interface,select,case,defer,Go,map,Struct,chan,else,Goto,   
package, Switch, const,fallthrough,if,range,Type,continue,for,import,return,Var .
#Variables-  variable is memory location to store a value of a specific type. variable declared so many ways ex-
var (keyword) age  int //variable declaration but 0 value
age=32         or    var age int =32  //var declaration and initial value.    or we can try  the below 
var (  
       name="bsr"
      age=32      //name and age are short hand declaration
      height  int
)
- All variable must be used,other wise through the error, initialize var but not use. Pascal or camelCase (HTTP,URL).
-Type conversion  destination Type(variable)   if  int convert string use or import "strconv"
*Boolean type- true or false   var n bool=true     or var n bool=false
n:=1 ==1         or var n bool (o/p   false it will run)
n:=2 ==2
*int 8  -128 to 127    int16  -32768 to 32767   int32   -2147483648 to  2147483647                                                                                                               	int64 - 9223372036854775808  to  9223372036854775807
*yvygvb njan ;.uint(unsigned integer)   var n uint16=42    uint8 0-255  uint16 0-65535   uint32 0- 4294967295
ex- //variables
package main    //package declaration
import "fmt"   //import fmt pkg
func main() {
    var a int =10
    var b int8 =4
    fmt.Println("a+b:",a+b)
    
}
thrown a error like mismatched type int and int8 we can chage int(b)
//variables
package main    //package declaration
import "fmt"   //import fmt pkg
func main() {
    var a int =10
    var b int8 =4
    fmt.Println("a+b:",a+int(b))
    
}

o/p=14
*bit operator  and &, or| , ^exclusive , &^ and not       mainly use for condition time
if we use in integer we got answer like system related(binary 0011,1100)
//variablesOperator
package main    //package declaration
import "fmt"   //import fmt pkg
func main() {
    var a int =10
    var b int =3

    fmt.Println("a&b:",a & b)
    fmt.Println("a|b:",a | b)
    fmt.Println("a^b:",a ^ b)
    fmt.Println("a&^b:",a &^ b)
    
}
0/p-      2 11 9 8
*Bit shifting -   << >>        //bit shiftOperator
package main    //package declaration
import "fmt"   //import fmt pkg
func main() {
    var a int =8
    fmt.Println("a<<3:",a << 3)  //2^3* 2^3=2^6
    fmt.Println("a>>b:",a >> 3) //2^3/ 2^3=2^0
}
o/p-    64   1
*floating point numbers-  float +/-1.18E - +/-3.4E38 follow IEEE standards Decimal no store.
package main    //package declaration
import "fmt"   //import fmt pkg
func main() {
    n :=3.14
    n=13.7e72
    n=2.1E14
    fmt.Println("Print n:",n)  
}
o/p - 2.1e+14
package main    //package declaration
import "fmt"   //import fmt pkg
func main() {
    //n :=3.14
    //n=13.7e72
    //n=2.1E14
    var n float64=3.14
    n=13.7e72
    n=2.1E14
    fmt.Println("Print n:",n)  
}
o/p-  
*Operator -  +,-,*,/
*Complex numbers-  complex64, complex128  ex- var a complex64 =1+2i
* Text type- simple string type   s="bsr" string stored in data byte.immutable.
*rune- rune is type of variable 32bit  value(utf unicode)   r:='a'  o/p 97, int32  var r rune='a',string is different than rune, rune is also an alias for int32,rune is character('' 0r back ticks ``)
#Constant- const keyword, Constants refer to fixed values that the program may not alter during its execution. basic data types like an integer constant, a floating constant, a character constant, or a string literal. Constants are treated just like regular variables except that their values cannot be modified after their definition..
#Variables-
//variable is store the memory location
package main
import "fmt"
func main() {
    var a int=14
    var b float64=4.12
    const c string="Sreenivasa Byka"
    //variable short hand declaration
    d:=10
    e:=5.12
    f:="Byka"
    //mutiple declarations...
    x,y,z:=10,20,"sree"
    fmt.Println("Hello Go Programming")
    fmt.Println("a values:",a)
    fmt.Println("b value:",b)
    fmt.Println("c value:",c)
    fmt.Println("d value:",d)
    fmt.Println("e value:",e)
    fmt.Println("f value:",f)
    fmt.Println("x,y,z Value:",x,y,z)
    *************
    Array   collection elememnts of single type or special variable accessed "Index" index starts with '0'
    *******************
var a [5]int 
b := [5]int{1, 2, 3, 4, 5} 

fmt.Println("array is:", b)
double array
array := [2][4]int{{1,2,3,4}, {5,6,7}}
***************************************************************************************************************
Slice -looklike Array but more power full and flexblnle than array
Slices dont store data, they are sections/slices of an array
***********************************************************************************************************
primes := [6]int{2,3,5,7,11,13} #array of prime numbers
var s []int = primes[1:4] # slice of prime array, values are from 1 to 4 index
var s []int = primes[:4] # slice, values from 0 to 4 index
len(s) # length of slice
cap(s) # max length of slice
***************************************************************************
Map - Key and Vlues combination (Python Dictionary)
****************************************************************************
m=make[string] int         m=["Byka"]=33
greet:=[string]stringn{
"bsr":"good morning",
"srinu":"Breafast??",
"other":"NA",
}
Methods-  delete,copy etc...

Conditionals (if/for/switch)
******* ***********************************************************************************************
For   go only 1 loop stmt that for we wire in 3 different ways intializtion(i:=1) condition(i<=5)increment(i++) we can in different lines
same line or 
*****

for i:=1;i<=5;i++ {
    fmt.Println(1)
    or 
    
 i := 1
   for i <= 5 {
   fmt.Println(i)
   i = i + 1   or i++
 }

 for j := 7; j <= 9; j++ {
   fmt.Println(j)
 }

 for {
   fmt.Println("continous")
 break
 }
****************************************************************************
if else if else   like conditions or conditional stmt
***************************************************************************
if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
*********************************************************************
Switch Case
********************************************************************
a := 2
    fmt.Print("write ", a, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
*****************************************************************************
Functions
*****************************************************************************
Group of statements or code reuse or lot of short code(min your code)


func add(a, b int) int {
    return a + b
}

func main () {
  result := add(1, 2)
  fmt.Println("1+2 =", result)
}

## 1+2 = 3


multi value return functions

func SumProduct(a, b int) (added, multiply int) {
  added = a+b
  multiplied = a*b
  return added, multiply
}


