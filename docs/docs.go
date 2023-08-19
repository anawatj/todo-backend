// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/tasks": {
            "get": {
                "description": "Getting list task from database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get list task",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tasks.ListResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Adding new task to database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create task",
                "parameters": [
                    {
                        "description": "Task Data",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/tasks.TaskValidator"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/tasks.ListResponse"
                        }
                    }
                }
            }
        },
        "/tasks/{id}": {
            "get": {
                "description": "getting task by id from database",
                "consumes": [
                    "application/json"
                ],
                "summary": "get task by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tasks.ListResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "update task to database",
                "consumes": [
                    "application/json"
                ],
                "summary": "update task",
                "parameters": [
                    {
                        "description": "Task Data",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/tasks.TaskValidator"
                        }
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tasks.ListResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete task from database",
                "consumes": [
                    "application/json"
                ],
                "summary": "delete task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tasks.ListResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "tasks.ListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/tasks.TaskResponse"
                    }
                }
            }
        },
        "tasks.TaskResponse": {
            "type": "object",
            "properties": {
                "createAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "tasks.TaskValidator": {
            "type": "object",
            "required": [
                "description",
                "image",
                "status",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Zog_festiv eCommerce API",
	Description:      "API for ecommerce website",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	//LeftDelim:        "{{",
	//RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
