// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://developers.ctd.com.ar/es_ar/terminos-y-condiciones",
        "contact": {
            "name": "API Support",
            "url": "https://developers.ctd.com.ar/support"
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
        "/odontologos/all": {
            "get": {
                "description": "Retrieve all odontologos",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "Get all odontologos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Odontologo"
                            }
                        }
                    }
                }
            }
        },
        "/odontologos/create": {
            "post": {
                "description": "Create a new odontologo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "Create a new odontologo",
                "parameters": [
                    {
                        "description": "Odontologo object",
                        "name": "odontologo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Odontologo"
                        }
                    },
                    {
                        "type": "string",
                        "description": "TOKEN",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Odontologo"
                        }
                    }
                }
            }
        },
        "/odontologos/delete/{id}": {
            "delete": {
                "description": "Delete an odontologo by ID",
                "tags": [
                    "Odontologo"
                ],
                "summary": "Delete an odontologo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Odontologo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "TOKEN",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Odontologo deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/odontologos/update/{id}": {
            "put": {
                "description": "Update an existing odontologo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "Update an existing odontologo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Odontologo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Odontologo object",
                        "name": "odontologo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Odontologo"
                        }
                    },
                    {
                        "type": "string",
                        "description": "TOKEN",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated",
                        "schema": {
                            "$ref": "#/definitions/domain.Odontologo"
                        }
                    }
                }
            }
        },
        "/odontologos/updateField/{id}": {
            "patch": {
                "description": "Update a specific field of an odontologo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "Update a specific field of an odontologo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Odontologo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Field name",
                        "name": "field_name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "New value",
                        "name": "value",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "TOKEN",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Field updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/odontologos/{id}": {
            "get": {
                "description": "Retrieve an odontologo by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "Get an odontologo by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Odontologo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Odontologo"
                        }
                    }
                }
            }
        },
        "/pacientes/all": {
            "get": {
                "description": "Retrieve all patients",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Get all patients",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Paciente"
                            }
                        }
                    }
                }
            }
        },
        "/pacientes/create": {
            "post": {
                "description": "Create a new patient",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Create a new patient",
                "parameters": [
                    {
                        "description": "Patient object",
                        "name": "patient",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    },
                    {
                        "type": "string",
                        "description": "TOKEN",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    }
                }
            }
        },
        "/pacientes/delete/{id}": {
            "delete": {
                "description": "Delete a patient by ID",
                "tags": [
                    "Paciente"
                ],
                "summary": "Delete a patient",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "TOKEN",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Paciente deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pacientes/update/{id}": {
            "put": {
                "description": "Update an existing patient",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Update an existing patient",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Patient object",
                        "name": "patient",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    },
                    {
                        "type": "string",
                        "description": "TOKEN",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated",
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    }
                }
            }
        },
        "/pacientes/updateField/{id}": {
            "patch": {
                "description": "Update a specific field of a patient",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Update a specific field of a patient",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Field name",
                        "name": "field_name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "New value",
                        "name": "value",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "TOKEN",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Field updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pacientes/{id}": {
            "get": {
                "description": "Retrieve a patient by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Get a patient by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    }
                }
            }
        },
        "/turnos/all": {
            "get": {
                "description": "Retrieve all turnos",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Get all turnos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Turno"
                            }
                        }
                    }
                }
            }
        },
        "/turnos/byDNI": {
            "get": {
                "description": "Retrieve turnos by paciente's DNI",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Get turnos by paciente's DNI",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Paciente DNI",
                        "name": "dni",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/turnos/create": {
            "post": {
                "description": "Create a new turno",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Create a new turno",
                "parameters": [
                    {
                        "description": "Turno object",
                        "name": "turno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Turno"
                        }
                    },
                    {
                        "type": "string",
                        "description": "TOKEN",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Turno"
                        }
                    }
                }
            }
        },
        "/turnos/createByDniAndMatricula": {
            "post": {
                "description": "Create a new turno by Paciente's DNI and Odontologo's Matricula",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Create a new turno by Paciente's DNI and Odontologo's Matricula",
                "parameters": [
                    {
                        "description": "Turno data",
                        "name": "turno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/turnos.CreateTurnoData"
                        }
                    },
                    {
                        "type": "string",
                        "description": "TOKEN",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Turno"
                        }
                    }
                }
            }
        },
        "/turnos/update/{id}": {
            "put": {
                "description": "Update an existing turno",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Update an existing turno",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Turno ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Turno object",
                        "name": "turno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Turno"
                        }
                    },
                    {
                        "type": "string",
                        "description": "TOKEN",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated",
                        "schema": {
                            "$ref": "#/definitions/domain.Turno"
                        }
                    }
                }
            }
        },
        "/turnos/{id}": {
            "get": {
                "description": "Retrieve a turno by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Get a turno by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Turno ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Turno"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a turno by ID",
                "tags": [
                    "Turno"
                ],
                "summary": "Delete a turno",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Turno ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "TOKEN",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Turno deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/turnos/{id}/update-field": {
            "put": {
                "description": "Update a specific field of a turno by ID",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Update a specific field of a turno",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Turno ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authentication Token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Name of the field to update",
                        "name": "field_name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "New value for the field",
                        "name": "value",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "type": "string",
                        "description": "TOKEN",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Field updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.CustomTime": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        },
        "domain.Odontologo": {
            "type": "object",
            "properties": {
                "apellido": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "matricula": {
                    "type": "string"
                },
                "nombre": {
                    "type": "string"
                }
            }
        },
        "domain.Paciente": {
            "type": "object",
            "properties": {
                "apellido": {
                    "type": "string",
                    "example": "Doe"
                },
                "dni": {
                    "type": "string",
                    "example": "12345678"
                },
                "domicilio": {
                    "type": "string",
                    "example": "123 Main St"
                },
                "fecha_alta": {
                    "$ref": "#/definitions/domain.CustomTime"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "nombre": {
                    "type": "string",
                    "example": "John"
                }
            }
        },
        "domain.Turno": {
            "type": "object",
            "properties": {
                "descripcion": {
                    "type": "string"
                },
                "fecha": {
                    "$ref": "#/definitions/domain.CustomTime"
                },
                "id": {
                    "type": "integer"
                },
                "id_odontologo": {
                    "type": "integer"
                },
                "id_paciente": {
                    "type": "integer"
                }
            }
        },
        "turnos.CreateTurnoData": {
            "type": "object",
            "properties": {
                "descripcion": {
                    "type": "string"
                },
                "dni": {
                    "type": "string"
                },
                "fecha": {
                    "$ref": "#/definitions/domain.CustomTime"
                },
                "matricula": {
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Certified Tech Developer - Back End III",
	Description:      "This API handles Patients and Dentists.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
