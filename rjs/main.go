package main
import(
	"../engine"
	"os"
	"flag"
	//"io"
	"fmt"
	colors "github.com/issue9/term/colors"
	"../c/build_about"
)
var std = colors.New(colors.Black,colors.White)
func main(){
	std.Println("Welcome to RJS.")
	std.Println("Ruixue builded this version on:",build_about.GetBuildTime())
	//fmt.Print(delete_interface("[sss]"))

	file := flag.String("file","","The R.JS source file")
	/*		获取部分限定参数		*/
	rjs := engine.RJSEngine{}
	RunPoiontMode := flag.Bool("runmode",false,"If on Running Mode")

	temp_file_max := flag.Int("opened_file_max",10,"The MAX number of OPENING FILES")

	flag.Parse()
	rjs.OPENED_FILE_MAX = *temp_file_max
	rjs.RunPoiontMode = *RunPoiontMode
	if *file == ""{
		fmt.Print("ERROR FILE NAME")
		os.Exit(0)
	}
	rjs.Init()
	JavaScript,read_err := engine.ReadAll(*file)
	if read_err != nil{
		fmt.Print("ERROR:",read_err)
		os.Exit(0)
	}
	value,err := rjs.Js.Run(JavaScript)

	if value.IsDefined(){
		fmt.Println(value)
	}
	if err != nil{
		fmt.Println(err)
	}
	return
	//rjs.Js.Run()
	//engine.OneLineRun()
}