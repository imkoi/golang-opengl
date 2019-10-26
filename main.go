package main

import (
	. "./voxcake"
)

func main() {
	Window.CaptureMouse(true)
	Window.SetGrab(true)
	Window.SetRelativeMouseMode(true)

	cameraController := new(CameraController)
	volume := Volume("field.vxl")
	Engine.AddComponent(cameraController)
	Engine.AddComponent(volume)

	Engine.Run()
}
