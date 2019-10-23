/* constant  fixed variable can't change
'const' keyword to use*/
 package main
 import ( "fmt"
		"math" 
)
 const s int=10
 const a string="sree"
 func main() {
	 fmt.Println("value of s",s,"value of a",a)
 //constant variable
	const n=500000000
    const b=3e20 /n
    fmt.Printf("value of b %T\n",b)
    fmt.Println(int64(b))
    fmt.Printf("type a %T\n ",s)
	fmt.Printf("type of a %T\n",a)
	fmt.Println(math.Sin(n))

 }

