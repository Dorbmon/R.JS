package pkg_channel
import (
	"fmt"
	"github.com/Dorbmon/otto"
	"os"
	"../pkg_list"
)
var js *otto.Otto
type channel_struct struct{
	used bool
	name string
	List pkg_list.List
}
var channel_map map[string]*channel_struct
func Swap_data(js_engine *otto.Otto){
	*js = *js_engine
	channel_map = make(map[string]*channel_struct,1)
	js.Set("CHANNEL_START",func(call otto.FunctionCall)otto.Value{	//第一个参数为通道名称
		if call.Argument(0).IsNull(){
			fmt.Print("Data is not enough for CHANNEL_START on:",call.CallerLocation())
			os.Exit(0)
		}
		channel_name,err := call.Argument(0).ToString()
		if err != nil{
			return otto.FalseValue()
		}
		//在MAP中创建
		if channel_map[channel_name].used{
			//已经被占用
			return otto.FalseValue()
		}
		channel_map[channel_name].used = true
		channel_map[channel_name].List.List = make([]interface{},1)
		channel_map[channel_name].name = channel_name
		//创建完成
		return otto.TrueValue()
	})
	js.Set("CHANNEL_APPEND",func(call otto.FunctionCall)otto.Value{
		if call.Argument(1).IsNull(){
			fmt.Print("Data is not enough for CHANNEL_APPEND on:",call.CallerLocation())
			os.Exit(0)
		}
		//判断通道是否存在
		channel_name,err := call.Argument(0).ToString()
		if err != nil{
			return otto.FalseValue()
		}
		if !channel_map[channel_name].used{
			return otto.FalseValue()
		}

		//输入数据
		data := call.Argument(1)
		channel_map[channel_name].List.Append(data)
		return otto.TrueValue()
	})
	js.Set("CHANNEL_OUT",func(call otto.FunctionCall)otto.Value{
		if call.Argument(0).IsNull(){
			fmt.Print("Data is not enough for CHANNEL_OUT on:",call.CallerLocation())
			os.Exit(0)
		}
		//判断通道是否存在
		channel_name,err := call.Argument(0).ToString()
		if err != nil{
			return otto.FalseValue()
		}
		if !channel_map[channel_name].used{
			return otto.FalseValue()
		}

		//输出数据
		value,err := otto.ToValue(channel_map[channel_name].List.Out())
		if err != nil{
			return otto.FalseValue()
		}
		return value
	})
}
