package mytv

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"net/http"
)

func Lan() (lanIp string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {

			if strings.HasPrefix(i.Name, "VMnet") || strings.HasPrefix(i.Name, "VirtualBox") {
				continue
			}

			var addrs []net.Addr
			addrs, err = i.Addrs()
			if err != nil {
				log.Println(err)
				continue
			}

			for _, a := range addrs {
				if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && !ipNet.IP.IsLinkLocalUnicast() && ipNet.IP.To4() != nil {
					lanIp = ipNet.IP.String()
					break
				}
			}
		}
	}
	return
}

func Internet() (ip string) {
	resp, err := http.Get("https://api64.ipify.org?format=text")
	if err != nil {
		return ""
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	_, _ = fmt.Fscanf(resp.Body, "%s", &ip)
	return
}
