
.section .text.tinygo_scanCurrentStack
.global tinygo_scanCurrentStack
tinygo_scanCurrentStack:
    // TODO: save callee saved registers on the stack
    movi a3, tinygo_scanstack
    jx a3

