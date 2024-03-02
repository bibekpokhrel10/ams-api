package middleware

import (
	"github.com/gin-gonic/gin"
)

// const (
// 	authorizationHeaderKey  = "authorization"
// 	authorizationTypeBearer = "bearer"
// 	authorizationPayloadKey = "authorization_payload"
// )

// func AuthMiddleware(tokenMaker token.IMaker, s service.IService, cache cache.ICache) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
// 		if len(authorizationHeader) == 0 {
// 			err := errors.New("authorization is not provided")
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
// 			return
// 		}
// 		fields := strings.Fields(authorizationHeader)
// 		if len(fields) < 2 {
// 			err := errors.New("invalid authorization header format")
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
// 			return
// 		}
// 		authorizationType := strings.ToLower(fields[0])
// 		if authorizationType != authorizationTypeBearer {
// 			err := fmt.Errorf("unsupported authorization type %s ", authorizationType)
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
// 			return
// 		}
// 		accessToken := fields[1]
// 		payloadKey := "payload-" + accessToken
// 		var value token.Payload
// 		if ok := cache.Get(payloadKey, &value); ok {
// 			ctx.Set(authorizationPayloadKey, value)
// 			ctx.Next()
// 			return
// 		}
// 		payload, err := tokenMaker.VerifyToken(accessToken)
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
// 			return
// 		}
// 		if payload.AgentId != 0 {
// 			agent, _, err := s.GetAgent(payload.AgentId)
// 			if err != nil {
// 				err = fmt.Errorf("agent not found :: %w", err)
// 				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
// 				return
// 			}
// 			if !agent.IsActive {
// 				err = fmt.Errorf("agent not active :: %w", err)
// 				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
// 				return
// 			}
// 		}
// 		if payload.PlayerId != 0 {
// 			player, err := s.GetPlayer(payload.PlayerId)
// 			if err != nil {
// 				err = fmt.Errorf("player not found :: %w", err)
// 				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
// 				return
// 			}
// 			if !player.IsActive {
// 				err = fmt.Errorf("player blocked :: %w", err)
// 				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ERROR(err))
// 				return
// 			}
// 		}
// 		cache.Set(payloadKey, payload)
// 		ctx.Set(authorizationPayloadKey, payload)
// 		ctx.Next()
// 	}
// }

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
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-agent-code")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}