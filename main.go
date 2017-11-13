package main

import (
	"github.com/chenxisdp/nexmo_middleware/nccologic"
)

func main() {
	//	beego.Router("/phone", &controllers.PhoneController{})
	//	beego.Run()

	answerMessage := nccologic.AnswerMessage{}
	answerMessage.Action = "talk"
	answerMessage.VoiceName = "Russell"
	answerMessage.Text = "Hi, this is Russell. You are listening to a text-to-speech Call made with Nexmo's Voice API"
	nccologic.BuildAnswerURL(answerMessage)
}
