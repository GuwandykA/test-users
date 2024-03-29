// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users/add": {
            "post": {
                "description": "create and update data  user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "parameters": [
                    {
                        "description": "Create and Update JSON",
                        "name": "users",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.ReqUser"
                        }
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
        "/users/delete": {
            "post": {
                "description": "delete data  user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "parameters": [
                    {
                        "description": "Delete JSON",
                        "name": "users",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.ReqIdDTO"
                        }
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
        "/users/get-all": {
            "post": {
                "description": "all data users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "parameters": [
                    {
                        "description": "Get All JSON",
                        "name": "users",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.PaginationDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/users.DataDTO"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "appresult.AppError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "appresult.AppSuccess": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "users.DataDTO": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/users.UserDTO"
                    }
                }
            }
        },
        "users.PaginationDTO": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                }
            }
        },
        "users.ReqIdDTO": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer"
                }
            }
        },
        "users.ReqUser": {
            "type": "object",
            "required": [
                "name",
                "surname"
            ],
            "properties": {
                "Id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "users.UserDTO": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer"
                },
                "age": {
                    "type": "integer"
                },
                "country": {},
                "createdAt": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "probability": {
                    "type": "number"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8082",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Users APIs",
	Description:      "Testing Swagger APIs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
