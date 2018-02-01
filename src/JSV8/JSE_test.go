package JSV8

import "testing"
//import "fmt"

func BenchmarkSendingAndReceiving100000(b *testing.B){
	J:=NewJSV8()
	go J.Serve()
        <-J.Started
	for i:=0;i<b.N;i++{
        	J.SendingAndReceiving(10)
	}
}
