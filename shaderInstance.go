package main

//func NewGLMesh() GLMesh {
//	for bindingIndex, buffer := range material.vertexBuffers {
//		for _, element := range buffer.layout.elements {
//			if location := gl.GetAttribLocation(material.shader.GetId(), gl.Str(element.Name+"\x00")); location != -1 {
//				shaderAttr := shaderTypeLayouts[element.ShaderDataType]
//				gl.EnableVertexArrayAttrib(uint32(material.vertexArray.GetId()), uint32(location))
//				gl.VertexArrayAttribFormat(uint32(material.vertexArray.GetId()), uint32(location), shaderAttr.n, shaderAttr.glType, false, element.Offset)
//				gl.VertexArrayAttribBinding(uint32(material.vertexArray.GetId()), uint32(location), uint32(bindingIndex))
//			}
//		}
//	}
//}
