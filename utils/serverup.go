package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinServerUp(listenAddr string, router *gin.Engine) error {

	cfg := *GetServerConfig()
	fmt.Println("[TLS.1.2]:", cfg.Servertls12client)
	srv := &http.Server{
		Addr:              listenAddr,
		Handler:           router,
		TLSConfig:         GetServerTlsConfig(),
		ReadTimeout:       cfg.ReadTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
		MaxHeaderBytes:    cfg.MaxHeaderBytes,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
	}

	if cfg.Servertls12client == "ON" {
		return srv.ListenAndServeTLS(cfg.Serverfilepubkey, cfg.Serverfileprivatekey)
	}
	return srv.ListenAndServe()
}
