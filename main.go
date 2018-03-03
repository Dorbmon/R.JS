package main

import (
	"fmt"
	otto "github.com/robertkrimen/otto"
	"flag"
	"os"
	"io/ioutil"
	//"strconv"
	"unsafe"
	"path/filepath"
	"os/exec"
	"errors"
	"strings"
	//"io"
)
var golang_path,_ = getCurrentPath()	//没有/
var js_source_path string	//没有/
var js_source_path_with_file_name string
const (
	 THE_THING_BETWING_DIR = "\\"
)
func main(){
		file := flag.String("file","","The R.JS source file")
		flag.Parse()
		if *file == ""{
			fmt.Print("ERROR FILE NAME")
			os.Exit(0)
		}
		JavaScript,read_err := ReadAll(*file)
		if read_err != nil{
			print("ERROR:",read_err)
			os.Exit(0)
		}
		//优先在当前目录搜索该文件
		//fmt.Print(golang_path + THE_THING_BETWING_DIR + *file)
		//fmt.Print(golang_path)
		//os.Exit(0)
		if checkFileIsExist( golang_path + THE_THING_BETWING_DIR + *file){
			js_source_path = golang_path
			js_source_path_with_file_name = golang_path + THE_THING_BETWING_DIR + *file
		}else{
			js_source_path_with_file_name = *file
			js_source_path = Substr(js_source_path_with_file_name,0,strings.LastIndex(js_source_path_with_file_name,THE_THING_BETWING_DIR))
		}
		//fmt.Print(js_source_path)
		//os.Exit(0)
		js := otto.New()
		init_Java_Script_Const(js)
		/*		IO部分		*/
		js.Set("output", func(call otto.FunctionCall) otto.Value {
			n := 0
			for {
				value := call.Argument(n)
				if value.String() == "undefined"{
					break
				}
				n++
				fmt.Print(value)
			}
			//return otto.Value{temp,"output"}
			return otto.Value{}
			//return otto.Value{}
		})
		js.Set("IO_fopen",func(call otto.FunctionCall) otto.Value{
			if !check_data(call,2){
				error_("ERROR DATA FOR fopen")
				os.Exit(0)
			}
			//打开一个文件并且返回文件号
			//优先在JS程序目录中寻找
			file_name := call.Argument(0).String()
			//temp,err := call.Argument(1).ToString()

			//file_mode,err := strconv.Atoi(temp)
			//fmt.Println("1:",call.Argument(1))
			temp,err := call.Argument(1).ToInteger()
			if err != nil {
				//error_ (string(temp) + " is wrong")
				//os.Exit(0)
				return otto.FalseValue()
			}
			//fmt.Println("temp:",temp)
			file_mode := int(temp)
			//temp,err := call.Argument(1).ToInteger()
			//fmt.Println(js_source_path + THE_THING_BETWING_DIR + file_name)
			if checkFileIsExist(js_source_path + THE_THING_BETWING_DIR + file_name){
				//是相对于JS程序的路径
				file_name = js_source_path + THE_THING_BETWING_DIR + file_name
			}
			result,err := os.OpenFile(file_name,file_mode,0)
			if err != nil{
				return otto.FalseValue()
			}
			defer result.Close()
			//int_pointer := *result
			//dd := int64(result)
			strPointerHex := string(fmt.Sprintf("%x",*result))
			//fmt.Println("length:",unsafe.Sizeof(*result))
			strPointerHex = Substr(strPointerHex,1,len(strPointerHex))
			strPointerHex = Substr(strPointerHex,0,len(strPointerHex) - 1)
			fmt.Println("1:",strPointerHex)
			fmt.Println("2",result)
			//os.Exit(0)
			data,err := otto.ToValue(strPointerHex)
			if err != nil{
				return otto.FalseValue()
			}
			return data
		})
		js.Set("IO_write",func(call otto.FunctionCall) otto.Value{
			if !check_data_(call,2){
				error_("DATA IS NOT ENOUGH FOR IO_write")
				os.Exit(0)
			}
			var file *os.File
			//file = (*os.File)(unsafe.Pointer(&call.ArgumentList[0]))
			//file := (*os.File)(unsafe.Pointer(&call.ArgumentList[0]))
			//file := (*os.File)call.Argument(0).ToString()
			//fmt.Print("\n",unsafe.Pointer(&call.ArgumentList[0]))
			//fmt.Print(&call.ArgumentList[0])
			address := call.ArgumentList[0].String()
			fmt.Println("address:",address)
			file = (*os.File)(unsafe.Pointer(&address))
			//gob.Register()
			fmt.Print("file:",file)
			//os.Exit(0)
			data := []byte(call.Argument(1).String())
			temp := *file
			_,err := temp.Write(data)
			if err != nil{
				return otto.FalseValue()
			}
			fmt.Print("Wrote.")
			return otto.TrueValue()
		})

		js.Set("IO_fcreate",func(call otto.FunctionCall) otto.Value{
			if !check_data(call,1){
				error_("DATA IS NOT ENOUGH FOR fcreate")
				os.Exit(0)
			}
			//创建文件 实际上就是以create权限打开文件
			file_name,err := call.Argument(0).ToString()
			if err != nil{
				error_e(err)
				os.Exit(0)
			}
			if checkFileIsExist(file_name){
				//文件存在，返回false
				result2,err1 := otto.ToValue(false)
				if err1 != nil{
					error_e(err1)
					os.Exit(0)
				}
				return result2
			}
			_,create_err := os.OpenFile(file_name,os.O_CREATE,0)
			if create_err != nil{
				//文件存在，返回false
				result2,err1 := otto.ToValue(false)
				if err1 != nil{
					error_e(err1)
					os.Exit(0)
				}
				return result2
			}
			result,err := otto.ToValue(true)
			if err != nil{
				error_e(err)
				os.Exit(0)
			}
			return result
		})
		js.Set("IO_fclose",func(call otto.FunctionCall) otto.Value{
			if !check_data(call,1){
				error_("DATA IS NOT ENOUGH FOR fclose")
				os.Exit(0)
			}
			//temp := call.Argument(0)
			//file := (*file)(call.Argument(0))
			var file *os.File
			file = (*os.File)(unsafe.Pointer(&call.ArgumentList[0]))
			err := file.Close()
			//file = call.Argument(0).ToInteger()
			result,err2 := otto.ToValue(false)
			if err2 != nil{
				error_e(err2)
				os.Exit(0)
			}
			if err == nil{
				result,err2 = otto.ToValue(true)
				if err2 != nil{
					error_e(err2)
					os.Exit(0)
				}
			}
			return result
		})

		/*		源代码处理部分		*/
		js.Set("include",func(call otto.FunctionCall) otto.Value{
			if !check_data(call,1){
				error_("DATA IS NOT ENOUGH FOR include")
				os.Exit(0)
			}
			n := 0
			for {
				value := call.Argument(n)
				if value.String() == "undefined"{
					break
				}
				n++
				//fmt.Print(value)
				//读取文件
				Code,read_err := ReadAll(*file)
				if read_err != nil{
					print("ERROR:",read_err)
					os.Exit(0)
				}
				js.Run(Code)
			}
			return otto.Value{}
		})
		js.Set("include_c",func(call otto.FunctionCall) otto.Value{
			//从系统目录下包含
			if !check_data(call,1){
				error_("DATA IS NOT ENOUGH FOR include")
				os.Exit(0)
			}
			n := 0
			for {
				value := call.Argument(n)
				if value.String() == "undefined"{
					break
				}
				n++
				//fmt.Print(value)
				//读取文件
				SourceName := value.String()
				Code,read_err := ReadAll(getCurrentDirectory() + THE_THING_BETWING_DIR + SourceName + THE_THING_BETWING_DIR + SourceName + ".js")
				if read_err != nil{
					error_e(read_err)
					os.Exit(0)
				}
				js.Run(Code)
			}
			return otto.Value{}
		})
		_,err := js.Run(string(JavaScript))
		if err != nil{
			error_e(err)
		}
		//fmt.Print(value)
}
func init_Java_Script_Const(vm *otto.Otto){
	/*		文件权限相关底层常量		*/
	vm.Set("IO_ReadWrite",os.O_RDWR)
	//syscall.O_RDWR
	vm.Set("IO_ReadOnly",os.O_RDONLY)
	vm.Set("IO_Create",os.O_CREATE)
	vm.Set("IO_Append",os.O_APPEND)	//追加方式


	return
	//JavaScript_const_var += "var IO_ReadWrite = " + string(os.O_RDWR)
	//JavaScript_const_var += "var IO_ReadOnly = " + string(os.O_RDONLY)
	//JavaScript_const_var += "var IO_Create = " + string(os.O_CREATE)
	//JavaScript_const_var += "\n"
}
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
func check_data(call otto.FunctionCall,number int)bool{
	number--
	if call.Argument(number).IsUndefined(){
		return false
	}
	return true
}
func check_data_(call otto.FunctionCall,must_number int) bool{
	must_number --
	if call.Argument(must_number).IsUndefined(){
		return false;
	}
	must_number += 2
	if call.Argument(must_number).IsUndefined(){
		return true
	}
	return false
}
func error_(message string){
	fmt.Print("ERROR:",message)
}
func error_e(message error){
	fmt.Print(message)
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
	return string(path[0 : i]), nil
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