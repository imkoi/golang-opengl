package voxcake

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	. "github.com/go-gl/mathgl/mgl32"
)

type MeshStruct struct {
	program      uint32
	vao          uint32
	vbo          uint32
	cbo          uint32
	ebo          uint32
	projectionID int32
	viewID       int32
	modelID      int32

	Vertices []float32
	Colors   []uint8
	Indices  []uint32
}

func Mesh(vertices []float32, colors []uint8, indices []uint32) *MeshStruct {
	var mesh MeshStruct
	mesh.Vertices = vertices
	mesh.Colors = colors
	mesh.Indices = indices
	return &mesh
}

func (mesh *MeshStruct) Start() {
	if len(mesh.Vertices) == 0 {
		return
	}
	gl.GenVertexArrays(1, &mesh.vao)
	gl.GenBuffers(1, &mesh.vbo)
	gl.GenBuffers(1, &mesh.cbo)
	gl.GenBuffers(1, &mesh.ebo)

	gl.BindVertexArray(mesh.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(mesh.Vertices), gl.Ptr(mesh.Vertices), gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.cbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(mesh.Vertices), gl.Ptr(mesh.Colors), gl.STATIC_DRAW)
	gl.VertexAttribPointer(1, 3, gl.UNSIGNED_BYTE, true, 0, nil)
	gl.EnableVertexAttribArray(1)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, mesh.ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(mesh.Indices), gl.Ptr(mesh.Indices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ARRAY_BUFFER, uint32(0))
	gl.BindVertexArray(uint32(0))

	mesh.program = Program

	mesh.projectionID = gl.GetUniformLocation(mesh.program, gl.Str("projection\x00"))
	mesh.viewID = gl.GetUniformLocation(mesh.program, gl.Str("view\x00"))
	mesh.modelID = gl.GetUniformLocation(mesh.program, gl.Str("model\x00"))
}

//Draw mesh
func (mesh *MeshStruct) Draw(model Mat4) {
	if len(mesh.Vertices) == 0 {
		return
	}

	gl.UniformMatrix4fv(mesh.projectionID, 1, false, &Camera.projection[0])
	gl.UniformMatrix4fv(mesh.viewID, 1, false, &Camera.view[0])
	gl.UniformMatrix4fv(mesh.modelID, 1, false, &model[0])

	gl.BindVertexArray(mesh.vao)
	gl.DrawElements(gl.TRIANGLES, int32(len(mesh.Indices)), gl.UNSIGNED_INT, nil)
}

func (mesh *MeshStruct) Update() {

}

func (mesh *MeshStruct) Name() string {
	return "Mesh"
}

func (mesh *MeshStruct) Return() interface{} {
	return mesh
}
