package request

import (
	"github.com/gin-gonic/gin"
	"sdm_demo_todolist/shared/resp"
)

type ProjectUri struct {
	PId int64 `uri:"p_id" binding:"required"`
}

func GetProjectUri(ctx *gin.Context) (*ProjectUri, error) {
	var uri ProjectUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		resp.Abort400BadUri(ctx, err)
		return nil, err
	}
	return &uri, nil
}

type Project struct {
	PName string `json:"p_name" binding:"required,lte=256"`
}

type TaskUri struct {
	TId int64 `uri:"t_id" binding:"required"`
}

func GetTaskUri(ctx *gin.Context) (*TaskUri, error) {
	var uri TaskUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		resp.Abort400BadUri(ctx, err)
		return nil, err
	}
	return &uri, nil
}

type NewTask struct {
	TSubject string `json:"t_subject" binding:"required,lte=256"`
}
