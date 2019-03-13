package controller

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/pibigstar/go-todo/middleware"
	"github.com/pibigstar/go-todo/models"
	"github.com/pibigstar/go-todo/utils"
)

func init(){
	s := g.Server()
	s.BindHandler("/collect",collectFormId)
}

type CollectionFormID struct {
	FormID string `json:"formId"`
}

// 收集用户的formId
func collectFormId(r *ghttp.Request){
	collection := new(CollectionFormID)
	r.GetToStruct(collection)
	openID, _ := middleware.GetOpenID(r)
	err := models.CollectFormID(openID,collection.FormID)
	if err!=nil {
		r.Response.WriteJson(utils.ErrorResponse(err.Error()))
		r.Exit()
	}
	r.Response.WriteJson(utils.SuccessResponse("OK"))
}