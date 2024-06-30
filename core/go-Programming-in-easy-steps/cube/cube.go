package cube

// Dims exportable.
type Dims struct{ width, length, height int }

func (d *Dims) area() int {
	return d.width * d.length
}

// SetSize exportable.
func (d *Dims) SetSize(w, l, h int) {
	d.width = w
	d.length = l
	d.height = h
}

// GetVolume exportable.
func (d *Dims) GetVolume() int {
	return d.width * d.length * d.height
}

// GetArea exportable.
func (d *Dims) GetArea() int {
	return d.area()
}
