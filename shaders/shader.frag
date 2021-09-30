#version 410 core

out vec4 frag_colour;
in vec2 texCoords;
uniform sampler2D u_texture;

uniform float u_ratio;
// uniform bool offset;


uniform vec2 u_center;
uniform float u_force;

uniform bool u_circle;

void main(){
    // frag_colour = vec4(1,0,0,1);


    vec2 scaledUV = (texCoords.yx - vec2(0.5, 0.0))/vec2( u_ratio,1.0) + vec2(0.5, 0.0);
        vec2 disp = normalize(texCoords-u_center)*u_force;
    if (u_circle){
        disp = normalize(scaledUV - u_center)*u_force;
    }
    // frag_colour = texture(u_texture, texCoords - vec2(0.5, 0.5));
    // frag_colour = vec4(texCoords-disp, 0,1);
    frag_colour = texture(u_texture, texCoords-disp);
    // frag_colour = vec4(texCoords, 0, 1);
    // frag_colour = vec4( texCoords, 0,1 );
    
}