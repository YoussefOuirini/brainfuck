package model

type Brainfuck struct {
	Contents []byte
}

func (b *Brainfuck) Add(operation byte) {
	b.Contents = append(b.Contents, operation)
}

func (b *Brainfuck) Remove() {
	b.Contents = b.Contents[:len(b.Contents)-1]
}
