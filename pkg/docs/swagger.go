package docs

import (
	"github.com/swaggo/swag"
)

// @title E-Commerce API
// @version 1.0
// @description E-Commerce API Documentation
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func init() {
	swag.Register(swag.Name, &s{})
}

type s struct{}

func (s *s) ReadDoc() string {
	return `{
		"swagger": "2.0",
		"info": {
			"description": "E-Commerce API Documentation",
			"version": "1.0",
			"title": "E-Commerce API",
			"contact": {
				"email": "support@swagger.io"
			},
			"license": {
				"name": "Apache 2.0",
				"url": "http://www.apache.org/licenses/LICENSE-2.0.html"
			}
		},
		"host": "localhost:8080",
		"basePath": "/api/v1",
		"schemes": ["http", "https"],
		"securityDefinitions": {
			"ApiKeyAuth": {
				"type": "apiKey",
				"name": "Authorization",
				"in": "header"
			}
		}
	}`
}
