package mytv

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Mytv struct {
	port     int
	channels []byte
}

type RespUpload struct {
	Code int `json:"code"`
	Data struct {
		LocalUrl    string `json:"local_url"`
		LanUrl      string `json:"lan_url"`
		InternetUrl string `json:"internet_url"`
	} `json:"data"`
}

func NewMytv(port int) (m *Mytv, err error) {
	m = &Mytv{
		port: port,
	}

	return
}

func (m *Mytv) Upload(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	m.channels = body
	var resp = new(RespUpload)
	resp.Data.LocalUrl = fmt.Sprintf("%s:%d/0", "127.0.0.1", m.port)
	resp.Data.LanUrl = fmt.Sprintf("%s:%d/0", Lan(), m.port)
	resp.Data.InternetUrl = fmt.Sprintf("%s:%d/0", Internet(), m.port)
	data, _ := json.Marshal(resp)
	_, _ = w.Write(data)
}

func (m *Mytv) Channels(w http.ResponseWriter, _ *http.Request) {
	if len(m.channels) == 0 {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}
	_, _ = w.Write(m.channels)
}
