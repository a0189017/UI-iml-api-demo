// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
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
        "/api/userDel": {
            "delete": {
                "description": "使用者刪除",
                "produces": [
                    "text/json"
                ],
                "tags": [
                    "userDel"
                ],
                "summary": "userDel",
                "operationId": "6",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/userDetail": {
            "get": {
                "description": "取得使用者資訊",
                "produces": [
                    "text/json"
                ],
                "tags": [
                    "UserInfo"
                ],
                "summary": "UserDetail",
                "operationId": "3",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/userList/{page}": {
            "get": {
                "description": "取得使用者列表",
                "produces": [
                    "text/json"
                ],
                "tags": [
                    "UserInfo"
                ],
                "summary": "userList",
                "operationId": "2",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/userSearch/{key}": {
            "get": {
                "description": "搜尋使用者",
                "produces": [
                    "text/json"
                ],
                "tags": [
                    "UserInfo"
                ],
                "summary": "userSearch",
                "operationId": "4",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "key",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/userSignin": {
            "post": {
                "description": "使用者登入",
                "produces": [
                    "text/json"
                ],
                "tags": [
                    "Signin/Signup"
                ],
                "summary": "userSignin",
                "operationId": "1",
                "parameters": [
                    {
                        "type": "string",
                        "description": "acct",
                        "name": "acct",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "pwd",
                        "name": "pwd",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/userSignup": {
            "post": {
                "description": "使用者註冊",
                "produces": [
                    "text/json"
                ],
                "tags": [
                    "Signin/Signup"
                ],
                "summary": "userSignup",
                "operationId": "5",
                "parameters": [
                    {
                        "type": "string",
                        "description": "acct",
                        "name": "acct",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "pwd",
                        "name": "pwd",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "fullname",
                        "name": "fullname",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/userUpdate": {
            "put": {
                "description": "使用者資訊更新",
                "produces": [
                    "text/json"
                ],
                "tags": [
                    "userUpdate"
                ],
                "summary": "userUpdate",
                "operationId": "7",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "pwd",
                        "name": "pwd",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "fullname",
                        "name": "fullname",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/userUpdateFullname": {
            "patch": {
                "description": "使用者fullname更新",
                "produces": [
                    "text/json"
                ],
                "tags": [
                    "userUpdate"
                ],
                "summary": "userUpdateFullname",
                "operationId": "8",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "fullname",
                        "name": "fullname",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:3000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "mux swagger",
	Description:      "mux swagger",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
