package iptables

import (
	"fmt"
	iptables_tool "istio.io/istio/offmesh-tools/iptables/go-iptables"
)

// AddIPTableRedirect = iptables -t nat -A PREROUTING -p tcp -s ${srcIP} -j DNAT --to-destination ${proxyIP}:${proxyPort}
func AddIPTableRedirect(srcIP string, proxyIP string, proxyPort string) error {
	iptable, err := iptables_tool.New()
	if err != nil {
		return err
	}
	err = iptable.AppendUnique("nat", "PREROUTING", "-p", "tcp", "-s", srcIP, "-j", "DNAT", "--to-destination", fmt.Sprintf("%s:%s", proxyIP, proxyPort))
	if err != nil {
		return err
	}
	return nil
}

// DeleteIPTableRedirect = iptables -t nat -D PREROUTING -p tcp -s ${srcIP} -j DNAT --to-destination ${proxyIP}:${proxyPort}
func DeleteIPTableRedirect(srcIP string, proxyIP string, proxyPort string) error {
	iptable, err := iptables_tool.New()
	if err != nil {
		return err
	}
	err = iptable.Delete("nat", "PREROUTING", "-p", "tcp", "-s", srcIP, "-j", "DNAT", "--to-destination", fmt.Sprintf("%s:%s", proxyIP, proxyPort))
	if err != nil {
		return err
	}
	return nil
}

//must run in sudo mod
//func main() {
//	cmdType := os.Args[1]
//	srcIP := os.Args[2]
//	proxyIP := os.Args[3]
//	proxyPort := os.Args[4]
//	if cmdType == `ADD` {
//		err := AddIPTableRedirect(srcIP, proxyIP, proxyPort)
//		if err != nil {
//			log.Error(err)
//		}
//	} else if cmdType == `DELETE` {
//		err := DeleteIPTableRedirect(srcIP, proxyIP, proxyPort)
//		if err != nil {
//			log.Error(err)
//		}
//	}
//}
