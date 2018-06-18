package pkg_web

import(
	"github.com/robertkrimen/otto"
	"fmt"
	"os"
	"net/http"
)

type RjsWebMoudle struct{
	js *otto.Otto
}
func (this RjsWebMoudle)Init(js *otto.Otto){
	this.js = js
	js.Set("WebServer",func()*otto.Object{
		obj,err := js.Object("{}")
		if err != nil{
			fmt.Println(err)
			os.Exit(0)
		}
		obj.Set("Band",func(call otto.FunctionCall)otto.Value{
			//设置回调函数	绑定
			if call.Argument(1).IsUndefined(){
				fmt.Println("ERROR.When call WebServer.Band(). on:",call.CallerLocation())
				OnStrictMode(js)
			}
			BandAddress := call.Argument(0).String()
			BandFunction := call.Argument(1)
			http.HandleFunc(BandAddress,func(w http.ResponseWriter,r *http.Request){
				Writer,_ := js.Object("Writer")
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
				Request,_ := js.Object("Request")
				Request.Set("GetForm",func(call otto.FunctionCall)otto.Value{
					if call.Argument(0).IsUndefined(){
						fmt.Println("ERROR.When call Request.GetForm(). on:",call.CallerLocation())
						OnStrictMode(js)
					}
					key := call.Argument(0).String()
					value,_ := otto.ToValue(r.Form.Get(key))
					return value
				})
				Request.Set("RequestUrl",r.URL.Path)
				Request.Set("GetPostForm",func(call otto.FunctionCall)otto.Value{
					if call.Argument(0).IsUndefined(){
						fmt.Println("ERROR.When call Request.GetForm(). on:",call.CallerLocation())
						OnStrictMode(js)
					}
					key := call.Argument(0).String()
					value,_ := otto.ToValue(r.PostForm.Get(key))
					return value
				})
				BandFunction.Call(BandFunction,Writer)
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