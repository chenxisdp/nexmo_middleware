package main

import (
	"github.com/astaxie/beego"
	"github.com/chenxisdp/nexmo_middleware/controllers"
)

/*
func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/shorten", &controllers.ShortController{})
	beego.Router("/v1/expand", &controllers.ExpandController{})
	beego.Run()
}
*/

/*
type AnswerMessage struct {
	Action    string
	VoiceName string
	Text      string
}

func main() {
	answerMessage := &AnswerMessage{}
	answerMessage.Action = "talk"
	answerMessage.VoiceName = "Russell"
	answerMessage.Text = "Hi, this is Russell. You are listening to a text-to-speech Call made with Nexmo's Voice API"

	jsonData, _ := json.Marshal(answerMessage)
	fmt.Println(string(jsonData))

	timeNow := time.Now().Unix()
	timeNowStr := strconv.FormatInt(timeNow, 10)
	outfile, err := os.Create(timeNowStr + ".json")
	defer outfile.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = outfile.Write(jsonData)
	if err != nil {
		log.Fatal(err)
		return
	}
}
*/

func main() {
	beego.Router("/phone", &controllers.PhoneController{})
	beego.Run()
}
