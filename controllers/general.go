package controllers

import (
	"middleware-experience/constants"
	"middleware-experience/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type (
	ControllerStruct struct {
		Name       string
		Tracer     opentracing.Tracer
		ResDefault models.ResHttpGeneral
	}

	ControllerInterface interface {
		CtrlLoginAuth(c *gin.Context)
		CtrlProfileUsername(c *gin.Context)
	}
)

func InitController(tracer opentracing.Tracer) ControllerInterface {

	return &ControllerStruct{
		Name:   "Controller Middleware - ",
		Tracer: tracer,
		ResDefault: models.ResHttpGeneral{
			Code: http.StatusBadRequest,
			ResGeneral: models.ResGeneral{
				RC:  constants.CODE_REJECT,
				MSG: constants.CODE_REJECT_MSG,
			},
		},
	}
}

func InitRes() models.ResData {
	return models.ResData{
		ResponseCode:    constants.ERROR_ID_FAILED,
		ResponseMessage: constants.ERROR_TEXT_FAILED,
		ResponseData:    nil,
	}

}
