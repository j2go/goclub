// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"go-club/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/banner",
			beego.NSInclude(
				&controllers.BannerController{},
			),
		),

		beego.NSNamespace("/block",
			beego.NSInclude(
				&controllers.BlockController{},
			),
		),

		beego.NSNamespace("/story",
			beego.NSInclude(
				&controllers.StoryController{},
			),
		),

		beego.NSNamespace("/story_reply",
			beego.NSInclude(
				&controllers.StoryReplyController{},
			),
		),

		beego.NSNamespace("/tab",
			beego.NSInclude(
				&controllers.TabController{},
			),
		),

		beego.NSNamespace("/topic",
			beego.NSInclude(
				&controllers.TopicController{},
			),
		),

		beego.NSNamespace("/topic_reply",
			beego.NSInclude(
				&controllers.TopicReplyController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
