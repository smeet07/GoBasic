package main
import "fmt"
type gasEngine struct{
	mpg uint8
	gallons uint8
	owner
	int
}
type owner struct{
	name string
}
type electricEngine struct{
	mpkwh uint8
	kwh uint8
}
func (e electricEngine) milesLeft() uint8{
	return e.kwh*e.mpkwh
}

//function belongs to the object of gasEngine type
func (e gasEngine) milesLeft() uint8{
	return e.gallons*e.mpg
}

type engine interface{
	milesLeft() uint8
}
func canMakeIt(e engine,miles uint8){
	if miles<=e.milesLeft(){
		fmt.Println("you can make it ")
	}else{
		fmt.Println("need to fuel up first")
	}
}
func main(){
	var myEngine gasEngine=gasEngine{30,10,owner{"Alex"},10}
	myEngine.mpg=20
	var myEngine2=struct{
		mpg uint8
		gallons uint8
	}{25,15}
	fmt.Println(myEngine.mpg,myEngine.gallons,myEngine.name)
	fmt.Println(myEngine2.mpg,myEngine2.gallons)
	fmt.Printf("Total miles left are: %v\n",myEngine.milesLeft())
	canMakeIt(myEngine,50)

}