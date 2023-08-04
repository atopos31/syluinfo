// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "hakcerxiao",
            "url": "http://www.hackerxiao.online"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/coskey": {
            "get": {
                "tags": [
                    "auth相关接口"
                ],
                "summary": "获取COS临时密钥接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "1000": {
                        "description": "code=1000,msg=\"success\",data里面是cos临时密钥数据",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    },
                    "1005": {
                        "description": "code=1000+，msg里面是错误信息,data=null",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth相关接口"
                ],
                "summary": "登录接口",
                "parameters": [
                    {
                        "description": "登录参数,必填",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamLogin"
                        }
                    }
                ],
                "responses": {
                    "1000": {
                        "description": "code=1000,msg=\"success\",成功返回token,若绑定教务信息，也会包含教务学生信息",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseLoginData"
                        }
                    },
                    "1001": {
                        "description": "请求错误参数,code=1000+，msg里面是错误信息",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        },
        "/auth/recoverpass": {
            "post": {
                "tags": [
                    "auth相关接口"
                ],
                "summary": "找回密码验证接口",
                "parameters": [
                    {
                        "description": "使用邮箱验证码新密码重置",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamReCover"
                        }
                    }
                ],
                "responses": {
                    "1000": {
                        "description": "code=1000,msg=\"success\",data=null",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    },
                    "1005": {
                        "description": "code=1000+，msg里面是错误信息,data=null",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        },
        "/auth/resetpass": {
            "post": {
                "tags": [
                    "auth相关接口"
                ],
                "summary": "重置密码验证接口",
                "parameters": [
                    {
                        "description": "使用邮箱旧密码新密码重置",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamReSet"
                        }
                    }
                ],
                "responses": {
                    "1000": {
                        "description": "code=1000,msg=\"success\",data=null",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    },
                    "1005": {
                        "description": "code=1000+，msg里面是错误信息,data=null",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        },
        "/auth/sendemail": {
            "post": {
                "tags": [
                    "auth相关接口"
                ],
                "summary": "邮箱验证接口",
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "sign",
                            "recoverpass"
                        ],
                        "type": "string",
                        "name": "mode",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "1000": {
                        "description": "code=1000,msg=\"success\",data=null",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    },
                    "1005": {
                        "description": "code=1000+，msg里面是错误信息,data=null",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth相关接口"
                ],
                "summary": "注册接口",
                "parameters": [
                    {
                        "description": "注册参数,必填",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamSignUp"
                        }
                    }
                ],
                "responses": {
                    "1000": {
                        "description": "code=1000,msg=\"success\",data=null",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    },
                    "1001": {
                        "description": "code=1000+，msg里面是错误信息,data=null",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        },
        "/edu/bind": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sylu相关接口"
                ],
                "summary": "绑定接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "绑定参数,必填",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamBind"
                        }
                    }
                ],
                "responses": {
                    "1000": {
                        "description": "code=1000,msg=\"success\",",
                        "schema": {
                            "$ref": "#/definitions/models.ReqBind"
                        }
                    },
                    "1001": {
                        "description": "请求错误参数,code=1000+，msg里面是错误信息",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        },
        "/edu/cookie": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sylu相关接口"
                ],
                "summary": "获取cookie接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "1000": {
                        "description": "code=1000,msg=\"success\",",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "1001": {
                        "description": "请求错误参数,code=1000+，msg里面是错误信息",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        },
        "/edu/courses": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sylu相关接口"
                ],
                "summary": "获取课表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "课表参数,必填,其中semester为3或12表示某学期，例如year=2022 semester=3 表示2022-2023学年第一学期",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamCourse"
                        }
                    }
                ],
                "responses": {
                    "1000": {
                        "description": "code=1000,msg=\"success\",",
                        "schema": {
                            "$ref": "#/definitions/models.ReqCourse"
                        }
                    },
                    "1001": {
                        "description": "请求错误参数,code=1000+，msg里面是错误信息",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        },
        "/edu/grade/detaile": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sylu相关接口"
                ],
                "summary": "获取成绩详情接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "成绩参数,必填，其中semester为3或12表示某学期，例如year=2022 semester=3 表示2022-2023学年第一学期",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamGradeDetaile"
                        }
                    }
                ],
                "responses": {
                    "1000": {
                        "description": "code=1000,msg=\"success\",",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ResGradeDetail"
                            }
                        }
                    },
                    "1001": {
                        "description": "请求错误参数,code=1000+，msg里面是错误信息",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        },
        "/edu/grades": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sylu相关接口"
                ],
                "summary": "获取成绩接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "成绩参数,必填，其中semester为3或12表示某学期，例如year=2022 semester=3 表示2022-2023学年第一学期",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamGrades"
                        }
                    }
                ],
                "responses": {
                    "1000": {
                        "description": "code=1000,msg=\"success\",",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.JsonGrades"
                            }
                        }
                    },
                    "1001": {
                        "description": "请求错误参数,code=1000+，msg里面是错误信息",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ResCode": {
            "type": "integer",
            "enum": [
                1000,
                1001,
                1002,
                1003,
                1004,
                1005,
                1006,
                1007,
                1008,
                1009,
                1010,
                1011,
                1012
            ],
            "x-enum-varnames": [
                "CodeSuccess",
                "CodeInvalidParam",
                "CodeEmailExist",
                "CodeEmailNotExist",
                "CodeInvalidPassword",
                "CodeInvalidCaptcha",
                "CodeCaptchaNotExistOrTimeOut",
                "CodeUnbound",
                "CodeServerBusy",
                "CodeInvalidToken",
                "CodeNeedLogin",
                "CodeInvalidCookie",
                "CodeNotData"
            ]
        },
        "controller.ResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/controller.ResCode"
                },
                "data": {},
                "msg": {}
            }
        },
        "controller._ResponseLoginData": {
            "type": "object",
            "properties": {
                "code": {
                    "default": 1005,
                    "allOf": [
                        {
                            "$ref": "#/definitions/controller.ResCode"
                        }
                    ]
                },
                "data": {
                    "$ref": "#/definitions/models.ReqLogin"
                },
                "msg": {
                    "type": "string",
                    "default": "success"
                }
            }
        },
        "models.JsonCourse": {
            "type": "object",
            "properties": {
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "teacher": {
                    "type": "string"
                },
                "time": {
                    "type": "integer"
                },
                "weekday": {
                    "type": "integer"
                },
                "weeks": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "models.JsonGrades": {
            "type": "object",
            "properties": {
                "classid": {
                    "type": "string"
                },
                "credits": {
                    "description": "学分",
                    "type": "number"
                },
                "fraction": {
                    "type": "number"
                },
                "gpa": {
                    "description": "绩点",
                    "type": "number"
                },
                "grade": {
                    "type": "string"
                },
                "gradepoints": {
                    "description": "学分绩点",
                    "type": "number"
                },
                "isdegree": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "teacher": {
                    "type": "string"
                }
            }
        },
        "models.ParamBind": {
            "type": "object",
            "required": [
                "password",
                "studentID"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "studentID": {
                    "type": "string"
                }
            }
        },
        "models.ParamCourse": {
            "type": "object",
            "required": [
                "semester",
                "year"
            ],
            "properties": {
                "cookie": {
                    "type": "string"
                },
                "semester": {
                    "type": "integer",
                    "enum": [
                        3,
                        12
                    ]
                },
                "year": {
                    "type": "string"
                }
            }
        },
        "models.ParamGradeDetaile": {
            "type": "object",
            "required": [
                "semester",
                "year"
            ],
            "properties": {
                "classid": {
                    "type": "string"
                },
                "cookie": {
                    "type": "string"
                },
                "semester": {
                    "type": "integer",
                    "enum": [
                        3,
                        12
                    ]
                },
                "year": {
                    "type": "string"
                }
            }
        },
        "models.ParamGrades": {
            "type": "object",
            "required": [
                "semester",
                "year"
            ],
            "properties": {
                "cookie": {
                    "type": "string"
                },
                "semester": {
                    "type": "integer",
                    "enum": [
                        3,
                        12
                    ]
                },
                "year": {
                    "type": "string"
                }
            }
        },
        "models.ParamLogin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "models.ParamReCover": {
            "type": "object",
            "required": [
                "captcha",
                "email",
                "newpassword",
                "renewpassword"
            ],
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "newpassword": {
                    "type": "string",
                    "minLength": 8
                },
                "renewpassword": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "models.ParamReSet": {
            "type": "object",
            "required": [
                "email",
                "newpassword",
                "password",
                "renewpassword"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "newpassword": {
                    "type": "string",
                    "minLength": 8
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "renewpassword": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "models.ParamSignUp": {
            "type": "object",
            "required": [
                "captcha",
                "email",
                "password",
                "repassword",
                "username"
            ],
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "repassword": {
                    "type": "string",
                    "minLength": 8
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.ReqBind": {
            "type": "object",
            "properties": {
                "cookie": {
                    "type": "string"
                },
                "syluinfo": {
                    "$ref": "#/definitions/models.ReqSyluInfo"
                }
            }
        },
        "models.ReqCourse": {
            "type": "object",
            "properties": {
                "courses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.JsonCourse"
                    }
                },
                "starttime": {
                    "type": "string"
                }
            }
        },
        "models.ReqLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "syluinfo": {
                    "$ref": "#/definitions/models.ReqSyluInfo"
                },
                "token": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.ReqSyluInfo": {
            "type": "object",
            "properties": {
                "college": {
                    "description": "学院",
                    "type": "string",
                    "default": "信息科学与工程学院"
                },
                "grade": {
                    "description": "年级",
                    "type": "string",
                    "default": "2022"
                },
                "major": {
                    "description": "专业",
                    "type": "string",
                    "default": "计算机科学与技术(0305)"
                },
                "reusername": {
                    "description": "真实姓名",
                    "type": "string",
                    "default": "肖嘉兴"
                },
                "studentID": {
                    "description": "学号",
                    "type": "string",
                    "default": "2203050212"
                }
            }
        },
        "models.ResGradeDetail": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "score": {
                    "type": "string"
                },
                "weight": {
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
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "sylu项目接口文档",
	Description:      "致力于为同学们提供校园服务(忽略状态码，所有响应都是200)",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
