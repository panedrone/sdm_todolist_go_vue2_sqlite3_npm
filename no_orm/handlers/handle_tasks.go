package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sdm_demo_todolist/no_orm/dbal"
	"sdm_demo_todolist/no_orm/dbal/dto"
	"sdm_demo_todolist/shared"
	"sdm_demo_todolist/shared/datetime"
	"sdm_demo_todolist/shared/request"
	"sdm_demo_todolist/shared/resp"
)

type taskHandlers struct {
	dao *dbal.TasksDao
}

func NewTaskHandlers() shared.TaskHandlers {
	return &taskHandlers{
		dao: dbal.NewTasksDao(),
	}
}

func (h *taskHandlers) TaskCreate(ctx *gin.Context) {
	uri, err := request.GetProjectUri(ctx)
	if err != nil {
		return
	}
	var inTask request.NewTask
	if err := ctx.ShouldBindJSON(&inTask); err != nil {
		resp.Abort400hBadRequest(ctx, err.Error())
		return
	}
	t := dto.Task{}
	t.PId = uri.PId
	t.TSubject = inTask.TSubject
	t.TPriority = 1
	t.TDate = datetime.NowLocalString()
	if err := h.dao.CreateTask(ctx, &t); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusCreated)
}

func (h *taskHandlers) TasksReadByProject(ctx *gin.Context) {
	uri, err := request.GetProjectUri(ctx)
	if err != nil {
		return
	}
	tasks, err := h.dao.ReadByProject(ctx, uri.PId)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, tasks)
}

func (h *taskHandlers) TaskRead(ctx *gin.Context) {
	uri, err := request.GetTaskUri(ctx)
	if err != nil {
		return
	}
	task, err := h.dao.ReadTask(ctx, uri.TId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			resp.Abort404(ctx, err)
			return
		}
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, task)
}

func (h *taskHandlers) TaskUpdate(ctx *gin.Context) {
	uri, err := request.GetTaskUri(ctx)
	if err != nil {
		return
	}
	var req dto.Task
	if err := request.BindJSON(ctx, &req); err != nil {
		return
	}
	if _, err := datetime.Parse(req.TDate); err != nil {
		resp.Abort400hBadRequest(ctx, fmt.Sprintf("Date format expected '%s': %s", datetime.TimeFormat, err.Error()))
		return
	}
	if len(req.TSubject) == 0 {
		resp.Abort400hBadRequest(ctx, fmt.Sprintf("Subject required"))
		return
	}
	if req.TPriority <= 0 {
		resp.Abort400hBadRequest(ctx, fmt.Sprintf("Invalid Priority: %d", req.TPriority))
		return
	}
	t, err := h.dao.ReadTask(ctx, uri.TId)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	t.TSubject = req.TSubject
	t.TPriority = req.TPriority
	t.TDate = req.TDate
	t.TComments = req.TComments
	if _, err = h.dao.UpdateTask(ctx, t); err != nil {
		resp.Abort500(ctx, err)
		return
	}
}

func (h *taskHandlers) TaskDelete(ctx *gin.Context) {
	uri, err := request.GetTaskUri(ctx)
	if err != nil {
		return
	}
	t := dto.Task{TId: uri.TId}
	if _, err := h.dao.DeleteTask(ctx, &t); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
