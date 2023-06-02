// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (

	"github.com/swaggo/swag"
)

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Sidarta Silva",
            "url": "http://www.sidartaoss.com",
            "email": "atendimento@sidartaoss.com"
        },
        "license": {
            "name": "Full Cycle License",
            "url": "http://www.fullcycle.com.br"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/drivers": {
            "get": {
                "description": "List all drivers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "drivers"
                ],
                "summary": "List drivers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Driver"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/drivers/{id}": {
            "get": {
                "description": "Get a driver by ID (uuid)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "drivers"
                ],
                "summary": "Get a driver",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "driver ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Driver"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Driver": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Sidarta S"
                },
                "uuid": {
                    "type": "string",
                    "example": "127a2285-0fd8-4df9-b4eb-485292f3e345"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "host.docker.internal:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Maratona Full Cycle Driver API",
	Description:      "Driver API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
