package main

import (
	. "./voxcake"
)

func main() {
	Window.CaptureMouse(true)
	Window.SetGrab(true)
	Window.SetRelativeMouseMode(true)

	cameraController := new(CameraController)
	volume := Volume("Map.vcmap")
	Engine.AddComponent(cameraController)
	Engine.AddComponent(volume)

	Engine.Run()
}
