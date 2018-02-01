package JSCGO

/*
int add(int x){
	return x+10;
}
*/
import "C"
//import "fmt"

func AddCGO(x int){
	a:=C.int(x)
	_=C.add(a)
	//fmt.Println(f)
}

