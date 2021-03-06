// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Ares"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/novel/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "小说（Novel）"
                ],
                "summary": "创建",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.NovelCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Succeed"
                        }
                    }
                }
            }
        },
        "/novel/delete": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "小说（Novel）"
                ],
                "summary": "删除",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.NovelDelete"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Succeed"
                        }
                    }
                }
            }
        },
        "/novel/get-detail": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "小说（Novel）"
                ],
                "summary": "详情",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.NovelDetail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/novelDTO.NovelDTO"
                        }
                    }
                }
            }
        },
        "/novel/get-list": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "小说（Novel）"
                ],
                "summary": "列表",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.NovelList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/novelDTO.NovelAdminListDTO"
                        }
                    }
                }
            }
        },
        "/novel/modify": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "小说（Novel）"
                ],
                "summary": "修改",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.NovelModify"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Succeed"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Novel": {
            "type": "object",
            "properties": {
                "chapterPreview": {
                    "type": "string"
                },
                "createdAt": {
                    "$ref": "#/definitions/utils.Datetime"
                },
                "id": {
                    "type": "integer"
                },
                "isDisplayIndex": {
                    "type": "integer"
                },
                "isHot": {
                    "type": "integer"
                },
                "isNew": {
                    "type": "integer"
                },
                "isRecommend": {
                    "type": "integer"
                },
                "mainRole": {
                    "type": "string"
                },
                "novelAuthor": {
                    "type": "string"
                },
                "novelBannerCover": {
                    "type": "string"
                },
                "novelCategory": {
                    "$ref": "#/definitions/model.NovelCategory"
                },
                "novelCategoryId": {
                    "type": "integer"
                },
                "novelCover": {
                    "type": "string"
                },
                "novelDescription": {
                    "type": "string"
                },
                "novelNavigation": {
                    "$ref": "#/definitions/model.NovelNavigation"
                },
                "novelNavigationId": {
                    "type": "integer"
                },
                "novelStatus": {
                    "type": "integer"
                },
                "novelThemes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.NovelTheme"
                    }
                },
                "novelTitle": {
                    "type": "string"
                },
                "sortNum": {
                    "type": "integer"
                },
                "updatedAt": {
                    "$ref": "#/definitions/utils.Datetime"
                },
                "viewNum": {
                    "type": "integer"
                }
            }
        },
        "model.NovelCategory": {
            "type": "object",
            "properties": {
                "categoryIcon": {
                    "type": "string"
                },
                "categoryName": {
                    "type": "string"
                },
                "createdAt": {
                    "$ref": "#/definitions/utils.Datetime"
                },
                "id": {
                    "type": "integer"
                },
                "isDisplayIndex": {
                    "type": "integer"
                },
                "novelNavigation": {
                    "$ref": "#/definitions/model.NovelNavigation"
                },
                "novelNavigationId": {
                    "type": "integer"
                },
                "novels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Novel"
                    }
                },
                "sortNum": {
                    "type": "integer"
                },
                "updatedAt": {
                    "$ref": "#/definitions/utils.Datetime"
                }
            }
        },
        "model.NovelNavigation": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "$ref": "#/definitions/utils.Datetime"
                },
                "id": {
                    "type": "integer"
                },
                "isDisplayIndex": {
                    "type": "integer"
                },
                "navigationName": {
                    "type": "string"
                },
                "novelCategories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.NovelCategory"
                    }
                },
                "novels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Novel"
                    }
                },
                "sortNum": {
                    "type": "integer"
                },
                "updatedAt": {
                    "$ref": "#/definitions/utils.Datetime"
                }
            }
        },
        "model.NovelTheme": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "$ref": "#/definitions/utils.Datetime"
                },
                "id": {
                    "type": "integer"
                },
                "novels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Novel"
                    }
                },
                "sortNum": {
                    "type": "integer"
                },
                "themeCover": {
                    "type": "string"
                },
                "themeDescription": {
                    "type": "string"
                },
                "themeName": {
                    "type": "string"
                },
                "themeSubtitle": {
                    "type": "string"
                },
                "updatedAt": {
                    "$ref": "#/definitions/utils.Datetime"
                }
            }
        },
        "novelDTO.CategoryDTO": {
            "type": "object",
            "properties": {
                "categoryName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "novelDTO.NavigationDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "navigationName": {
                    "type": "string"
                }
            }
        },
        "novelDTO.NovelAdminListDTO": {
            "type": "object",
            "properties": {
                "list": {
                    "description": "列表数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/novelDTO.NovelCopy"
                    }
                },
                "total": {
                    "description": "总数",
                    "type": "integer"
                }
            }
        },
        "novelDTO.NovelCopy": {
            "type": "object",
            "properties": {
                "chapterPreview": {
                    "type": "string"
                },
                "createdAt": {
                    "$ref": "#/definitions/utils.Datetime"
                },
                "id": {
                    "type": "integer"
                },
                "isDisplayIndex": {
                    "type": "integer"
                },
                "isHot": {
                    "type": "integer"
                },
                "isNew": {
                    "type": "integer"
                },
                "isRecommend": {
                    "type": "integer"
                },
                "mainRole": {
                    "type": "string"
                },
                "novelAuthor": {
                    "type": "string"
                },
                "novelBannerCover": {
                    "type": "string"
                },
                "novelCategory": {
                    "$ref": "#/definitions/model.NovelCategory"
                },
                "novelCategoryId": {
                    "type": "integer"
                },
                "novelCover": {
                    "type": "string"
                },
                "novelDescription": {
                    "type": "string"
                },
                "novelNavigation": {
                    "$ref": "#/definitions/model.NovelNavigation"
                },
                "novelNavigationId": {
                    "type": "integer"
                },
                "novelStatus": {
                    "type": "integer"
                },
                "novelThemes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.NovelTheme"
                    }
                },
                "novelTitle": {
                    "type": "string"
                },
                "sortNum": {
                    "type": "integer"
                },
                "updatedAt": {
                    "$ref": "#/definitions/utils.Datetime"
                },
                "viewNum": {
                    "type": "integer"
                }
            }
        },
        "novelDTO.NovelDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isDisplayIndex": {
                    "type": "integer"
                },
                "isHot": {
                    "type": "integer"
                },
                "isHotTitle": {
                    "type": "string"
                },
                "isNew": {
                    "type": "integer"
                },
                "isNewTitle": {
                    "type": "string"
                },
                "isRecommend": {
                    "type": "integer"
                },
                "isRecommendTitle": {
                    "type": "string"
                },
                "mainRole": {
                    "type": "string"
                },
                "novelAuthor": {
                    "type": "string"
                },
                "novelCategory": {
                    "$ref": "#/definitions/novelDTO.CategoryDTO"
                },
                "novelNavigation": {
                    "$ref": "#/definitions/novelDTO.NavigationDTO"
                },
                "novelStatus": {
                    "type": "integer"
                },
                "novelStatusTitle": {
                    "type": "string"
                },
                "novelThemes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/novelDTO.ThemeDTO"
                    }
                },
                "novelTitle": {
                    "type": "string"
                }
            }
        },
        "novelDTO.ThemeDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "themeName": {
                    "type": "string"
                }
            }
        },
        "request.NovelCreate": {
            "type": "object",
            "required": [
                "chapterPreview",
                "mainRole",
                "novelAuthor",
                "novelCategoryId",
                "novelCover",
                "novelDescription",
                "novelNavigationId",
                "novelThemeIds",
                "novelTitle"
            ],
            "properties": {
                "chapterPreview": {
                    "description": "章节预览",
                    "type": "string",
                    "example": "xxxx"
                },
                "isDisplayIndex": {
                    "description": "是否首页展示：1-是 0-否",
                    "type": "integer"
                },
                "isHot": {
                    "description": "是否热门：0-否 1-是",
                    "type": "integer"
                },
                "isNew": {
                    "description": "是否是最新：0-否 1-是",
                    "type": "integer"
                },
                "isRecommend": {
                    "description": "是否推荐：0-否 1-是",
                    "type": "integer"
                },
                "mainRole": {
                    "description": "小说角色",
                    "type": "string",
                    "example": "张三丰"
                },
                "novelAuthor": {
                    "description": "小说作者",
                    "type": "string",
                    "example": "天蚕土豆"
                },
                "novelBannerCover": {
                    "description": "小说banner",
                    "type": "string",
                    "example": "http://baidu.com/a.jpg"
                },
                "novelCategoryId": {
                    "description": "小说分类ID",
                    "type": "integer"
                },
                "novelCover": {
                    "description": "小说小封面",
                    "type": "string",
                    "example": "http://baidu.com/a.jpg"
                },
                "novelDescription": {
                    "description": "小说描述",
                    "type": "string",
                    "example": "这是一段美好的描述"
                },
                "novelNavigationId": {
                    "description": "小说导航ID",
                    "type": "integer",
                    "example": 1
                },
                "novelStatus": {
                    "description": "小说状态：1-连载中 2-已完结",
                    "type": "integer",
                    "example": 1
                },
                "novelThemeIds": {
                    "description": "小说专题id数组",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "novelTitle": {
                    "description": "小说标题",
                    "type": "string",
                    "example": "盘龙"
                },
                "sortNum": {
                    "description": "排序值",
                    "type": "integer"
                }
            }
        },
        "request.NovelDelete": {
            "type": "object",
            "required": [
                "novelId"
            ],
            "properties": {
                "novelId": {
                    "description": "小说ID",
                    "type": "integer"
                }
            }
        },
        "request.NovelDetail": {
            "type": "object",
            "required": [
                "novelId"
            ],
            "properties": {
                "novelId": {
                    "description": "小说ID",
                    "type": "integer"
                }
            }
        },
        "request.NovelList": {
            "type": "object",
            "properties": {
                "novelStatus": {
                    "description": "小说状态",
                    "type": "integer"
                },
                "novelTitle": {
                    "description": "小说标题",
                    "type": "string"
                },
                "page": {
                    "description": "当前页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页显示的条数",
                    "type": "integer"
                }
            }
        },
        "request.NovelModify": {
            "type": "object",
            "required": [
                "chapterPreview",
                "mainRole",
                "novelAuthor",
                "novelCategoryId",
                "novelCover",
                "novelDescription",
                "novelId",
                "novelNavigationId",
                "novelThemeIds",
                "novelTitle"
            ],
            "properties": {
                "chapterPreview": {
                    "description": "章节预览",
                    "type": "string"
                },
                "isDisplayIndex": {
                    "description": "是否首页展示：1-是 0-否",
                    "type": "integer"
                },
                "isHot": {
                    "description": "是否热门：0-否 1-是",
                    "type": "integer"
                },
                "isNew": {
                    "description": "是否是最新：0-否 1-是",
                    "type": "integer"
                },
                "isRecommend": {
                    "description": "是否推荐：0-否 1-是",
                    "type": "integer"
                },
                "mainRole": {
                    "description": "小说角色",
                    "type": "string",
                    "example": "张三丰"
                },
                "novelAuthor": {
                    "description": "小说作者",
                    "type": "string",
                    "example": "天蚕土豆"
                },
                "novelBannerCover": {
                    "description": "小说banner",
                    "type": "string",
                    "example": "http://baidu.com/a.jpg"
                },
                "novelCategoryId": {
                    "description": "小说分类ID",
                    "type": "integer"
                },
                "novelCover": {
                    "description": "小说小封面",
                    "type": "string",
                    "example": "http://baidu.com/a.jpg"
                },
                "novelDescription": {
                    "description": "小说描述",
                    "type": "string",
                    "example": "这是一段美好的描述"
                },
                "novelId": {
                    "description": "小说ID",
                    "type": "integer",
                    "example": 1
                },
                "novelNavigationId": {
                    "description": "小说导航ID",
                    "type": "integer"
                },
                "novelStatus": {
                    "description": "小说状态：1-连载中 2-已完结",
                    "type": "integer"
                },
                "novelThemeIds": {
                    "description": "小说专题id数组",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "novelTitle": {
                    "description": "小说标题",
                    "type": "string",
                    "example": "盘龙"
                },
                "sortNum": {
                    "description": "排序值",
                    "type": "integer"
                }
            }
        },
        "response.Succeed": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "utils.Datetime": {
            "type": "object",
            "properties": {
                "time.Time": {
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
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "http://example.com",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Gin小说网站接口实战",
	Description: "Gin小说网站接口文档 这里只做示例",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

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
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
