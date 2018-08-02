package pkg_thread

import "github.com/Dorbmon/otto"
import (
	"math/rand"
	"fmt"
)
type thread struct{
	stop chan int
	has bool
	status bool	//If it is true,the thread is running.
	otto *otto.Otto
}
type Pkg_thread struct{
	master *otto.Otto
	threads map[int]*thread
}

func (this Pkg_thread)SwapJS(js *otto.Otto){
	this.master = js
	this.threads = make(map[int]*thread)
	js.Set("RunThread",func(call otto.FunctionCall)otto.Value{
		ThreadFunction := call.Argument(0)
		if !ThreadFunction.IsFunction(){
			return otto.FalseValue()
		}
		SuceedFunction := call.Argument(1)
		if !SuceedFunction.IsFunction(){
			return otto.FalseValue()
		}
		ErrorFunction := call.Argument(2)
		if !ErrorFunction.IsFunction(){
			return otto.FalseValue()
		}
		//Get Arguments
		var Arugments []otto.Value
		n := 2
		for {
			n++
			fmt.Println("read")
			if call.Argument(n).IsNull() || !call.Argument(n).IsDefined(){
				break
			}
			Arugments[n - 3] = call.Argument(n)
		}
		ThreadID,err := this.RunThread(ThreadFunction,ErrorFunction,SuceedFunction,Arugments)
		if err != nil{
			return otto.FalseValue()
		}
		value,_ := otto.ToValue(ThreadID)
		return value
	})
	js.Set("DeleteThread",func(call otto.FunctionCall)otto.Value{
		T_ThreadID := call.Argument(0)
		if T_ThreadID.IsNull(){
			return otto.FalseValue()
		}
		T2_ThreadID,err := T_ThreadID.ToInteger()
		if err != nil{
			return otto.FalseValue()
		}
		ThreadID := int(T2_ThreadID)
		value := otto.ToValueNoERROR(this.DeleteThread(ThreadID))
		return value
	})
	js.Set("PauseThread",func(call otto.FunctionCall)otto.Value{
		T_ThreadID := call.Argument(0)
		if T_ThreadID.IsNull(){
			return otto.FalseValue()
		}
		T2_ThreadID,err := T_ThreadID.ToInteger()
		if err != nil{
			return otto.FalseValue()
		}
		ThreadID := int(T2_ThreadID)
		value,_ := otto.ToValue(this.PauseThread(ThreadID))
		return value
	})
	js.Set("ContinueThread",func(call otto.FunctionCall)otto.Value{
		T_ThreadID := call.Argument(0)
		if T_ThreadID.IsNull(){
			return otto.FalseValue()
		}
		T2_ThreadID,err := T_ThreadID.ToInteger()
		if err != nil{
			return otto.FalseValue()
		}
		ThreadID := int(T2_ThreadID)
		value,_ := otto.ToValue(this.ContinueThread(ThreadID))
		return value
	})
	return
}
func (this Pkg_thread)DeleteThread(ThreadID int)bool{
	if !this.threads[ThreadID].has{
		return false
	}
	if this.threads[ThreadID].status{
		this.threads[ThreadID].stop <- 1
	}
	close(this.threads[ThreadID].otto.Runtime.ContinueChan)
	delete(this.threads,ThreadID)
	return true
}
func (this Pkg_thread)StopThread(ThreadID int)bool{
	if _,ok := this.threads[ThreadID];!ok{
		return false
	}
	this.threads[ThreadID].stop <- 1
	this.threads[ThreadID].status = false
	return true
}
func (this Pkg_thread)PauseThread(ThreadID int)bool{
	//Check Map
	if _,ok := this.threads[ThreadID];!ok{
		return false
	}
	this.threads[ThreadID].otto.Runtime.Pause = true
	return true
}
func (this Pkg_thread)ContinueThread(ThreadID int)bool{
	if _,ok := this.threads[ThreadID];!ok{
		return false
	}
	this.threads[ThreadID].otto.Runtime.ContinueChan <- 1
	return true
}
func (this Pkg_thread)RunThread(Function otto.Value,ERRORFunction otto.Value,SuceedFunction otto.Value,Arguments []otto.Value,)(int,error){

	/*value := ""
	fmt.Println(value)
	return 0,nil	*/
	again:
	threadid := rand.Int()
	if _,ok := this.threads[threadid];ok{
		goto again
	}
	stopchan := make(chan int)
	this.threads[threadid] = &thread{
		has : true,
		stop : stopchan,
		status : false,
	}
	go func(a chan int,threadid int,Function otto.Value,Master *otto.Otto,Arguments []otto.Value){
		select {
		case <- a:
			return
		default:
			NewOTTO,err := otto.NewThread(Master,Function)
			if err != nil{
				fmt.Println(err)
				return
			}
			this.threads[threadid].otto = NewOTTO
			this.threads[threadid].status = true
			NewOTTO.Runtime.FPSFunction = func(){

			}
			var InterfaceGroup []interface{}
			for i,v := range Arguments{
				InterfaceGroup[i] = v
			}
			/*
			Function.SetRuntime(NewOTTO.Runtime)

			value2,err := Function.Call(Function,InterfaceGroup...)*/
			NewOTTO.Run(Function.String())
			FunctionName := NewOTTO.Runtime.MainFunctionList[0].GetName()
			rf,err := NewOTTO.Get(FunctionName)
			if err != nil{
				fmt.Println(err)
				return
			}
			if !rf.IsFunction(){
				return
			}
			value2,err := rf.Call(rf,InterfaceGroup...)
			this.threads[threadid].status = false
			//this.master.Runtime.Pause = true
			if err != nil {
				ERRORFunction.Call(ERRORFunction,err.Error())
				this.StopThread(threadid)
				return
			}
			//runtime.Gosched()
			SuceedFunction.Call(SuceedFunction,value2)
			//this.master.Runtime.ContinueChan <- 1
			//close(a)
			return
		}
		//Function.Object().Keys()
	}(stopchan,threadid,Function,this.master,Arguments)
	fmt.Println("sss")
	return threadid,nil
}