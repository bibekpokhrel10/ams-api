package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ams-api/internal/auth/token"
	"github.com/ams-api/internal/response"
	"github.com/ams-api/internal/service"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.IMaker, s service.IService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s ", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
			return
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

// func AgentCodeMiddleware(s service.IService, cache cache.ICache) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		agentCode := ctx.GetHeader("x-agent-code")
// 		if agentCode == "" {
// 			err := fmt.Errorf("x-agent-code header is missing")
// 			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ERROR(err))
// 			return
// 		}
// 		payloadKey := "agent-" + agentCode
// 		var value string
// 		if ok := cache.Get(payloadKey, &value); ok {
// 			ctx.Set("agentCode", value)
// 			ctx.Next()
// 			return
// 		}
// 		agent, err := s.GetAgentByCode(agentCode)
// 		if err != nil {
// 			err = fmt.Errorf("agent not found :: %w", err)
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
// 			return
// 		}
// 		if !agent.IsActive {
// 			err = fmt.Errorf("agent is inactive :: %w", err)
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
// 			return
// 		}
// 		cache.Set(payloadKey, agentCode)
// 		ctx.Set("agentCode", agentCode)
// 		ctx.Next()
// 	}
// }

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		// c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Credentials", "true")
		// c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-agent-code")
		c.Header("Access-Control-Allow-Headers", "*")
		// c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")
		c.Header("Access-Control-Allow-Methods", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
