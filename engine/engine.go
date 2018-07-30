package engine
/*		欢迎来到瑞雪		*/
/*			我们致力于创建最好的代码			*/

import (
	"fmt"
	"github.com/Dorbmon/otto"
	//"flag"
	"io/ioutil"
	"os"
	//"strconv"
	"errors"
	"go/build"
	"os/exec"
	"path/filepath"
	"strings"
	"unsafe"
	//"io"
	"math/rand"
	"runtime"
	//"math"
	//"github.com/mattn/go-sqlite3"
	//"database/sql"
)

//此处为RJS库
import (
	"../go_pkg/pkg_load"
	"../go_pkg/pkg_math"
	"../go_pkg/pkg_network"
	"../go_pkg/pkg_os"
	"../go_pkg/pkg_secret"
	"../go_pkg/pkg_stack"
	"../go_pkg/pkg_web"
	//"../go_pkg/pkg_ffi"
	//"log"

	"net/http"
	"time"
)

//runtime.GOMAXPROCS(runtime)
/*
                   _ooOoo_
                  o8888888o
                  88" . "88
                  (| -_- |)
                  O\  =  /O
               ____/`---'\____
             .'  \\|     |//  `.
            /  \\|||  :  |||//  \
           /  _||||| -:- |||||-  \
           |   | \\\  -  /// |   |
           | \_|  ''\---/''  |   |
           \  .-\__  `-`  ___/-. /
         ___`. .'  /--.--\  `. . __
      ."" '<  `.___\_<|>_/___.'  >'"".
     | | :  `- \`.;`\ _ /`;.`/ - ` : | |
     \  \ `-.   \_ __\ /__ _/   .-` /  /
======`-.____`-.___\_____/___.-`____.-'======
                   `=---='
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
         佛祖保佑       永无BUG
*/
func main(){

}
func HHHH(){
	fmt.Println("sss")
}
var opened_file_map map[int]opened_file //储存打开的文件
type opened_file struct {
	File     *os.File
	Chmode   int
	Is_alive bool
}

var ( //默认引擎的设置
	js_source_path_with_file_name string
	OPENED_FILE_NUMBER            = 10 //初始化时给opened_file map的个数
	OPENED_FILE_MAX               int  //最大打开文件数
	THE_THING_BETWING_DIR         = string(os.PathSeparator)
	js                            = otto.New()
	Golang_path, _                = getCurrentPath() //没有/
	golang_path, _                = getCurrentPath() //没有/
	js_source_path                string             //没有/
)

func GetJs() *otto.Otto {
	return js
}
func OneLineRun(line string) (value otto.Value, err error) {
	return js.Run(line)
}
func Run(file *string) {
	JavaScript, read_err := ReadAll(*file)
	if read_err != nil {
		fmt.Print("ERROR:", read_err)
		os.Exit(0)
	}

	//include_network.
	//优先在当前目录搜索该文件
	//fmt.Print(golang_path + THE_THING_BETWING_DIR + *file)
	//fmt.Print(golang_path)
	//os.Exit(0)
	opened_file_map = make(map[int]opened_file, OPENED_FILE_NUMBER)
	if checkFileIsExist(golang_path + THE_THING_BETWING_DIR + *file) {
		js_source_path = golang_path
		js_source_path_with_file_name = golang_path + THE_THING_BETWING_DIR + *file
	} else {
		js_source_path_with_file_name = *file
		js_source_path = Substr(js_source_path_with_file_name, 0, strings.LastIndex(js_source_path_with_file_name, THE_THING_BETWING_DIR))
	}
	//fmt.Print(js_source_path)
	//os.Exit(0)
	engine := RJSEngine{}
	engine.Init()
	//pkg_secret.SwapData(engine.Js)
	_, err := engine.Js.Run(string(JavaScript))
	if err != nil {
		error_e(err)
	}
	//fmt.Print(JavaScript)
}

