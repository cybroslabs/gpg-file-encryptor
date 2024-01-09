// Package openapi Code generated by swaggo/swag. DO NOT EDIT
package openapi

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/cybroslabs/gpg-file-encryptor/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/encrypt": {
            "post": {
                "description": "Encrypt and sign a file supplied in the request body using the public key supplied in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GPG"
                ],
                "summary": "Encrypt and sign a file",
                "parameters": [
                    {
                        "description": "A JSON struct with base64 encoded file (bin) and strings per public key, private key and passphrase",
                        "name": "encryptRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.EncryptRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns JSON with base64-encoded encrypted and signed file",
                        "schema": {
                            "$ref": "#/definitions/handlers.EncryptResponse"
                        }
                    },
                    "400": {
                        "description": "Error response, bad request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Error response, error while processing request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrResponse"
                        }
                    }
                }
            }
        },
        "/v1/sign": {
            "post": {
                "description": "Sign a file supplied in the request body using the public key supplied in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GPG"
                ],
                "summary": "Sign data",
                "parameters": [
                    {
                        "description": "A JSON struct with base64-encoded file (bin) and strings per private key and passphrase",
                        "name": "signOnlyRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.SignRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns JSON with GPG signature (in plain text)",
                        "schema": {
                            "$ref": "#/definitions/handlers.SignResponse"
                        }
                    },
                    "400": {
                        "description": "Error response, bad request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Error while signing file",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.EncryptRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "gpg_passphrase": {
                    "type": "string"
                },
                "gpg_private_key": {
                    "type": "string"
                },
                "gpg_public_key": {
                    "type": "string"
                }
            }
        },
        "handlers.EncryptResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "base64-encoded binary data",
                    "type": "string"
                }
            }
        },
        "handlers.ErrResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "error_message": {
                    "type": "string"
                }
            }
        },
        "handlers.SignRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "gpg_passphrase": {
                    "type": "string"
                },
                "gpg_private_key": {
                    "type": "string"
                }
            }
        },
        "handlers.SignResponse": {
            "type": "object",
            "properties": {
                "gpg_signature": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "GPG Encryptor Service API",
	Description:      "This is a small, lightweight service to encrypt and sign or just sign files using GPG.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}