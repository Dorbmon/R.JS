package main

import "net/http"
import (
	"github.com/Dorbmon/R.JS/engine"
	"github.com/Dorbmon/otto"
	"math/rand"
	"time"
	//"encoding/base64"
	"fmt"
	"encoding/base64"
)
func main(){
	//创建服务器
	http.HandleFunc("/",Run)
	http.ListenAndServe("0.0.0.0:8887",nil)
}
func Run(w http.ResponseWriter, r *http.Request){
	//获得程序
	w.Header().Set("Access-Control-Allow-Origin", "*")//允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers","Content-Type")//header的类型
	w.Header().Set("content-type","application/json")//返回数据格式是json
	//code,err := base64Decode([]byte(r.Form.Get("code")))
	err := r.ParseForm()
	if err != nil{
		fmt.Println(err)
		return
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(r.Form.Get("code"))
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}
	code := string(decodeBytes)
	//code := r.Form.Get("code")
	//code := r.PostForm.Get("code")
	fmt.Println(code)
	//创建RJS引擎
	var engine engine.RJSEngine
	engine.Init()
	resultString := ""
	engine.Js.Set("output", func(call otto.FunctionCall) otto.Value {
		//js.Call(call.Argument(0).String(),"")
		n := 0
		for {
			value := call.Argument(n)
			if !value.IsDefined() {
				break
			}
			n++
			resultString += value.String()
		}
		//return otto.Value{temp,"output"}
		return otto.Value{}
		//return otto.TrueValue()
	})
	engine.SetsandBoxMode(true,string(rand.Int()))
	//end := make(chan int)
	ok := false
	finished := make(chan int)
	go func(){
		select{
			case <-time.After(time.Second * 5):
				finished <- 1
				return
			default:
				_,err := engine.Js.Run(code)
				ok = true
				//resultString += value.String()
				if err != nil{
					resultString = err.Error()
				}
				if resultString == ""{
					resultString = "Okay,No output."
				}
				w.Write([]byte(resultString))
				finished <- 1
				fmt.Println(resultString)
				return
		}
	}()
	<- finished
	close(finished)
	if resultString == ""{
		resultString = "Timeout,No Output."
	}
	w.Write([]byte(resultString))
	return
}