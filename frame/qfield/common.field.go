package qfield

type Id struct {
	Id *uint `where:"" example:"1" v:"required|integer#id必填|id必须为整数"`
}

func (f *Id) GetId() uint {
	if *f.Id == 0 {
		panic("Id 不允许为空")
	}
	return *f.Id
}

type IdParam interface {
	GetId() uint
}
