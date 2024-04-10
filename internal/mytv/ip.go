package mytv

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Ip struct {
	port     int
	channels []byte
}

type RespIp struct {
	Code int `json:"code"`
	Data struct {
		LocalUrl    string `json:"local_url"`
		LanUrl      string `json:"lan_url"`
		InternetUrl string `json:"internet_url"`
	} `json:"data"`
}

func NewIp(port int) (i *Ip, err error) {
	i = &Ip{
		port: port,
	}

	return
}

func (i *Ip) Info(w http.ResponseWriter, _ *http.Request) {
	var resp = new(RespIp)
	resp.Data.LocalUrl = fmt.Sprintf("%s:%d", "127.0.0.1", i.port)
	resp.Data.LanUrl = fmt.Sprintf("%s:%d", Lan(), i.port)
	resp.Data.InternetUrl = fmt.Sprintf("%s:%d", Internet(), i.port)
	data, _ := json.Marshal(resp)
	_, _ = w.Write(data)
}
