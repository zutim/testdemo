package handler

import (
	"github.com/gin-gonic/gin"
	"ztm-demo/internal/application"
	"ztm-demo/network/resp"
)

type DemoHandler struct {
	service application.DemoInterface
}

func NewDemoHandler(service application.DemoInterface) *DemoHandler {
	return &DemoHandler{service: service}
}

func (u *DemoHandler) Add(ctx *gin.Context) {
	panic("add ")
	resp.SuccessResponse(ctx, "")
}

func (u *DemoHandler) Search(ctx *gin.Context) {
	panic("get ")
	resp.SuccessResponse(ctx, "")
}

func (u *DemoHandler) Update(ctx *gin.Context) {
	panic("update")
	resp.SuccessResponse(ctx, "")

}

func (u *DemoHandler) Delete(ctx *gin.Context) {
	panic("delete")
	resp.SuccessResponse(ctx, "")
}
