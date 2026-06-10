package main
import "fmt"
func main(){
	var intNum int= 32767
	intNum=intNum+1
	fmt.Println(intNum)
	var floatNum float64=12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32=10.1
	var intNum32 int32=2
	var result float32=floatNum32 +float32(intNum32)
	fmt.Println(result)

	var num1 int=2
	var num2 int=4
	fmt.Println(num1/num2)
	fmt.Println(num1%num2)

	var myString string="Hello" + " " + "World"
	fmt.Println(myString)
	//Prints number of bytes not characters
	fmt.Println(len(myString))

	//will print 97
	var myRune rune='a'
	fmt.Println(myRune)

	var myBool bool=true
	fmt.Println(myBool)

	var intNum3 rune
	fmt.Println(intNum3)

	var myVar="text"
	fmt.Println(myVar)

	myvar:="text"
	fmt.Println(myvar)

	var1,var2:=10,20
	fmt.Println(var1,var2)

	//must be initialized with a value  and cannot be changed later
	const pi float64=3.14159
	fmt.Println(pi)


}