func init_Java_Script_Const(vm *otto.Otto) {
	/*		文件权限相关底层常量		*/
	vm.Set("IfRJSRunTime", true)
	vm.Set("IO_ReadWrite", os.O_RDWR)
	//syscall.O_RDWR
	vm.Set("IO_ReadOnly", os.O_RDONLY)
	vm.Set("IO_Create", os.O_CREATE)
	vm.Set("IO_Append", os.O_APPEND)       //追加方式
	vm.Set("RJS_CONFIG_STRIT_MODE", false) //严格模式下内存错误等会导致程序退出
	vm.Set("PathSeparator", os.PathSeparator)

	//Time
	vm.Set("Minute",time.Minute)
	vm.Set("Second",time.Second)
	return
	//JavaScript_const_var += "var IO_ReadWrite = " + string(os.O_RDWR)
	//JavaScript_const_var += "var IO_ReadOnly = " + string(os.O_RDONLY)
	//JavaScript_const_var += "var IO_Create = " + string(os.O_CREATE)
	//JavaScript_const_var += "\n"
}
func delete_interface(data string) string {
	data = Substr(data, 1, len(data))
	data = Substr(data, 0, len(data)-1)
	return data
}
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
func check_data(call otto.FunctionCall, number int) bool {
	number--
	if call.Argument(number).IsUndefined() {
		return false
	}
	return true
}
func check_data_(call otto.FunctionCall, must_number int) bool {
	must_number--
	if call.Argument(must_number).IsUndefined() {
		return false
	}
	must_number += 2
	if call.Argument(must_number).IsUndefined() {
		return true
	}
	return false
}
func error_(message string) {
	fmt.Println("ERROR:", message)
}
func error_e(message error) {
	fmt.Println(message)
}
func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		error_e(err)
		os.Exit(0)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0:i]), nil
}

//截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
func load_outside_progarm(js *otto.Otto, this *RJSEngine) {
	//加载编译时包含的库

	//pkg_stack.Set_JS_Stack(js)//栈库
	pkg_os.Swap_data(js)
	//pkg_load.SwapJS(js)
	this.Load_.SwapJS(js)
	this.Stack_.Set_JS_Stack(js)
	pkg_secret.SwapData(js)
}

type RJSEngine struct {
	Js *otto.Otto
	//环境变量
	RunPoiontMode                     bool
	VMRootPath                        string
	js_source_path                    string //没有/
	js_source_path_with_file_name     string
	OPENED_FILE_NUMBER                int //初始化时给opened_file map的个数
	OPENED_FILE_MAX                   int //最大打开文件数
	THE_THING_BETWING_DIR             string
	OnSandMode                        bool
	Realjs_source_path                string //没有/
	Realjs_source_path_with_file_name string
	//下列为库
	Stack_ pkg_stack.JS_StackEngine
	Load_  pkg_load.JSLoader
	Web pkg_web.RjsWebMoudle
	//FfiSystem pkg_ffi.FfiSystem
}

