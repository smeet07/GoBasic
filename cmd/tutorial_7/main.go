package main
import (
	"fmt"
)
func main(){
	var p *int32=new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n",i)
	//*p=10
	p=&i
	*p=1
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n",i)

	//slice internally store pointers not the actual variables 
	var slice=[]int32{1,2,6}
	var sliceCopy=slice
	sliceCopy[1]=3
	fmt.Printf("The value of slice is %v\n",slice)
	fmt.Printf("The value of slicecopy is %v\n",sliceCopy)
	var thing1=[5]float64{1,2,3,4,5}
	var result[5]float64=square(&thing1)
	fmt.Println(result)
	fmt.Println(thing1)

}
func square(thing2 *[5]float64) [5]float64{
	for i:=range thing2{
		thing2[i]=thing2[i]*thing2[i]
	}
	return *thing2
}