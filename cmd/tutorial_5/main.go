package main
import "fmt"
import "strings"
func main(){
	var myString string="résumé"
	fmt.Println(myString)
	var indexed=myString[0]
	fmt.Printf("%v, %T\n",indexed,indexed)
	for i,v:=range myString{
		fmt.Println(i,v)
	}
	//This is because of utf-8 encoding. each character is represented in bytes of variable lengths
	//which was the length prints the number of bytes
	//easier way to avoid this to use runes
	var myString2=[]rune("résumé")
	var idx2=myString2[1]
	fmt.Printf("%v %T \n",idx2,idx2)
	for i,v:=range myString{
		fmt.Println(i,v)
	}
	var strSlice=[]string{"s","u","b"}
	var strBuilder strings.Builder
	for i:=range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr=strBuilder.String()
	fmt.Printf("\n%v",catStr)


}