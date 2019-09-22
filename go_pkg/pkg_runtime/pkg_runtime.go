package pkg_runtime

import "github.com/Dorbmon/otto"

type Pkg_runtime struct {
	js *otto.Otto
}

func (this Pkg_runtime) SwapJs (Engine *otto.Otto) {
	this.js = Engine
	Engine.Set ("ExitRJSProgram",func () {
		Engine.Runtime.Pause = true
	})
}
