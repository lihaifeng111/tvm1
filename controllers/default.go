
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ipfs/go-ipfs-api"
	"os"
	"strings"
	"time"
)

type MainController struct {
	beego.Controller
}

type Transaction struct{
	Type string `json:"type"`
	FromAddress string `json:"from_address"`
	ToAddress string `json:"to_address"`
	Amount int `json:"amount"`
	Time int64 `json:"time"`
}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) GetTransaction(){
	c.TplName="writeTransaction.html"
}

func (c *MainController) PushTransaction() string{
	//获取页面提交的数据
	typeName:=c.GetString("typeName")
	from:=c.GetString("from_address")
	to:=c.GetString("to_address")
	amount,_:=c.GetInt("amount")
	now:=time.Now().Unix()
	//formatNow := now.Format("2006-01-02 15:04:05")

	transaction:=Transaction{typeName,from,to,amount,now}
	//beego.Info(transaction)
	tJson, err := json.Marshal(transaction)
	if err!=nil{
		beego.Info("JSON转换错误")
		os.Exit(1)
	}
	transactionJson:=string(tJson)
	beego.Info(transactionJson)

	resultFlag,resultMsg:=cacheJsonInTangle(transactionJson)
	if resultFlag{
		return "ok"
	}else{
		return resultMsg
	}


}

func cacheJsonInTangle(transactionJson string) (bool,string){
	sh := shell.NewShell(beego.AppConfig.String("ipfsport"))
	fmt.Println(beego.AppConfig.String("ipfsport"))
	cid, err := sh.Add(strings.NewReader(transactionJson))

	if err != nil {
		//fmt.Fprintf(os.Stderr, "error: %s", err)
		fmt.Println(err)
		os.Exit(1)
	}

	if cid[:2]=="Qm"{
    fmt.Println(cid)
    return true,"ok"
	}else{
		return false,"上传ipfs错误"
	}

}

func (c *MainController)GetExecuteContract(){
	c.TplName="tvm-executeContract.html"
}

func (c *MainController)ExecuteContract(){
	typeName:=c.GetString("typeName")
	from:=c.GetString("from_address")
	to:=c.GetString("to_address")
	amount,_:=c.GetInt("amount")
	now:=time.Now().Unix()
	//formatNow := now.Format("2006-01-02 15:04:05")

	transaction:=Transaction{typeName,from,to,amount,now}
	//beego.Info(transaction)
	tJson, err := json.Marshal(transaction)
	if err!=nil{
		beego.Info("JSON转换错误")
		os.Exit(1)
	}
	demo:=string(tJson)
	beego.Info(demo)
}