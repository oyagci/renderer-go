package main

import (
	_ "embed"
	"fmt"
	"log"
	"runtime"
	"strings"
	"unsafe"

	"github.com/oyagci/renderer-go/program"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/oyagci/renderer-go/glBuffers"
	"github.com/oyagci/renderer-go/renderer"
)

//go:embed shaders/simple.vs.glsl
var VERTEX_SHADER_SRC string

//go:embed shaders/simple.fs.glsl
var FRAGMENT_SHADER_SRC string

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func glDebugLog(source uint32, glType uint32, id uint32, severity uint32,
	length int32, message string, userParam unsafe.Pointer) {

	srcStr := func(source uint32) string {
		switch source {
		case gl.DEBUG_SOURCE_API:
			return "API"
		case gl.DEBUG_SOURCE_WINDOW_SYSTEM:
			return "WINDOW SYSTEM"
		case gl.DEBUG_SOURCE_SHADER_COMPILER:
			return "SHADER COMPILER"
		case gl.DEBUG_SOURCE_THIRD_PARTY:
			return "THIRD PARTY"
		case gl.DEBUG_SOURCE_APPLICATION:
			return "APPLICATION"
		case gl.DEBUG_SOURCE_OTHER:
			return "OTHER"
		default:
			return ""
		}
	}(source)

	typeStr := func(glType uint32) string {
		switch glType {
		case gl.DEBUG_TYPE_ERROR:
			return "ERROR"
		case gl.DEBUG_TYPE_DEPRECATED_BEHAVIOR:
			return "DEPRECATED_BEHAVIOR"
		case gl.DEBUG_TYPE_UNDEFINED_BEHAVIOR:
			return "UNDEFINED_BEHAVIOR"
		case gl.DEBUG_TYPE_PORTABILITY:
			return "PORTABILITY"
		case gl.DEBUG_TYPE_PERFORMANCE:
			return "PERFORMANCE"
		case gl.DEBUG_TYPE_MARKER:
			return "MARKER"
		case gl.DEBUG_TYPE_OTHER:
			return "OTHER"
		default:
			return ""
		}
	}(glType)

	severityStr := func(severity uint32) string {
		switch severity {
		case gl.DEBUG_SEVERITY_NOTIFICATION:
			return "NOTIFICATION"
		case gl.DEBUG_SEVERITY_LOW:
			return "LOW"
		case gl.DEBUG_SEVERITY_MEDIUM:
			return "MEDIUM"
		case gl.DEBUG_SEVERITY_HIGH:
			return "HIGH"
		default:
			return ""
		}
	}(severity)

	log.Printf("[%v:%v:%v:%v] %v\n", srcStr, typeStr, severityStr, id, message)
}

func main() {

	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(800, 600, "Testing", nil, nil)

	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err = gl.Init(); err != nil {
		panic(err)
	}

	gl.Enable(gl.DEBUG_OUTPUT)
	gl.DebugMessageCallback(glDebugLog, nil)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.1, 0.1, 0.1, 1.0)

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	shaderProgramBuilder := program.NewGLShaderProgramBuilder()

	if err := shaderProgramBuilder.AddVertexShaderFromSource(VERTEX_SHADER_SRC); err != nil {
		log.Fatalf("%v", err)
	}

	if err := shaderProgramBuilder.AddFragmentShaderFromSource(FRAGMENT_SHADER_SRC); err != nil {
		log.Fatalf("%v", err)
	}

	shaderProgram, err := shaderProgramBuilder.Build()
	if err != nil {
		log.Fatalf("%v", err)
	}

	type Position3 struct {
		X float32
		Y float32
		Z float32
	}

	vertices := []renderer.TriangleVertex{
		{Position: renderer.Position3{X: -0.5, Y: -0.5, Z: 0.0}, Color: renderer.Position3{X: 1.0, Y: 0.0, Z: 0.0}},
		{Position: renderer.Position3{X: 0.5, Y: -0.5, Z: 0.0}, Color: renderer.Position3{X: 0.0, Y: 1.0, Z: 0.0}},
		{Position: renderer.Position3{X: 0.0, Y: 0.5, Z: 0.0}, Color: renderer.Position3{X: 0.0, Y: 0.0, Z: 1.0}},
	}
	indices := []uint32{
		0, 1, 2,
	}

	bufferLayout := glBuffers.NewBufferLayout([]glBuffers.BufferElement{
		{
			Name:           "position",
			ShaderDataType: glBuffers.Vector3f,
			Size:           3 * 4,
			Normalized:     false,
			Offset:         uint32(unsafe.Offsetof(vertices[0].Position)),
		},
		{
			Name:           "color",
			ShaderDataType: glBuffers.Vector3f,
			Size:           3 * 4,
			Normalized:     false,
			Offset:         uint32(unsafe.Offsetof(vertices[0].Color)),
		},
	})
	triangleMesh := renderer.NewOpenGLMesh(vertices, indices)
	bufferObject := glBuffers.CreateBufferObject(bufferLayout, triangleMesh)
	defer bufferObject.Delete()

	vertexArray := glBuffers.CreateVertexArrayObject()
	defer vertexArray.Delete()

	vertexArray.AddBufferObject(&bufferObject)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		program.UseGLProgram(*shaderProgram)

		gl.BindVertexArray(uint32(vertexArray.GetId()))
		gl.DrawElementsWithOffset(gl.TRIANGLES, int32(len(triangleMesh.GetData().Indices)), gl.UNSIGNED_INT, uintptr(triangleMesh.GetIndicesOffset()))

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func newProgram(vertexShaderSource string, fragmentShaderSource string) (uint32, error) {

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
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
