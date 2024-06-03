package builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	// getHouse 获取产品
	getHouse() House
}
