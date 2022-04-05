package main

import (
	"fmt"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

type Position3 struct {
	X float32
	Y float32
	Z float32
}

type TriangleVertex struct {
	Position Position3
	Color    Position3
}

type OpenGLMeshData struct {
}

type Material struct {
	shader        GLShaderProgram
	vertexArray   VertexArray
	vertexBuffers []BufferObject
}

func NewMaterial(shader GLShaderProgram, vertexArray VertexArray, vertexBuffers []BufferObject) Material {

	for bindingIndex, buffer := range vertexBuffers {
		for _, element := range buffer.layout.elements {
			if location := gl.GetAttribLocation(shader.GetId(), gl.Str(element.Name+"\x00")); location != -1 {
				shaderAttr := shaderTypeLayouts[element.ShaderDataType]
				gl.EnableVertexArrayAttrib(uint32(vertexArray.GetId()), uint32(location))
				gl.VertexArrayAttribFormat(uint32(vertexArray.GetId()), uint32(location), shaderAttr.n, shaderAttr.glType, false, element.Offset)
				gl.VertexArrayAttribBinding(uint32(vertexArray.GetId()), uint32(location), uint32(bindingIndex))
			} else {
				fmt.Println("Location for %v not found.", element.Name)
			}
		}
	}

	return Material{
		shader:        shader,
		vertexArray:   vertexArray,
		vertexBuffers: vertexBuffers,
	}
}

type OpenGLMesh struct {
	Vertices []TriangleVertex
	Indices  []uint32

	verticesSize   int
	verticesOffset int
	indicesSize    int
	indicesOffset  int

	vertexArray  VertexArray
	vertexBuffer BufferObject

	material Material
}

func NewOpenGLMesh(vertices []TriangleVertex, indices []uint32, material Material) OpenGLMesh {
	vertexSize := uint64(unsafe.Sizeof(vertices[0]))
	verticesSize := int(vertexSize) * len(vertices)

	mesh := OpenGLMesh{
		Vertices:       vertices,
		Indices:        indices,
		verticesSize:   verticesSize,
		verticesOffset: 0,
		indicesSize:    4 * len(indices),
		indicesOffset:  verticesSize,
		material:       material,
	}

	vao := CreateVertexArrayObject()
	vbo := CreateBufferObject(layout, mesh)

	mesh.vertexArray = vao
	mesh.vertexBuffer = vbo

	return mesh
}

func (mesh OpenGLMesh) GetVerticesSize() int {
	return mesh.verticesSize
}

func (mesh OpenGLMesh) GetVerticesOffset() int {
	return mesh.verticesOffset
}

func (mesh OpenGLMesh) GetIndicesSize() int {
	return mesh.indicesSize
}

func (mesh OpenGLMesh) GetIndicesOffset() int {
	return mesh.indicesOffset
}
