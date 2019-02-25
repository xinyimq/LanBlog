package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["MyBlog/controllers:AdminController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:AdminController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Options",
            Router: `/login`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:LabelController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:LabelController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:LabelController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:LabelController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:LabelController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:LabelController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:LabelController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:LabelController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:LabelController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:LabelController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:ObjectController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:ObjectController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:ObjectController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:ObjectController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:ObjectController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:TopicController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:TopicController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:TopicController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:TopicController"],
        beego.ControllerComments{
            Method: "GetByCate",
            Router: `/cate/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:TopicController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:TopicController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:TopicController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:TopicController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:TopicController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:TopicController"],
        beego.ControllerComments{
            Method: "GetByVagueName",
            Router: `/find/:vname`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:TopicController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:TopicController"],
        beego.ControllerComments{
            Method: "GetByLabel",
            Router: `/label/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:TopicController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:TopicController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:TopicController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:TopicController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
