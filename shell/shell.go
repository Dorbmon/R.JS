package main
import (
	"runtime"
	"fmt"
	"../c/build_about"
	"os"
	"bufio"
	"../engine"
	"github.com/robertkrimen/otto"
	"os/exec"
	"strings"
	"path/filepath"
	"errors"
)
var now_user string
var now_path string
var path_between = "/"	// 有可能为/
//var shell_path =
func main(){
	//初始化环境
	fmt.Println("CPU:",runtime.GOARCH)
	fmt.Println("RSHELL builded on :",build_about.GetBuildTime())
	fmt.Println("By Ruixue.")

	//初始化shell
	Init_Special_Func(engine.GetJs())

	for {
		fmt.Print(now_user + " " + now_path + ":")
		reader := bufio.NewReader(os.Stdin)
		temp, _, _ := reader.ReadLine()
		order := string(temp)
		return_value,err := engine.OneLineRun(order)
		if err != nil{
			fmt.Print(err)
			os.Exit(0)
		}
		fmt.Print(return_value)
		continue
	}
}
func Init_Special_Func(obj *otto.Otto){
	obj.Set("about","Ruixues build this program.https://rxues.site.Welcome to Ruixue!\n")
	obj.Set("cd",cd)
}
/*		以下为Shell特殊JS函数		*/

func cd(call otto.FunctionCall)otto.Value{
	if call.Argument(0).IsNull(){
		result,_ := otto.ToValue("DATA IS NOT ENOUGH FOR cd")
		return result
	}

	/*		检查文件/目录		*/
	//判断是否为相对路径
	dir_name,err := call.Argument(0).ToString()
	if err != nil{
		result,_ := otto.ToValue("ERROR DATA'S TYPE FOR cd")
		return result
	}
	if stat,_ := os.Stat(dir_name);!stat.IsDir(){
		//非绝对目录 判断是否为相对目录
		real_path,_ := getCurrentPath()
		real_path +=  path_between
		real_path += dir_name
		result,_ := otto.ToValue(dir_name + " is not a dir name.Please check")
		return result
	}
	return otto.Value{}
}
func getCurrentPath() (string, error) {	//没有/
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	path_between = "/"
	if i < 0 {
		i = strings.LastIndex(path, "\\")
		path_between = "\\"
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i]), nil
}