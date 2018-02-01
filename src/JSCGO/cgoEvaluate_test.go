package JSCGO

import "testing"

func BenchmarkAddCGO1000(t *testing.B){
	for i:=0;i<t.N;i++{
		AddCGO(i)
	}
	
}

