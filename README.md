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
var (
	prDao = dbal.NewProjectsDao()
)

func ProjectCreate(ctx *gin.Context) {
	var req request.Project
	if err := request.BindJSON(ctx, &req); err != nil {
		return
	}
	if err := prDao.CreateProject(ctx, &m.Project{PName: req.PName}); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusCreated)
}

func ProjectsReadAll(ctx *gin.Context) {
	all, err := prDao.ReadAll(ctx)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, all)
}

func ProjectRead(ctx *gin.Context) {
	uri, err := request.GetProjectUri(ctx)
	if err != nil {
		return
	}
	pr, err := prDao.ReadProject(ctx, uri.PId)
	if err != nil {
		resp.Abort500(ctx, err)
		return
	}
	resp.JSON(ctx, http.StatusOK, pr)
}

func ProjectUpdate(ctx *gin.Context) {
	uri, err := request.GetProjectUri(ctx)
	if err != nil {
		return
	}
	var req request.Project
	if err := request.BindJSON(ctx, &req); err != nil {
		return
	}
	pr := &m.Project{PId: uri.PId, PName: req.PName}
	if _, err := prDao.UpdateProject(ctx, pr); err != nil {
		resp.Abort500(ctx, err)
	}
}

func ProjectDelete(ctx *gin.Context) {
	uri, err := request.GetProjectUri(ctx)
	if err != nil {
		return
	}
	if _, err := prDao.DeleteProject(ctx, &m.Project{PId: uri.PId}); err != nil {
		resp.Abort500(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
```
