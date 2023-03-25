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

type Image2Image struct{
	taskcontrol.TaskModuleCommon
	Text string
	ImageData string
}

type Image2ImageTrans struct {
	Prompt 		string 		`json:"prompt"`
	InitImages 	[]string	`json:"init_images"`
}

func (image2Image Image2Image)Handle() string{
	logrus.Info("[Image2Image][Handle] Start!")
	url := fmt.Sprintf("%s/sdapi/v1/img2img", conf.Cfg.WorkerConf["ai.painter.image2image"]["api_url"])
	data := Image2ImageTrans{
		Prompt: image2Image.Text,
		InitImages: []string{image2Image.ImageData},
	}
	b, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Authorization", "Basic YWRtaW46SGVsbG9BSSExMjM=")
	req.Header.Set("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		logrus.Errorf("[Image2Image][Handle][Image to Image] %s", err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}