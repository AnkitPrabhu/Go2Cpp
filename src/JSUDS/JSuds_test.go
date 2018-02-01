package JSUDS

import "testing"
//import "fmt"

func BenchmarkSendingAndReceiving100(b *testing.B){
	J:=NewJSUDS()
        go J.Serve()
        <-J.Started
        for i:=0;i<b.N;i++{
                J.SendingAndReceiving(10)
        }
	J.Started<-0
}
