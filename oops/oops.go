package main

import (
	"fmt"
)

// Composition over Inheritance
type Creature struct {
	Name string
	Real bool
}

func (c Creature) toString() string {
	return fmt.Sprintf("Creature: %s ; Real : %v", c.Name, c.Real)
}

func (c Creature) Foo() string {
	return fmt.Sprintf("Creature: %s's Foo", c.Name)
}

type FlyingCreature struct {
	Creature
	WingSpan int
}

// syntax of method of a struct
// func (<object-name/this> <ClassName>) <func-name> () <return-type>
func (c FlyingCreature) toString() string {
	return fmt.Sprintf("Flying Creature: %s ; Real : %v ; WingSpan : %d, Foo: %s", c.Name, c.Real, c.WingSpan, c.Foo())
}

// Polymorphism
type IShape interface {
	Area() float32
}

type Rectangle struct {
	width, height float32
}

// You can make a Rectangle an IShape
// without knowing about IShape or Rectangle extending IShape before.
func (r Rectangle) Area() float32 {
	return r.width * r.height
}

// this class is in some other library whose code we don't control.
type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return 3.14 * c.radius * c.radius
}

func SomeWork(s IShape) {
	fmt.Println("Some work on IShape: Area is: ", s.Area())
}

func main() {
	c := Creature{Name: "Lion", Real: true}
	fmt.Println(c.toString())
	f := FlyingCreature{Creature: Creature{Name: "Bat", Real: true}, WingSpan: 1}
	fmt.Println(f.toString())
	fmt.Printf("Rectangle : %f\n", Rectangle{width: 2, height: 3}.Area())
	fmt.Printf("Circle : %f\n", Circle{radius: 3}.Area())
	rect := Rectangle{width: 2, height: 6}
	SomeWork(rect) // this is polymorphism.

	// some other library which takes BoostThread
	bt := BoostThread{}
	StartThread(bt)
	st := StdThread{}
	StartThread(st)
}

type IThread interface {
	Start()
}

type BoostThread struct {
}

func (BoostThread) Start() {
	fmt.Println("BoostThread start...")
}

func StartThread(bt IThread) {
	fmt.Println("I was written for boost thread: ")
	bt.Start()
}

// now std threds come from another library which we can't change
// but StdThread has methods that we need and may be more functions as well.
type StdThread struct {
}

func (StdThread) Start() {
	fmt.Println("StdThread start...")
}

func (StdThread) MoreMethods() {
	fmt.Println("StdThread start...")
}