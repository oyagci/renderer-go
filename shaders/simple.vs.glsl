#version 460 core

in vec3 vert;
in vec3 color;

out vec3 out_color;

void main() {
    gl_Position = vec4(vert, 1.0f);
    out_color = color;
}