#version 460 core

in vec3 position;
in vec3 color;

out vec3 out_color;

void main() {
    gl_Position = vec4(position, 1.0f);
    out_color = color;
}