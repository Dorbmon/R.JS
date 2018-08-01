package pkg_thread

import "github.com/dorbmon/otto"
import "errors"
type pkg_thread struct{
	master *otto.Otto
}

func (this pkg_thread)SwapJS(js *otto.Otto){
	this.master = js
}
func (this pkg_thread)StopThread(ThreadID int){

}
func (this pkg_thread)RunThread(Function otto.Value)(int,error){
	NewOTTO,err := otto.NewThread(this.master,Function)
	if err != nil{
		return 0,err
	}
	go func(){
		Function.Call()
		NewOTTO.Run(Function.String())
		//Function.Object().Keys()
	}()
}