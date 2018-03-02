package main

import (
	"fmt"
	otto "github.com/robertkrimen/otto"
	"flag"
	"os"
	"io/ioutil"
	"strconv"
	"unsafe"
)
var JavaScript_const_var string
func main(){
		file := flag.String("file","","The R.JS source file")
		flag.Parse()
		print(*file)
		if *file == ""{
			fmt.Print("ERROR FILE NAME")
			os.Exit(0)
		}
		JavaScript,read_err := ReadAll(*file)
		if read_err != nil{
			print("ERROR:",read_err)
			os.Exit(0)
		}

		js := otto.New()
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
		js.Set("fopen",func(call otto.FunctionCall) otto.Value{
			if !check_data(call,2){
				error_("ERROR DATA FOR fopen")
			}
			//打开一个文件并且返回文件号
			file_name := call.Argument(0).String()
			temp,err := call.Argument(1).ToString()
			if err != nil{
				error_ (string(temp) + " is wrong")
				os.Exit(0)
			}
			file_mode,err := strconv.Atoi(temp)
			if err != nil{
				error_ (string(file_mode) + " is wrong")
				os.Exit(0)
			}
			result,err := os.OpenFile(file_name,file_mode,0)
			if err != nil{
				error_e(err)
				os.Exit(0)
			}
			defer result.Close()
			data,err := otto.ToValue(result)
			if err != nil{
				error_e (err)
				os.Exit(0)
			}
			return data
		})
		js.Set("fcreate",func(call otto.FunctionCall) otto.Value{
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
		js.Set("fclose",func(call otto.FunctionCall) otto.Value{
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
		js.Run(JavaScript_const_var + string(JavaScript))
		//fmt.Print(value)
}
func init_Java_Script_Const(){
	/*		文件权限相关		*/
	JavaScript_const_var += "var IO_ReadWrite = " + string(os.O_RDWR)
	JavaScript_const_var += "var IO_ReadOnly = " + string(os.O_RDONLY)
	JavaScript_const_var += "var IO_Create = " + string(os.O_CREATE)
	JavaScript_const_var += "\n"
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
	if call.Argument(number).String() == "undefined"{
		return false
	}
	return true
}
func error_(message string){
	fmt.Print("ERROR:",message)
}
func error_e(message error){
	fmt.Print(message)
}
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}