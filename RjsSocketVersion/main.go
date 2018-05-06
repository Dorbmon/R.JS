//Only Linux
package main
import(
	//"github.com/gorilla/websocket"
	rr "../c/runtime"
	"os"
	"fmt"
	"os/exec"
)

func main(){
	//此程序为监视进程。
	//启动RJS后台监听程序
	path := rr.GetGolangPath()
	//判断主程序是否存在
	if _,err := os.Stat(path+string(os.PathSeparator)+"RJSWeb");err != nil{
		//文件损坏
		fmt.Print("Error file.")
		os.Exit(0)
	}
	//启动
	exec.Command("nohup",path+string(os.PathSeparator)+"RJSWeb","&")
	return
}