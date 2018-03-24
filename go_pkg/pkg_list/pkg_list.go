package pkg_list

import (
	"./../pkg_stack"
	errors "errors"
)
const (
	DEFAULT_MAX = 100
)
type list struct{
	data map[int]string
	unused_number pkg_stack.Stack
	max_number int
}
func (this list)New(max int)list{
	if max == 0{
		max = DEFAULT_MAX
	}
	this.max_number = max
	this.data = make(map[int]string,1)	//初始化队列
	//this.unused_number = make([]int,max)
	for i:=0;i<max;i++{
		this.unused_number.Push(i)
	}
	return this
}
func (this list)add(data interface{}){

}
func (this list)get(number int)(interface{},error){
	if number > this.max_number{
		return nil,errors.New("Too large number")
	}
	return this.data[]
}