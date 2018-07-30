package pkg_http

//HTTP子包
//瑞雪Dorbmon
import(
	"github.com/Dorbmon/otto"
	"net/http"
	"fmt"
	"os"
)
type Pkg_http struct{
	//HTTP包
	engine *otto.Otto
	BandFunctions map[string] *bandFunctions
}
type bandFunctions struct{
	ServerMux *http.ServeMux
	Server http.Server
	BandFunctions map[string] func(w http.ResponseWriter,r *http.Request)
	has bool
}
func (this *Pkg_http)Init(js *otto.Otto){
	this.engine = js
	//交换Engine
	//初始化内存
	this.BandFunctions = make(map[string]*bandFunctions)
	js.Set("NETWORK_HTTP_NEW_SERVER",func(call otto.FunctionCall)otto.Value{	//第一个参数为实例名称，第二个参数为访问的地址，第三个参数为回调函数()
	/*
		回调函数格式
		只有一个参数
		message{
			url 请求的地址
			method 请求方式 GET 或 POST
			get function 获取指定key的value
			free 释放自身
		}
	 */
		if call.Argument(2).IsNull(){
			fmt.Print("Data is not enough for NETWORK_AJAX on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		HttpName,err := call.Argument(0).ToString()
		if err != nil{
			//出错
			OnStrictMode(this.engine)
			return otto.FalseValue()
		}
		if this.BandFunctions[HttpName].has{
			OnStrictMode(this.engine)
			return otto.FalseValue()
		}
		//创建新的监听
		Address,err := call.Argument(1).ToString()
		if err != nil{
			//出错
			OnStrictMode(this.engine)
			return otto.FalseValue()
		}
		CallBackFunction := call.Argument(2)
		if !CallBackFunction.IsFunction(){
			fmt.Print(CallBackFunction.String(),"is not a Function On:",call.CallerLocation())
			OnStrictMode(this.engine)
			return otto.FalseValue()
		}
		this.BandFunctions[HttpName].BandFunctions[Address] = func(w http.ResponseWriter,r *http.Request){
			//调用回调函数，传递数据
			obj,_ := this.engine.Object("CallBack_" + HttpName)
			obj.Set("url",r.URL.Path)
			obj.Set("method",r.Method)
			obj.Set("get",func(call otto.FunctionCall)otto.Value{
				//获取指定字段
				key,err := call.Argument(0).ToString()
				if err != nil{
					return otto.FalseValue()
				}
				value,err := otto.ToValue(r.Form.Get(key))
				if err != nil{
					OnStrictMode(this.engine)
					fmt.Println(err," on:",call.CallerLocation())
					return otto.FalseValue()
				}
				return value
			})
			obj.Set("write",func(call otto.FunctionCall)otto.Value{
				if call.Argument(0).IsNull(){
					fmt.Print("Data is not enough for NETWORK_AJAX on:",call.CallerLocation())
					OnStrictMode(js)
					return otto.FalseValue()
				}
				data := call.Argument(0).String()
				_,errin := w.Write([]byte(data))
				if errin != nil{
					return otto.FalseValue()
				}
				return otto.TrueValue()
			})
			obj.Set("free",func(){
				//释放自身
				this.engine.Set("CallBack_" + HttpName,"")
			})
			CallBackFunction.Call(CallBackFunction,obj)
		}
		//http.HandleFunc(Address,this.BandFunctions[HttpName].BandFunctions[Address])
		this.BandFunctions[HttpName].ServerMux.HandleFunc(Address,this.BandFunctions[HttpName].BandFunctions[Address])
		this.BandFunctions[HttpName].Server.Handler = this.BandFunctions[HttpName].ServerMux
		this.BandFunctions[HttpName].has = true

		//err := this.BandFunctions[HttpName].Server.ListenAndServe()
		return otto.TrueValue()
	})
	js.Set("NETWORK_HTTP_SET_CALLBACK_FUNCTION",func(call otto.FunctionCall)otto.Value{
		return otto.TrueValue()
	})
	js.Set("NETWORK_HTTP_LISTEN",func(call otto.FunctionCall)otto.Value{	//第一个参数为HTTP名称，第二个参数为端口
		HttpName,err := call.Argument(0).ToString()
		if err != nil{
			fmt.Print("It is not a string on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		if !this.BandFunctions[HttpName].has{
			return otto.FalseValue()
		}
		Address,err := call.Argument(1).ToString()
		if err != nil{
			fmt.Print("It is not a string on:",call.CallerLocation())
			OnStrictMode(js)
			return otto.FalseValue()
		}
		this.BandFunctions[HttpName].Server.Addr = Address
		go this.BandFunctions[HttpName].Server.ListenAndServe()
		return otto.TrueValue()
	})
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