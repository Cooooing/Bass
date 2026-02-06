package tests

import (
	"bytes"
	"common/pkg/model"
	"fmt"
	"html/template"
	"testing"

	"github.com/google/uuid"
)

func TestNotificationTemplate(t *testing.T) {
	content := `
    <p>Hello {{.}}</p> {{.User.Name}} 发布文章 {{.Article.Title}}
`
	// 解析指定文件生成模板对象
	tmpl, err := template.New("").Parse(content)

	if err != nil {
		fmt.Println("create content failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	user := model.User{Name: "name222"}
	w := &bytes.Buffer{}
	err = tmpl.Execute(w, map[string]any{"User": user, "Article": map[string]any{"Title": "title2222"}})
	if err != nil {
		fmt.Println("execute failed, err:", err)
		return
	}
	fmt.Println(w.String())
}

/*
    <p>Hello map[Article:map[author_user:&lt;nil&gt; bounty_points:0 commentable:true content:esse
 content_render: cover_image_url:&lt;nil&gt; created_at:2025-12-09T14:22:41.577765Z created_by:1 created_by_name:admin edges:map[] id:14 last_replied_at:&lt;nil&gt; last_reply_user:&lt;nil&gt; listable:true reward_content_render:&lt;nil&gt; reward_points:0 status:3 title:test updated_at:2025-12-09T14:22:41.577765Z updated_by:1 updated_by_name:admin] User:map[avatar_url:https://treble.sxisa.com/v1/user/avatar/admin created_at:2025-12-09T13:42:30.412511Z email:uclgm3.gc3@sohu.com enable_email_subscribe:true enable_web_notify:true id:1 language:zh-CN mobile_theme:default name:admin nickname:admin password:$2a$10$JUfPISgMcWKHnXH2r8OnPuIlugREdpUbO2./QqhZiHSQIJpiWRDHS public_articles:true public_comments:true public_followers:true public_location:true public_online_status:true public_points:true theme:default timezone:Asia/Shanghai updated_at:2025-12-09T13:42:30.412511Z]]</p>  发布文章

*/

func TestName(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(uuid.New().String())
	}
}