func (this *RJSEngine) Init() {
	this.Js = otto.New()
	this.OPENED_FILE_MAX = 10 //默认为10
	this.THE_THING_BETWING_DIR = string(os.PathSeparator)
	js := this.Js
	load_outside_progarm(js, this) //加载库
	/*	init Including Setting	*/
	pkg_network.Swap_Data_From_Main(js)
	this.Web.Init(this.Js)
	//include_network.
	init_Java_Script_Const(js)
	js.SetFPSFunction(func(){
		//Do Nothing...

	})//防止野指针
	js.Set("瑞雪",func(){
		fmt.Println("Ruixue is a really strong team.")
		fmt.Println("We try to build everything.")
	})
	js.Set("HSR",func(){
		fmt.Println("Now,You know she is my girlfriend.")
		return
	})
	js.Set("output", func(call otto.FunctionCall) otto.Value {
		//js.Call(call.Argument(0).String(),"")
		n := 0
		for {
			value := call.Argument(n)
			if !value.IsDefined() {
				break
			}
			n++
			fmt.Print(value)
		}
		//return otto.Value{temp,"output"}
		return otto.Value{}
		//return otto.TrueValue()
	})
	if this.RunPoiontMode {
		goto jump_step
	}
	{
		/*		IO部分		*/
		//js.Set("")
		js.Set("call", func(call otto.Value) {
			//call.Argument(0).Call()
			//sjs.Eval(call.Argument(0))
			//_,err := js.Call(call.Argument(0).String(),nil)
			//f,_ := js.Get(call.Argument(0).String())

			//js.Run(f)
			//js.Call(f.String(),nil)
			//call.
			//fmt.Print(call.Call(call,nil))

		})
		/*		初始化SqlLite3		*/
		//db,err := sql.Open("","./RjsSystem.db")
		/*if err != nil{
			//发生错误
			fmt.Print(err)
			this.OnStrictMode()
		}*/
		//db.Ping()
		js.Set("OnlyRand", func() otto.Value { //生成唯一的随机ID
			value, _ := otto.ToValue(pkg_math.UniqueId())
			return value
		})
		//JsOS, err := js.Object("JsOS")
		JsOS, err := js.Object(`({})`)
		if err != nil {
			fmt.Println(111, err)
			os.Exit(0)
		}
		JsOS.Set("OS_SET_MAXPROCS", func(call otto.FunctionCall) otto.Value {
			procs, err := call.Argument(0).ToInteger()
			if err != nil {
				fmt.Print("ERROR Type Of Data For OS_SET")
				this.OnStrictMode()
				return otto.FalseValue()
			}
			//fmt.Print("porcs:",procs)
			runtime.GOMAXPROCS(int(procs))
			//runtime.SetCPUProfileRate()
			return otto.TrueValue()
		})
		JsOS.Set("OS_GET_MAXPROCS", func() otto.Value {
			value, _ := otto.ToValue(runtime.NumCPU())
			return value
		})
		JsOS.Set("IO_fopen", func(call otto.FunctionCall) otto.Value {
			if !check_data(call, 2) {
				error_("ERROR DATA FOR fopen")
				os.Exit(0)
			}
			//打开一个文件并且返回文件号
			//优先在JS程序目录中寻找
			file_name := call.Argument(0).String()
			//temp,err := call.Argument(1).ToString()

			//file_mode,err := strconv.Atoi(temp)
			//fmt.Println("1:",call.Argument(1))
			temp, err := call.Argument(1).ToInteger()
			if err != nil {
				//error_ (string(temp) + " is wrong")
				//os.Exit(0)
				return otto.FalseValue()
			}
			//fmt.Println("temp:",temp)
			file_mode := int(temp)
			//temp,err := call.Argument(1).ToInteger()
			//fmt.Println(js_source_path + THE_THING_BETWING_DIR + file_name)
			if checkFileIsExist(js_source_path + THE_THING_BETWING_DIR + file_name) {
				//是相对于JS程序的路径
				file_name = js_source_path + THE_THING_BETWING_DIR + file_name
			}
		again_rand:
			rand_id := rand.Intn(OPENED_FILE_MAX)
			if opened_file_map[rand_id].Is_alive {
				goto again_rand
			}
			//fmt.Println("rand_id:",rand_id)
			//rand_id := rand.Int()
			//opened_file_map[rand_id].File,err = os.OpenFile(file_name,file_mode,0)
			//temp_opened_file := opened_file{}
			//temp_opened_file.Chmode = file_mode
			//temp_opened_file.File,err = os.OpenFile(file_name,file_mode,0)
			if err != nil {
				temp, _ := otto.ToValue(0)
				return temp
			}
			//temp_opened_file.Is_alive = true
			//opened_file_map[rand_id] = temp_opened_file
			temp_address, _ := os.OpenFile(file_name, file_mode, 0)
			opened_file_map[rand_id] = opened_file{temp_address, file_mode, true}
			/*defer func(){opened_file_map[rand_id].File.Close()
				fmt.Print("ssClosed","err:")

			}()*/
			if _, err = opened_file_map[rand_id].File.Stat(); err != nil {
				temp, _ := otto.ToValue(0)
				return temp
			}
			//int_pointer := *result
			//dd := int64(result)
			//strPointerHex := string(fmt.Sprintf("%x",*opened_file_map[rand_id].File))
			//fmt.Println("length:",unsafe.Sizeof(*result))
			//strPointerHex = Substr(strPointerHex,1,len(strPointerHex))
			//strPointerHex = Substr(strPointerHex,0,len(strPointerHex) - 1)
			//fmt.Println("1:",strPointerHex)
			//fmt.Println("2",result)
			//os.Exit(0)
			data, err_c := otto.ToValue(rand_id)
			if err_c != nil {
				//fmt.Print("reason:",err)
				temp, _ := otto.ToValue(0)
				return temp
			}
			return data
		})
		JsOS.Set("IO_write", func(call otto.FunctionCall) otto.Value {
			if !check_data_(call, 2) {
				error_("DATA IS NOT ENOUGH FOR IO_write")
				os.Exit(0)
			}
			var file *os.File
			//file = (*os.File)(unsafe.Pointer(&call.ArgumentList[0]))
			//file := (*os.File)(unsafe.Pointer(&call.ArgumentList[0]))
			//file := (*os.File)call.Argument(0).ToString()
			//fmt.Print("\n",unsafe.Pointer(&call.ArgumentList[0]))
			//fmt.Print(&call.ArgumentList[0])

			address, err := call.ArgumentList[0].ToInteger()
			if err != nil {
				//fmt.Print("error:",err)
				return otto.FalseValue()
			}
			file = opened_file_map[int(address)].File
			//fmt.Println("address:",address)
			//file = (*os.File)(unsafe.Pointer(&address))
			//_,err = file.Stat()
			//stat,_ := opened_file_map[int(address)].File.Stat()
			//fmt.Print("stat",stat)
			//gob.Register()
			//fmt.Print("file:",file)
			//os.Exit(0)
			data := []byte(delete_interface(call.Argument(1).String()))
			//fmt.Print(file.Stat())
			//data := []byte()
			//temp := *file
			//_,err := temp.Write(data)
			_, err = file.Write(data)
			if err != nil {
				return otto.FalseValue()
			}
			return otto.TrueValue()
		})
		JsOS.Set("IO_fcreate", func(call otto.FunctionCall) otto.Value {
			if !check_data(call, 1) {
				error_("DATA IS NOT ENOUGH FOR fcreate")
				os.Exit(0)
			}
			//创建文件 实际上就是以create权限打开文件
			file_name, err := call.Argument(0).ToString()
			if err != nil {
				error_e(err)
				os.Exit(0)
			}
			if checkFileIsExist(file_name) {
				//文件存在，返回false
				result2, err1 := otto.ToValue(false)
				if err1 != nil {
					error_e(err1)
					os.Exit(0)
				}
				return result2
			}
			_, create_err := os.OpenFile(file_name, os.O_CREATE, 0)
			if create_err != nil {
				//文件存在，返回false
				result2, err1 := otto.ToValue(false)
				if err1 != nil {
					error_e(err1)
					os.Exit(0)
				}
				return result2
			}
			result, err := otto.ToValue(true)
			if err != nil {
				error_e(err)
				os.Exit(0)
			}
			return result
		})
		JsOS.Set("IO_fclose", func(call otto.FunctionCall) otto.Value {
			if !check_data(call, 1) {
				error_("DATA IS NOT ENOUGH FOR fclose")
				os.Exit(0)
			}
			//temp := call.Argument(0)
			//file := (*file)(call.Argument(0))
			var file *os.File
			file = (*os.File)(unsafe.Pointer(&call.ArgumentList[0]))
			err := file.Close()
			//file = call.Argument(0).ToInteger()
			result, err2 := otto.ToValue(false)
			if err2 != nil {
				error_e(err2)
				os.Exit(0)
			}
			if err == nil {
				result, err2 = otto.ToValue(true)
				if err2 != nil {
					error_e(err2)
					os.Exit(0)
				}
			}
			return result
		})
		/*		源代码处理部分		*/
		js.Set("include", func(call otto.FunctionCall) otto.Value {
			if !check_data(call, 1) {
				error_("DATA IS NOT ENOUGH FOR include on:" + call.CallerLocation())
				os.Exit(0)
			}
			n := 0
			for {
				value := call.Argument(n)
				if value.String() == "undefined" {
					break
				}
				n++
				//fmt.Print(value)
				//读取文件
				//Code,read_err := ReadAll(*file)
				Code, read_err := ReadAll(value.String())
				if read_err != nil {
					print("ERROR:", read_err)
					os.Exit(0)
				}
				js.Run(Code)
			}
			return otto.Value{}
		})
		js.Set("LoadModule", func(call otto.FunctionCall) otto.Value { //第一个参数为是否忽略错误执行,第二个参数为是否为临时下载
			//Address := call.Argument(0).String()
			if !check_data(call, 0) || !check_data(call, 1) || !check_data(call, 2) {
				value, _ := otto.ToValue("ERROR DATA for LoadMould on " + call.CallerLocation())
				return value
			}
			if !call.ArgumentList[0].IsBoolean() {
				value, _ := otto.ToValue("ERROR DATA Type for LoadMould on " + call.CallerLocation())
				return value
			}
			if !call.ArgumentList[1].IsBoolean() {
				value, _ := otto.ToValue("ERROR DATA Type for LoadMould on " + call.CallerLocation())
				return value
			}
			IfIgnoreMistake, _ := call.ArgumentList[0].ToBoolean()
			IfTempMode, _ := call.ArgumentList[1].ToBoolean()
			var err error = nil
			if IfIgnoreMistake {
				n := 3
				for {
					value := call.Argument(n)
					if !value.IsDefined() {
						break
					}
					n++
					Rdata := value.String()
					this.LoadModule(Rdata, IfTempMode)
				}
			} else {
				n := 3
				for {
					value := call.Argument(n)
					if !value.IsDefined() {
						break
					}
					n++
					Rdata := value.String()
					err = this.LoadModule(Rdata, IfTempMode)
					if err != nil {
						break
					}
				}
			}
			value, _ := otto.ToValue(err)
			return value
		})
		js.Set("include_c", func(call otto.FunctionCall) otto.Value {
			//从系统目录下包含
			if !check_data(call, 1) {
				error_("DATA IS NOT ENOUGH FOR include")
				os.Exit(0)
			}
			n := 0
			for {
				value := call.Argument(n)
				if value.String() == "undefined" {
					break
				}
				n++
				//fmt.Print(value)
				//读取文件
				SourceName := value.String()
				Code, read_err := ReadAll(getCurrentDirectory() + THE_THING_BETWING_DIR + SourceName + THE_THING_BETWING_DIR + SourceName + ".js")
				if read_err != nil {
					error_e(read_err)
					os.Exit(0)
				}
				js.Run(Code)
			}
			return otto.Value{}
		})
		js.Set("exit", func(call otto.FunctionCall) otto.Value {
			if !check_data(call, 1) {
				error_("DATA IS NOT ENOUGH FOR include")
				os.Exit(0)
			}
			temp, _ := call.Argument(0).ToInteger()
			fmt.Print(call.CallerLocation())
			os.Exit(int(temp))
			return otto.Value{}
		})
		js.Set("IO_Start_TCP_Server", func(call otto.FunctionCall) otto.Value { //TCP监听IP 监听端口 客户连接回调函数 收到数据回调函数 错误回调函数
			if !check_data(call, 5) {
				error_("DATA IS NOT ENOUGH FOR IO_Start_TCP_Server")
				os.Exit(0)
			}
			TCP_IP, err := call.Argument(0).ToString()
			if err != nil {
				error_("ERROR TCP_IP FOR IO_Start_TCP_Server")
				os.Exit(0)
			}
			Port, err := call.Argument(1).ToInteger()
			if err != nil {
				error_("ERROR TCP_IP FOR IO_Start_TCP_Server")
				os.Exit(0)
			}
			//fmt.Print(TCP_IP)
			var temp_recall_message pkg_network.TCP_LISTENER
			temp_recall_message.On_Connect_func, err = call.Argument(2).ToString()
			if err != nil {
				error_("ERROR At IO_Start_TCP_Server")
				os.Exit(0)
			}
			temp_recall_message.On_Data_func, err = call.Argument(3).ToString()
			if err != nil {
				error_("ERROR At IO_Start_TCP_Server")
				os.Exit(0)
			}
			temp_recall_message.On_ERROR, err = call.Argument(4).ToString()
			if err != nil {
				error_("ERROR At IO_Start_TCP_Server")
				os.Exit(0)
			}
			pkg_network.Start_Listen(TCP_IP, int(Port), temp_recall_message)
			return otto.Value{}
		})
		JsRuntime, _ := js.Object("({})")
		JsRuntime.Set("GetLine", func(call otto.FunctionCall) otto.Value {
			value, _ := otto.ToValue(call.CallerLocation())
			return value
		})
		JsRuntime.Set("SetFPSFunction",func(call otto.FunctionCall)otto.Value{
			if !call.Argument(0).IsFunction(){
				return otto.FalseValue()
			}
			js.SetFPSFunction(func(){
				call.Argument(0).Call(call.Argument(0))
			})
			return otto.TrueValue()
		})
		JsRuntime.Set("Gc",func(){
			runtime.GC()	//World Stop.
		})
		JsRuntime.Set("Wait",func(call otto.FunctionCall)otto.Value{
			time1,err := call.Argument(0).ToInteger()
			if err != nil{
				fmt.Println("ERROR TYPE OF TIME For runtime.Wait().on ",call.CallerLocation())
				this.OnStrictMode()
				return otto.FalseValue()
			}
			fmt.Println("Wait ",time1)

			//time.After(time.Minute)
			time.Sleep(time.Duration(time1))
			//time.Sleep(time.Minute)
			return otto.TrueValue()

		})
		JsRuntime.Set("LockOSThread",func(call otto.FunctionCall){
			runtime.LockOSThread()
		})
		js.Set("runtime", JsRuntime)
	}
jump_step:
	return
}
func (this *RJSEngine) SetsandBoxMode(status bool, vmName string) (string, error) {
	//是否启用沙盒模式。如果使用沙盒模式，任何操作不会在物理机上生效，如果关闭沙盒操作，之前的操作也不会生效
	//建立虚拟目录系统.

	if status { //开启
		os.MkdirAll(golang_path+this.THE_THING_BETWING_DIR+"RJSVM", 0777) //创建虚拟根目录
		//创建一个虚拟系统
		err := os.Mkdir(golang_path+this.THE_THING_BETWING_DIR+"RJSVM"+this.THE_THING_BETWING_DIR+vmName, 0777)
		if err != nil {
			return "", errors.New("ERROR When create VM root path.")
		}
		//模拟环境	设置初始根目录
		this.VMRootPath = golang_path + THE_THING_BETWING_DIR + "RJSVM" + THE_THING_BETWING_DIR + vmName + THE_THING_BETWING_DIR
		//设置虚拟根目录
		this.Realjs_source_path = this.js_source_path
		this.js_source_path = this.VMRootPath + this.js_source_path
		this.Realjs_source_path_with_file_name = this.js_source_path_with_file_name
		this.js_source_path_with_file_name = this.VMRootPath + this.js_source_path_with_file_name
		this.OnSandMode = true
		return this.VMRootPath, nil //返回虚拟根目录 方便移动相关脚本文件
	}
	//去除虚拟文件路径.
	if this.OnSandMode { //曾开启
		//删除虚拟路径
		this.VMRootPath = ""
		this.js_source_path_with_file_name = this.Realjs_source_path_with_file_name
		this.js_source_path = this.Realjs_source_path
	}
	this.OnSandMode = false
	return "", nil
}
func (this RJSEngine) OnStrictMode() {
	IfStrictMode, err := this.Js.Get("RJS_CONFIG_STRIT_MODE")
	if err != nil {
		fmt.Print("ERRO CONFIG.RJS_CONFIG_STRIT_MODE")
		os.Exit(0)
	}
	confident, err := IfStrictMode.ToBoolean()
	if err != nil {
		fmt.Print("ERRO CONFIG.RJS_CONFIG_STRIT_MODE")
		os.Exit(0)
	}
	if confident {
		os.Exit(0)
	}
	return
}

