package pkg_stack
import (
	"errors"
)
type Stack struct{
	data []interface{}
	now int
}
func New()Stack{
	var temp Stack
	temp.data = make([]interface{},1)
	return temp
}
func (this Stack)Pop()(interface{}){
	this.now--
	if this.now == 0{
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