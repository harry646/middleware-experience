package helpers

import (
	"context"
	"middleware-experience/host/accounting_portal"
	"middleware-experience/models"

	"github.com/opentracing/opentracing-go"
)

type (
	RoutingHostStruct struct {
		Span              opentracing.Span
		Ctx               context.Context
		Accounting_Portal accounting_portal.HostAccountingPortalInterface
	}

	RoutingHostInterface interface {
		//Account
		LoginAccount(req models.ReqLoginAccount) (interface{}, bool, error)
		ProfileByUsername(req models.ProfileByUsername) (interface{}, bool, error)
	}
)

func InitRoutingHost(ctx context.Context, span opentracing.Span) RoutingHostInterface {
	return &RoutingHostStruct{
		Ctx:               ctx,
		Span:              span,
		Accounting_Portal: accounting_portal.InitHostAccountingPortal(ctx, span),
	}
}

func DummyCredentialTracing(ctx context.Context) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Account Credential")
	defer span.Finish()
	return "00", nil
}

func (rh *RoutingHostStruct) LoginAccount(req models.ReqLoginAccount) (interface{}, bool, error) {
	_, err := DummyCredentialTracing(rh.Ctx)
	if err != nil {
		return nil, false, err
	}
	pending := false

	dataLandingPage, pending, err := rh.Accounting_Portal.LoginAccount(req)

	return dataLandingPage.Response_Data, pending, err
}

func (rh *RoutingHostStruct) ProfileByUsername(req models.ProfileByUsername) (interface{}, bool, error) {
	_, err := DummyCredentialTracing(rh.Ctx)
	if err != nil {
		return nil, false, err
	}
	pending := false

	dataLandingPage, pending, err := rh.Accounting_Portal.ProfileByUsername(req)

	return dataLandingPage.Response_Data, pending, err
}
