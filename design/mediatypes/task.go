// ./design/mediatypes/task.go

package mediatypes

import (
	goa "github.com/goadesign/goa/design"
	dsl "github.com/goadesign/goa/design/apidsl"
)

// Task はタスクリソースのメディアタイプ
var Task = dsl.MediaType("application/x-learning-goa+json", func() {
	dsl.Description("タスク")
	dsl.Attributes(func() {
		dsl.Attribute("id", goa.Integer, "タスクID", func() {
			dsl.Example(12345)
		})
		dsl.Attribute("title", goa.String, "タスクのタイトル", func() {
			dsl.Example("my task")
		})
		dsl.Attribute("done", goa.Boolean, "タスクの完了状態", func() {
			dsl.Example(true)
		})
		dsl.Attribute("created_at", goa.DateTime, "タスクの作成日時")
		dsl.Attribute("updated_at", goa.DateTime, "タスクの更新日時")

		dsl.Required("id", "title", "done", "created_at", "updated_at")
	})

	dsl.View("default", func() {
		dsl.Attribute("id")
		dsl.Attribute("title")
		dsl.Attribute("done")
		dsl.Attribute("created_at")
		dsl.Attribute("updated_at")
	})
})
