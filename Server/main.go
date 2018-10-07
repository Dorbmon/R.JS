package main

import "net/http"
import "encoding/base64"
import (
	"github.com/dorbmon/R.JS/engine"
	"github.com/dorbmon/otto"
	"math/rand"
	"time"
)
const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)
var coder = base64.NewEncoding(base64Table)
func main(){
	//创建服务器
	http.HandleFunc("/",Run)
	http.ListenAndServe("0.0.0.0:8887",nil)
}
func Run(w http.ResponseWriter, r *http.Request){
	//获得程序
	code,err := base64Decode([]byte(r.Form.Get("code")))
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}
	//创建RJS引擎
	var engine engine.RJSEngine
	engine.Init()
	resultString := ""
	output := func(call otto.FunctionCall){
		for n := 0;;n ++ {
			Data := call.Argument(n)
			if Data.IsNull(){
				break
			}
			resultString += Data.String()
		}
		return
	}
	engine.Js.Set("output",output)
	engine.SetsandBoxMode(true,string(rand.Int()))
	end := make(chan int)
	go func(){
		select{
			case <-end:
				return
			default:
				value,err := engine.Js.Run(code)
				resultString += value.String()
				if err != nil{
					resultString = err.Error()
				}
				return
		}
	}()
	time.After(time.Second * 5 )
	if resultString == ""{
		resultString = "Timeout,No Output."
	}
	w.Write([]byte(resultString))
	return
}
func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}