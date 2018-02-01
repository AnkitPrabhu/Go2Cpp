package JSV8

import (
	//"time"
	"encoding/json"
	//"encoding/binary"
	"bufio"
	//"bytes"
	"net"
	"os/exec"
	"strings"
	"fmt"
	)

type JSV8 struct{
	Conn net.Conn
	connReader *bufio.Reader
	Started chan interface{}
	outp *bufio.Reader
	tmp []byte
}
func NewJSV8() *JSV8{
	return &JSV8{Started: make(chan interface{}),
			tmp:make([]byte,256)}
}
func (J *JSV8)Serve(){
	lis,err := net.Listen("tcp","localhost:0")
        if err!=nil{
        }
        port:=strings.Split(fmt.Sprintf("%v",lis.Addr()),":")[1]
        cmd := exec.Command("../../simple",port)
        go J.handleRequest(lis)
        cmd.Start()
        err=cmd.Wait()
}

func (J *JSV8)handleRequest(lis net.Listener){
	J.Conn,_=lis.Accept()
	J.connReader=bufio.NewReader(J.Conn)
	J.Started<-10 
}

func (J *JSV8)SendingAndReceiving(x int){
    var buffer bytes.Buffer
    opcode,_:=json.Marshal(x)
    binary.Write(&buffer, binary.LittleEndian, opcode)
    binary.Write(J.Conn, binary.LittleEndian, buffer.Bytes())
    tmp := make([]byte, 256)
 	_, _ = J.Conn.Read(tmp)
	//binary.Read(J.Conn, binary.LittleEndian, tmp)
}
