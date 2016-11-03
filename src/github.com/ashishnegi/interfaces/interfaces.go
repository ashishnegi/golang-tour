package interfaces

import (
	"fmt"
)

type ISeq interface {
	first()
	rest()
}

type SeqString struct {
	data string
}

func (iSeq *SeqString) first() byte {
	return iSeq.data[0]
}

func (iSeq *SeqString) rest() string {
	return iSeq.data[1:]
}

func TestInterfaces() {
	seq := SeqString{"s..seqs rocks!!!"}
	fmt.Println("first char: ", seq.first(), "\n rest seq:", seq.rest())
}
