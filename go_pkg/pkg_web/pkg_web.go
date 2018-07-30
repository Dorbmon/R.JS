package pkg_web

import(
	"github.com/Dorbmon/otto"
	"fmt"
	"os"
	"net/http"
	"net/url"
)

type RjsWebMoudle struct{
	js *otto.Otto
}
func (this RjsWebMoudle)Init(js *otto.Otto){
	this.js = js
	js.Set("WebServer",func()*otto.Object{
		obj,err := js.Object("({})")
		if err != nil{
			fmt.Println(err)
			os.Exit(0)
		}
		obj.Set("Bind",func(call otto.FunctionCall)otto.Value{
			//设置回调函数	绑定
			if call.Argument(1).IsUndefined(){
				fmt.Println("ERROR.When call WebServer.Bind(). on:",call.CallerLocation())
				OnStrictMode(js)
			}
			BindAddress := call.Argument(0).String()
			BindFunction := call.Argument(1)
			//fmt.Println("Address:",BindAddress)
			//defer fmt.Println("Address:",BindAddress," Has been defined.")
			http.HandleFunc(BindAddress,func(w http.ResponseWriter,r *http.Request){
				//fmt.Println("sss")
				Writer,_ := js.Object("({})")
				queryGetForm, err := url.ParseQuery(r.URL.RawQuery)	//只有Get
				if err == nil && len(queryGetForm["id"]) > 0 {
					fmt.Fprintln(w, queryGetForm["id"][0])
				}
				Writer.Set("Write",func(call otto.FunctionCall)otto.Value{
					data := []byte(call.Argument(0).String())
					i,err := w.Write(data)
					if err != nil{
						result,_ := otto.ToValue(err)
						return result
					}
					result,_ := otto.ToValue(i)
					return result
				})
				Request,_ := js.Object("({})")
				Request.Set("GetForm",func(call otto.FunctionCall)otto.Value{
					if call.Argument(0).IsUndefined(){
						fmt.Println("ERROR.When call Request.GetForm(). on:",call.CallerLocation())
						OnStrictMode(js)
					}
					key := call.Argument(0).String()
					value,err := otto.ToValue(r.Form.Get(key))
					if err != nil{
						fmt.Println(err)
						return otto.FalseValue()
					}
					//fmt.Println(r.URL.String())
					fmt.Println(r.Form.Encode())

					return value
				})
				Request.Set("RequestUrl",r.URL.Path)
				Request.Set("RemoteAddr",r.RemoteAddr)
				Request.Set("Host",r.Host)
				Request.Set("GetPostForm",func(call otto.FunctionCall)otto.Value{
					if call.Argument(0).IsUndefined(){
						fmt.Println("ERROR.When call Request.GetForm(). on:",call.CallerLocation())
						OnStrictMode(js)
					}
					key := call.Argument(0).String()
					//value,_ := otto.ToValue(r.PostForm.Get(key))
					valuet := queryGetForm.Get(key)
					value,_ := otto.ToValue(valuet)
					return value
				})
				BindFunction.Call(BindFunction,Writer,Request)
			})
			return otto.TrueValue()
		})
		obj.Set("Listen",func(call otto.FunctionCall)otto.Value{
			addr := call.Argument(0).String()
			err := http.ListenAndServe(addr,nil)
			result,_ := otto.ToValue(err)
			return result
		})
		return obj
	})
}
func OnStrictMode(this *otto.Otto){
	IfStrictMode,err := this.Get("RJS_CONFIG_STRIT_MODE")
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