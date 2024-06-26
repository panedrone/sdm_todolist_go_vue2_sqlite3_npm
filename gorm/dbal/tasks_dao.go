package dbal

// Code generated by a tool. DO NOT EDIT.
// https://sqldalmaker.sourceforge.net/

import (
	"context"
	"sdm_demo_todolist/gorm/dbal/models"
)

type TasksDao struct {
	ds DataStore
}

// (C)RUD: tasks
// Generated/AI values are passed to DTO/model.

func (dao *TasksDao) CreateTask(ctx context.Context, item *models.Task) error {
	return dao.ds.Create(ctx, "tasks", item)
}

// C(R)UD: tasks

func (dao *TasksDao) ReadTaskList(ctx context.Context) (res []*models.Task, err error) {
	err = dao.ds.ReadAll(ctx, "tasks", &res)
	return
}

// C(R)UD: tasks

func (dao *TasksDao) ReadTask(ctx context.Context, tID int64) (*models.Task, error) {
	res := &models.Task{}
	err := dao.ds.Read(ctx, "tasks", res, tID)
	if err == nil {
		return res, nil
	}
	return nil, err
}

// CR(U)D: tasks

func (dao *TasksDao) UpdateTask(ctx context.Context, item *models.Task) (rowsAffected int64, err error) {
	rowsAffected, err = dao.ds.Update(ctx, "tasks", item)
	return
}

// CRU(D): tasks

func (dao *TasksDao) DeleteTask(ctx context.Context, item *models.Task) (rowsAffected int64, err error) {
	rowsAffected, err = dao.ds.Delete(ctx, "tasks", item)
	return
}

func (dao *TasksDao) RawProjectTasks(ctx context.Context, pID int64) (res []*models.TaskLi, err error) {
	sql := `select t_id, t_priority, t_date, t_subject from tasks where p_id =? 
		order by t_id`
	err = dao.ds.Select(ctx, sql, &res, pID)
	return
}
