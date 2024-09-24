package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"middleware-experience/constants"
	"middleware-experience/models"
	"middleware-experience/services"
	"middleware-experience/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

func (ctrl *ControllerStruct) CtrlLoginAuth(c *gin.Context) {
	logName := "CtrlAccountLogin"
	res := InitRes()
	Tracer, Closer, err := utils.InitJaeger()
	if err != nil {
		utils.LogData(logName, constants.ERROR_ID_INIT_JEAGER, constants.LEVEL_LOG_WARNING, err.Error())
	}
	defer Closer.Close()

	utils.LogFmtTemp("Start " + logName)
	spanCtx, _ := Tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	span := Tracer.StartSpan("", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	// Get and check the Authorization header
	x_signature := c.Request.Header.Get("X-SIGNATURE")
	if x_signature == "" {
		utils.LogData(logName, "Missing Authorization Header", constants.LEVEL_LOG_WARNING, "Authorization header is required")
		res.ResponseMessage = "Authorization header is required"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// x_signature := "test"

	x_timestamps := c.Request.Header.Get("X-TIMESTAMPS")

	reqHttpMethod := c.Request.Method

	var req models.ReqLoginAccount
	dataBody, err := c.GetRawData()
	utils.LogData(logName, "Request", constants.LEVEL_LOG_INFO, string(dataBody))
	if err != nil {
		utils.LogData(logName, "GetRawData", constants.LEVEL_LOG_ERROR, err.Error())
		res.ResponseMessage = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err = json.Unmarshal(dataBody, &req); err != nil {
		utils.LogData(logName, "json.Unmarshal", constants.LEVEL_LOG_ERROR, err.Error())
		res.ResponseMessage = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	span.LogFields(
		log.Object("Request ", req),
	)
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		res.ResponseMessage = fmt.Sprint(utils.GenerateMessageErrorValidate(err))
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	services.InitMiddlewareServices(ctx, span).SvcAccountLogin(req, x_signature, x_timestamps, reqHttpMethod, &res)
	span.LogFields(
		log.Object("Response ", res),
	)
	dataRes, _ := json.Marshal(res)
	utils.LogData(logName, "Response", constants.LEVEL_LOG_INFO, string(dataRes))
	c.JSON(http.StatusOK, res)
	utils.LogFmtTemp("End " + logName)
}

func (ctrl *ControllerStruct) CtrlProfileUsername(c *gin.Context) {
	logName := "Get Profile by Username"
	res := InitRes()
	Tracer, Closer, err := utils.InitJaeger()
	if err != nil {
		utils.LogData(logName, constants.ERROR_ID_INIT_JEAGER, constants.LEVEL_LOG_WARNING, err.Error())
	}
	defer Closer.Close()
	utils.LogFmtTemp("Start" + logName)
	spanCtx, _ := Tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	span := Tracer.StartSpan("", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	var req models.ProfileByUsername
	dataBody, err := c.GetRawData()
	utils.LogData(logName, "Request", constants.LEVEL_LOG_INFO, string(dataBody))
	if err != nil {
		utils.LogData(logName, "GetRawData", constants.LEVEL_LOG_ERROR, err.Error())
		res.ResponseMessage = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err = json.Unmarshal(dataBody, &req); err != nil {
		utils.LogData(logName, "json.Unmarshal", constants.LEVEL_LOG_ERROR, err.Error())
		res.ResponseMessage = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	span.LogFields(
		log.Object("Request ", req),
	)

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		res.ResponseMessage = fmt.Sprint(utils.GenerateMessageErrorValidate(err))
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	services.InitMiddlewareServices(ctx, span).SvcProfileUsername(req, &res)
	span.LogFields(
		log.Object("Response ", res),
	)
	dataRes, _ := json.Marshal(res)
	utils.LogData(logName, "Response", constants.LEVEL_LOG_INFO, string(dataRes))
	c.JSON(http.StatusOK, res)
	utils.LogFmtTemp("End", logName)
}
