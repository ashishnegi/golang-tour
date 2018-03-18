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

func DynamicTypingIsIt() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func newton(zn, x float64) float64 {
	return zn - (zn*zn-x)/(2*zn)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1, ErrNegativeSqrt(x)
	}
	sqrtValue := newton(x, x)
	oldSqrtValue := x
	change := oldSqrtValue - sqrtValue
	for change > 0.01 {
		oldSqrtValue = sqrtValue
		sqrtValue = newton(sqrtValue, x)
		change = oldSqrtValue - sqrtValue
	}
	return sqrtValue, nil
}

func CheckSqrt() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
