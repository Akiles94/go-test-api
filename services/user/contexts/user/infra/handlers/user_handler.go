package handlers

import (
	"strings"

	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/dto"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/ports/inbound"
	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/infra/shared_handlers"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	loginUseCase    inbound.LoginUseCasePort
	registerUseCase inbound.RegisterUseCasePort
	jwtService      shared_ports.JWTServicePort
}

func NewUserHandler(loginUseCase inbound.LoginUseCasePort, registerUseCase inbound.RegisterUseCasePort) *UserHandler {
	return &UserHandler{
		loginUseCase:    loginUseCase,
		registerUseCase: registerUseCase,
	}
}

func (uh *UserHandler) Login(c *gin.Context) {
	var loginDto dto.LoginRequestDto
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.Error(shared_handlers.ErrInvalidPayload)
		return
	}
	loginResponse, err := uh.loginUseCase.Execute(c.Request.Context(), loginDto)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, loginResponse)
}

func (uh *UserHandler) Register(c *gin.Context) {
	var registerDto dto.RegisterRequestDto
	if err := c.ShouldBindJSON(&registerDto); err != nil {
		c.Error(shared_handlers.ErrInvalidPayload)
		return
	}
	registerResponse, err := uh.registerUseCase.Execute(c.Request.Context(), registerDto)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, registerResponse)
}

func (uh *UserHandler) ValidateToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.Status(401)
		return
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		c.Status(401)
		return
	}

	claims, err := uh.jwtService.Parse(bearerToken[1])
	if err != nil {
		c.Status(401)
		return
	}

	c.Header("X-User-ID", claims.ID)
	c.Header("X-User-Role", string(claims.Role))
	c.Status(200)
}
