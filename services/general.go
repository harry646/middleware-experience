package services

import (
	"context"
	"middleware-experience/helpers"
	"middleware-experience/models"

	"github.com/opentracing/opentracing-go"
)

type (
	MiddlewareServicesStruct struct {
		Ctx     context.Context
		Routing helpers.RoutingHostInterface
		Span    opentracing.Span
	}

	MiddlewareServiceInterface interface {
		// Account
		SvcAccountLogin(req models.ReqLoginAccount, signatureHeader string, timestampsHeader string, reqHttpMethod string, res *models.ResData)
		SvcProfileUsername(req models.ProfileByUsername, res *models.ResData)
	}
)

func InitMiddlewareServices(ctx context.Context, span opentracing.Span) MiddlewareServiceInterface {
	return &MiddlewareServicesStruct{
		Ctx:     ctx,
		Routing: helpers.InitRoutingHost(ctx, span),
		Span:    span,
	}
}
