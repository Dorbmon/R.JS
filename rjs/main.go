package main
import(
	"../engine"
	"os"
	"flag"
	//"io"
	"fmt"

)
func main(){
	//fmt.Print(delete_interface("[sss]"))
	file := flag.String("file","","The R.JS source file")
	/*		获取部分限定参数		*/
	engine.OPENED_FILE_MAX = *flag.Int("opened_file_max",10,"The MAX number of OPENING FILES")
	flag.Parse()
	if *file == ""{
		fmt.Print("ERROR FILE NAME")
		os.Exit(0)
	}
	engine.Run(file)
	//engine.OneLineRun()
}