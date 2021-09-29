#version 410 core
layout (location = 0) in vec4 aData;


out vec2 texCoords;

void main(){
    gl_Position = vec4(aData.xy, 0, 1);
    texCoords = aData.zw;
}