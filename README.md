# brainfuck
Brainfuck interpreter in Go

# How to use
Run `go run main.go "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."`

Replace bf with any other bf instruction

# How to use library
1. Run `go get github.com/youssefouirini/brainfuck`
2. Create Bf using `model.Brainfuck` in `github.com/youssefouirini/brainfuck/model`
3. Use as `cmd.ExecuteBf(model.Brainfuck{})` or rename package as package `bf github.com/youssefouirini/brainfuck/cmd` and use as `bf.ExecuteBf(model.Brainfuck{})`
