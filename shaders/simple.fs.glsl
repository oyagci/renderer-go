#version 460 core

out vec4 fragColor;

in vec3 out_color;

void main() {
	fragColor = vec4(out_color, 1.0);
}