package pkg_ffi

import(
	"syscall"
	"../../engine"
)

type FfiSystem struct{
	vm *engine.RJSEngine

	//CallBacks..
	SetVar uintptr
	GetVar uintptr
	CallFunc uintptr
}
func (this FfiSystem)Init(vm *engine.RJSEngine){
	this.vm = vm
	this.SetVar = syscall.NewCallback(vm.SetVar)
	this.GetVar = syscall.NewCallback(vm.GetVar)
	this.CallFunc = syscall.NewCallback(vm.CallFunc)
	return
}

