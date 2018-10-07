package pkg_network

import(
	"fmt"
	"net"
	//"math/rand"
	"github.com/Dorbmon/otto"
	"github.com/Dorbmon/R.JS/go_pkg/pkg_stack"
	"net/url"
	"net/http"
	//"io/ioutil"
	//"log"
	"os"
	"strings"
)
import(
	//"./pkg_http"
)
//var TCP_LISTENER_MAP = make(map[int]TCP_LISTENER)
//var TCP_LISTENER_MAP map[int]TCP_LISTENER
type AJAX_LIST struct{
	url string
	value url.Values
	has bool
	OnErrorCallBackFuntion otto.Value
	OnMessageCallBackFunction otto.Value
	mode string
	hr *http.Request
	request http.Request
}
var AJAX_OBJECT_LIST map[string]*AJAX_LIST
var js *otto.Otto
var listener_number = 0	//tcp
var Ajax_listener_number = 0
var listen_list pkg_stack.Stack
var error_ = func(err error){
	fmt.Print(err)
}
const (
	MAX_TCP_LISTENER = 100
	MAX_AJAX = 100
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
		listener_number ++
		if Ajax_listener_number > MAX_AJAX{
			//超出最大AJAX数
			fmt.Println("ERROR,Ajax_listen_number > MAX_AJAX")
			Ajax_listener_number--
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
		urlArgument,err := call.Argument(0).ToString()
		if err != nil{
			//类型错误
			fmt.Print("The type of the url of NETWORK_AJAX is wrong on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		Mode,err := call.Argument(0).ToString()
		if err != nil{
			//类型错误
			fmt.Print("The type of the url of NETWORK_AJAX is wrong on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		if AJAX_OBJECT_LIST[obj_name].has{	//已经存在该名称
			return otto.FalseValue()
		}
		AJAX_OBJECT_LIST[obj_name].mode = Mode
		//AJAX_OBJECT_LIST[obj_name].url,err = url.Parse(urlArgument)
		AJAX_OBJECT_LIST[obj_name].hr.ParseForm()
		if err != nil{
			fmt.Print("ERROR URL on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		//AJAX_OBJECT_LIST[obj_name].hr = http.NewRequest(Mode,urlArgument,)
		//AJAX_OBJECT_LIST[obj_name].value = AJAX_OBJECT_LIST[obj_name].url.Query()
		AJAX_OBJECT_LIST[obj_name].url = urlArgument
		value,_ := otto.ToValue(obj_name)
		return value
	})
	js.Set("NET_AJAX_ON_ERROR_FUNCTION",func(CallBackFunction otto.Value,call otto.FunctionCall)otto.Value{//参数2为AJAX名称。参数1为函数 非名称
		if call.Argument(1).IsNull(){
			fmt.Print("Data is not enough for NETWORK_AJAX_SET on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		ajaxName := call.Argument(0).String()
		if !AJAX_OBJECT_LIST[ajaxName].has{
			fmt.Print("ERROR ajaxName:\"" + ajaxName +"\" on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		OnErrorCallBackFunction := CallBackFunction
		if !OnErrorCallBackFunction.IsFunction(){
			fmt.Print(OnErrorCallBackFunction," is not a Function on:",call.CallerLocation()," Please Check it.")
			OnStrictMode(js)
			return otto.FalseValue()
		}
		AJAX_OBJECT_LIST[ajaxName].OnErrorCallBackFuntion = CallBackFunction
		return otto.TrueValue()
	})
	js.Set("NET_AJAX_ON_MESSAGE_FUNCTION",func(CallBackFunction otto.Value,call otto.FunctionCall)otto.Value{//参数2为AJAX名称。参数1为函数 非名称
		if call.Argument(1).IsNull(){
			fmt.Print("Data is not enough for NETWORK_AJAX_SET on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		ajaxName := call.Argument(0).String()
		if !AJAX_OBJECT_LIST[ajaxName].has{
			fmt.Print("ERROR ajaxName:\"" + ajaxName +"\" on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		OnErrorCallBackFunction := CallBackFunction
		if !OnErrorCallBackFunction.IsFunction(){
			fmt.Print(OnErrorCallBackFunction," is not a Function on:",call.CallerLocation()," Please Check it.")
			OnStrictMode(js)
			return otto.FalseValue()
		}
		AJAX_OBJECT_LIST[ajaxName].OnMessageCallBackFunction = CallBackFunction
		return otto.TrueValue()
	})
	js.Set("NETWORK_AJAX_SET",func(call otto.FunctionCall)otto.Value{//参数1为AJAX名称。参数2为参数名称，参数3为参数值
		if call.Argument(2).IsNull(){
			fmt.Print("Data is not enough for NETWORK_AJAX_SET on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		ajaxName := call.Argument(0).String()
		if !AJAX_OBJECT_LIST[ajaxName].has{
			fmt.Print("ERROR ajaxName:\"" + ajaxName +"\" on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		name := call.Argument(1).String()
		value := call.Argument(2).String()
		//AJAX_OBJECT_LIST[ajaxName].value.Set(name,value)
		//AJAX_OBJECT_LIST[ajaxName].hr.Form.Add(name,value)
		AJAX_OBJECT_LIST[ajaxName].hr.Form.Set(name,value)
		return otto.TrueValue()
	})
	js.Set("NETWORK_AJAX_SET_REQUEST_HEADER",func(call otto.FunctionCall)otto.Value{
		if call.Argument(3).IsNull(){
			fmt.Print("Data is not enough for NETWORK_AJAX_SET_REQUEST_HEADER on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		ajaxName := call.Argument(0).String()
		if !AJAX_OBJECT_LIST[ajaxName].has{
			fmt.Print("ERROR ajaxName:\"" + ajaxName +"\" on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		key := call.Argument(1).String()
		value := call.Argument(2).String()
		AJAX_OBJECT_LIST[ajaxName].hr.Header.Set(key,value)
		return otto.TrueValue()
	})
	js.Set("NETWORK_AJAX_SEND",func(call otto.FunctionCall)otto.Value{	//参数1为AJAX名称，参数2为是否为异步
		if call.Argument(1).IsNull(){
			fmt.Print("Data is not enough for NETWORK_AJAX_SEND on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		ajaxName := call.Argument(0).String()
		ifAsynchronous,err := call.Argument(1).ToBoolean()
		if err != nil{
			fmt.Print("ERROR TYPE OF asynchronous for NETWORK_AJAX_SEND on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		if !AJAX_OBJECT_LIST[ajaxName].has{
			fmt.Print("ERROR ajaxName:\"" + ajaxName +"\" on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		//解析信息
		//AJAX_OBJECT_LIST[ajaxName].url.RawQuery = AJAX_OBJECT_LIST[ajaxName].value.Encode()
		AJAX_OBJECT_LIST[ajaxName].hr,err = http.NewRequest(AJAX_OBJECT_LIST[ajaxName].mode,AJAX_OBJECT_LIST[ajaxName].url,strings.NewReader(strings.TrimSpace(AJAX_OBJECT_LIST[ajaxName].hr.Form.Encode())))
		//res,err := http.Get(AJAX_OBJECT_LIST[ajaxName].url.String())
		if err != nil{	//创建出错
			AJAX_OBJECT_LIST[ajaxName].OnErrorCallBackFuntion.Call(AJAX_OBJECT_LIST[ajaxName].OnErrorCallBackFuntion,err)
		}
		if ifAsynchronous{
			//异步
			go func(engine *otto.Otto){
				//AJAX_OBJECT_LIST[ajaxName].url.
				resp,err := http.DefaultClient.Do(AJAX_OBJECT_LIST[ajaxName].hr)
				//_,err := http.Get(AJAX_OBJECT_LIST[ajaxName].url.String())
				if err != nil{
					//请求出错
					AJAX_OBJECT_LIST[ajaxName].OnErrorCallBackFuntion.Call(AJAX_OBJECT_LIST[ajaxName].OnErrorCallBackFuntion,err)
				}
				//请求完成
				//回调成功函数
				Code := resp.StatusCode
				AJAX_OBJECT_LIST[ajaxName].OnMessageCallBackFunction.Call(AJAX_OBJECT_LIST[ajaxName].OnMessageCallBackFunction,Code,resp.Body)

			}(js)
		}else{	//同步
			resp,err := http.DefaultClient.Do(AJAX_OBJECT_LIST[ajaxName].hr)
			//_,err := http.Get(AJAX_OBJECT_LIST[ajaxName].url.String())
			if err != nil{
				//请求出错
				AJAX_OBJECT_LIST[ajaxName].OnErrorCallBackFuntion.Call(AJAX_OBJECT_LIST[ajaxName].OnErrorCallBackFuntion,err)
				OnStrictMode(js)
				return otto.FalseValue()
			}
			//请求完成
			Code := resp.StatusCode
			AJAX_OBJECT_LIST[ajaxName].OnMessageCallBackFunction.Call(AJAX_OBJECT_LIST[ajaxName].OnMessageCallBackFunction,Code,resp.Body)
		}
		return otto.TrueValue()
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