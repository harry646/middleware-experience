package utils

import (
	"io"
	"middleware-experience/constants"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

// initJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func InitJaeger() (opentracing.Tracer, io.Closer, error) {
	service := EnvString("Tracing.Service", "LandingPage")
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           EnvBool("Tracing.Logs", true),
			LocalAgentHostPort: EnvString("Tracing.Host", "localhost:6831"),
			QueueSize:          EnvInt("Tracing.MaxQueue", 100),
		},
	}
	tracer, closer, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
		//config.Metrics(metricsFactory),
	)
	if err != nil {
		LogData(PkgName+"InitJaeger", "cfg.NewTracer", constants.LEVEL_LOG_WARNING, err.Error())
		return tracer, closer, err
	}
	opentracing.InitGlobalTracer(tracer)
	return tracer, closer, err
}

// OpenTracer - middleware that addes opentracing
func OpenTracer(operationPrefix []byte) gin.HandlerFunc {
	if operationPrefix == nil {
		operationPrefix = []byte("api-request-")
	}
	return func(c *gin.Context) {
		// all before request is handled
		var span opentracing.Span
		if cspan, ok := c.Get("tracing-context"); ok {
			span = StartSpanWithParent(cspan.(opentracing.Span).Context(), string(operationPrefix)+c.Request.RequestURI, c.Request.Method, c.Request.URL.Path)

		} else {
			span = StartSpanWithHeader(&c.Request.Header, string(operationPrefix)+c.Request.RequestURI, c.Request.Method, c.Request.URL.Path)
		}
		defer span.Finish()            // after all the other defers are completed.. finish the span
		c.Set("tracing-context", span) // add the span to the context so it can be used for the duration of the request.
		c.Next()

		span.SetTag(string(ext.HTTPStatusCode), c.Writer.Status())
	}
}

func StartSpanWithParent(parent opentracing.SpanContext, operationName, method, path string) opentracing.Span {
	options := []opentracing.StartSpanOption{
		opentracing.Tag{Key: ext.SpanKindRPCServer.Key, Value: ext.SpanKindRPCServer.Value},
		opentracing.Tag{Key: string(ext.HTTPMethod), Value: method},
		opentracing.Tag{Key: string(ext.HTTPUrl), Value: path},
		opentracing.Tag{Key: "current-goroutines", Value: runtime.NumGoroutine()},
	}

	if parent != nil {
		options = append(options, opentracing.ChildOf(parent))
	}

	return opentracing.StartSpan(operationName, options...)
}

func StartSpanWithHeader(header *http.Header, operationName, method, path string) opentracing.Span {
	var wireContext opentracing.SpanContext
	if header != nil {
		wireContext, _ = opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(*header))
	}
	span := StartSpanWithParent(wireContext, operationName, method, path)
	span.SetTag("current-goroutines", runtime.NumGoroutine())
	return span
	// return StartSpanWithParent(wireContext, operationName, method, path)
}

func AddSpanHeader(span opentracing.Span, url string, method string, head http.Header) (opentracing.Span, http.Header) {
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, method)
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(head),
	)
	return span, head
}
