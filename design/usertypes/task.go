// ./design/usertypes/task.go

package usertypes

import (
	goa "github.com/goadesign/goa/design"
	dsl "github.com/goadesign/goa/design/apidsl"
)

// TaskPayload はリクエストパラメータの定義
var TaskPayload = dsl.Type("TaskPayload", func() {
	dsl.Attribute("title", goa.String, "タスクのタイトル", func() {
		dsl.MinLength(2)
		dsl.MaxLength(128)
		dsl.Example("example task title")
	})

	dsl.Attribute("done", goa.Boolean, "タスクの完了状態", func() {
		dsl.Default(false)
		dsl.Example(false)
	})

	dsl.Required("title")
})
