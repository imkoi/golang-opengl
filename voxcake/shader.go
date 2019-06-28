package voxcake

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

var Program uint32

const (
	Vs = `
		#version 410
		layout(location = 0) in vec3 vertexPosition;
		layout(location = 1) in vec3 vertexColor;
		out vec4 fragColor;

		uniform mat4 projection;
		uniform mat4 view;
		uniform mat4 model;

		void main() {
			gl_Position = projection * view * model * vec4(vertexPosition, 1.0);
			fragColor = vec4(vertexColor, 1.0);
		}
	` + "\x00"

	Fs = `
		#version 410
		in vec4 fragColor;
		out vec4 Color;

		void main() {
			Color = fragColor;
		}
	` + "\x00"
)

func NewProgram(vertexShaderSource string, fragmentShaderSource string) uint32 {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)

	if err != nil {
		panic(err)
	}
	return prog
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
