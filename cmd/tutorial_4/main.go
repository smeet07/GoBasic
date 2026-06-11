package main
import "fmt"
func main(){
	var intArr [3]int32
	intArr[1]=132
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
	arr2:=[...]int32{1,2,3}
	fmt.Println(arr2[0:3])
	//slices
	var intSlice []int32=[]int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v \n",len(intSlice),cap(intSlice))
	intSlice=append(intSlice,7)
	fmt.Printf("The length is %v with capacity %v\n",len(intSlice),cap(intSlice))
	var intSlice2 []int32=[]int32{8,9,10}
	intSlice=append(intSlice,intSlice2...)
	fmt.Println(intSlice)
	fmt.Printf("\nThe length is %v with capacity %v\n",len(intSlice),cap(intSlice))
	var intSlice3 []int32=make([]int32,3,8)
	fmt.Println(intSlice3)

	//maps in go
	//var myMap map[string]uint8=make(map[string]uint8)
	var myMap2 =map[string]uint8{"Adam":23,"Sarah":45,"Bobi":46}
	fmt.Println(myMap2["Adam"])
	//if the key doesnt exist then it returns the default value of the value type in this case 0
	fmt.Println(myMap2["Jason"])
	var age,ok=myMap2["Jason"]
	if ok==false{
		fmt.Println("Key doesn't exist")
	}else{
		fmt.Println(age)
	}
	//delets by reference
	delete(myMap2,"Adam")

	//traversing the map
	for name,age:=range myMap2{
		fmt.Printf("Name: %v, Age: %v \n",name,age)
	}

	for i,v:= range intArr{
		fmt.Printf("The element at index %v is %v \n",i,v)
	}
	//while loops
	var i int=0
	for i<10{
		fmt.Println(i)
		i=i+1
	}
	//can also do this way
	i=0
	for {
		if i>=10{
			break
		}
		fmt.Println(i)
		i=i+1
	}
	for i=0;i<10;i++{
		fmt.Println(i)
	}

	


}