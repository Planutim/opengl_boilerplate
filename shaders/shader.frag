#version 410 core

out vec4 frag_colour;
in vec2 texCoords;
uniform sampler2D u_texture;

// uniform bool offset;

uniform vec2 offset;

void main(){
    // frag_colour = vec4(1,0,0,1);
    vec2 modifiedTexCoords = texCoords;
        modifiedTexCoords = texCoords +offset;


    if (floor(mod(texCoords.x*4,2)) ==1){
        modifiedTexCoords.y += 0.25;
    }else{
        modifiedTexCoords.y -= 0.25;
    }
    frag_colour = texture(u_texture, modifiedTexCoords);
    // frag_colour = vec4( texCoords, 0,1 );
}