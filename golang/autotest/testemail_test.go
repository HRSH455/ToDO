package autotest

import("testing")

func testisemail(t *testing.T){
	-,err := IsEmail("yagoo")
	if err == nil{
		t.Error("it is not a email")
	}
}
