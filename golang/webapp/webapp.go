package main

import(
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type toDoL struct{
	count int
	todos []string
}
func errorCheck(err error){
	if err !=nil{
		log.Fatal(err)
	}
}
func write(writer http.ResponseWriter, msg string){
	_,err :=writer.Write([]byte(msg))
	errorCheck(err)
}
func getString(fileName string) []string{
	var lines []string
	file,err :=os.Open(fileName)
	if os.IsNotExist(err){
		return nil
	}
	errorCheck(err)
	defer file.Close()
	scanner :=bufio.NewScanner(file)
	for scanner.Scan(){
		lines =append(lines,scanner.Text())
	}
	errorCheck(scanner.Err())
	return lines
}
func engHand(writer http.ResponseWriter, request *http.Request){
	write(writer,"HELLO")
}
func intHand(writer http.ResponseWriter, request *http.Request){
	todoVal := getString("to_do.txt")
	fmt.Printf("%#v\n",todoVal)
	tmp,err :=template.ParseFiles("view.html")
	errorCheck(err)
	todos1 :=toDoL{
		count : len(todoVal),todos : todoVal,

	}
	err =tmp.Execute(writer,todos1)

}
func newHand(writer http.ResponseWriter, request *http.Request){
	tmp,err :=template.ParseFiles("new.html")
	errorCheck(err)
	err =tmp.Execute(writer,nil)
}
func createHand(writer http.ResponseWriter, request *http.Request){
	todo :=request.FormValue("todo")
	options :=os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file , err := os.OpenFile("to_do.txt",options,os.FileMode(0600))
	errorCheck(err)
	_,err =fmt.Fprintln(file ,todo)
	errorCheck(err)
	err =file.Close()
	errorCheck(err)
	http.Redirect(writer ,request,"/interact",http.StatusFound)

}

func main(){
	http.HandleFunc("/hello",engHand)
	http.HandleFunc("/interact",intHand)
	http.HandleFunc("/new" ,newHand)
	http.HandleFunc("/create" ,createHand)
	err :=http.ListenAndServe("localhost:8080",nil)
	log.Fatal(err)
}