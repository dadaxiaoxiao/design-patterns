package builder

// Director 指挥者
type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

// buildHouse 指挥构建顺序
func (d *Director) buildHouse() House {
	// 延迟执行某些步骤而不会影响最终产品
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
