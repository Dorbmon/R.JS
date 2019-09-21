package pkg_gtk

import (
	"github.com/Dorbmon/otto"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"os"
	"github.com/mattn/go-gtk/glib"
	)

type Pkg_gtk struct{
	js *otto.Otto
}

func(this Pkg_gtk)SwapJS(js *otto.Otto){
	this.js = js
	gtk.Init(&os.Args)
	js.Set("NewWindow",func(call otto.FunctionCall)otto.Value{
		if call.Argument(0).IsNull() || !call.Argument(0).IsNumber(){
			return otto.FalseValue()
		}
		WindowType,_ := call.Argument(0).ToInteger()
		//window := gtk.NewWindow(gtk.WindowType(WindowType))
		window := gtk.NewWindow(gtk.WindowType(WindowType))
		obj,_ := js.Object("({})")
		obj.Set("_window",window)
		obj.Set("_gtk",true)
		obj.Set("_obj",window)
		obj.Set("_gtkname","window")
		obj.Set("SetIconFromFile",func(call otto.FunctionCall)otto.Value{
			Address,err := call.Argument(0).ToString()
			if err != nil{
				return otto.FalseValue()
			}
			window.SetIconFromFile(Address)
			return otto.TrueValue()
		})
		obj.Set("SetPosition",func(call otto.FunctionCall)otto.Value{
			if call.Argument(0).IsNull() || !call.Argument(0).IsNumber(){
				return otto.FalseValue()
			}
			Position,_ := call.Argument(0).ToInteger()
			window.SetPosition(gtk.WindowPosition(Position))
			return otto.TrueValue()
		})
		obj.Set("SetTitle",func(call otto.FunctionCall)otto.Value{
			if call.Argument(0).IsNull() == true{
				return otto.FalseValue()
			}
			title := call.Argument(0).String()
			window.SetTitle(title)
			return otto.TrueValue()
		})
		obj.Set("SetSizeRequest",func(call otto.FunctionCall)otto.Value{
			if call.Argument(0).IsNull() == true || call.Argument(1).IsNull() == true{
				return otto.FalseValue()
			}
			if call.Argument(0).IsNumber() != true || call.Argument(1).IsNumber() != true{
				return otto.FalseValue()
			}
			temp,_ := call.Argument(0).ToInteger()
			x := int(temp)
			temp,_ = call.Argument(1).ToInteger()
			y := int(temp)
			window.SetSizeRequest(x,y)
			return otto.TrueValue()
		})
		obj.Set("Add",func(call otto.FunctionCall)otto.Value{
			Context := call.Argument(0)
			if !IfGTK(Context){
				return otto.FalseValue()
			}
			gtkobj,err := Context.Object().Get("_obj")
			if err != nil{
				return otto.FalseValue()
			}
			obj,err := gtkobj.Export()
			if err != nil{
				value,_ := otto.ToValue(err)
				return value
			}
			switch obj.(type){
			case *gtk.Fixed:
				window.Add(obj.(*gtk.Fixed))
			case *gtk.HBox:
				window.Add(obj.(*gtk.HBox))
			default:
				return otto.FalseValue()
			}
			return otto.TrueValue()

		})
		obj.Set("ShowAll",func(){
			window.ShowAll()
		})
		obj.Set("GetSizeRequest",func()otto.Value{
			x,y := window.GetSizeRequest()
			result,_ := js.Object("({})")
			result.Set("x",x)
			result.Set("y",y)
			value,_ := otto.ToValue(result)
			return value
		})
		obj.Set("Connect",func(call otto.FunctionCall)otto.Value{
			event,err := call.Argument(0).ToString()
			if err != nil{
				return otto.FalseValue()
			}
			CallbackFunction := call.Argument(1)
			if !CallbackFunction.IsFunction(){
				return otto.FalseValue()
			}
			if GetArgumentNumber(call) == 2{	//无参数类型
				window.Connect(event, func(Function *glib.CallbackContext) {
					Function.Data().(otto.Value).Call(Function.Data().(otto.Value))
					return
				},CallbackFunction)
			}else{
				OtherDatas := make([]interface{},1)
				OtherDatas[0] = CallbackFunction
				for n := 2;n < GetArgumentNumber(call);n++{
					OtherDatas[n - 1] = call.Argument(n)
				}
				window.Connect(event,func(datas ...interface{}){
					datas[0].(otto.Value).Call(datas[0].(otto.Value),datas[1:]...)
					return
				},OtherDatas...)
			}
			return otto.TrueValue()
		})
		obj.Set("Show",func(){
			window.Show()
		})

		value,_ := otto.ToValue(obj)
		return value
	})
	js.Set("NewGtkFixed",func(call otto.FunctionCall)otto.Value{
		fixed := gtk.NewFixed()
		obj,_ := js.Object("({})")
		obj.Set("_gtkname","gtklayout")
		obj.Set("_gtk",true)
		obj.Set("_obj",fixed)
		obj.Set("layout",fixed)
		obj.Set("SetSizeRequest",func(call otto.FunctionCall)otto.Value{
			if call.Argument(0).IsNull() == true || call.Argument(1).IsNull() == true{
				return otto.FalseValue()
			}
			if call.Argument(0).IsNumber() != true || call.Argument(1).IsNumber() != true{
				return otto.FalseValue()
			}
			temp,_ := call.Argument(0).ToInteger()
			x := int(temp)
			temp,_ = call.Argument(1).ToInteger()
			y := int(temp)
			fixed.SetSizeRequest(x,y)
			return otto.TrueValue()
		})
		obj.Set("Put",func(call otto.FunctionCall)otto.Value{
			gtkobj := call.Argument(0).Object()
			if !IfGTK(call.Argument(0)){
				return otto.FalseValue()
			}
			if call.Argument(1).IsNull() == true || call.Argument(2).IsNull() == true{
				return otto.FalseValue()
			}
			if call.Argument(1).IsNumber() != true || call.Argument(2).IsNumber() != true{
				return otto.FalseValue()
			}
			temp,_ := call.Argument(0).ToInteger()
			x := int(temp)
			temp,_ = call.Argument(1).ToInteger()
			y := int(temp)
			obj,err := gtkobj.Get("_obj")
			if err != nil{
				return otto.FalseValue()
			}
			robj,err := obj.Export()
			if err != nil{
				return otto.FalseValue()
			}

			switch robj.(type) {
			case *gtk.Button:
				fixed.Put(robj.(*gtk.Button),x,y)
			case *gtk.Label:
				fixed.Put(robj.(*gtk.Label),x,y)
			case *gtk.Image:
				fixed.Put(robj.(*gtk.Image),x,y)
			case *gtk.ProgressBar:
				fixed.Put(robj.(*gtk.ProgressBar),x,y)
			case *gtk.Entry:
				fixed.Put(robj.(*gtk.Entry),x,y)
			default:
				return otto.FalseValue()
			}
			return otto.TrueValue()
		})
		obj.Set("Move",func(call otto.FunctionCall)otto.Value{
			gtkobj := call.Argument(0).Object()
			if !IfGTK(call.Argument(0)){
				return otto.FalseValue()
			}
			if call.Argument(1).IsNull() == true || call.Argument(2).IsNull() == true{
				return otto.FalseValue()
			}
			if call.Argument(1).IsNumber() != true || call.Argument(2).IsNumber() != true{
				return otto.FalseValue()
			}
			temp,_ := call.Argument(0).ToInteger()
			x := int(temp)
			temp,_ = call.Argument(1).ToInteger()
			y := int(temp)
			obj,err := gtkobj.Get("_obj")
			if err != nil{
				return otto.FalseValue()
			}
			robj,err := obj.Export()
			if err != nil{
				return otto.FalseValue()
			}

			switch robj.(type) {
			case *gtk.Button:
				fixed.Move(robj.(*gtk.Button),x,y)
			case *gtk.Label:
				fixed.Move(robj.(*gtk.Label),x,y)
			case *gtk.Image:
				fixed.Move(robj.(*gtk.Image),x,y)
			case *gtk.ProgressBar:
				fixed.Move(robj.(*gtk.ProgressBar),x,y)
			case *gtk.Entry:
				fixed.Move(robj.(*gtk.Entry),x,y)
			default:
				return otto.FalseValue()
			}
			return otto.TrueValue()
		})
		value,_ := otto.ToValue(obj)
		return value
	})
	js.Set("NewLable",func(call otto.FunctionCall)otto.Value{
		label_name,err := call.Argument(0).ToString()
		if err != nil{
			return otto.FalseValue()
		}
		label := gtk.NewLabel(label_name)
		obj,_ := js.Object("({})")
		obj.Set("_obj",label)
		obj.Set("_gtkname","label")
		obj.Set("_gtk",true)
		obj.Set("SetSizeRequest",func(call otto.FunctionCall)otto.Value{
			if call.Argument(0).IsNull() == true || call.Argument(1).IsNull() == true{
				return otto.FalseValue()
			}
			if call.Argument(0).IsNumber() != true || call.Argument(1).IsNumber() != true{
				return otto.FalseValue()
			}
			temp,_ := call.Argument(0).ToInteger()
			x := int(temp)
			temp,_ = call.Argument(1).ToInteger()
			y := int(temp)
			label.SetSizeRequest(x,y)
			return otto.TrueValue()
		})
		obj.Set("Connect",func(call otto.FunctionCall)otto.Value{
			event,err := call.Argument(0).ToString()
			if err != nil{
				return otto.FalseValue()
			}
			CallbackFunction := call.Argument(1)
			if !CallbackFunction.IsFunction(){
				return otto.FalseValue()
			}
			if GetArgumentNumber(call) == 2{	//无参数类型
				label.Connect(event, func(Function *glib.CallbackContext) {
					Function.Data().(otto.Value).Call(Function.Data().(otto.Value))
					return
				},CallbackFunction)
			}else{
				OtherDatas := make([]interface{},1)
				OtherDatas[0] = CallbackFunction
				for n := 2;n < GetArgumentNumber(call);n++{OtherDatas[n - 1] = call.Argument(n)
				}
				label.Connect(event,func(datas *glib.CallbackContext){
					//datas[0].(otto.Value).Call(datas[0].(otto.Value),datas[1:]...)
					data := datas.Data().([]interface{})
					data[0].(otto.Value).Call(data[0].(otto.Value),data[1:])
					return
				},OtherDatas)
			}
			return otto.TrueValue()
		})
		obj.Set("SetText",func(call otto.FunctionCall)otto.Value{
			lable_name,err := call.Argument(0).ToString()
			if err != nil{
				return otto.FalseValue()
			}
			label.SetText(lable_name)
			return otto.TrueValue()
		})
		obj.Set("GetSizeRequest",func()otto.Value{
			x,y := label.GetSizeRequest()
			result,_ := js.Object("({})")
			result.Set("x",x)
			result.Set("y",y)
			value,_ := otto.ToValue(result)
			return value
		})
		value,_ := otto.ToValue(obj)
		return value
	})
	js.Set("NewProgressBar",func()otto.Value{
		ProgressBar := gtk.NewProgressBar()
		obj,_ := js.Object("({})")
		obj.Set("_gtk",true)
		obj.Set("_obj",ProgressBar)
		obj.Set("_gtkname","progressbar")
		obj.Set("SetFraction",func(call otto.FunctionCall)otto.Value{
			value,err := call.Argument(0).ToFloat()
			if err != nil{
				return otto.FalseValue()
			}
			if value > 1 || value < 0{
				return otto.FalseValue()
			}
			ProgressBar.SetFraction(value)
			return otto.TrueValue()
		})
		obj.Set("GetFraction",func()otto.Value{
			value,_ := otto.ToValue(ProgressBar.GetFraction())
			return value
		})
		obj.Set("SetText",func(call otto.FunctionCall)otto.Value{
			Text,err := call.Argument(0).ToString()
			if err != nil{
				return otto.FalseValue()
			}
			ProgressBar.SetText(Text)
			return otto.TrueValue()
		})
		obj.Set("GetText",func()otto.Value{
			value,_ := otto.ToValue(ProgressBar.GetText())
			return value
		})
		obj.Set("SetSizeRequest",func(call otto.FunctionCall)otto.Value{
			if call.Argument(0).IsNull() == true || call.Argument(1).IsNull() == true{
				return otto.FalseValue()
			}
			if call.Argument(0).IsNumber() != true || call.Argument(1).IsNumber() != true{
				return otto.FalseValue()
			}
			temp,_ := call.Argument(0).ToInteger()
			x := int(temp)
			temp,_ = call.Argument(1).ToInteger()
			y := int(temp)
			ProgressBar.SetSizeRequest(x,y)
			return otto.TrueValue()
		})
		obj.Set("GetSize",func()otto.Value{
			x,y := ProgressBar.GetSizeRequest()
			result,_ := js.Object("({})")
			result.Set("x",x)
			result.Set("y",y)
			value,_ := otto.ToValue(result)
			return value
		})
		obj.Set("Connect",func(call otto.FunctionCall)otto.Value{
			event,err := call.Argument(0).ToString()
			if err != nil{
				return otto.FalseValue()
			}
			CallbackFunction := call.Argument(1)
			if !CallbackFunction.IsFunction(){
				return otto.FalseValue()
			}
			if GetArgumentNumber(call) == 2{	//无参数类型
				ProgressBar.Connect(event, func(Function otto.Value) {
					Function.Call(Function)
					return
				},CallbackFunction)
			}else{
				OtherDatas := make([]interface{},1)
				OtherDatas[0] = CallbackFunction
				for n := 2;n < GetArgumentNumber(call);n++{
					OtherDatas[n - 1] = call.Argument(n)
				}
				ProgressBar.Connect(event,func(datas ...interface{}){
					datas[0].(otto.Value).Call(datas[0].(otto.Value),datas[1:]...)
					return
				},OtherDatas...)
			}
			return otto.TrueValue()
		})
		value,_ := otto.ToValue(obj)
		return value
	})
	js.Set("NewImage",func()otto.Value{
		image := gtk.NewImage()
		obj,_ := js.Object("({})")
		obj.Set("_obj",image)
		obj.Set("_gtkname","image")
		obj.Set("_gtk",true)
		obj.Set("SetSizeRequest",func(call otto.FunctionCall)otto.Value{
			if call.Argument(0).IsNull() == true || call.Argument(1).IsNull() == true{
				return otto.FalseValue()
			}
			if call.Argument(0).IsNumber() != true || call.Argument(1).IsNumber() != true{
				return otto.FalseValue()
			}
			temp,_ := call.Argument(0).ToInteger()
			x := int(temp)
			temp,_ = call.Argument(1).ToInteger()
			y := int(temp)
			image.SetSizeRequest(x,y)
			return otto.TrueValue()
		})
		obj.Set("GetSizeRequest",func()otto.Value{
			x,y := image.GetSizeRequest()
			result,_ := js.Object("({})")
			result.Set("x",x)
			result.Set("y",y)
			value,_ := otto.ToValue(result)
			return value
		})
		obj.Set("Connect",func(call otto.FunctionCall)otto.Value{
			event,err := call.Argument(0).ToString()
			if err != nil{
				return otto.FalseValue()
			}
			CallbackFunction := call.Argument(1)
			if !CallbackFunction.IsFunction(){
				return otto.FalseValue()
			}
			if GetArgumentNumber(call) == 2{	//无参数类型
				image.Connect(event, func(Function otto.Value) {
					Function.Call(Function)
					return
				},CallbackFunction)
			}else{
				OtherDatas := make([]interface{},1)
				OtherDatas[0] = CallbackFunction
				for n := 2;n < GetArgumentNumber(call);n++{
					OtherDatas[n - 1] = call.Argument(n)
				}
				image.Connect(event,func(datas ...interface{}){
					datas[0].(otto.Value).Call(datas[0].(otto.Value),datas[1:]...)
					return
				},OtherDatas...)
			}
			return otto.TrueValue()
		})
		obj.Set("SetFromPixbuf",func(call otto.FunctionCall)otto.Value{
			if call.Argument(0).IsNull(){
				return otto.FalseValue()
			}
			if !IfGDK(call.Argument(0)){
				return otto.FalseValue()
			}
			pixbuf := call.Argument(0).Object()
			value,_ := pixbuf.Get("_gdkname")
			if value.String() != "pixbuf"{
				return otto.FalseValue()
			}
			rpixbuf,err := pixbuf.Get("_obj")
			if err != nil{
				return otto.FalseValue()
			}
			data,err := rpixbuf.Export()
			if err != nil{
				return otto.FalseValue()
			}
			image.SetFromPixbuf(data.(*gdkpixbuf.Pixbuf))
			return otto.TrueValue()
		})
		value,_ := otto.ToValue(obj)
		return value
	})
	js.Set("NewPixbuf",func(call otto.FunctionCall)otto.Value{
		Address,err := call.Argument(0).ToString()
		if err != nil{
			return otto.FalseValue()
		}
		//判断参数个数
		var pixbuf *gdkpixbuf.Pixbuf
		obj,_ := js.Object("({})")
		obj.Set("_gdk",true)
		obj.Set("_gdkname","pixbuf")
		if GetArgumentNumber(call) != 1{
			//需要确定大小的版本
			x,err := call.Argument(1).ToInteger()
			if err != nil{
				return otto.FalseValue()
			}
			y,err := call.Argument(2).ToInteger()
			if err != nil{
				return otto.FalseValue()
			}
			SavePastSize,err := call.Argument(3).ToBoolean()
			if err != nil{
				return otto.FalseValue()
			}
			var err2 *glib.Error
			pixbuf,err2 = gdkpixbuf.NewPixbufFromFileAtScale(Address,int(x),int(y),SavePastSize)
			if err2 != nil{
				return otto.FalseValue()
			}
			obj.Set("_obj",pixbuf)
		}else{
			//Only File
			var err *glib.Error
			pixbuf,err = gdkpixbuf.NewPixbufFromFile(Address)
			if err != nil{
				//fmt.Println(err.Error())
				return otto.FalseValue()
			}
			obj.Set("_obj",pixbuf)
		}
		obj.Set("Unref",func(){
			pixbuf.Unref()
		})
		value,_ := otto.ToValue(obj)
		return value
	})
	js.Set("NewButton",func()otto.Value{
		button := gtk.NewButton()
		obj,_ := js.Object("({})")
		obj.Set("_obj",button)
		obj.Set("_gtkname","button")
		obj.Set("_gtk",true)
		obj.Set("SetLable",func(call otto.FunctionCall)otto.Value{
			LableName,err := call.Argument(0).ToString()
			if err != nil{
				return otto.FalseValue()
			}
			button.SetLabel(LableName)
			return otto.TrueValue()
		})
		obj.Set("SetSizeRequest",func(call otto.FunctionCall)otto.Value{
			if call.Argument(0).IsNull() == true || call.Argument(1).IsNull() == true{
				return otto.FalseValue()
			}
			if call.Argument(0).IsNumber() != true || call.Argument(1).IsNumber() != true{
				return otto.FalseValue()
			}
			temp,_ := call.Argument(0).ToInteger()
			x := int(temp)
			temp,_ = call.Argument(1).ToInteger()
			y := int(temp)
			button.SetSizeRequest(x,y)
			return otto.TrueValue()
		})
		obj.Set("Connect",func(call otto.FunctionCall)otto.Value{
			event,err := call.Argument(0).ToString()
			if err != nil{
				return otto.FalseValue()
			}
			CallbackFunction := call.Argument(1)
			if !CallbackFunction.IsFunction(){
				return otto.FalseValue()
			}
			if GetArgumentNumber(call) == 2{	//无参数类型
				button.Connect(event, func(Function *glib.CallbackContext) {
					Function.Data().(otto.Value).Call(Function.Data().(otto.Value))
					return
				},CallbackFunction)
			}else{
				OtherDatas := make([]interface{},1)
				OtherDatas[0] = CallbackFunction
				for n := 2;n < GetArgumentNumber(call);n++{OtherDatas[n - 1] = call.Argument(n)
				}
				button.Connect(event,func(datas *glib.CallbackContext){
					//datas[0].(otto.Value).Call(datas[0].(otto.Value),datas[1:]...)
					data := datas.Data().([]interface{})
					data[0].(otto.Value).Call(data[0].(otto.Value),data[1:])
					return
				},OtherDatas)
			}
			return otto.TrueValue()
		})
		obj.Set("SetSensitive",func(call otto.FunctionCall)otto.Value{
			Sensitive,err := call.Argument(0).ToBoolean()
			if err != nil{
				return otto.FalseValue()
			}
			button.SetSensitive(Sensitive)
			return otto.TrueValue()
		})
		obj.Set("SetImage",func(call otto.FunctionCall)otto.Value{
			if !call.Argument(0).IsObject(){
				return otto.FalseValue()
			}
			temp1 := call.Argument(0).Object()
			if !IfGTK(call.Argument(0)){
				return otto.FalseValue()
			}
			if v,err := temp1.Get("_gtkname");err != nil || v.String() != "image"{
				return otto.FalseValue()
			}
			temp2,err := temp1.Get("_obj")
			if err != nil{
				return otto.FalseValue()
			}
			temp3,err := temp2.Export()
			if err != nil{
				return otto.FalseValue()
			}
			button.SetImage(temp3.(*gtk.Image))
			return otto.TrueValue()
		})
		obj.Set("GetSizeRequest",func()otto.Value{
			x,y := button.GetSizeRequest()
			result,_ := js.Object("({})")
			result.Set("x",x)
			result.Set("y",y)
			value,_ := otto.ToValue(result)
			return value
		})
		value,_ := otto.ToValue(obj)
		return value
	})
	js.Set("NewEntry",func()otto.Value{
		obj,_ := js.Object("({})")
		Entry := gtk.NewEntry()
		obj.Set("_gtk",true)
		obj.Set("_obj",Entry)
		obj.Set("_gtkname","entry")
		obj.Set("GetSizeRequest",func()otto.Value{
			x,y := Entry.GetSizeRequest()
			result,_ := js.Object("({})")
			result.Set("x",x)
			result.Set("y",y)
			value,_ := otto.ToValue(result)
			return value
		})
		obj.Set("SetText",func(call otto.FunctionCall)otto.Value{
			Text,err := call.Argument(0).ToString()
			if err != nil{
				return otto.FalseValue()
			}
			Entry.SetText(Text)
			return otto.TrueValue()
		})
		obj.Set("SetVisibility",func(call otto.FunctionCall)otto.Value{
			Visibility,err := call.Argument(0).ToBoolean()
			if err != nil{
				return otto.FalseValue()
			}
			Entry.SetVisibility(Visibility)
			return otto.TrueValue()
		})
		obj.Set("SetEditable",func(call otto.FunctionCall)otto.Value{
			Visibility,err := call.Argument(0).ToBoolean()
			if err != nil{
				return otto.FalseValue()
			}
			Entry.SetEditable(Visibility)
			return otto.TrueValue()
		})
		obj.Set("SetSizeRequest",func(call otto.FunctionCall)otto.Value{
			if call.Argument(0).IsNull() == true || call.Argument(1).IsNull() == true{
				return otto.FalseValue()
			}
			if call.Argument(0).IsNumber() != true || call.Argument(1).IsNumber() != true{
				return otto.FalseValue()
			}
			temp,_ := call.Argument(0).ToInteger()
			x := int(temp)
			temp,_ = call.Argument(1).ToInteger()
			y := int(temp)
			Entry.SetSizeRequest(x,y)
			return otto.TrueValue()
		})
		obj.Set("Connect",func(call otto.FunctionCall)otto.Value{
			event,err := call.Argument(0).ToString()
			if err != nil{
				return otto.FalseValue()
			}
			CallbackFunction := call.Argument(1)
			if !CallbackFunction.IsFunction(){
				return otto.FalseValue()
			}
			if GetArgumentNumber(call) == 2{	//无参数类型
				Entry.Connect(event, func(Function *glib.CallbackContext) {
					Function.Data().(otto.Value).Call(Function.Data().(otto.Value))
					return
				},CallbackFunction)
			}else{
				OtherDatas := make([]interface{},1)
				OtherDatas[0] = CallbackFunction
				for n := 2;n < GetArgumentNumber(call);n++{
					OtherDatas[n - 1] = call.Argument(n)
				}
				Entry.Connect(event,func(datas ...interface{}){
					datas[0].(otto.Value).Call(datas[0].(otto.Value),datas[1:]...)
					return
				},OtherDatas...)
			}
			return otto.TrueValue()
		})
		obj.Set("GetText",func()otto.Value{
			value,_ := otto.ToValue(Entry.GetText())
			return value
		})
		value,_ := otto.ToValue(obj)
		return value
	})
	js.Set("NewHBox",func(call otto.FunctionCall)otto.Value{
		homogeneous,err := call.Argument(0).ToBoolean()
		if err != nil{
			return otto.FalseValue()
		}
		spacing,err := call.Argument(1).ToInteger()
		if err != nil{
			return otto.FalseValue()
		}
		HBox := gtk.NewHBox(homogeneous,int(spacing))
		obj,_ := js.Object("({})")
		obj.Set("_gtk",true)
		obj.Set("_obj",HBox)
		obj.Set("_gtkname","hbox")

		/*homogeneous是一个布尔值，为TRUE时，强制盒中的构件都占用相同大小的空间，不管每个空间的大小。
		spacing是以像素为单位设置的构件之间的间距。*/
	})
	js.Set("MainQuit",func(){
		gtk.MainQuit()
	})
	js.Set("GTKMain",func(){
		gtk.Main()
	})
	js.Set("ImportGTKConst",func()otto.Value{
		obj,_ := js.Object("({})")
		obj.Set("WIN_POS_CENTER",gtk.WIN_POS_CENTER)
		obj.Set("WINDOW_TOPLEVEL",gtk.WINDOW_TOPLEVEL)
		obj.Set("GTK_QUITE",gtk.MainQuit)
		value,_ := otto.ToValue(obj)
		return value
	})
	return
}
func IfGTK(object otto.Value)bool{
	obj := object.Object()
	value,err := obj.Get("_gtk")
	if err != nil{
		return false
	}
	if !value.IsBoolean(){
		return false
	}
	bvalue,_ := value.ToBoolean()
	return bvalue
}
func IfGDK(object otto.Value)bool{
	obj := object.Object()
	value,err := obj.Get("_gdk")
	if err != nil{
		return false
	}
	if !value.IsBoolean(){
		return false
	}
	bvalue,_ := value.ToBoolean()
	return bvalue
}
func GetArgumentNumber(call otto.FunctionCall)int{
	n := 0
	for ;;n++{
		if call.Argument(n).IsNull() || call.Argument(n).IsUndefined(){
			break
		}
	}
	return n
}