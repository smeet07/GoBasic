package main
import "fmt"
import "time"
import "math/rand"
// func main(){
// 	var c=make(chan int)
// 	go process(c)
// 	for i:=range c{
// 		fmt.Println(i)
// 	}

// }
// func process(c chan int){
// 	defer close(c)
// 	for i:=0;i<5;i++{
// 		c<-i
// 	}

// }
// func main(){
// 	var c=make(chan int,5)
// 	go process(c)
// 	for i:=range c{
// 		fmt.Println(i)
// 		time.Sleep(time.Second*1)
// 	}

// }
// func process(c chan int){
// 	defer close(c)
// 	for i:=0;i<5;i++{
// 		c<-i
// 	}
// 	fmt.Println("Exiting process")

// }

var MAX_PRICE float32=5
func main(){
	var pizzaChannel=make(chan string)
	var websites=[]string{"walmart.com","costco.com","wholefoods.com"}
	for i:=range websites{
		go checkpizzaPrices(websites[i],pizzaChannel)
	}
	sendMessage(pizzaChannel)
}
func checkpizzaPrices(website string,pizzaChannel chan string){
	for{
		time.Sleep(time.Second*1)
		var pizzaPrice=rand.Float32()*20
		if pizzaPrice<=MAX_PRICE{
			pizzaChannel<-website
			break
		}
	}
}
func sendMessage(pizzaChannel chan string){
	fmt.Printf("\n Found a deal on pizza at %s",<-pizzaChannel)
}
