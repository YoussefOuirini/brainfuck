# brainfuck
Brainfuck interpreter in Go

# How to use
Run `go run main.go "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."`

Replace bf with any other bf instruction

# How to use library
1. Run `go get github.com/youssefouirini/brainfuck/cmd`
2. Use as `cmd.ExecuteBf(bytes)` or rename package as package `bf github.com/youssefouirini/brainfuck/cmd` and use as `bf.ExecuteBf(bytes)`
