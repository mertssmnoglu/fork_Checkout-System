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
        "/api/v1/cart/{cartId}": {
            "get": {
                "description": "Display a cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Display cart",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cart ID",
                        "name": "cartId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Cart displayed successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Reset a cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Reset cart",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cart ID",
                        "name": "cartId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Cart reset successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/cart/{cartId}/promotion/{promotionId}": {
            "post": {
                "description": "Apply a promotion to a cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Apply promotion",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cart ID",
                        "name": "cartId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Promotion ID",
                        "name": "promotionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Promotion applied successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/category": {
            "post": {
                "description": "Create a new category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Create a category",
                "parameters": [
                    {
                        "description": "Category object",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category created successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/category/list": {
            "get": {
                "description": "Get a list of categories",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "List categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Category"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/category/{id}": {
            "get": {
                "description": "Get a category by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get a category by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category ID",
                        "name": "id",
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
        "/api/v1/item": {
            "post": {
                "description": "Create a new item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Create an item",
                "parameters": [
                    {
                        "description": "Item object",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Item"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Item created successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/item/list": {
            "get": {
                "description": "Get a list of items",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "List items",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int",
                        "description": "Cart ID",
                        "name": "cart_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Item"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/item/{id}": {
            "get": {
                "description": "Get an item by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Get an item by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int",
                        "description": "Item ID",
                        "name": "id",
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
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an item by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Delete an item by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/promotion": {
            "post": {
                "description": "Create a new promotion",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Promotion"
                ],
                "summary": "Create a promotion",
                "parameters": [
                    {
                        "description": "Promotion object",
                        "name": "promotion",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Promotion"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Promotion created successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/promotion/list": {
            "get": {
                "description": "Get a list of promotions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Promotion"
                ],
                "summary": "List promotions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Promotion"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/promotion/{id}": {
            "get": {
                "description": "Get a promotion by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Promotion"
                ],
                "summary": "Get a promotion by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Promotion ID",
                        "name": "id",
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
        "/api/v1/vasitem": {
            "post": {
                "description": "Create a new vas item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "VasItem"
                ],
                "summary": "Create a vas item",
                "parameters": [
                    {
                        "description": "Vas Item object",
                        "name": "vas_item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.VasItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Vas Item created successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/vasitem/list": {
            "get": {
                "description": "Get a list of vas items",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "VasItem"
                ],
                "summary": "List vas items",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int",
                        "description": "Item ID",
                        "name": "item_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.VasItem"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/vasitem/{id}": {
            "get": {
                "description": "Get a vas item",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "VasItem"
                ],
                "summary": "Get vas item",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int",
                        "description": "Vas Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.VasItem"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a vas item",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "VasItem"
                ],
                "summary": "Delete a vas item",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int",
                        "description": "Vas Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Vas Item deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Category": {
            "type": "object",
            "required": [
                "itemType",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "itemType": {
                    "type": "integer",
                    "enum": [
                        1,
                        2
                    ]
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.CategoryPromotionDiscount": {
            "type": "object",
            "required": [
                "categoryID",
                "discountRate"
            ],
            "properties": {
                "categoryID": {
                    "type": "integer"
                },
                "discountRate": {
                    "type": "number",
                    "maximum": 100,
                    "minimum": 0
                }
            }
        },
        "entity.Item": {
            "type": "object",
            "required": [
                "cartId",
                "categoryId",
                "price",
                "quantity",
                "sellerId"
            ],
            "properties": {
                "cartId": {
                    "type": "integer"
                },
                "categoryId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "itemType": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer",
                    "maximum": 10
                },
                "sellerId": {
                    "type": "integer"
                }
            }
        },
        "entity.Promotion": {
            "type": "object",
            "required": [
                "promotionType"
            ],
            "properties": {
                "categoryPromotion": {
                    "$ref": "#/definitions/entity.CategoryPromotionDiscount"
                },
                "id": {
                    "type": "integer"
                },
                "promotionType": {
                    "enum": [
                        1,
                        2,
                        3
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/entity.PromotionType"
                        }
                    ]
                },
                "sameSellerPromotion": {
                    "$ref": "#/definitions/entity.SameSellerPromotionDiscount"
                },
                "totalPricePromotions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.TotalPricePromotionDiscount"
                    }
                }
            }
        },
        "entity.PromotionType": {
            "type": "integer",
            "enum": [
                1,
                2,
                3,
                1,
                2
            ],
            "x-enum-varnames": [
                "SameSellerPromotion",
                "CategoryPromotion",
                "TotalPricePromotion",
                "DigitalItem",
                "DefaultItem"
            ]
        },
        "entity.SameSellerPromotionDiscount": {
            "type": "object",
            "required": [
                "discountRate"
            ],
            "properties": {
                "discountRate": {
                    "type": "number",
                    "maximum": 100,
                    "minimum": 0
                }
            }
        },
        "entity.TotalPricePromotionDiscount": {
            "type": "object",
            "required": [
                "discountAmount",
                "priceRangeEnd",
                "priceRangeStart"
            ],
            "properties": {
                "discountAmount": {
                    "type": "number"
                },
                "priceRangeEnd": {
                    "type": "number"
                },
                "priceRangeStart": {
                    "type": "number"
                }
            }
        },
        "entity.VasItem": {
            "type": "object",
            "required": [
                "categoryId",
                "itemId",
                "price",
                "quantity",
                "sellerId"
            ],
            "properties": {
                "categoryId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "itemId": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer",
                    "maximum": 3,
                    "minimum": 1
                },
                "sellerId": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
