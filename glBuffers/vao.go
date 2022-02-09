package glBuffers

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

type VertexArrayObjectID uint32

type VertexArrayObject struct {
	id      VertexArrayObjectID
	layout  *BufferLayout
	buffers []*BufferObject
}

func CreateVertexArrayObject() VertexArrayObject {
	return VertexArrayObject{
		id: generateVertexArray(),
	}
}

func (vao *VertexArrayObject) Delete() {
	gl.DeleteVertexArrays(1, (*uint32)(&vao.id))
}

func (vao *VertexArrayObject) GetId() VertexArrayObjectID {
	return vao.id
}

func generateVertexArray() VertexArrayObjectID {
	var id uint32
	gl.CreateVertexArrays(1, &id)

	return VertexArrayObjectID(id)
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

func (vao *VertexArrayObject) AddBufferObject(buffer *BufferObject) {

	vbBindingIdx := uint32(len(vao.buffers))

	vao.buffers = append(vao.buffers, buffer)

	gl.VertexArrayElementBuffer(uint32(vao.id), uint32(buffer.Id()))
	gl.VertexArrayVertexBuffer(uint32(vao.id), vbBindingIdx, uint32(buffer.Id()), 0, buffer.GetLayout().GetStride())
	bufferElements := buffer.GetLayout().GetElements()

	for idx, element := range bufferElements {
		shaderAttr := shaderTypeLayouts[element.ShaderDataType]
		gl.EnableVertexArrayAttrib(uint32(vao.id), uint32(idx))
		gl.VertexArrayAttribFormat(uint32(vao.id), uint32(idx), shaderAttr.n, shaderAttr.glType, false, uint32(element.Offset))
		gl.VertexArrayAttribBinding(uint32(vao.id), uint32(idx), vbBindingIdx)
	}
}
