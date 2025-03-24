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
        "/api/orders": {
            "get": {
                "description": "Find All Orders",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "FindAll Orders",
                "responses": {
                    "200": {
                        "description": "Order list",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Order"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Add and register new Order",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Create an Order",
                "parameters": [
                    {
                        "description": "create order payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PayloadCreateOrder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ResponseOrder"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/orders/checkout": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Checkout Order and Change Status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Checkout Order",
                "responses": {
                    "200": {
                        "description": "Order created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ResponseOrder"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/orders/save": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Save Order Detail ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Save Order Detail ID",
                "responses": {
                    "200": {
                        "description": "Order created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ResponseOrder"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/orders/{id}": {
            "get": {
                "description": "Find Order By ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "FindById an Order",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"67cdcb62a50a990a870d928f\"",
                        "description": "address id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Address found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ResponseOrder"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update Order By ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Update an Order",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"67cdcb62a50a990a870d928f\"",
                        "description": "order id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update order payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PayloadUpdateOrder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order updated",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ResponseOrder"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete Order By ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "DeleteById an Order",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"67cdcb62a50a990a870d928f\"",
                        "description": "order id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order deleted",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/order-details": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Add and register new order detail",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OrderDetail"
                ],
                "summary": "Create an Order Detail",
                "parameters": [
                    {
                        "description": "create order detail payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PayloadCreateOrderDetail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order detail created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ResponseOrderDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/order-details/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Find order detail by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OrderDetail"
                ],
                "summary": "Find an Order Detail by ID",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"67cdcb62a50a990a870d928f\"",
                        "description": "order detail id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order detail found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ResponseOrderDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update order detail by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OrderDetail"
                ],
                "summary": "Update an Order Detail",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"67cdcb62a50a990a870d928f\"",
                        "description": "order detail id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update order detail payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PayloadUpdateOrderDetail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order detail updated",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ResponseOrderDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete order detail by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OrderDetail"
                ],
                "summary": "Delete an Order Detail",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"67cdcb62a50a990a870d928f\"",
                        "description": "order detail id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order detail deleted",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.WebResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {},
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Order": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "driver_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "order_date": {
                    "type": "string"
                },
                "order_detail_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "order_status": {
                    "$ref": "#/definitions/model.OrderStatus"
                },
                "shipping_status": {
                    "$ref": "#/definitions/model.ShippingStatus"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_shipping": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.OrderStatus": {
            "type": "string",
            "enum": [
                "draft",
                "pending",
                "paid",
                "rejected"
            ],
            "x-enum-varnames": [
                "OrderStatusDraft",
                "OrderStatusPending",
                "OrderStatusPaid",
                "OrderStatusRejected"
            ]
        },
        "model.PayloadCreateOrder": {
            "description": "PayloadCreateOrder details.",
            "type": "object",
            "properties": {
                "note": {
                    "type": "string",
                    "example": "mohon hati hati karena terdapat barang pecah belah"
                }
            }
        },
        "model.PayloadCreateOrderDetail": {
            "type": "object",
            "required": [
                "destination_address_id",
                "origin_address_id",
                "recycle_hub_id",
                "waste_weight"
            ],
            "properties": {
                "destination_address_id": {
                    "type": "string"
                },
                "origin_address_id": {
                    "type": "string"
                },
                "recycle_hub_id": {
                    "type": "string"
                },
                "waste_weight": {
                    "type": "number"
                }
            }
        },
        "model.PayloadUpdateOrder": {
            "description": "PayloadUpdateOrder details.",
            "type": "object",
            "properties": {
                "note": {
                    "type": "string",
                    "example": "40.7128"
                },
                "orderDetailIDs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.PayloadUpdateOrderDetail": {
            "type": "object",
            "required": [
                "destination_address_id",
                "origin_address_id",
                "waste_weight"
            ],
            "properties": {
                "destination_address_id": {
                    "type": "string"
                },
                "origin_address_id": {
                    "type": "string"
                },
                "waste_weight": {
                    "type": "number"
                }
            }
        },
        "model.ResponseOrder": {
            "description": "ResponseOrder details.",
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "driver_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "order_date": {
                    "type": "string"
                },
                "order_detail_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "order_status": {
                    "$ref": "#/definitions/model.OrderStatus"
                },
                "shipping_status": {
                    "$ref": "#/definitions/model.ShippingStatus"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_shipping": {
                    "type": "string"
                }
            }
        },
        "model.ResponseOrderDetail": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "delivery_price": {
                    "type": "number"
                },
                "destination_address_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "origin_address_id": {
                    "type": "string"
                },
                "recycle_hub_id": {
                    "type": "string"
                },
                "sub_total": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "waste_weight": {
                    "type": "number"
                }
            }
        },
        "model.ShippingStatus": {
            "type": "string",
            "enum": [
                "unassigned",
                "pickup",
                "delivery",
                "finish"
            ],
            "x-enum-varnames": [
                "ShippingStatusUnassigned",
                "ShippingStatusPickup",
                "ShippingStatusDelivery",
                "ShippingStatusFinish"
            ]
        },
        "model.WebResponse": {
            "description": "represents the standard API response format.",
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Operation Service API",
	Description:      "This is the documentation of Operation Service API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
