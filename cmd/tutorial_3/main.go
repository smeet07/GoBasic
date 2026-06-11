package main
import (
	"fmt"
	"errors"
)
func main(){
	var printValue string="Hello World"
	printMe(printValue)
	var numerator,denominator int =9,2
	result,remainder,err:=intDivision(numerator,denominator)
	//if err!=nil{
		//fmt.Printf(err.Error())
	//} else if remainder==0{
		//fmt.Printf("Result of the integer division is %v",result)
	//}else{
	//fmt.Printf("The result of the division is %v with remainder %v",result,remainder)
	//}"""
	switch {
		case err!=nil:
			fmt.Printf(err.Error())
		case remainder==0:
			fmt.Printf("Result of the integer division is %v",result)
		default:
			fmt.Printf("The result of the division is %v with remainder %v",result,remainder)
	}
	switch remainder{
	case 0:
		fmt.Printf("\nDivision was exact")
	case 1,2:
		fmt.Printf("\nDivision was close")
	default:
		fmt.Printf("\nTry a closer value")
	}

}
func printMe(printValue string){
	fmt.Println(printValue)
}
func intDivision(num1 int, num2 int) (int,int,error) {
	var err error
	if num2==0{
		err=errors.New("Cannot divide by 0")
		return 0,0,err
	}
	var result int=num1/num2
	var remainder int=num1%num2
	return result,remainder,nil
}