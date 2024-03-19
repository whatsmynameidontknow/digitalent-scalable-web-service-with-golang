// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/orders": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "list all orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-array_dto_OrderResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "create a new order",
                "parameters": [
                    {
                        "description": "required body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.OrderRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-dto_OrderCreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    }
                }
            }
        },
        "/orders/{orderId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "get a specific order by their ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-dto_OrderResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "update an order by their ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "required body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.OrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "delete order by their ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.Response-any"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ItemRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "a thing I need to order"
                },
                "itemCode": {
                    "type": "string",
                    "example": "123"
                },
                "quantity": {
                    "type": "integer",
                    "example": 69
                }
            }
        },
        "dto.ItemResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "item_code": {
                    "type": "string"
                },
                "item_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "dto.OrderCreateResponse": {
            "type": "object",
            "properties": {
                "order_id": {
                    "type": "integer"
                }
            }
        },
        "dto.OrderRequest": {
            "type": "object",
            "properties": {
                "customerName": {
                    "type": "string",
                    "example": "John Doe"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ItemRequest"
                    }
                },
                "orderedAt": {
                    "type": "string",
                    "example": "2019-11-09T21:21:46+00:00"
                }
            }
        },
        "dto.OrderResponse": {
            "type": "object",
            "properties": {
                "customer_name": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ItemResponse"
                    }
                },
                "order_id": {
                    "type": "integer"
                },
                "ordered_at": {
                    "type": "string"
                }
            }
        },
        "helper.Response-any": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "helper.Response-array_dto_OrderResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.OrderResponse"
                    }
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "helper.Response-dto_OrderCreateResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dto.OrderCreateResponse"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "helper.Response-dto_OrderResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dto.OrderResponse"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Hacktiv8-Golang assignment-02",
	Description:      "submission for assignment-02",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}