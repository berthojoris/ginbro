package handlers

import (
	"github.com/berthojoris/ginbro/models"
	"github.com/gin-gonic/gin"
)

func init() {
	groupApi.GET("order", orderAll)
	groupApi.GET("order/:id", orderOne)
	groupApi.POST("order", orderCreate)
	groupApi.PATCH("order", orderUpdate)
	groupApi.DELETE("order/:id", orderDelete)
}

//All
func orderAll(c *gin.Context) {
	mdl := models.Order{}
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
func orderOne(c *gin.Context) {
	var mdl models.Order
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
func orderCreate(c *gin.Context) {
	var mdl models.Order
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
func orderUpdate(c *gin.Context) {
	var mdl models.Order
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
func orderDelete(c *gin.Context) {
	var mdl models.Order
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
