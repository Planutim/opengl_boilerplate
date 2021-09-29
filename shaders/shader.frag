#version 410 core

out vec4 frag_colour;
in vec2 texCoords;
uniform sampler2D u_texture;

void main(){
    // frag_colour = vec4(1,0,0,1);
    frag_colour = texture(u_texture, texCoords);
    // frag_colour = vec4( texCoords, 0,1 );
}