package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"gvp/pkg/e"
	"net/http"
	log  "gvp/pkg/logging"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		log.Info("BindAndValid-bind-err:",err.Error())
		return http.StatusBadRequest, e.INVALID_PARAMS
	}



	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		log.Info("BindAndValid-err:",err.Error())
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}