type RJSEngineVersion struct {
	Version float32
	BuildOS string
	message string
}

func Version() RJSEngineVersion {
	version := RJSEngineVersion{}
	version.Version = 0.1
	BuildContext := build.Default
	version.BuildOS = BuildContext.GOOS
	version.message = `Ruixue Build This Version for everyone for free.`
	//fmt.Print(BuildContext.GOARCH)
	//build.
	return version
}
func (this RJSEngine) LoadModule(Address string, tempMode bool) error {
	//The Address Mustn't have anything like 'http'
	//本地寻找
	Exists, _ := file_exists(Address)
	if !Exists {
		//进入Web模式
		res, err := http.Get(Address)
		if err != nil {
			return err
		}
		if tempMode {
			//No write
			_, err = this.Js.Run(res.Body)
			return err
		}
		//Write
		//分离文件名
		var data []byte
		res.Body.Read(data)
		position := strings.LastIndex(Address, "/")

		if position == -1 {
			//直接写入文件
			if is, _ := dir_exists(Address); is {
				goto write
			}
			err = os.Mkdir(Address, os.ModePerm)
			if err != nil {
				return err
			}
		write:
			if ex, _ := file_exists(Address + string(os.PathSeparator) + "index.rjs"); ex {
				//删除原有文件
				err = os.Remove(Address + string(os.PathSeparator) + "index.rjs")
				if err != nil {
					return err
				}
			}
			file, err := os.Create(Address + string(os.PathSeparator) + "index.rjs")
			if err != nil {
				return err
			}
			_, err = file.Write(data)
			if err != nil {
				return err
			}
			_, err = this.Js.Run(string(data))
			return err
		}
		//os.MkdirAll(Address)
		//have /
		// Got Dir
		Rdata := strings.Split(Address, "/")
		l := len(Rdata)
		FileName := Rdata[l-1]
		var Dir string
		for n := 0; n < (l - 1); n++ {
			Dir += Rdata[n]
		}
		if i, _ := dir_exists(Dir); i {
			goto write2
		}
		err = os.MkdirAll(Dir, os.ModePerm)
		if err != nil {
			return err
		}
	write2:
		//Write To File
		if i, _ := file_exists(Dir + string(os.PathSeparator) + FileName); i {
			err = os.Remove(Dir + string(os.PathSeparator) + FileName)
			if err != nil {
				return err
			}
		}
		//Create File
		file, err := os.Create(Dir + string(os.PathSeparator) + FileName)
		if err != nil {
			return err
		}
		_, err = file.Write(data)
		return err
	}
	//读取
	data, err := ReadAll(Address)
	if err != nil {
		return err
	}
	_, err = this.Js.Run(string(data))
	if err != nil {
		return err
	}
	return nil
}
func file_exists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err == nil {
		if !stat.IsDir() {
			return true, nil
		} else {
			return false, nil //It's a dir.
		}
	}
	return true, err
}
func dir_exists(dir string) (bool, error) {
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err == nil {
		if stat.IsDir() {
			return true, nil
		} else {
			return false, nil //It's a dir.
		}
	}
	return false, err
}

// These functions that are below are set for C++
func (this RJSEngine) SetVar(VarName string, Value interface{}) string {
	error := this.Js.Set(VarName, Value)
	return error.Error()
}
func (this RJSEngine) GetVar(VarName string) interface{} {
	value, err := this.Js.Get(VarName)
	if err != nil {
		return err
	}
	return value
}
func (this RJSEngine) CallFunc(FuncName string, data []interface{}) interface{} {
	result, err := this.Js.Call(FuncName, data[0:])
	if err != nil {
		return err
	}
	return result
	//result := this.Js.Call()
}
