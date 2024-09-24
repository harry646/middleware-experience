package accounting_portal

import (
	"context"
	"encoding/json"
	"middleware-experience/constants"
	"middleware-experience/models"
	"middleware-experience/utils"
	"net/http"

	"github.com/opentracing/opentracing-go"
)

var (
	Host                  string
	PkgName               = "Host LandingPage - "
	ExternalRoutingHeader []models.HeaderHostItem
	ByteHeader            []byte
	V1                    string
)

type (
	HostAccountingPortalStruct struct {
		Span  opentracing.Span
		Ctx   context.Context
		Retry bool
	}

	HostAccountingPortalInterface interface {
		LoginAccount(req models.ReqLoginAccount) (ResProcess, bool, error)
		ProfileByUsername(req models.ProfileByUsername) (ResProcess, bool, error)
	}
)

func InitHostAccountingPortal(ctx context.Context, span opentracing.Span) HostAccountingPortalInterface {
	return &HostAccountingPortalStruct{
		Ctx:   ctx,
		Retry: false,
		Span:  span,
	}
}

func init() {
	Host = utils.EnvString("ExternalRouting.AccountingPortal.Host", "127.0.0.1:8666")
	ByteHeader, _ = json.Marshal(utils.EnvInterface("ExternalRouting.AccountingPortal.Header", nil))
	if len(ByteHeader) > 0 {
		err := json.Unmarshal(ByteHeader, &ExternalRoutingHeader)
		if err != nil {
			utils.LogData(PkgName+"Init", "json.Unmarshal ", constants.LEVEL_LOG_INFO, err.Error())
		}
	}
}

func HeaderGenerate() http.Header {
	Header := make(http.Header)
	for _, item := range ExternalRoutingHeader {
		Header.Add(item.Path, item.Value)
	}
	Header.Add("access-token", constants.AccessToken)
	return Header
}

func HeaderGenerateDownload() http.Header {
	Header := make(http.Header)
	for _, item := range ExternalRoutingHeaderDownload {
		Header.Add(item.Path, item.Value)
	}
	Header.Add("access-token", constants.AccessToken)
	return Header
}

type HeaderItem struct {
	Path  string `json:"Path"`
	Value string `json:"Value"`
}

var ExternalRoutingHeaderDownload = []HeaderItem{
	{
		Path:  "Content-Description",
		Value: "File Transfer",
	},
	{
		Path:  "Content-Transfer-Encoding",
		Value: "binary",
	},
	{
		Path:  "Content-Type",
		Value: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	},
}
