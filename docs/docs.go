// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/json": {
            "post": {
                "description": "get username\u0026password from user input",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GeneralHandler"
                ],
                "summary": "solve username\u0026password",
                "parameters": [
                    {
                        "description": "input message",
                        "name": "inputMessage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTP404Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTP5xxError"
                        }
                    }
                }
            }
        },
        "/redis/list/inQueue": {
            "post": {
                "description": "add elements to a redis db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HttpRedis"
                ],
                "summary": "Http redis inQueue",
                "parameters": [
                    {
                        "description": "message",
                        "name": "massage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListInQueueStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.FmtResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "type": "integer"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTP404Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTP5xxError"
                        }
                    }
                }
            }
        },
        "/redis/list/len": {
            "post": {
                "description": "show redis db len by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HttpRedis"
                ],
                "summary": "Http redis len",
                "parameters": [
                    {
                        "description": "db name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListLenOutQueueStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.FmtResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "integer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTP404Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTP5xxError"
                        }
                    }
                }
            }
        },
        "/redis/list/outQueue": {
            "post": {
                "description": "remove elements from a redis db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HttpRedis"
                ],
                "summary": "Http redis outQueue",
                "parameters": [
                    {
                        "description": "db name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListLenOutQueueStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.FmtResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTP404Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTP5xxError"
                        }
                    }
                }
            }
        },
        "/tools/file/upload": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "upload file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message\":200,\"size\": 4}",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httputil.HTTP404Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 404
                },
                "message": {
                    "type": "string",
                    "example": "Not find"
                }
            }
        },
        "httputil.HTTP5xxError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "message": {
                    "type": "string",
                    "example": "Server internal err"
                }
            }
        },
        "httputil.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "model.FmtResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "model.ListInQueueStruct": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "element": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        2,
                        3
                    ]
                },
                "name": {
                    "type": "string",
                    "example": "list name"
                }
            }
        },
        "model.ListLenOutQueueStruct": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "list name"
                }
            }
        },
        "model.ResponseMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "format": "string",
                    "example": "success"
                },
                "size": {
                    "type": "integer",
                    "example": 4
                }
            }
        },
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "123"
                },
                "user_name": {
                    "type": "string",
                    "example": "xiaoming"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "www.hyh.com:8000",
	BasePath:    "/v1/api",
	Schemes:     []string{},
	Title:       "HTTP redis queue API",
	Description: "This is a http redis queue API",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
