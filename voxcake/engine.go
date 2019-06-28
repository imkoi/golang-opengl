package voxcake

import "github.com/go-gl/gl/v4.1-core/gl"

var Engine EngineStruct

type EngineStruct struct {
	components []IComponent
}

func init() {
	Window = *NewWindow("VoxCake", 1280, 720)
	Camera = *NewCamera()
	Program = NewProgram(Vs, Fs)
	Engine.components = make([]IComponent, 0)
}

func (engine *EngineStruct) AddComponent(component IComponent) {
	engine.components = append(engine.components, component)
}

// Run is methid that execute the program
func (engine *EngineStruct) Run() {
	engine.execStart()
	engine.execUpdate()
}

// Quit is methid that stop the program
func (engine *EngineStruct) Quit() {
	Window.isOpen = false
}

func (engine *EngineStruct) execStart() {
	for c := 0; c < len(engine.components); c++ {
		engine.components[c].Start()
	}
}

func (engine *EngineStruct) execUpdate() {
	for Window.IsOpen() {
		Window.pollEvent()
		gl.UseProgram(Program)

		Camera.Update()
		for c := 0; c < len(engine.components); c++ {
			engine.components[c].Update()
		}

		Window.glSwap()
	}
}
