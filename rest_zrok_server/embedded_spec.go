// Code generated by go-swagger; DO NOT EDIT.

package rest_zrok_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/zrok.v1+json"
  ],
  "produces": [
    "application/zrok.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "zrok client access",
    "title": "zrok",
    "version": "1.0.0"
  },
  "paths": {
    "/account": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "createAccount",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/accountRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "account created",
            "schema": {
              "$ref": "#/definitions/accountResponse"
            }
          },
          "400": {
            "description": "account not created (already exists)"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/enable": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "enable",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/enableRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "environment enabled",
            "schema": {
              "$ref": "#/definitions/enableResponse"
            }
          },
          "404": {
            "description": "account not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/version": {
      "get": {
        "tags": [
          "metadata"
        ],
        "operationId": "version",
        "responses": {
          "200": {
            "description": "retrieve the current server version",
            "schema": {
              "$ref": "#/definitions/version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "accountRequest": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "accountResponse": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "enableRequest": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "enableResponse": {
      "type": "object",
      "properties": {
        "cfg": {
          "type": "string"
        },
        "identity": {
          "type": "string"
        }
      }
    },
    "version": {
      "type": "object",
      "properties": {
        "version": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/zrok.v1+json"
  ],
  "produces": [
    "application/zrok.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "zrok client access",
    "title": "zrok",
    "version": "1.0.0"
  },
  "paths": {
    "/account": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "createAccount",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/accountRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "account created",
            "schema": {
              "$ref": "#/definitions/accountResponse"
            }
          },
          "400": {
            "description": "account not created (already exists)"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/enable": {
      "post": {
        "tags": [
          "identity"
        ],
        "operationId": "enable",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/enableRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "environment enabled",
            "schema": {
              "$ref": "#/definitions/enableResponse"
            }
          },
          "404": {
            "description": "account not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/version": {
      "get": {
        "tags": [
          "metadata"
        ],
        "operationId": "version",
        "responses": {
          "200": {
            "description": "retrieve the current server version",
            "schema": {
              "$ref": "#/definitions/version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "accountRequest": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "accountResponse": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "enableRequest": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "enableResponse": {
      "type": "object",
      "properties": {
        "cfg": {
          "type": "string"
        },
        "identity": {
          "type": "string"
        }
      }
    },
    "version": {
      "type": "object",
      "properties": {
        "version": {
          "type": "string"
        }
      }
    }
  }
}`))
}
