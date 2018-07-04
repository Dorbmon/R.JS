package main
import(
	"../engine"
	"github.com/bitly/go-simplejson"
	"github.com/Unknwon/goconfig"
	"net"
	"os"
	"flag"
	//"io"
	"fmt"
	//"../c/build_about"
)
func main(){
	if iss,_ := exists("RJS.ini");!iss{
		os.Create("RJS.ini")
		DefaultData := &goconfig.ConfigFile{}
		//DefaultData.SetSectionComments("Main","RunRJSServer",)
		DefaultData.SetKeyComments("Main","RunRJSServer","yes")
		DefaultData.SetKeyComments("Main","RJSServerPort","1225")
		DefaultData.SetKeyComments("Main","RJSServerNeedLisence","no")
		goconfig.SaveConfigFile(DefaultData,"RJS.ini")
	}
	//判断是否为加载RJS web库命令
	WebDownload := flag.String("get","","If Download From Web")
	//std.Println("Welcome to RJS.")
	//std.Println("Ruixue builded this version on:",build_about.GetBuildTime())
	//fmt.Print(delete_interface("[sss]"))
	file := flag.String("file","","The R.JS source file")
	/*		获取部分限定参数		*/
	rjs := engine.RJSEngine{}
	RunPoiontMode := flag.Bool("runmode",false,"If on Running Mode")
	temp_file_max := flag.Int("opened_file_max",10,"The MAX number of OPENING FILES")
	IfVersion := flag.Bool("v",false,"Show Version and something else about RJS engine.")
	flag.Parse()
	if *IfVersion{
		//介绍RJS
		ShowAboutEngine()
		os.Exit(0)
	}
	if *WebDownload != ""{
		InstallMoudle(*WebDownload)
		os.Exit(0)
	}
	rjs.OPENED_FILE_MAX = *temp_file_max
	rjs.RunPoiontMode = *RunPoiontMode
	if *file == ""{
		//挂载后台RJS通讯进程。
		RJSListener()
		fmt.Println("RJS is running...")
		//os.Exit(0)
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
func ShowAboutEngine(){
	fmt.Println("Welcome to use RJS.")
	version := engine.Version()
	fmt.Println("RJS Version:",version.Version)
	fmt.Println("This RJS engine build on :",version.BuildOS)
	fmt.Println("Ruixue:https://rxues.site")
	fmt.Println("If you want to join us.E-mail at admin@rxues.site")
	fmt.Println("Welcome To Ruixue!")
	fmt.Println("My Girfriend -- HSR...")
	fmt.Println("My heart is full of you...")
	return
}
func InstallMoudle(WebSite string){
	//首先Get下文件

}
func RJSListener(){
	//获取配置信息
	config,err := goconfig.LoadConfigFile("RJS.ini")
	if err != nil{
		fmt.Println(err)
		os.Exit(0)
	}
	Port := config.GetKeyComments("Main","RJSServerPort")

	netListen, err := net.Listen("tcp", "localhost:" + Port)
	if err != nil{
		fmt.Println(err)
		os.Exit(0)
	}
	defer netListen.Close()
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		go DealWithConn(conn)
		//处理连接

	}
}
func DealWithConn(conn net.Conn){
	config,err := goconfig.LoadConfigFile("RJS.ini")
	if err != nil{
		fmt.Println(err)
		os.Exit(0)
	}
	tNeedLisence := config.GetKeyComments("Main","RJSServerNeedLisence")
	NeedLisence := true
	if tNeedLisence == "yes"{
		NeedLisence = true
	}else{
		NeedLisence = false
	}
	buffer := make([]byte, 2048)
	Lisence := false
	for {

		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		data := string(buffer[:n])
		//处理data
		json,err := simplejson.NewJson(buffer[:n])
		if err != nil{
			//JSON解析错误
			fmt.Println(err)
			continue
		}
		dtype,err := json.Get("Type").String()
		if err != nil{
			fmt.Println(err)
			continue
		}
		switch dtype {
		case "register":
			//注册VM
			//判断是否需要授权再注册
			//exists()
			//获取信息
			Name, err := json.Get("Name").String()
			if err != nil {
				fmt.Println("ERROR JSON.Lost Name")
				continue
			}
			About, err := json.Get("About").String()
			if err != nil {
				fmt.Println("ERROR JSON.Lost About")
				continue
			}
			if NeedLisence {
				//需要授权
				if !Lisence(Name, About) {
					conn.Write([]byte(`
						type:'Respond',
						data:"ERROR LISENCE.User didn't agree."
					`))
					break
				}
			}
			//授权成功
			Lisence = true

		case "load":
			if !Lisence{
				//未授权
				conn.Write([]byte(`
						type:'Respond',
						data:"ERROR LISENCE.You don't have lisence'."
				`))
				break
			}
			//加载到RJS引擎中
			

		default:
			//Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
	}
	}

}
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}
func Lisence(Name string,About string)bool{
	return true
}