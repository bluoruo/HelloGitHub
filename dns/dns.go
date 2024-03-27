package dns

import (
	"context"
	"fmt"
	"HelloGitHub/lib"
	"net"
	"time"
)

func NsLookUp(domain, dnsServer string) (string, bool) {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: 5 * time.Second,
			}
			return d.DialContext(ctx, "udp", dnsServer+":53")
		},
	}
	ips, err := r.LookupHost(context.Background(), domain)
	if err == nil && len(ips) > 0 {
		return ips[0], true
	}
	fmt.Println("查询DNS出错", err)
	lib.Logger.Println("[DNS查询] 出错:", err)
	return "", false
}
