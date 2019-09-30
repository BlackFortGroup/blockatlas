// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-09-30 17:34:28.798536 -0300 -03 m=+0.048014579

package docs

import (
	"bytes"
	"encoding/json"

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
        "/metrics": {
            "get": {
                "description": "Get application metrics",
                "tags": [
                    "metrics"
                ],
                "summary": "Get Metrics",
                "operationId": "metrics"
            }
        },
        "/observer/v1/status": {
            "get": {
                "description": "Get coin status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "observer",
                    "subscriptions"
                ],
                "summary": "Get coin status",
                "operationId": "coin_status",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer test",
                        "description": "Bearer authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.CoinStatus"
                        }
                    }
                }
            }
        },
        "/observer/v1/webhook/register": {
            "post": {
                "description": "Create a webhook for addresses transactions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "observer",
                    "subscriptions"
                ],
                "summary": "Create a webhook",
                "operationId": "create_webhook",
                "parameters": [
                    {
                        "description": "Accounts subscriptions",
                        "name": "subscriptions",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.Webhook"
                        }
                    },
                    {
                        "type": "string",
                        "default": "Bearer test",
                        "description": "Bearer authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ObserverResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a webhook for addresses transactions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "observer",
                    "subscriptions"
                ],
                "summary": "Delete a webhook",
                "operationId": "delete_webhook",
                "parameters": [
                    {
                        "description": "Accounts subscriptions",
                        "name": "subscriptions",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.Webhook"
                        }
                    },
                    {
                        "type": "string",
                        "default": "Bearer test",
                        "description": "Bearer authorization header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ObserverResponse"
                        }
                    }
                }
            }
        },
        "/v1/{coin}/xpub/{xpub}": {
            "get": {
                "description": "Get transactions from xpub address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platform",
                    "tx"
                ],
                "summary": "Get xpub transactions",
                "operationId": "xpub",
                "parameters": [
                    {
                        "type": "string",
                        "default": "bitcoin",
                        "description": "the coin name",
                        "name": "coin",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC",
                        "description": "the xpub address",
                        "name": "xpub",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/blockatlas.TxPage"
                        }
                    }
                }
            }
        },
        "/v1/{coin}/{address}": {
            "get": {
                "description": "Get transactions from the address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platform",
                    "tx"
                ],
                "summary": "Get Transactions",
                "operationId": "tx_v1",
                "parameters": [
                    {
                        "type": "string",
                        "default": "tezos",
                        "description": "the coin name",
                        "name": "coin",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q",
                        "description": "the query address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/v2/staking/delegations": {
            "post": {
                "description": "Get Stake Delegations for multiple coins",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platform",
                    "staking"
                ],
                "summary": "Get Multiple Stake Delegations",
                "operationId": "batch_delegations",
                "parameters": [
                    {
                        "description": "Validators addresses and coins",
                        "name": "delegations",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.AddressesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/blockatlas.DelegationsBatchPage"
                        }
                    }
                }
            }
        },
        "/v2/{coin}/collections/{address}": {
            "get": {
                "description": "Get all collections from the address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platform",
                    "collection"
                ],
                "summary": "Get Collections",
                "operationId": "collections",
                "parameters": [
                    {
                        "type": "string",
                        "default": "ethereum",
                        "description": "the coin name",
                        "name": "coin",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
                        "description": "the query address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/blockatlas.CollectionPage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/v2/{coin}/collections/{address}/collection/": {
            "get": {
                "description": "Get a collection from the address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platform",
                    "collection"
                ],
                "summary": "Get Collection",
                "operationId": "collection",
                "parameters": [
                    {
                        "type": "string",
                        "default": "ethereum",
                        "description": "the coin name",
                        "name": "coin",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
                        "description": "the query address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "0x06012c8cf97bead5deae237070f9587f8e7a266d",
                        "description": "the query collection",
                        "name": "collection_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/blockatlas.CollectionPage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/v2/{coin}/staking/delegations/{address}": {
            "get": {
                "description": "Get stake delegations from the address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platform",
                    "staking"
                ],
                "summary": "Get Stake Delegations",
                "operationId": "delegations",
                "parameters": [
                    {
                        "type": "string",
                        "default": "tron",
                        "description": "the coin name",
                        "name": "coin",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "TPJYCz8ppZNyvw7pTwmjajcx4Kk1MmEUhD",
                        "description": "the query address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/blockatlas.DocsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/v2/{coin}/staking/validators": {
            "get": {
                "description": "Get validators from the address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platform",
                    "staking"
                ],
                "summary": "Get Validators",
                "operationId": "validators",
                "parameters": [
                    {
                        "type": "string",
                        "default": "cosmos",
                        "description": "the coin name",
                        "name": "coin",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/blockatlas.DocsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/v2/{coin}/tokens/{address}": {
            "get": {
                "description": "Get tokens from the address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platform",
                    "token"
                ],
                "summary": "Get Tokens",
                "operationId": "tokens",
                "parameters": [
                    {
                        "type": "string",
                        "default": "ethereum",
                        "description": "the coin name",
                        "name": "coin",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
                        "description": "the query address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/blockatlas.CollectionPage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/v2/{coin}/transactions/{address}": {
            "get": {
                "description": "Get transactions from the address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "platform",
                    "tx"
                ],
                "summary": "Get Transactions",
                "operationId": "tx_v2",
                "parameters": [
                    {
                        "type": "string",
                        "default": "tezos",
                        "description": "the coin name",
                        "name": "coin",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q",
                        "description": "the query address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/blockatlas.TxPage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.AddressBatchRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "coin": {
                    "type": "integer"
                }
            }
        },
        "api.AddressesRequest": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "address": {
                        "type": "string"
                    },
                    "coin": {
                        "type": "integer"
                    }
                }
            }
        },
        "api.ApiError": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_message": {
                    "type": "string"
                }
            }
        },
        "api.CoinStatus": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                }
            }
        },
        "api.ObserverResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "api.Webhook": {
            "type": "object",
            "properties": {
                "subscriptions": {
                    "type": "object"
                },
                "webhook": {
                    "type": "string"
                },
                "xpub_subscriptions": {
                    "type": "object"
                }
            }
        },
        "blockatlas.Collection": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "category_address": {
                    "type": "string"
                },
                "coin": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "external_link": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nft_version": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "blockatlas.CollectionPage": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "address": {
                        "type": "string"
                    },
                    "category_address": {
                        "type": "string"
                    },
                    "coin": {
                        "type": "integer"
                    },
                    "description": {
                        "type": "string"
                    },
                    "external_link": {
                        "type": "string"
                    },
                    "id": {
                        "type": "string"
                    },
                    "image_url": {
                        "type": "string"
                    },
                    "name": {
                        "type": "string"
                    },
                    "nft_version": {
                        "type": "string"
                    },
                    "slug": {
                        "type": "string"
                    },
                    "symbol": {
                        "type": "string"
                    },
                    "total": {
                        "type": "integer"
                    },
                    "type": {
                        "type": "string"
                    }
                }
            }
        },
        "blockatlas.Delegation": {
            "type": "object",
            "properties": {
                "coin": {
                    "type": "object",
                    "$ref": "#/definitions/coin.ExternalCoin"
                },
                "delegator": {
                    "type": "object",
                    "$ref": "#/definitions/blockatlas.StakeValidator"
                },
                "metadata": {
                    "type": "object"
                },
                "status": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "blockatlas.DelegationsBatch": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "coin": {
                    "type": "integer"
                },
                "delegations": {
                    "type": "object",
                    "$ref": "#/definitions/blockatlas.DelegationsPage"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "blockatlas.DelegationsBatchPage": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "address": {
                        "type": "string"
                    },
                    "coin": {
                        "type": "integer"
                    },
                    "delegations": {
                        "type": "object",
                        "$ref": "#/definitions/blockatlas.DelegationsPage"
                    },
                    "error": {
                        "type": "string"
                    }
                }
            }
        },
        "blockatlas.DelegationsPage": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "coin": {
                        "type": "object",
                        "$ref": "#/definitions/coin.ExternalCoin"
                    },
                    "delegator": {
                        "type": "object",
                        "$ref": "#/definitions/blockatlas.StakeValidator"
                    },
                    "metadata": {
                        "type": "object"
                    },
                    "status": {
                        "type": "string"
                    },
                    "value": {
                        "type": "string"
                    }
                }
            }
        },
        "blockatlas.DocsResponse": {
            "type": "object",
            "properties": {
                "docs": {
                    "type": "object"
                }
            }
        },
        "blockatlas.StakeValidator": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "info": {
                    "type": "object",
                    "$ref": "#/definitions/blockatlas.StakeValidatorInfo"
                },
                "locktime": {
                    "type": "integer"
                },
                "minimum_amount": {
                    "type": "string"
                },
                "reward": {
                    "type": "object",
                    "$ref": "#/definitions/blockatlas.StakingReward"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "blockatlas.StakeValidatorInfo": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "blockatlas.StakingReward": {
            "type": "object",
            "properties": {
                "annual": {
                    "type": "number"
                }
            }
        },
        "blockatlas.Tx": {
            "type": "object",
            "properties": {
                "block": {
                    "description": "Height of the block the transaction was included in",
                    "type": "integer"
                },
                "coin": {
                    "description": "SLIP-44 coin index of the platform",
                    "type": "integer"
                },
                "date": {
                    "description": "Unix timestamp of the block the transaction was included in",
                    "type": "integer"
                },
                "direction": {
                    "description": "Transaction Direction",
                    "type": "string"
                },
                "error": {
                    "description": "Empty if the transaction was successful,\nelse error explaining why the transaction failed (optional)",
                    "type": "string"
                },
                "fee": {
                    "description": "Transaction fee (native currency)",
                    "type": "string"
                },
                "from": {
                    "description": "Address of the transaction sender",
                    "type": "string"
                },
                "id": {
                    "description": "Unique identifier",
                    "type": "string"
                },
                "inputs": {
                    "description": "Input addresses",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/blockatlas.TxOutput"
                    }
                },
                "memo": {
                    "description": "Meta data object",
                    "type": "string"
                },
                "metadata": {
                    "type": "object"
                },
                "outputs": {
                    "description": "Output addresses",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/blockatlas.TxOutput"
                    }
                },
                "sequence": {
                    "description": "Transaction nonce or sequence",
                    "type": "integer"
                },
                "status": {
                    "description": "Status of the transaction",
                    "type": "string"
                },
                "to": {
                    "description": "Address of the transaction recipient",
                    "type": "string"
                },
                "type": {
                    "description": "Type of metadata",
                    "type": "string"
                }
            }
        },
        "blockatlas.TxOutput": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "blockatlas.TxPage": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "block": {
                        "description": "Height of the block the transaction was included in",
                        "type": "integer"
                    },
                    "coin": {
                        "description": "SLIP-44 coin index of the platform",
                        "type": "integer"
                    },
                    "date": {
                        "description": "Unix timestamp of the block the transaction was included in",
                        "type": "integer"
                    },
                    "direction": {
                        "description": "Transaction Direction",
                        "type": "string"
                    },
                    "error": {
                        "description": "Empty if the transaction was successful,\nelse error explaining why the transaction failed (optional)",
                        "type": "string"
                    },
                    "fee": {
                        "description": "Transaction fee (native currency)",
                        "type": "string"
                    },
                    "from": {
                        "description": "Address of the transaction sender",
                        "type": "string"
                    },
                    "id": {
                        "description": "Unique identifier",
                        "type": "string"
                    },
                    "inputs": {
                        "description": "Input addresses",
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/blockatlas.TxOutput"
                        }
                    },
                    "memo": {
                        "description": "Meta data object",
                        "type": "string"
                    },
                    "metadata": {
                        "type": "object"
                    },
                    "outputs": {
                        "description": "Output addresses",
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/blockatlas.TxOutput"
                        }
                    },
                    "sequence": {
                        "description": "Transaction nonce or sequence",
                        "type": "integer"
                    },
                    "status": {
                        "description": "Status of the transaction",
                        "type": "string"
                    },
                    "to": {
                        "description": "Address of the transaction recipient",
                        "type": "string"
                    },
                    "type": {
                        "description": "Type of metadata",
                        "type": "string"
                    }
                }
            }
        },
        "coin.ExternalCoin": {
            "type": "object",
            "properties": {
                "coin": {
                    "type": "integer"
                },
                "decimals": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "symbol": {
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
var SwaggerInfo = swaggerInfo{Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
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
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
