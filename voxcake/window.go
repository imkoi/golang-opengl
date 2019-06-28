package voxcake

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)

var Window windowStruct

type windowStruct struct {
	Title     string
	Width     int32
	Height    int32
	isOpen    bool
	sdlWindow *sdl.Window
	context   sdl.GLContext
}

func NewWindow(title string, width int32, height int32) *windowStruct {
	var window windowStruct
	var err error

	window.Title = title
	window.Width = width
	window.Height = height
	window.isOpen = true

	sdl.Init(sdl.INIT_EVERYTHING)
	window.sdlWindow, err = sdl.CreateWindow(window.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, window.Width, window.Height, sdl.WINDOW_OPENGL)
	window.context, err = window.sdlWindow.GLCreateContext()

	Time.lastTick = 0
	Time.DeltaTime = 0

	gl.Init()

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)
	gl.ClearColor(0.84, 0.93, 0.93, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)
	gl.Viewport(0, 0, width, height)

	if err != nil {
		panic(err)
	}
	return &window
}

func (window *windowStruct) GetPositionX() int32 {
	x, _ := window.sdlWindow.GetPosition()
	return x
}
func (window *windowStruct) GetPositionY() int32 {
	_, y := window.sdlWindow.GetPosition()
	return y
}

func (window *windowStruct) SetCursor(x int32, y int32) {
	sdl.WarpMouseGlobal(x, y)
}

func (window *windowStruct) CenterCursor(x int32, y int32) {
	window.sdlWindow.WarpMouseInWindow(x, y)
}

func (window *windowStruct) CaptureMouse(state bool) {
	sdl.CaptureMouse(state)
}

func (window *windowStruct) SetGrab(state bool) {
	window.sdlWindow.SetGrab(state)
}

func (window *windowStruct) SetRelativeMouseMode(state bool) {
	sdl.SetRelativeMouseMode(state)
}

func (window *windowStruct) ShowCursor(show bool) {
	if show {
		sdl.ShowCursor(1)
	} else {
		sdl.ShowCursor(0)
	}
}

func (window *windowStruct) SetFullscreen(state bool) {
	if state {
		window.sdlWindow.SetFullscreen(1)
	} else {
		window.sdlWindow.SetFullscreen(0)
	}
}

func (window *windowStruct) IsOpen() bool {
	return window.isOpen
}

func (window *windowStruct) glSwap() {
	window.sdlWindow.GLSwap()
}

func (window *windowStruct) pollEvent() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.MouseButtonEvent:
			Input.mouse.SetMouseButton(t.Button, t.State)
		case *sdl.MouseMotionEvent:
			Input.mouse.SetMousePosition(t.X, t.Y)
		case *sdl.KeyboardEvent:
			Input.keyboard.SetKeyboardInput(t.State, uint8(t.Keysym.Sym))
		case *sdl.QuitEvent:
			window.isOpen = false
		}
	}
	Time.setDeltaTime(sdl.GetTicks())
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (window *windowStruct) Destroy() {
	window.isOpen = false
	defer window.sdlWindow.Destroy()
	defer sdl.GLDeleteContext(window.context)
	defer sdl.Quit()
}
