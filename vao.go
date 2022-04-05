package main

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

type VertexArrayId uint32

type VertexArray struct {
	id      VertexArrayId
	layout  *BufferLayout
	buffers []*BufferObject
}

func CreateVertexArrayObject() VertexArray {
	return VertexArray{
		id: generateVertexArray(),
	}
}

func (vao *VertexArray) Delete() {
	gl.DeleteVertexArrays(1, (*uint32)(&vao.id))
}

func (vao *VertexArray) GetId() VertexArrayId {
	return vao.id
}

func generateVertexArray() VertexArrayId {
	var id uint32
	gl.CreateVertexArrays(1, &id)

	return VertexArrayId(id)
}

type ShaderAttribFormat struct {
	n      int32
	glType uint32
	size   int
}

var shaderTypeLayouts = map[EShaderDataType]ShaderAttribFormat{
	Int32:    {1, gl.INT, 4},
	Float32:  {1, gl.FLOAT, 4},
	Float64:  {1, gl.DOUBLE, 8},
	Vector2f: {2, gl.FLOAT, 4 * 2},
	Vector3f: {3, gl.FLOAT, 4 * 3},
	Vector4f: {4, gl.FLOAT, 4 * 4},
}

func (vao *VertexArray) AddBufferObject(buffer *BufferObject) {

	bindingIndex := uint32(len(vao.buffers))
	vao.buffers = append(vao.buffers, buffer)

	gl.VertexArrayVertexBuffer(uint32(vao.id), bindingIndex, uint32(buffer.Id()), 0, buffer.GetLayout().GetStride())
}

func (vao VertexArray) AddElementBuffer(buffer BufferObject) {
	gl.VertexArrayElementBuffer(uint32(vao.id), uint32(buffer.Id()))
}
