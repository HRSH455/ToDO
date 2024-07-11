package main
import ( "fmt"
convr "example/project/mypackage"
)

func main(){
	fmt.Println("Hello",convr.Name)
	intArr := []int{2,3,5,7,11}
	strArr := convr.IntToStr(intArr)
	fmt.Println(strArr)
}
