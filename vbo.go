package main

import (
	"unsafe"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type BufferObjectID uint32

type IBufferObject interface {
	Bind()
}

type BufferObject struct {
	id     BufferObjectID
	layout BufferLayout
}

func CreateBufferObject(layout BufferLayout, mesh OpenGLMesh) BufferObject {
	bufferObject := BufferObject{
		id:     generateBufferObjectObject(),
		layout: layout,
	}

	meshData := mesh.GetData()
	verticesSize := mesh.GetVerticesSize()
	indicesSize := mesh.GetIndicesSize()

	gl.NamedBufferStorage(uint32(bufferObject.id), verticesSize+indicesSize, gl.Ptr(meshData.Vertices), gl.DYNAMIC_STORAGE_BIT)
	gl.NamedBufferSubData(uint32(bufferObject.id), mesh.verticesOffset, verticesSize, gl.Ptr(meshData.Vertices))
	gl.NamedBufferSubData(uint32(bufferObject.id), mesh.indicesOffset, indicesSize, gl.Ptr(meshData.Indices))

	return bufferObject
}

func (vbo BufferObject) Delete() {
	gl.DeleteBuffers(1, (*uint32)(&vbo.id))
}

func (vbo BufferObject) BufferData(data unsafe.Pointer, dataSize int, usage uint32) {
	gl.NamedBufferData(uint32(vbo.id), dataSize, data, usage)
}

func generateBufferObjectObject() BufferObjectID {
	var id uint32
	gl.CreateBuffers(1, &id)

	return BufferObjectID(id)
}

func (vbo BufferObject) Id() BufferObjectID {
	return vbo.id
}

func (vbo BufferObject) GetLayout() *BufferLayout {
	return &vbo.layout
}
