package mock

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"taskulu/pkg"
	"taskulu/internal"
	"io/ioutil"
	"fmt"
)

type Server struct {
	engine     *gin.Engine
	authorized *gin.RouterGroup
	handler    *Handler
}

type Option struct {
	Address string
	User    string
	Pass    string
}

func New(log *pkg.Logger, option Option) *Server {
	engine := gin.Default()
	auth := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		option.User: option.Pass,
		//"user2": "pass2", // user:user2 password:pass2
	}))
	s := Server{
		engine:     engine,
		authorized: auth,
		handler:    NewHandler(log)}
	s.setupRouter()

	go func(address string) {
		err := s.engine.Run(address)
		if err != nil {
			log.Fatal(err)
		}
	}(option.Address)

	return &s
}

type Handler struct {
	log *pkg.Logger
}

func NewHandler(log *pkg.Logger) *Handler {
	return &Handler{log}
}

func (h *Handler) GetActivities(c *gin.Context) {
	fmt.Println("=========", c.Query("app_key"))
	if c.Query("app_key") == "" &&  c.Query("session_key") == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	fmt.Println(":::::::::::")
	c.String(http.StatusOK, Activities)
	return
}

func (h *Handler) CreateSession(c *gin.Context) {
	c.String(http.StatusCreated, Session)
	return
}

func (h *Handler) BaleIntegration(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Error(err)
		return
	}
	c.String(http.StatusOK, string(b))
	return
}

func (s *Server) setupRouter() {
	s.engine.POST(internal.GetTaskuluApi().CreateSession(), s.handler.CreateSession)
	s.engine.GET(internal.GetTaskuluApi().GetActivities("123456"), s.handler.GetActivities)
	s.engine.GET("/v1/webhooks/", s.handler.BaleIntegration)
}
