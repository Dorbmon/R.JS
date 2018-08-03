package pkg_gtk

import (
	"github.com/Dorbmon/otto"
	"github.com/mattn/go-gtk/gtk"
	"os"
	)

type Pkg_gtk struct{
	js *otto.Otto

}

func(this Pkg_gtk)SwapJS(js *otto.Otto){
	this.js = js
	gtk.Init(&os.Args)
	js.Set("ImportGTKConst",func()otto.Value{
		obj,_ := js.Object("({})")
		obj.Set("WIN_POS_CENTER",gtk.WIN_POS_CENTER)
		obj.Set("WINDOW_TOPLEVEL",gtk.WINDOW_TOPLEVEL)
		value,_ := otto.ToValue(obj)
		return value
	})
	js.Set("NewWindow",func(call otto.FunctionCall)otto.Value{
		if call.Argument(0).IsNull() || !call.Argument(0).IsNumber(){
			return otto.FalseValue()
		}
		WindowType,_ := call.Argument(0).ToInteger()
		window := gtk.NewWindow(gtk.WindowType(WindowType))
		obj,_ := js.Object("({})")
		obj.Set("_window",window)
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
		obj.Set("Show",func(){
			window.Show()
		})

		value,_ := otto.ToValue(obj)
		return value
	})
	js.Set("NewLayout",func(call otto.FunctionCall)otto.Value{
		
	})
	js.Set("GTKMain",func(){
		gtk.Main()
	})



}