package services

import (
	"middleware-experience/constants"
	"middleware-experience/helpers"
	"middleware-experience/models"
	"middleware-experience/utils"

	"github.com/go-playground/validator"
)

func (svc *MiddlewareServicesStruct) SvcAccountLogin(req models.ReqLoginAccount, signatureHeader string, timestampsHeader string, reqHttpMethod string, res *models.ResData) {
	validate := validator.New()
	if checkStruct := validate.Struct(req); checkStruct != nil {
		res.ResponseMessage = checkStruct.Error()
		return
	}

	check_header := helpers.GenerateSignatureLogin(req, timestampsHeader, reqHttpMethod)
	if check_header == signatureHeader {
		data, pending, err := svc.Routing.LoginAccount(req)

		if err != nil || pending {
			utils.LogFmtTemp(err.Error())
			res.ResponseCode = constants.ERROR_ID_FAILED
			res.ResponseMessage = err.Error()
			return
		}

		res.ResponseCode = constants.CODE_SUCCESS
		res.ResponseMessage = constants.CODE_SUCCESS_MSG
		res.ResponseData = data
	} else {
		res.ResponseCode = constants.CODE_SIGNATURE_ERR
		res.ResponseMessage = constants.CODE_SIGNATURE_ERR_MSG
	}

}

func (svc *MiddlewareServicesStruct) SvcProfileUsername(req models.ProfileByUsername, res *models.ResData) {
	data, pending, err := svc.Routing.ProfileByUsername(req)

	if err != nil || pending {
		utils.LogFmtTemp(err.Error())
		res.ResponseCode = constants.ERROR_ID_FAILED
		res.ResponseMessage = err.Error()
		return
	}

	res.ResponseCode = constants.CODE_SUCCESS
	res.ResponseMessage = constants.CODE_SUCCESS_MSG
	res.ResponseData = data

}
