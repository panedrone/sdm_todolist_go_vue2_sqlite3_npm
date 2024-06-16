package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sdm_demo_todolist/gorm/dbal"
	"sdm_demo_todolist/gorm/dbal/models"
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
	var req request.NewTask
	if err := request.BindJSON(ctx, &req); err != nil {
		return
	}
	t := models.Task{}
	t.PId = uri.PId
	t.TSubject = req.TSubject
	t.TPriority = 1
	t.TDate = datetime.NowLocalString()
	if err := h.dao.CreateTask(ctx, &t); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusCreated)
}

func (h *taskHandlers) TaskRead(ctx *gin.Context) {
	uri, err := request.GetTaskUri(ctx)
	if err != nil {
		return
	}
	task, err := h.dao.ReadTask(ctx, uri.TId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Abort404(ctx, err)
			return
		}
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, task)
}

func (h *taskHandlers) TasksReadByProject(ctx *gin.Context) {
	uri, err := request.GetProjectUri(ctx)
	if err != nil {
		return
	}
	// tasks, err := tDao.RawProjectTasks(ctx, uri.PId)
	tasks, err := h.dao.ReadProjectTasks(ctx, uri.PId)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, tasks)
}

func (h *taskHandlers) TaskUpdate(ctx *gin.Context) {
	uri, err := request.GetTaskUri(ctx)
	if err != nil {
		return
	}
	dao := h.dao
	t, err := dao.ReadTask(ctx, uri.TId)
	if err != nil {
		resp.Abort400hBadRequest(ctx, err.Error())
		return
	}
	var req models.Task
	if err := request.BindJSON(ctx, &req); err != nil {
		return
	}
	_, err = datetime.Parse(req.TDate)
	if err != nil {
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
	t.TSubject = req.TSubject
	t.TPriority = req.TPriority
	t.TDate = req.TDate
	t.TComments = req.TComments
	if _, err = dao.UpdateTask(ctx, t); err != nil {
		resp.Abort500(ctx, err)
		return
	}
}

func (h *taskHandlers) TaskDelete(ctx *gin.Context) {
	uri, err := request.GetTaskUri(ctx)
	if err != nil {
		return
	}
	if _, err := h.dao.DeleteTask(ctx, &models.Task{TId: uri.TId}); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
