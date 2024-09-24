package accounting_portal

import (
	"encoding/json"
	"errors"
	"middleware-experience/constants"
	"middleware-experience/models"
	"middleware-experience/utils"
)

func (hosts *HostAccountingPortalStruct) LoginAccount(req models.ReqLoginAccount) (ResProcess, bool, error) {
	var res ResProcess
	ctrlName := PkgName + "LoginAccount"
	headNew := HeaderGenerate()

	url := utils.EnvString("ExternalRouting.AccountingPortal.Endpoint.EAccPortalAccount.AccountLogin.Url", "/e-accounting/account/login")
	method := utils.EnvString("ExternalRouting.AccountingPortal.Endpoint.EAccPortalAccount.LoAccountLogingin.Method", constants.HttpMethodPost)
	_, headNew = utils.AddSpanHeader(hosts.Span, Host+url, method, headNew)
	data, err := utils.SendHttpRequest(method, Host+url, headNew, req, hosts.Retry)
	if err != nil {
		utils.LogData(ctrlName, "utils.SendHttpRequest", constants.LEVEL_LOG_WARNING, err.Error())
		return res, false, err
	}
	err = json.Unmarshal(data, &res)

	if err != nil {
		utils.LogData(ctrlName, "json.Unmarshal", constants.LEVEL_LOG_WARNING, err.Error())
		return res, true, err
	}
	if res.Response_Code == constants.CODE_PENDING {
		return res, true, nil
	} else if res.Response_Code == constants.CODE_SUCCESS || res.Response_Code == constants.CODE_CREATED {
		return res, false, nil
	} else {
		return res, false, errors.New(res.Response_Message)

	}
}

func (hosts *HostAccountingPortalStruct) ProfileByUsername(req models.ProfileByUsername) (ResProcess, bool, error) {
	var res ResProcess
	ctrlName := PkgName + "ProfileByUsername"
	headNew := HeaderGenerate()

	url := utils.EnvString("ExternalRouting.AccountingPortal.Endpoint.EAccPortalAccount.ProfileByUsername.Url", "/e-accounting/account/profile/username")
	method := utils.EnvString("ExternalRouting.AccountingPortal.Endpoint.EAccPortalAccount.ProfileByUsername.Method", constants.HttpMethodPost)
	_, headNew = utils.AddSpanHeader(hosts.Span, Host+url, method, headNew)
	data, err := utils.SendHttpRequest(method, Host+url, headNew, req, hosts.Retry)
	if err != nil {
		utils.LogData(ctrlName, "utils.SendHttpRequest", constants.LEVEL_LOG_WARNING, err.Error())
		return res, false, err
	}
	err = json.Unmarshal(data, &res)

	if err != nil {
		utils.LogData(ctrlName, "json.Unmarshal", constants.LEVEL_LOG_WARNING, err.Error())
		return res, true, err
	}
	if res.Response_Code == constants.CODE_PENDING {
		return res, true, nil
	} else if res.Response_Code == constants.CODE_SUCCESS || res.Response_Code == constants.CODE_CREATED {
		return res, false, nil
	} else {
		return res, false, errors.New(res.Response_Message)

	}
}
