package autotest

import ( "fmt"
        "regexp"
)

func isEmail(s string) (string,error){
	r,_ := regexp.Compile(`[\w._%+]`)
	if r.MatchString(s){
		return "valid" , nil
	}else {
		return "",fmt.Errorf("not valid")
	}
}