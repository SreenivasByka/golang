/* for loop specifies execution of a block,
iteration,condition/range ,increment/decrement*/
 package main
 import ( "fmt"
		//"math" 
)
func main() {
	//intial condition and increment different ways
	fmt.Printf("alphafunction")
	alpha()
	i:=1
	for i<=10{
		fmt.Println(i)
		i++    //i=i+1
	}
	for a:=2;a<=10;a=a+1{
		fmt.Println(a)
	}
	for {
		fmt.Println("Hello looping")
		break
	}
	for b:=0;b<=5;b++{
		if b%2==0{
			continue
		}
		fmt.Println(b)
	}
}
func alpha(){
	for ss:='a';ss<='z';ss++{
		fmt.Printf("%3c",ss)
	}
}

