package painter

import (
	"SuperArchWorker/conf"
	"SuperArchWorker/middleware/taskcontrol"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Txt2Image struct{
	taskcontrol.TaskModuleCommon
	Text string
}

type Txt2ImageTrans struct{
	Prompt 	string				`json:"prompt"`
}

func (txt2Image Txt2Image)Handle() string{
	logrus.Info("[Txt2Image][Handle] Start!")
	url := fmt.Sprintf("%s/sdapi/v1/txt2img", conf.Cfg.WorkerConf["ai.painter.txt2image"]["api_url"])
	data := Txt2ImageTrans{
		Prompt: txt2Image.Text,
	}
	b, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Authorization", "Basic YWRtaW46SGVsbG9BSSExMjM=")
	req.Header.Set("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		logrus.Errorf("[Txt2Image][Handle][Text to Image] %s", err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)

}