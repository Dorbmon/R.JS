package pkg_network

import(
	"fmt"
	"net"
	"math/rand"
	otto "github.com/robertkrimen/otto"
)
//var TCP_LISTENER_MAP = make(map[int]TCP_LISTENER)
var TCP_LISTENER_MAP map[int]TCP_LISTENER
var js *otto.Otto
var listener_number = 0
var error_ = func(err error){
	fmt.Print(err)
}
const (
	MAX_TCP_LISTENER = 100
)
func Swap_Data_From_Main(js_engine *otto.Otto){
	*js = *js_engine
	TCP_LISTENER_MAP = make(map[int]TCP_LISTENER)
	//error_ = error_func.(func())
}
type TCP_LISTENER struct {
	listener *net.Listener
	On_Data_func string
	On_Connect_func string
	On_DisConnect_func string
	On_ERROR string
	has bool
	on bool	//是否开启 如果开启则继续消息循环 如果关闭了消息循环自动调整has
}
type TCP_CONN struct{
	conn net.Conn
	has bool
	on bool	//此为信号量 通知消息循环结束
}
/*		error code		*/
//	001 创建监听失败
//	0011 创建失败 达到监听上限
//	002 客户连接失败
//
func Start_Listen(IP string,Port int,CallBackData TCP_LISTENER){
	if listener_number == MAX_TCP_LISTENER{
		//达到监听上限
		js.Call(CallBackData.On_ERROR,001)
		return
	}
	var err error
	//TCP_LISTENER_MAP[listener_num].On_Connect_func = ""
	again_rand:
	listener_num := rand.Intn(MAX_TCP_LISTENER)
	if TCP_LISTENER_MAP[listener_num].has{
		//已被占用
		goto again_rand
	}
	TCP_LISTENER_MAP[listener_num] = CallBackData
	*TCP_LISTENER_MAP[listener_num].listener,err = net.Listen("tcp",IP + ":" + string(Port))
	(TCP_LISTENER_MAP[listener_num]).has = true
	if err != nil{
		error_(err)
		//return err
	}
	//设置监听相关
	//*TCP_LISTENER_MAP[listener_num].
	go func(listener_num int,listener TCP_LISTENER){
		for
		{
			conn, err := (*TCP_LISTENER_MAP[listener_num].listener).Accept()
			if err != nil {
				error_(err)
				//投递错误信息
				js.Call(TCP_LISTENER_MAP[listener_num].On_ERROR,err)
				continue
			}
			//投递连接
			js.Call(listener.On_Connect_func,listener_num)	//投递虚ID
			//针对该客户进行消息循环
			go func(conn net.Conn,listener *TCP_LISTENER){
				for {
					defer func() {
						conn.Close()
						return
					}()
					buffer := make([]byte, 1024)
					n, err := conn.Read(buffer)
					if err != nil {
						return
					}
					data := string(buffer[:n]) //创建一个变量来输出缓冲区
					//投递数据
					js.Call(listener.On_Data_func, data)
				}
			}(conn,&listener)
			//

		}
	}(listener_num,TCP_LISTENER_MAP[listener_num])
}
