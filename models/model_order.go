package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

// Validation :
// https://github.com/gin-gonic/gin#model-binding-and-validation
// https://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Baked_In_Validators_and_Tags

//Order
type Order struct {
	Id          uint          `gorm:"column:id" form:"id" json:"id" comment:"" sql:"int(10),PRI"`
	Total       float64       `gorm:"column:total" form:"total" json:"total" comment:"" sql:"double"`
	OrderDetail []OrderDetail `gorm:"foreignkey:OrderID" json:"detail_order"`
}

//TableName
func (m *Order) TableName() string {
	return "order"
}

//One
func (m *Order) One() (one *Order, err error) {
	one = &Order{}
	err = crudOneRelated(m, one)
	return
}

//All
func (m *Order) All(q *PaginationQuery) (list *[]Order, total uint, err error) {
	list = &[]Order{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *Order) Update() (err error) {
	where := Order{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *Order) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *Order) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
