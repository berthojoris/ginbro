package handlers

import (
	"github.com/berthojoris/ginbro/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("order-detail", orderDetailAll)
	groupApi.GET("order-detail/:id", orderDetailOne)
	groupApi.POST("order-detail", orderDetailCreate)
	groupApi.PATCH("order-detail", orderDetailUpdate)
	groupApi.DELETE("order-detail/:id", orderDetailDelete)
}

//All
func orderDetailAll(c *gin.Context) {
	mdl := models.OrderDetail{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	jsonPagination(c, list, total, query)
}

//One
func orderDetailOne(c *gin.Context) {
	var mdl models.OrderDetail
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

//Create
func orderDetailCreate(c *gin.Context) {
	var mdl models.OrderDetail
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Create()
	if handleError(c, err) {
		return
	}
	jsonData(c, mdl)
}

//Update
func orderDetailUpdate(c *gin.Context) {
	var mdl models.OrderDetail
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Update()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

//Delete
func orderDetailDelete(c *gin.Context) {
	var mdl models.OrderDetail
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	err = mdl.Delete()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}
