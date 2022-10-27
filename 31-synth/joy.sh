S1='b4,b4,c5,d5,d5,c5,b4,a4,g4,g4,a4,b4,b4,a4,a4,_'
S2='b4,b4,c5,d5,d5,c5,b4,a4,g4,g4,a4,b4,a4,g4,g4,_'

A1='g4,g4,a4,g4,g4,a4,g4,f#4,d4,d4,f#4,g4,g4,f#4,f#4,_'
A2='g4,g4,a4,g4,g4,a4,g3,f#4,d4,d4,f#4,g4,g4,f#4,d4,_'

T1='d4,d4,c4,b3,e4,d4,d4,d4,b3,b3,d4,d4,d4,d4,d4,_'
T2='d4,d4,c4,b3,e4,d4,d4,d4,b3,b3,d4,d4,d4,d4,b3,_'

B1='g3,g3,g3,g3,e3,f#3,g3,d3,b3,b3,a3,g3,d3,d3,d3,_'
B2='g3,g3,g3,g3,e3,f#3,g3,d3,b3,b3,a3,g3,d3,d3,g2,_'

go run synth.go -r 48000 -t 0 -o /dev/stdout \
    "$S1,$S2" "$A1,$A2" "$T1,$T2" "$B1,$B2" |
  aplay -r 48000 -f S16_LE
