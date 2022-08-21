// Package docs GENERATED BY SWAG; DO NOT EDIT
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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/articles": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "Lists the top 50 records",
                "responses": {}
            }
        },
        "/articles/by/sourceid": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "Finds the articles based on the SourceID provided.  Returns the top 50.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Source ID UUID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/articles/by/tag": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "Finds the articles based on the SourceID provided.  Returns the top 50.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Tag name",
                        "name": "tag",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/articles/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "Returns an article based on defined ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/config/sources": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config",
                    "Source"
                ],
                "summary": "Lists the top 50 records",
                "responses": {}
            }
        },
        "/config/sources/by/source": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config",
                    "Source"
                ],
                "summary": "Lists the top 50 records based on the name given. Example: reddit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Source Name",
                        "name": "source",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/config/sources/new/reddit": {
            "post": {
                "tags": [
                    "Config",
                    "Source",
                    "Reddit"
                ],
                "summary": "Creates a new reddit source to monitor.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "url",
                        "name": "url",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/config/sources/new/twitch": {
            "post": {
                "tags": [
                    "Config",
                    "Source",
                    "Twitch"
                ],
                "summary": "Creates a new twitch source to monitor.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/config/sources/new/youtube": {
            "post": {
                "tags": [
                    "Config",
                    "Source",
                    "YouTube"
                ],
                "summary": "Creates a new youtube source to monitor.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "url",
                        "name": "url",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/config/sources/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config",
                    "Source"
                ],
                "summary": "Returns a single entity by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "tags": [
                    "Config",
                    "Source"
                ],
                "summary": "Deletes a record by ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/config/sources/{id}/disable": {
            "post": {
                "tags": [
                    "Config",
                    "Source"
                ],
                "summary": "Disables a source from processing.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/config/sources/{id}/enable": {
            "post": {
                "tags": [
                    "Config",
                    "Source"
                ],
                "summary": "Enables a source to continue processing.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/discord/queue": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Debug",
                    "Discord",
                    "Queue"
                ],
                "summary": "Returns the top 100 entries from the queue to be processed.",
                "responses": {}
            }
        },
        "/discord/webhooks": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config",
                    "Discord",
                    "Webhook"
                ],
                "summary": "Returns the top 100 entries from the queue to be processed.",
                "responses": {}
            }
        },
        "/discord/webhooks/byId": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config",
                    "Discord",
                    "Webhook"
                ],
                "summary": "Returns the top 100 entries from the queue to be processed.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/discord/webhooks/new": {
            "post": {
                "tags": [
                    "Config",
                    "Discord",
                    "Webhook"
                ],
                "summary": "Creates a new record for a discord web hook to post data to.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "url",
                        "name": "url",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Server name",
                        "name": "server",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Channel name",
                        "name": "channel",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/hello/{who}": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Debug"
                ],
                "summary": "Responds back with \"Hello x\" depending on param passed in.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Who",
                        "name": "who",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/helloworld": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Debug"
                ],
                "summary": "Responds back with \"Hello world!\"",
                "responses": {}
            }
        },
        "/ping": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Debug"
                ],
                "summary": "Sends back \"pong\".  Good to test with.",
                "responses": {}
            }
        },
        "/settings/{key}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Settings"
                ],
                "summary": "Returns a object based on the Key that was given.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Settings Key value",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/subscriptions": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config",
                    "Subscription"
                ],
                "summary": "Returns the top 100 entries from the queue to be processed.",
                "responses": {}
            }
        },
        "/subscriptions/byDiscordId": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config",
                    "Subscription"
                ],
                "summary": "Returns the top 100 entries from the queue to be processed.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/subscriptions/bySourceId": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config",
                    "Subscription"
                ],
                "summary": "Returns the top 100 entries from the queue to be processed.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/subscriptions/new/discordwebhook": {
            "post": {
                "tags": [
                    "Config",
                    "Source",
                    "Discord",
                    "Subscription"
                ],
                "summary": "Creates a new subscription to link a post from a Source to a DiscordWebHook.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "discordWebHookId",
                        "name": "discordWebHookId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "sourceId",
                        "name": "sourceId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "NewsBot collector",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
