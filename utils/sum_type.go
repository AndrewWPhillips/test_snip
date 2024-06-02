package utils

type (
	pictureType interface {
		circle | rectangle | translated[circle]
	}
	picture struct {
		val any
	}
	circle struct {
		radius float64
	}
	rectangle struct {
		width, height float64
	}
	translated[T pictureType] struct {
		xoffset, yoffset float64
		object           T
	}
)

func NewPicture[T pictureType](val T) picture {
	return picture{
		val: val,
	}
}

func NewCircle(radius float64) circle {
	return circle{radius: radius}
}

func (c circle) GetRadius() float64 {
	return c.radius
}

func (c *circle) SetRadius(radius float64) {
	c.radius = radius
}

func NewRectangle(width, height float64) rectangle {
	return rectangle{
		width:  width,
		height: height,
	}
}

func (r rectangle) GetWidth() float64 {
	return r.width

}
func (r *rectangle) SetWidth(width float64) {
	r.width = width
}

func (r rectangle) GetHeight() float64 {
	return r.height

}
func (r *rectangle) SetHeight(height float64) {
	r.height = height
}

func NewTranslated[T pictureType](xoffset, yoffset float64, val T) translated[T] {
	return translated[T]{
		xoffset: xoffset,
		yoffset: yoffset,
		object:  val,
	}
}
