// ./design/resources/task.go

package resources

import (
	"github.com/BcRikko/learning-goa/design/mediatypes"
	"github.com/BcRikko/learning-goa/design/usertypes"
	goa "github.com/goadesign/goa/design"
	dsl "github.com/goadesign/goa/design/apidsl"
)

var _ = dsl.Resource("Tasks", func() {
	dsl.DefaultMedia(mediatypes.Task)
	dsl.BasePath("/tasks")

	// GET http://localhost:8080/api/tasks
	dsl.Action("list", func() {
		dsl.Routing(dsl.GET(""))
		dsl.Description("タスク一覧を取得する。")
		dsl.Response(goa.OK, dsl.CollectionOf(mediatypes.Task))
	})

	// GET http://localhost:8080/api/tasks/:taskID
	dsl.Action("show", func() {
		dsl.Routing(dsl.GET("/:taskID"))
		dsl.Description("指定IDのタスクを取得する。")
		dsl.Params(func() {
			dsl.Param("taskID", goa.Integer, "タスクID")
		})
		dsl.Response(goa.OK)
		dsl.Response(goa.NotFound)
		dsl.Response(goa.BadRequest, goa.ErrorMedia)
	})

	// POST http://localhost:8080/api/tasks
	dsl.Action("create", func() {
		dsl.Routing(dsl.POST(""))
		dsl.Description("タスクを作成する。")
		dsl.Payload(usertypes.TaskPayload)
		dsl.Response(goa.Created, "/tasks/[0-9]+")
		dsl.Response(goa.BadRequest, goa.ErrorMedia)
	})

	// PUT http://localhost:8080/api/tasks/:taskID
	dsl.Action("update", func() {
		dsl.Routing(dsl.POST("/:taskID"))
		dsl.Description("指定IDのタスクを更新する。")
		dsl.Params(func() {
			dsl.Param("taskID", goa.Integer, "タスクID")
		})
		dsl.Payload(usertypes.TaskPayload)
		dsl.Response(goa.NoContent)
		dsl.Response(goa.NotFound)
		dsl.Response(goa.BadRequest, goa.ErrorMedia)
	})

	// DELETE http://localhost:8080/api/tasks/:taskID
	dsl.Action("delete", func() {
		dsl.Routing(dsl.DELETE("/:taskID"))
		dsl.Description("指定IDのタスクを削除する。")
		dsl.Params(func() {
			dsl.Param("taskID", goa.Integer, "タスクID")
		})
		dsl.Response(goa.NoContent)
		dsl.Response(goa.NotFound)
		dsl.Response(goa.BadRequest, goa.ErrorMedia)
	})
})
