package qfield

type Id struct {
	Id *uint `where:"" example:"1" v:"required|integer#id必填|id必须为整数"`
}
