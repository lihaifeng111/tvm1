package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
	//"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/upload_transaction",&controllers.MainController{},"Get:GetTransaction;Post:PushTransaction")
    beego.Router("/executeContract",&controllers.MainController{},"Get:GetExecuteContract;Post:ExecuteContract")
}
