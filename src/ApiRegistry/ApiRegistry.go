/*
Package registry provides functions as service registry.
1. Regist services from the application.
2. Load balance. */
package registry

var hostMap = map[string]string{
	"BPO": "localhost:8012",
	"FYU": "www.fyu.com:80",
	"ESB": "www.esb.com:80",
	"SDB": "www.sdb.com:80",
}

// HostAddress gets host address by the product ID.
func HostAddress(productID string) string {
	if hostAddress, ok := hostMap[productID]; ok {
		return hostAddress
	}
	panic("Unrecognized product, please check your request url.")
}
