package mytv

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

type Example struct {
	path     string
	host     string
	channels []byte
}

type Channel struct {
	Group string   `json:"group,omitempty"`
	Logo  string   `json:"logo,omitempty"`
	Name  string   `json:"name,omitempty"`
	Title string   `json:"title,omitempty"`
	Uris  []string `json:"uris,omitempty"`
}

func NewExample(path string, host string) (e *Example, err error) {
	e = &Example{
		path: path,
		host: host,
	}

	return
}

func (e *Example) Channels(w http.ResponseWriter, r *http.Request) {
	var channels = make([]Channel, 0)

	files, err := os.ReadDir(e.path + "/video")
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		name := file.Name()
		channel := Channel{
			Group: "本地視頻",
			Title: name[:strings.LastIndex(name, ".")],
			Uris:  []string{"http://" + r.Host + "/assets/video/" + name},
		}
		channels = append(channels, channel)
	}

	data, _ := json.Marshal(channels)
	_, _ = w.Write(data)
}
