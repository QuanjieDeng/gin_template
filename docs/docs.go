// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-03-20 14:09:16.011991577 +0800 CST m=+0.044354307

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/bind": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取多个绑定关系",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "游戏名称",
                        "name": "gameid",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "游戏名称",
                        "name": "serviceid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "添加绑定关系",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "游戏名称",
                        "name": "game",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "服务ID",
                        "name": "serviceid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "游戏ID",
                        "name": "gameid",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/bind/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取单个绑定信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除绑定关系",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "绑定关系ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/client/svc": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取服务商信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "游戏用户appid",
                        "name": "appid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "游戏用户appkey",
                        "name": "appkey",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "服务商名称",
                        "name": "svcname",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/game": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取多个游戏信息",
                "parameters": [
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "description": "账号状态0 锁定 1 正常",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "游戏名称",
                        "name": "game",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "添加游戏",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "游戏名称",
                        "name": "game",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "enum": [
                            1,
                            7
                        ],
                        "type": "integer",
                        "description": "游戏类型",
                        "name": "gametype",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "手机号码",
                        "name": "telnum",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "创建人",
                        "name": "createdby",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "description": "状态",
                        "name": "state",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/game/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取单个游戏信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改游戏信息",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "游戏ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "description": "游戏名称",
                        "name": "game",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            1,
                            2,
                            3,
                            4,
                            5,
                            6,
                            7
                        ],
                        "description": "游戏类型",
                        "name": "gametype",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "description": "手机号码",
                        "name": "telnum",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 100,
                        "description": "本次修改人",
                        "name": "modifiedby",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            "0",
                            "1"
                        ],
                        "description": "状态",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除游戏",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "游戏ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/svc": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取多个服务商信息",
                "parameters": [
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "description": "服务状态0 锁定 1 正常",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "服务商名称",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "添加服务商信息",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "服务商名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "服务商appid",
                        "name": "appid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "服务商appkey",
                        "name": "appkey",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "服务商userid",
                        "name": "userid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "服务商url",
                        "name": "url",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "maxLength": 300,
                        "type": "string",
                        "description": "服务商描述信息",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "本次修改人",
                        "name": "createdby",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "description": "状态",
                        "name": "state",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/svc/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取单个服务商信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改服务商信息",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "服务商ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "description": "服务商名称",
                        "name": "name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 100,
                        "description": "服务商appid",
                        "name": "appid",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 100,
                        "description": "服务商appkey",
                        "name": "appkey",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 100,
                        "description": "服务商userid",
                        "name": "userid",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 100,
                        "description": "服务商url",
                        "name": "url",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 300,
                        "description": "服务商描述信息",
                        "name": "description",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 100,
                        "description": "本次修改人",
                        "name": "modifiedby",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "description": "状态",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除服务商",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "服务商ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/auth": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "根据用户名和密码获取token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "失败",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
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
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "GVP API",
	Description: "gvp是语音平台的后台管理接口服务器",
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
