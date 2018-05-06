package main

import(
	"../engine"
	"net"
	"fmt"
	"os"
	"runtime"
	"encoding/json"
	"errors"
)
var DebuugerMap map[string] * Debugger
func main(){
	NetListener,err := net.Listen("tcp","127.0.0.1:5555")
	if err != nil{
		fmt.Print(err)
		os.Exit(0)
	}
	DebuugerMap = make(map[string] * Debugger)
	runtime.GOMAXPROCS(runtime.NumCPU())
	for {
		conn, err := NetListener.Accept()
		if err != nil {
			fmt.Print("A wrong client")
			continue //No client
		}
		go WaitingForMessage(conn)

	}
}
func WaitingForMessage(conn net.Conn){
	buffer := make([]byte, 2048)
	n,err := conn.Read(buffer)
	if err != nil {
		Log(conn.RemoteAddr().String(), " connection error: ", err)
		return
	}
	data := string(buffer[:n])		//创建一个变量来输出缓冲区
	var JsonData map[string]interface{}
	json.Unmarshal(buffer,&JsonData)
	//解析JSON
	if !check_map(JsonData,"id") || !check_map(JsonData,"request"){	//操作识别号不存在。错误请求。断开连接
		conn.Close()
		return
	}
	id := JsonData["id"].(int)
	RequestId := JsonData["request"].(string)
	switch id{
	case 100:
		//新建一个Debugger
		name := JsonData["debugger_name"].(string)
		err := NewDebugger(name)
		if err != nil{
			//生成出错
			conn.Write([]byte(`{
				"request": "` + RequestId + `",
				"if" : "0",
				"reason" : "` + err.Error() + ` "
			}`))
		}
		conn.Write([]byte(`{
			"request": "` + RequestId + `",
			"if" : "1"
		}`))
		break
	case 101:{
		//导入JS脚本
		
	}
	default:
		conn.Close()
		return
	}

	//解析JSON

}
func Log(data ...interface{}){
	fmt.Print(data)
}
func check_map(m map[string]interface{},key string)bool{
	_,err := m[key]
	return err
}
type Debugger struct{
	engine engine.RJSEngine
	has bool
}

func NewDebugger(name string)(error){
	if DebuugerMap[name].has{
		err := errors.New(name + " has been on")
		return err
	}
	//temp := Debugger{}
	//temp.engine.Init()
	DebuugerMap[name].has = true
	DebuugerMap[name].engine.Init()
	return nil
}