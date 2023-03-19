package __

import (
	"log"
	"math"
	"testing"
)

type graphical interface {
	Area() float64
	Perimeter() float64
}

type rect struct{ x, y float64 }

func NewRect() rect { return rect{1, 2} }

func (this *rect) Area() float64 { // const?
	return this.x * this.y
}
func (this *rect) Perimeter() float64 {
	return 2*this.x + 2*this.y
}

type circle struct{ radius float64 }

func (this *circle) Perimeter() float64 {
	return 2 * math.Pi * this.radius
}
func (this *circle) Area() float64 {
	return this.radius * this.radius * math.Pi
}

func TestArea(t *testing.T) {
	a := rect{2, 3}
	b := circle{1.5}

	log.Printf("Area of rect is %v\n", a.Area())
	log.Printf("Area of circle is %v\n", b.Area())
}

func TestInterface(t *testing.T) {
	a := rect{2, 3}
	b := circle{1.5}
	g := graphical(&a)

	log.Println(g)
	log.Println(a)
	log.Println(b)
}
