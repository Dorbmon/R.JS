package pkg_load
//RJS动态加载链接库

import(
	"plugin"
	"github.com/robertkrimen/otto"
	"fmt"
	"strings"
)
var js *otto.Otto
type LoadedFileFunc struct{
	//func_map map[string]*plugin.Plugin
	FilePointer *plugin.Plugin
	has bool
}
var LoadedFile map[string]*LoadedFileFunc
func SwapJS(engine *otto.Otto){
	js = engine
	LoadedFile = make(map[string]*LoadedFileFunc,1)
}

func Load(file_name string)bool{
	if LoadedFile[file_name].has{
		return true
	}
	temp,err := plugin.Open(file_name)
	if err != nil{
		return false
	}
	LoadedFile[file_name].has = true
	LoadedFile[file_name].FilePointer = temp
	FunctionList,err := temp.Lookup("RjsFunctionList")
	if err != nil{
		fmt.Print("不规范的RJS库")
		return false
	}
	FuncList := strings.SplitAfter(FunctionList.(func()string)()," ")
	var row string
	for _,row := range FuncList{
		LoadedFile[file_name].FilePointer
	}
	//提供回调接口给对方库
}
/*		RJS动态库规范		*/
/*		必须有一个函数为获取动态链接库信息	*/
/*		必须有一个名为RjsFunctionList函数返回所有挂载到RJS的函数列表，使用字符串，用空格分开		*/
/*		必须有一个名为RjsSetGetVariableFunction的函数让RJS来设置回调函数的地址。这个函数提供给库用来获取变量的值	，该函数只有一个参数，为函数地址。RJS提供的函数原型为
GetVariableValue(VariableName string)[]Byte
返回字节类型
*/