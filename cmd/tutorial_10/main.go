package main
import "fmt"

// func main(){
// 	var intSlice=[]int{1,2,3}
// 	fmt.Println(sumSlice[int](intSlice))
// 	var float32Slice=[]float32{1,2,3}
// 	fmt.Println(sumSlice[float32](float32Slice))
// }
// func sumSlice[T int | float32 | float64](slice []T)T{
// 	var sum T
// 	for _,v:= range slice{
// 		sum+=v
// 	}
// 	return sum
// }

type gasEngine struct{
	gallons float32
	mpg float32
}

type electricEngine struct{
	kwh float32
	mpkwh float32
}
type car [T gasEngine | electricEngine] struct{
	carMake string
	carModel string
	engine T
}
func main(){
	var gasCar=car[gasEngine]{
		carMake:"Honda",
		carModel:"Civic",
		engine:gasEngine{
			gallons:12.4,
			mpg:40,
		},
	}
	fmt.Println(gasCar)
}