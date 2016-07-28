package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["go-club/controllers:BannerController"] = append(beego.GlobalControllerRouter["go-club/controllers:BannerController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:BannerController"] = append(beego.GlobalControllerRouter["go-club/controllers:BannerController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:BannerController"] = append(beego.GlobalControllerRouter["go-club/controllers:BannerController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:BannerController"] = append(beego.GlobalControllerRouter["go-club/controllers:BannerController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:BannerController"] = append(beego.GlobalControllerRouter["go-club/controllers:BannerController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:BlockController"] = append(beego.GlobalControllerRouter["go-club/controllers:BlockController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:BlockController"] = append(beego.GlobalControllerRouter["go-club/controllers:BlockController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:BlockController"] = append(beego.GlobalControllerRouter["go-club/controllers:BlockController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:BlockController"] = append(beego.GlobalControllerRouter["go-club/controllers:BlockController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:BlockController"] = append(beego.GlobalControllerRouter["go-club/controllers:BlockController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:SectionController"] = append(beego.GlobalControllerRouter["go-club/controllers:SectionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:SectionController"] = append(beego.GlobalControllerRouter["go-club/controllers:SectionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:SectionController"] = append(beego.GlobalControllerRouter["go-club/controllers:SectionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:SectionController"] = append(beego.GlobalControllerRouter["go-club/controllers:SectionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:SectionController"] = append(beego.GlobalControllerRouter["go-club/controllers:SectionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:StoryController"] = append(beego.GlobalControllerRouter["go-club/controllers:StoryController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:StoryController"] = append(beego.GlobalControllerRouter["go-club/controllers:StoryController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:StoryController"] = append(beego.GlobalControllerRouter["go-club/controllers:StoryController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:StoryController"] = append(beego.GlobalControllerRouter["go-club/controllers:StoryController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:StoryController"] = append(beego.GlobalControllerRouter["go-club/controllers:StoryController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:StoryReplyController"] = append(beego.GlobalControllerRouter["go-club/controllers:StoryReplyController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:StoryReplyController"] = append(beego.GlobalControllerRouter["go-club/controllers:StoryReplyController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:StoryReplyController"] = append(beego.GlobalControllerRouter["go-club/controllers:StoryReplyController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:StoryReplyController"] = append(beego.GlobalControllerRouter["go-club/controllers:StoryReplyController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:StoryReplyController"] = append(beego.GlobalControllerRouter["go-club/controllers:StoryReplyController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TabController"] = append(beego.GlobalControllerRouter["go-club/controllers:TabController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TabController"] = append(beego.GlobalControllerRouter["go-club/controllers:TabController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TabController"] = append(beego.GlobalControllerRouter["go-club/controllers:TabController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TabController"] = append(beego.GlobalControllerRouter["go-club/controllers:TabController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TabController"] = append(beego.GlobalControllerRouter["go-club/controllers:TabController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TopicController"] = append(beego.GlobalControllerRouter["go-club/controllers:TopicController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TopicController"] = append(beego.GlobalControllerRouter["go-club/controllers:TopicController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TopicController"] = append(beego.GlobalControllerRouter["go-club/controllers:TopicController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TopicController"] = append(beego.GlobalControllerRouter["go-club/controllers:TopicController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TopicController"] = append(beego.GlobalControllerRouter["go-club/controllers:TopicController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TopicReplyController"] = append(beego.GlobalControllerRouter["go-club/controllers:TopicReplyController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TopicReplyController"] = append(beego.GlobalControllerRouter["go-club/controllers:TopicReplyController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TopicReplyController"] = append(beego.GlobalControllerRouter["go-club/controllers:TopicReplyController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TopicReplyController"] = append(beego.GlobalControllerRouter["go-club/controllers:TopicReplyController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:TopicReplyController"] = append(beego.GlobalControllerRouter["go-club/controllers:TopicReplyController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:UserController"] = append(beego.GlobalControllerRouter["go-club/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:UserController"] = append(beego.GlobalControllerRouter["go-club/controllers:UserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:UserController"] = append(beego.GlobalControllerRouter["go-club/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:UserController"] = append(beego.GlobalControllerRouter["go-club/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-club/controllers:UserController"] = append(beego.GlobalControllerRouter["go-club/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
