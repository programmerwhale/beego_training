package basic

import "github.com/beego/beego/v2/server/web"

type UserController struct {
	web.Controller
}

func (u *UserController) HelloWorld()  {
	u.Ctx.WriteString("hello, world")
}

func main() {
	// get http://127.0.0.1:8080/user/helloworld
	web.AutoRouter(&UserController{})
	web.Run()
}