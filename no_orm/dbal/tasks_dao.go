package dbal

// Code generated by a tool. DO NOT EDIT.
// Additions may be hand-coded in a separate go-file.
// https://sqldalmaker.sourceforge.net/

import (
	"context"
	"sdm_demo_todolist/no_orm/dbal/dto"
)

type TasksDao struct {
	ds DataStore
}

// (C)RUD: tasks
// Generated/AI values are passed to DTO/model.

func (dao *TasksDao) CreateTask(ctx context.Context, item *dto.Task) error {
	sql := `insert into tasks (p_id, t_priority, t_date, t_subject, t_comments) values (?, ?, ?, ?, ?)`
	row, err := dao.ds.Insert(ctx, sql, "t_id", item.PId, item.TPriority, item.TDate, item.TSubject, item.TComments)
	if err == nil {
		err = SetRes(&item.TId, row)
	}
	return err
}

// C(R)UD: tasks

func (dao *TasksDao) ReadTaskList(ctx context.Context) (res []*dto.Task, err error) {
	sql := `select * from tasks`
	errMap := make(map[string]int)
	_onRow := func(row map[string]interface{}) {
		item := dto.Task{}
		SetInt64(&item.TId, row, "t_id", errMap)
		SetInt64(&item.PId, row, "p_id", errMap)
		SetInt64(&item.TPriority, row, "t_priority", errMap)
		SetString(&item.TDate, row, "t_date", errMap)
		SetString(&item.TSubject, row, "t_subject", errMap)
		SetString(&item.TComments, row, "t_comments", errMap)
		res = append(res, &item)
	}
	err = dao.ds.QueryAllRows(ctx, sql, _onRow)
	if err == nil {
		err = ErrMapToErr(errMap)
	}
	return
}

// C(R)UD: tasks

func (dao *TasksDao) ReadTask(ctx context.Context, tId int64) (*dto.Task, error) {
	sql := `select * from tasks where t_id=?`
	row, err := dao.ds.QueryRow(ctx, sql, tId)
	if err != nil {
		return nil, err
	}
	item := dto.Task{}
	errMap := make(map[string]int)
	SetInt64(&item.TId, row, "t_id", errMap)
	SetInt64(&item.PId, row, "p_id", errMap)
	SetInt64(&item.TPriority, row, "t_priority", errMap)
	SetString(&item.TDate, row, "t_date", errMap)
	SetString(&item.TSubject, row, "t_subject", errMap)
	SetString(&item.TComments, row, "t_comments", errMap)
	err = ErrMapToErr(errMap)
	return &item, err
}

// CR(U)D: tasks

func (dao *TasksDao) UpdateTask(ctx context.Context, item *dto.Task) (rowsAffected int64, err error) {
	sql := `update tasks set p_id=?, t_priority=?, t_date=?, t_subject=?, t_comments=? where t_id=?`
	rowsAffected, err = dao.ds.Exec(ctx, sql, item.PId, item.TPriority, item.TDate, item.TSubject, item.TComments, item.TId)
	return
}

// CRU(D): tasks

func (dao *TasksDao) DeleteTask(ctx context.Context, item *dto.Task) (rowsAffected int64, err error) {
	sql := `delete from tasks where t_id=?`
	rowsAffected, err = dao.ds.Exec(ctx, sql, item.TId)
	return
}

func (dao *TasksDao) ReadByProject(ctx context.Context, pId int64) (res []*dto.TaskLi, err error) {
	sql := `select t_id, t_priority, t_date, t_subject from tasks where p_id =? 
		order by t_id`
	errMap := make(map[string]int)
	_onRow := func(row map[string]interface{}) {
		item := dto.TaskLi{}
		SetInt64(&item.TId, row, "t_id", errMap)
		SetInt64(&item.TPriority, row, "t_priority", errMap)
		SetString(&item.TDate, row, "t_date", errMap)
		SetString(&item.TSubject, row, "t_subject", errMap)
		res = append(res, &item)
	}
	err = dao.ds.QueryAllRows(ctx, sql, _onRow, pId)
	if err == nil {
		err = ErrMapToErr(errMap)
	}
	return
}

func (dao *TasksDao) DelByProject(ctx context.Context, pId string) (rowsAffected int64, err error) {
	sql := `delete from tasks where p_id=?`
	rowsAffected, err = dao.ds.Exec(ctx, sql, pId)
	return
}

func (dao *TasksDao) GetCount(ctx context.Context) (res int64, err error) {
	sql := `select count(*) from tasks`
	err = dao.ds.Query(ctx, sql, &res)
	return
}

func (dao *TasksDao) GetProjectTasks2(ctx context.Context, gId string) (res []*dto.TaskLi, err error) {
	sql := `delete from tasks where p_id=?`
	errMap := make(map[string]int)
	_onRow := func(row map[string]interface{}) {
		item := dto.TaskLi{}
		SetInt64(&item.TId, row, "t_id", errMap)
		SetInt64(&item.TPriority, row, "t_priority", errMap)
		SetString(&item.TDate, row, "t_date", errMap)
		SetString(&item.TSubject, row, "t_subject", errMap)
		res = append(res, &item)
	}
	err = dao.ds.QueryAllRows(ctx, sql, _onRow, gId)
	if err == nil {
		err = ErrMapToErr(errMap)
	}
	return
}

func (dao *TasksDao) GetTask(ctx context.Context, gId string) (*dto.Task, error) {
	sql := `delete from tasks where p_id=?`
	row, err := dao.ds.QueryRow(ctx, sql, gId)
	if err != nil {
		return nil, err
	}
	item := dto.Task{}
	errMap := make(map[string]int)
	SetInt64(&item.TId, row, "t_id", errMap)
	SetInt64(&item.PId, row, "p_id", errMap)
	SetInt64(&item.TPriority, row, "t_priority", errMap)
	SetString(&item.TDate, row, "t_date", errMap)
	SetString(&item.TSubject, row, "t_subject", errMap)
	SetString(&item.TComments, row, "t_comments", errMap)
	err = ErrMapToErr(errMap)
	return &item, err
}
