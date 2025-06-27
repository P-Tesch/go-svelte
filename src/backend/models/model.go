package models

type Model interface {
	FindById() bool
	Save() int64
	Delete() bool
	CreateDTO() any
}
