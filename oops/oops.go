package main

import (
	"fmt"
)

type Creature struct {
	Name string
	Real bool
}

func (c Creature) toString() string {
	return fmt.Sprintf("Creature: %s ; Real : %v", c.Name, c.Real)
}

type FlyingCreature struct {
	Creature
	WingSpan int
}

func (c FlyingCreature) toString() string {
	return fmt.Sprintf("Flying Creature: %s ; Real : %v ; WingSpan : %d", c.Name, c.Real, c.WingSpan)
}

type IShape interface {
	Area() float32
}

type Rectangle struct {
	width, height float32
}

type Circle struct {
	radius float32
}

func (r Rectangle) Area() float32 {
	return r.width * r.height
}

func (c Circle) Area() float32 {
	return 3.14 * c.radius * c.radius
}

func main() {
	c := Creature{Name: "Lion", Real: true}
	fmt.Println(c.toString())
	f := FlyingCreature{Creature: Creature{Name: "Bat", Real: true}, WingSpan: 1}
	fmt.Println(f.toString())
	fmt.Printf("Rectangle : %f\n", Rectangle{width: 2, height: 3}.Area())
	fmt.Printf("Circle : %f\n", Circle{radius: 3}.Area())
}
