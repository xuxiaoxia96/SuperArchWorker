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

type Image2Txt struct{
	taskcontrol.TaskModuleCommon
	ImageData map[string]string
}

type Image2TxtTrans struct{
	ImageId			string		`json:"image_id"`
	ImageTransData 	string		`json:"image_data"`
}

func (image2txt *Image2Txt) Handle() string{
	logrus.Info("[Image2Txt][Handle] Start!")

	var data Image2TxtTrans
	url := fmt.Sprintf("%s/blip/caption", conf.Cfg.WorkerConf["ai.painter.image2txt"]["api_url"])
	for hashId, imageData := range image2txt.ImageData{
		if ! utils.ImgMatchHash(imageData, hashId){
			logrus.Errorf("[Image2Txt][Handle] Image not match the hash")
			continue
		}
		data = Image2TxtTrans{
			ImageId: hashId,
			ImageTransData: imageData,
		}
		break
	}

	// handle python script
	// get image txt
	b, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Authorization", "Basic YWRtaW46SGVsbG9BSSExMjM=")
	req.Header.Set("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		logrus.Errorf("[Image2Txt][Handle][Get Image Caption] %s", err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
