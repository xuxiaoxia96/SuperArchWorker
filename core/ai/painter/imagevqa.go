package painter

import (
	"SuperArchWorker/conf"
	"SuperArchWorker/middleware/taskcontrol"
	"SuperArchWorker/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type ImageVQA struct{
	taskcontrol.TaskModuleCommon
	question  string
	ImageData map[string]string
}

type ImageVQATrans struct{
	ImageId			string		`json:"image_id"`
	ImageTransData 	string		`json:"image_data"`
	Question		string		`json:"question"`
}

func (imagevqa *ImageVQA) Handle() string{
	logrus.Info("[ImageVQA][Handle] Start!")

	var data ImageVQATrans
	url := fmt.Sprintf("%s/blip/vqa", conf.Cfg.WorkerConf["ai.painter.imagevqa"]["api_url"])
	for hashId,imageData := range imagevqa.ImageData{
		if ! utils.ImgMatchHash(imageData, hashId){
			logrus.Errorf("[ImageVQA][Handle] Image not match the hash")
			continue
		}
		data = ImageVQATrans{
			ImageId: hashId,
			ImageTransData: imageData,
			Question: imagevqa.question,
		}
		break
	}

	// handle python script
	// get image vqa
	b, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Authorization", "Basic YWRtaW46SGVsbG9BSSExMjM=")
	req.Header.Set("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		logrus.Errorf("[ImageScore][Handle][Get Image score] %s", err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
