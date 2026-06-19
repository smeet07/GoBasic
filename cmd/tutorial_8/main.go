package main
import (
	"fmt"
	//"math/rand"
	"time"
	"sync"
)
var wg=sync.WaitGroup{}
var dbData=[]string{"id1","id2","id3","id4","id5"}
var result=[]string{}
var m=sync.RWMutex{}
func main(){
	t0:=time.Now()
	for i:=0;i<len(dbData);i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\n Total execution time is %v",time.Since(t0))
	fmt.Printf("\n The result are %v",result)
}
func dbCall(i int){
	var delay float32=2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	//fmt.Println("the result from the database is: ",dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}
func save(results string){
	m.Lock()
	result=append(result,results)
	m.Unlock()
}
func log(){
	m.RLock()
	fmt.Printf("\n the current results are: %v",result)
	m.RUnlock()
}