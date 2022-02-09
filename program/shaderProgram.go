package program

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type GLShaderProgramBuilder struct {
	vertexShader   uint32
	fragmentShader uint32
}

type GLShaderProgram struct {
	id uint32
}

func NewGLShaderProgramBuilder() GLShaderProgramBuilder {
	return GLShaderProgramBuilder{}
}

func (builder *GLShaderProgramBuilder) AddVertexShaderFromSource(source string) error {
	shader, err := compileShader(source, gl.VERTEX_SHADER)

	if err != nil {
		return err
	}

	builder.vertexShader = shader

	return nil
}

func (builder *GLShaderProgramBuilder) AddFragmentShaderFromSource(source string) error {
	shader, err := compileShader(source, gl.FRAGMENT_SHADER)

	if err != nil {
		return err
	}

	builder.fragmentShader = shader

	return nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	source += "\x00"

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

func (builder *GLShaderProgramBuilder) Build() (*GLShaderProgram, error) {

	program := gl.CreateProgram()

	gl.AttachShader(program, builder.vertexShader)
	gl.AttachShader(program, builder.fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(builder.vertexShader)
	gl.DeleteShader(builder.fragmentShader)

	builder.vertexShader = 0
	builder.fragmentShader = 0

	return &GLShaderProgram{
		id: program,
	}, nil
}

func UseGLProgram(program GLShaderProgram) {
	gl.UseProgram(program.id)
}
