package JSUDS

import (
        //"time"
        "encoding/json"
        "encoding/binary"
        "bufio"
        "bytes"
        "net"
        "os/exec"
        //"strings"
	//"fmt"
        )

type JSUDS struct{
	OutPipe *bufio.Reader
        Conn net.Conn
	Started chan int
}

func NewJSUDS() *JSUDS{
        return &JSUDS{Started: make(chan int)}
}

func (J *JSUDS)Serve(){
        lis,err := net.Listen("unix","/tmp/go24.sock")
        if err!=nil{
        	//fmt.Println(err)
	}
        cmd := exec.Command("../../uds","/tmp/go24.sock")
        go J.handleRequest(lis)
        err=cmd.Start()
        err=cmd.Wait()
}

func (J *JSUDS)handleRequest(lis net.Listener){
        J.Conn,_=lis.Accept()
        J.Started<-10
	<-J.Started
	lis.Close()
}

func (J *JSUDS)SendingAndReceiving(x int){
         var buffer bytes.Buffer
         opcode,_:=json.Marshal(x)
         binary.Write(&buffer, binary.LittleEndian, opcode)
         J.Conn.Write(buffer.Bytes())
	
         tmp := make([]byte, 256)
         _, _= J.Conn.Read(tmp)
	//fmt.Println(string(tmp))
}
