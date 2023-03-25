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

type ImageScore struct{
	taskcontrol.TaskModuleCommon
	ImageData 		map[string]string
}

type ImageScoreTrans struct{
	ImageId			string		`json:"image_id"`
	ImageTransData 	string		`json:"image_data"`
}


func (imageScore *ImageScore) Handle() string{
	logrus.Info("[ImageScore][Handle] Start!")

	// get image score
	var data ImageScoreTrans
	url := fmt.Sprintf("%s/ai/painter/imagescore", conf.Cfg.WorkerConf["ai.painter.imagescore"]["api_url"])
	for hashId, imageData := range imageScore.ImageData{
		if ! utils.ImgMatchHash(imageData, hashId){
			logrus.Errorf("[ImageScore][Handle] Image not match the hash")
			continue
		}
		data = ImageScoreTrans{
			ImageId: hashId,
			ImageTransData: imageData,
		}
		break
	}

	b, _ := json.Marshal(data)

	resp, err := http.Post(url,
		"application/json",
		bytes.NewBuffer(b))
	if err != nil {
		logrus.Errorf("[ImageScore][Handle][Get Image score] %s", err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}