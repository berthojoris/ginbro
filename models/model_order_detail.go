package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

//OrderDetail
type OrderDetail struct {
	Id        uint    `gorm:"column:id" form:"id" json:"id" comment:"" sql:"int(10),PRI"`
	OrderId   int     `gorm:"column:order_id" form:"order_id" json:"order_id" comment:"" sql:"int(10),MUL"`
	ItemId    int     `gorm:"column:item_id" form:"item_id" json:"item_id" comment:"" sql:"int(10)"`
	ItemName  string  `gorm:"column:item_name" form:"item_name" json:"item_name" comment:"" sql:"varchar(100)"`
	ItemPrice float64 `gorm:"column:item_price" form:"item_price" json:"item_price" comment:"" sql:"double"`
}

//TableName
func (m *OrderDetail) TableName() string {
	return "order_detail"
}

//One
func (m *OrderDetail) One() (one *OrderDetail, err error) {
	one = &OrderDetail{}
	err = crudOne(m, one)
	return
}

//All
func (m *OrderDetail) All(q *PaginationQuery) (list *[]OrderDetail, total uint, err error) {
	list = &[]OrderDetail{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *OrderDetail) Update() (err error) {
	where := OrderDetail{Id: m.Id}
	m.Id = 0

	return crudUpdate(m, where)
}

//Create
func (m *OrderDetail) Create() (err error) {
	m.Id = 0

	return mysqlDB.Create(m).Error
}

//Delete
func (m *OrderDetail) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
