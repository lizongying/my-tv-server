package mytv

import (
	"log"
	"net"
)

func Lan() (lanIp string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			var addrs []net.Addr
			addrs, err = i.Addrs()
			if err != nil {
				log.Println(err)
				continue
			}

			for _, a := range addrs {
				if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					lanIp = ipNet.IP.String()
					break
				}
			}
		}
	}
	return
}
