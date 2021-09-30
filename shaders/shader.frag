#version 410 core

out vec4 frag_colour;
in vec2 texCoords;
uniform sampler2D u_texture;

uniform float u_ratio;
// uniform bool offset;


uniform vec2 u_center;
uniform float u_force;

uniform bool u_circle;
uniform bool use_mask;
uniform float u_size;

uniform float thickness;

void main(){
    // frag_colour = vec4(1,0,0,1);


    vec2 scaledUV = (texCoords.xy - vec2(0.5, 0.0))/vec2( u_ratio,1.0) + vec2(0.5, 0.0);
        vec2 disp = normalize(texCoords-u_center)*u_force;
    if (u_circle){
        disp = normalize(scaledUV - u_center)*u_force;
    }

    float mask = (1-smoothstep(u_size-0.1, u_size,length(texCoords-u_center)))*smoothstep(u_size-thickness-0.1, u_size-thickness,length(texCoords-u_center));
    if (u_circle){
        mask = length(scaledUV-u_center);
    }

    if (use_mask){
        disp = disp*mask;
    }
    // frag_colour = texture(u_texture, texCoords - vec2(0.5, 0.5));
    // frag_colour = vec4(texCoords-disp, 0,1);
    frag_colour = texture(u_texture, texCoords-disp);
    // frag_colour.rgb= vec3(mask);
    // frag_colour = vec4(texCoords, 0, 1);
    // frag_colour = vec4( texCoords, 0,1 );
    
}