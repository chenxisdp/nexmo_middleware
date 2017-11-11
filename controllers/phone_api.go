package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/astaxie/beego"
)

//phone message,Tos means phone number
type PhoneTTS struct {
	Tos     string `json:"tos"`
	Content string `json:"content"`
}

type PhoneController struct {
	beego.Controller
}

func (this *PhoneController) Post() {
	var body PhoneTTS
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	log.Println(body)
	fmt.Println(body)
	if err != nil {
		this.RenderError(err.Error())
	} else {
		this.RenderJson(body)
	}
}

func (this *PhoneController) RenderError(msg string) {
	resp := NewDto()
	resp.Code = "fail"
	resp.Msg = msg
	this.Data["json"] = resp
	this.Ctx.Output.SetStatus(400)
	this.ServeJSON()
}

func (this *PhoneController) RenderJson(json interface{}) {
	this.Data["json"] = json
	this.ServeJSON()
}

type Dto struct {
	Code        string `json:"code"`
	Msg         string `json:"message"`
	Host_id     string `json:"host_id"`
	Server_time string `json:"server_time"`
}

func NewDto() *Dto {
	host_name, _ := os.Hostname()
	return &Dto{
		Host_id:     host_name,
		Server_time: time.Now().Format(time.RFC3339),
	}
}
