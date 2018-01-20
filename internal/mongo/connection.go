package mongo

import (
	"crypto/tls"
	"net"
	"strings"

	"gopkg.in/mgo.v2"
)

// ParseURL parses connection string and returns dial info.
// it supports ssl=true flag and configures dialinfo to use SSL if flag is passed.
func ParseURL(connstring string) (*mgo.DialInfo, error) {

	ip := connstring

	hasSSL := false
	i := strings.Index(ip, "?")
	if i != -1 {
		ssl := strings.Index(ip[i:], "ssl=true")
		if ssl != -1 {
			hasSSL = true
			ip = strings.Replace(ip, "ssl=true", "", -1)
		}
	}

	dialInfo, err := mgo.ParseURL(ip)

	if hasSSL {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: false,
		}

		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
	}

	return dialInfo, err
}
