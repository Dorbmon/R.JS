package pkg_math

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"github.com/robertkrimen/otto"
	"fmt"
	"os"
)

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
//Bind Functions
func Init(engine *otto.Otto){
	engine.Set("LoadMath",func(){
		math,err := engine.Object("Math")
		if err != nil{
			fmt.Println(err)
			OnStrictMode(engine)
		}
		math.Set("Fibonacci",func(call otto.FunctionCall)otto.Value{
			number,err := call.Argument(0).ToInteger()
			if err != nil{
				fmt.Println(err)
				OnStrictMode(engine)
			}
			last := 1
			llast := 1
			now := 0
			nm := int(number)
			for n := 3; n <= nm;{
				now := last + llast
				llast = last
				last = now
				continue
			}
			value,_ := otto.ToValue(now)
			return value
		})
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