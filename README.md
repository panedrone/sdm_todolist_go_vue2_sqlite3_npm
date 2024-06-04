A project demonstrating how to use [SQL DAL Maker](https://github.com/panedrone/sqldalmaker) + Golang.
Front-end is written in Vue.js 2.7.

The following cases  are considered:

* [Part 1](./gorm): using "github.com/go-gorm/gorm"
* [Part 2](./sqlx): using "github.com/jmoiron/sqlx"
* [Part 3](./no_orm): using "database/sql" directly

![demo-go.png](demo-go.png)

sdm.xml

```xml

<sdm>

    <dto-class name="Project" ref="projects"/>

    <dto-class name="ProjectLi" ref="projects">
        <field type="int64$" column="p_tasks_count"/>
    </dto-class>

    <dto-class name="Task" ref="tasks"/>

    <dto-class name="TaskLi" ref="tasks">
        <field type="-" column="p_id"/>
        <field type="-" column="t_comments"/>
    </dto-class>

    <dao-class name="ProjectsDao">
        <crud dto="Project"/>
        <query-dto-list method="ReadAll" dto="ProjectLi" ref="get_projects.sql"/>
    </dao-class>

    <dao-class name="TasksDao">
        <crud table="tasks" dto="Task"/>
        <query-dto-list method="ReadByProject(pId)" ref="get_project_tasks.sql" dto="TaskLi"/>
    </dao-class>

</sdm>
```

Generated code in action:

```go
type projectHandlers struct {
	dao *dbal.ProjectsDao
}

func NewProjectHandlers() shared.ProjectHandlers {
	return &projectHandlers{
		dao: dbal.NewProjectsDao(),
	}
}

func (h *projectHandlers) ProjectCreate(ctx *gin.Context) {
	var req request.Project
	if err := request.BindJSON(ctx, &req); err != nil {
		return
	}
	if err := h.dao.CreateProject(ctx, &m.Project{PName: req.PName}); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusCreated)
}

func (h *projectHandlers) ProjectsReadAll(ctx *gin.Context) {
	all, err := h.dao.ReadAll(ctx)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, all)
}

func (h *projectHandlers) ProjectRead(ctx *gin.Context) {
	uri, err := request.GetProjectUri(ctx)
	if err != nil {
		return
	}
	pr, err := h.dao.ReadProject(ctx, uri.PId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Abort404(ctx, err)
			return
		}
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, pr)
}

func (h *projectHandlers) ProjectUpdate(ctx *gin.Context) {
	uri, err := request.GetProjectUri(ctx)
	if err != nil {
		return
	}
	var req request.Project
	if err := request.BindJSON(ctx, &req); err != nil {
		return
	}
	pr := &m.Project{PId: uri.PId, PName: req.PName}
	if _, err := h.dao.UpdateProject(ctx, pr); err != nil {
		resp.Abort500(ctx, err)
	}
}

func (h *projectHandlers) ProjectDelete(ctx *gin.Context) {
	uri, err := request.GetProjectUri(ctx)
	if err != nil {
		return
	}
	if _, err := h.dao.DeleteProject(ctx, &m.Project{PId: uri.PId}); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
```
