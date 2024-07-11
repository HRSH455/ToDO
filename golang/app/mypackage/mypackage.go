package convr
import(
	
	"strconv"
	
)

var Name  string = "hb"
func IntToStr(intArr[]int) []string{
	var strArr []string
	for _,i := range intArr{
		strArr = append(strArr,strconv.Itoa(i))

	}
	return strArr
}