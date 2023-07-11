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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/add": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Add a new associated with the provided information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Add an associated",
                "parameters": [
                    {
                        "description": "Associated data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/route.AssociatedRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/route.AssociatedResponse"
                        }
                    }
                }
            }
        },
        "/admin/delete": {
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Delete an associated based on the provided ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Delete an associated",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the associated to delete",
                        "name": "ID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Deleted successfully!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/update": {
            "put": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Update an existing associated with the provided information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Update an associated",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The ID of the associated",
                        "name": "ID",
                        "in": "query"
                    },
                    {
                        "description": "Updated associated data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/route.UpdateAssociatedRequest"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/route.UpdateAssociatedResponse"
                        }
                    }
                }
            }
        },
        "/user/get": {
            "get": {
                "description": "Get a list of associateds based on the provided order, locate, and amount.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get associateds",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The order to sort the associateds: 1 = Ascendant; 2 = By location; 3 = Descendant",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "The location of the associateds",
                        "name": "locate",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "The amount of associateds to return",
                        "name": "amount",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/route.AssociatedResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "route.AssociatedRequest": {
            "type": "object",
            "properties": {
                "asscdescription": {
                    "type": "string"
                },
                "asscname": {
                    "type": "string"
                },
                "contactnumber": {
                    "type": "string"
                },
                "descriptionaddr": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "logoimage": {
                    "type": "string"
                },
                "pix": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                }
            }
        },
        "route.AssociatedResponse": {
            "type": "object",
            "properties": {
                "asscdescription": {
                    "type": "string"
                },
                "asscname": {
                    "type": "string"
                },
                "contactnumber": {
                    "type": "string"
                },
                "descriptionaddr": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "logoimage": {
                    "type": "string"
                },
                "pix": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                }
            }
        },
        "route.UpdateAssociatedRequest": {
            "type": "object",
            "properties": {
                "asscdescription": {
                    "type": "string"
                },
                "asscname": {
                    "type": "string"
                },
                "contactnumber": {
                    "type": "string"
                },
                "descriptionaddr": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "logoimage": {
                    "type": "string"
                },
                "pix": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                }
            }
        },
        "route.UpdateAssociatedResponse": {
            "type": "object",
            "properties": {
                "asscdescription": {
                    "type": "string"
                },
                "asscname": {
                    "type": "string"
                },
                "contactnumber": {
                    "type": "string"
                },
                "descriptionaddr": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "logoimage": {
                    "type": "string"
                },
                "pix": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
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
