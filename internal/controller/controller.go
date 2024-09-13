package controller

import (
	"errors"
	"fmt"

	"github.com/ams-api/internal/auth/token"
	"github.com/ams-api/internal/config"
	"github.com/ams-api/internal/service"
	"github.com/ams-api/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	config     config.Config
	tokenMaker token.IMaker
	service    service.IService
	//cache      cache.ICache
	//mailer     mailer.IMailer
}

func NewServer(config config.Config, svc service.IService) (*Server, error) {
	tokenMaker, err := token.NewToken(config)
	if err != nil {
		return nil, fmt.Errorf("cannot make token maker : %w", err)
	}
	server := &Server{}
	server.router = gin.Default()
	server.service = svc
	server.config = config
	server.tokenMaker = tokenMaker
	// server.mailer = mailer.NewMailer(config)
	// server.cache = cache.NewCache(1 * time.Minute)
	server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	if address == "" {
		address = "8080"
	}
	return server.router.Run(fmt.Sprintf(":%s", address))
}

func (server *Server) getAuthPayload(ctx *gin.Context) (*token.Payload, error) {
	pl := ctx.Value("authorization_payload")
	var payload *token.Payload
	util.ConvertType(pl, &payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}
	return payload, nil
}

// func (server *Server) getHeaderXAgentCode(ctx *gin.Context) (string, error) {
// 	ac := ctx.Value("agentCode")
// 	agentCode, ok := ac.(string)
// 	if !ok {
// 		return "", errors.New("invaild agent code")
// 	}
// 	return agentCode, nil
// }
