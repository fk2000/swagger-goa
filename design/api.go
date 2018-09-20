// ./design/api.go
package design

// golint:should not use dot imports と怒られるので名前付きでimportする
import (
  goa "github.com/goadesign/goa/design"
  dsl "github.com/goadesign/goa/design/apidsl"
  // API Resources
  _ "github.com/BcRikko/learning-goa/design/resources"
)

// APIの定義(http://localhost:8080/api)
var _ = dsl.API("Task", func() {
  dsl.Title("タスク管理API")
  dsl.Description("タスク管理のAPIです。")
  dsl.Version("0.1")
  dsl.Scheme("http")
  dsl.Host("localhost:8080")
  dsl.BasePath("/api")
  dsl.Consumes("application/json") // リクエストのメディアタイプ
  dsl.Produces("application/json") // レスポンスのメディアタイプ

        // 202:Createdのときのレスポンステンプレートを定義
  dsl.ResponseTemplate(goa.Created, func(pattern string) {
    dsl.Description("タスク作成が完了しました。")
    dsl.Status(201)
    dsl.Headers(func() {
      dsl.Header("Location", goa.String, "作成したタスクのリンク", func() {
        dsl.Pattern(pattern)
      })
    })
  })
})

