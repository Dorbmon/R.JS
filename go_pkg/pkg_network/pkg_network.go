package pkg_network

import(
	"fmt"
	"net"
	//"math/rand"
	otto "github.com/robertkrimen/otto"
	"../pkg_stack"
	"net/url"
	"net/http"
	"io/ioutil"
	"log"
	"os"
)
//var TCP_LISTENER_MAP = make(map[int]TCP_LISTENER)
//var TCP_LISTENER_MAP map[int]TCP_LISTENER
type AJAX_LIST struct{
	url *url.URL
	value url.Values
	has bool
}
var AJAX_OBJECT_LIST map[string]*AJAX_LIST
var js *otto.Otto
var listener_number = 0
var listen_list pkg_stack.Stack
var error_ = func(err error){
	fmt.Print(err)
}
const (
	MAX_TCP_LISTENER = 100
)
func Swap_Data_From_Main(js_engine *otto.Otto){
	AJAX_OBJECT_LIST = make(map[string]*AJAX_LIST)
	//url,err := url.Parse("xx")
	js = js_engine
	//TCP_LISTENER_MAP = make(map[int]TCP_LISTENER)
	//error_ = error_func.(func())
	//初始化队列
	listen_list = pkg_stack.New()
	js.Set("NETWORK_AJAX",func(call otto.FunctionCall)otto.Value{	//参数1为对方地址，二为协议 GET或POST。必须大写。返回该对象名称，参数3为唯一名称
		//新建一个AJAX实例
		if call.Argument(2).IsNull(){
			fmt.Print("Data is not enough for NETWORK_AJAX on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		obj_name,err := call.Argument(2).ToString()
		if err != nil{
			//类型错误
			fmt.Print("The type of the name of NETWORK_AJAX is wrong on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		if AJAX_OBJECT_LIST[obj_name].has{	//已经存在该名称
			return otto.FalseValue()
		}
		urlArgument,err := call.Argument(0).ToString()
		if err != nil{
			//类型错误
			fmt.Print("The type of the url of NETWORK_AJAX is wrong on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		AJAX_OBJECT_LIST[obj_name].url,err = url.Parse(urlArgument)
		if err != nil{
			fmt.Print("ERROR URL on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		value,_ := otto.ToValue(obj_name)
		return value
	})
}
type TCP_LISTENER struct {
	listener *net.Listener
	On_Data_func string
	On_Connect_func string
	On_DisConnect_func string
	On_ERROR string
	has bool
	on bool	//是否开启 如果开启则继续消息循环 如果关闭了消息循环自动调整has
}
type TCP_CONN struct{
	conn net.Conn
	has bool
	on bool	//此为信号量 通知消息循环结束
}
/*		error code		*/
//	001 创建监听失败
//	0011 创建失败 达到监听上限
//	002 客户连接失败
//
func Start_Listen(IP string,Port int,CallBackData TCP_LISTENER){
	if listener_number == MAX_TCP_LISTENER{
		//达到监听上限
		js.Call(CallBackData.On_ERROR,0011)
		return
	}
	var err error
	//TCP_LISTENER_MAP[listener_num].On_Connect_func = ""
	temp_tcp_listener := TCP_LISTENER{}
	temp_tcp_listener = CallBackData
	*(temp_tcp_listener.listener),err = net.Listen("tcp",IP + ":" + string(Port))
	temp_listener_number := listen_list.Push(temp_tcp_listener)
	listener_number++
	if err != nil{
		error_(err)
		//return err
	}
	//设置监听相关
	//*TCP_LISTENER_MAP[listener_num].
	go func(listener_num int){
		temp,_ := listen_list.Get(listener_num)
		listener := temp.(TCP_LISTENER)
		for
		{
			//conn, err := (*TCP_LISTENER_MAP[listener_num].listener).Accept()
			conn,err := (*listener.listener).Accept()
			if err != nil {
				error_(err)
				//投递错误信息
				//js.Call(TCP_LISTENER_MAP[listener_num].On_ERROR,err)
				js.Call(listener.On_ERROR,err)
				continue
			}
			//投递连接
			js.Call(listener.On_Connect_func,listener_num)	//投递虚ID
			//针对该客户进行消息循环
			go func(conn net.Conn,listener *TCP_LISTENER){
				for {
					defer func() {
						conn.Close()
						return
					}()
					buffer := make([]byte, 1024)
					n, err := conn.Read(buffer)
					if err != nil {
						js.Call(listener.On_ERROR,err)
					}
					data := string(buffer[:n]) //创建一个变量来输出缓冲区
					//投递数据
					js.Call(listener.On_Data_func, data)
				}
			}(conn,&listener)
			//

		}
	}(temp_listener_number)
}
func OnStrictMode(vm *otto.Otto){
	IfStrictMode,err := vm.Get("RJS_CONFIG_STRIT_MODE")
	if err != nil{
		fmt.Print("ERRO CONFIG.RJS_CONFIG_STRIT_MODE")
		os.Exit(0)
	}
	confident,err :=  IfStrictMode.ToBoolean()
	if err != nil{
		fmt.Print("ERRO CONFIG.RJS_CONFIG_STRIT_MODE")
		os.Exit(0)
	}
	if confident{
		os.Exit(0)
	}
	return
}