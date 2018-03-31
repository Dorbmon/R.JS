package pkg_stack
import (
	"errors"
	otto "github.com/robertkrimen/otto"
	"fmt"
	"os"
)
type Stack struct{
	data []interface{}
	now int
}
func (this *Stack)init(){
	this.data = make([]interface{},1)
}
func New()Stack{
	var temp Stack
	temp.data = make([]interface{},1)
	return temp
}
func (this *Stack)Pop()(interface{}){
	this.now--
	if this.now == -1{
		this.now ++
		//var temp interface{}
		return nil
	}
	return_value := this.data[this.now]
	this.Remove(this.now)
	return return_value
}
func (this *Stack)Push(data interface{})int{	//返回数字从0开始数
	this.data[this.now] = data
	temp := this.now
	this.now++
	return temp
}
func (this Stack)Output(){
	fmt.Print(this.data)
}
func (this *Stack)Remove(key int){	//从0开始数
	n := make([]interface{},len(this.data))
	now_data_point := 0
	for i:=0;i<len(this.data);i++{
		if i != key{
			n[now_data_point] = this.data[i]
			now_data_point++
		}
	}
	this.data = n
}
func (this Stack)Get(key int)(interface{},error){	//从0开始数
	if key >= this.now{
		return nil,errors.New("Too large number")
	}
	return this.data[key],nil
}
/*		JS中的Stack支持		*/
type the_struct_of_js_stack struct{
	stack *Stack
	Used bool
}
var JS_Stack map[string] *the_struct_of_js_stack
func Set_JS_Stack(js *otto.Otto){
	JS_Stack = make(map[string]*the_struct_of_js_stack)
	//0表示参数错误
	//2表示该栈名已存在
	//1表示成功
	js.Set("NewStack",func(call otto.FunctionCall)otto.Value{
		if call.Argument(0).IsNull(){
			fmt.Print("No enough Arguments For NewStack")
			return otto.FalseValue()
		}
		stack_name,err := call.Argument(0).ToString()
		if err != nil{
			fmt.Print("Error data for NewStack")
			return otto.FalseValue()
		}
		//fmt.Print(JS_Stack[stack_name].Used)
		//return otto.Value{}
		if JS_Stack[stack_name] != nil {
			value,_ := otto.ToValue(2)
			return value
		}
		JS_Stack[stack_name] = new(the_struct_of_js_stack)
		JS_Stack[stack_name].Used = true
		JS_Stack[stack_name].stack = new(Stack)
		JS_Stack[stack_name].stack.init()
		return otto.TrueValue()
	})
	js.Set("Stack_Push",func(call otto.FunctionCall)otto.Value{
		if call.Argument(1).IsNull(){
			fmt.Print("No enough Arguments For Stack_Push")
			return otto.FalseValue()
		}
		//判断栈是否存在
		stack_name,err := call.Argument(0).ToString()
		if err != nil{
			fmt.Print("Error type of Data for Stack_Push on line:",call.CallerLocation())
			return otto.FalseValue()
		}
		if JS_Stack[stack_name] == nil{
			fmt.Print("Try to use a undefined Stack on line:",call.CallerLocation())
			return otto.FalseValue()
		}
		//入栈
		data := call.Argument(1)
		fmt.Print("Pushed : ",data)
		JS_Stack[stack_name].stack.Push(data)
		return otto.TrueValue()
	})
	js.Set("Stack_Pop",func(call otto.FunctionCall)otto.Value{
		if call.Argument(0).IsNull(){
			fmt.Print("No enough Arguments For Stack_Push")
			return otto.FalseValue()
		}
		//判断栈是否存在
		stack_name,err := call.Argument(0).ToString()
		if err != nil{
			fmt.Print("Error type of Data for Stack_Push on line:",call.CallerLocation())
			return otto.FalseValue()
		}
		if JS_Stack[stack_name] == nil{
			fmt.Print("Try to use a undefined Stack on line:",call.CallerLocation())
			return otto.FalseValue()
		}
		//判断是否有数据
		if JS_Stack[stack_name].stack.now == 0{
			return otto.FalseValue()
		}
		fmt.Print("Pop:",JS_Stack[stack_name].stack.Pop())
		value,err := otto.ToValue("s")
		if err != nil{
			fmt.Print("ERROR !!!!!!!RJS ERROR!!!!!On Stack_Pop.Please report this situation to Ruixue at https://Rxues.site")
			os.Exit(0)
		}
		return value
	})
}