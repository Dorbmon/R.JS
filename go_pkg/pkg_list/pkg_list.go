package pkg_list


type List struct{
	List []interface{}
	now int	//指向未使用的元素
}
func New()List{
	temp := List{}
	temp.List = make([]interface{},1)
	return temp
}
func(this *List) Append(data interface{}){
	this.List[this.now] = data
	this.now ++
	return
}
func(this *List)Out()interface{}{	//First input last output
	//全部前移一位
	if this.now == 0{
		return false
	}
	data := this.List[0]
	for n := 0;n < len(this.List);n++{
		this.List[n] = this.List[n + 1]
	}
	return data
}