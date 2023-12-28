package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"os"

	env "github.com/1704mori/certguardian/internal"
	"github.com/1704mori/certguardian/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const FRONTEND_DIR = "../frontend"
const INDEX = "/build/index.html"

type Server struct {
	db     *db.Database
	router *gin.Engine
}

func NewServer(db *db.Database) *Server {
	s := &Server{
		db:     db,
		router: gin.Default(),
	}

	s.routes()

	return s
}

func buildDir() string {
	path, _ := os.Getwd()
	fullDir := path + FRONTEND_DIR
	if os.Getenv("GIN_MODE") == "release" {
		fullDir = path + "/frontend"
	}

	return fullDir
}

func (s *Server) routes() {
	v1 := s.router.Group("/v1")
	setupRoutes(v1, s.db)

	sub, err := fs.Sub(os.DirFS(buildDir()), "build")
	if err != nil {
		panic(err)
	}

	s.router.GET("/_app/*filepath", func(c *gin.Context) {
		http.FileServer(http.FS(sub)).ServeHTTP(c.Writer, c.Request)
	})

	s.router.NoRoute(func(c *gin.Context) {
		executeTemplate(c)
	})

	s.router.Use(corsMiddleware())
}

func (s *Server) Run(addr string) {
	s.router.Run(addr)
}

func executeTemplate(c *gin.Context) {
	file, err := os.Open(buildDir() + "/build/index.html")
	if err != nil {
		log.Panic().Msgf("Failed to open index.html: %v", err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Panic().Msgf("Failed to read index.html: %v", err)
	}

	file.Close()

	mashaled, err := json.Marshal(env.Args)
	if err != nil {
		log.Panic().Msgf("Failed to marshal config: %v", err)
	}

	tmpl, err := template.New(buildDir() + "/build/index.html").Funcs(template.FuncMap{
		"env": func() template.JS {
			return template.JS(fmt.Sprintf(`{"env": %s}`, string(mashaled)))
		},
	}).Parse(string(bytes))
	if err != nil {
		log.Panic().Msgf("Failed to parse index.html: %v", err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Panic().Msgf("Failed to execute index.html: %v", err)
	}

	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		log.Panic().Msgf("Failed to execute index.html: %v", err)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
