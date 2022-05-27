package tools

import (
	"crypto/tls"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
)

// GetTLSConfig
// 用户获取TLS配置，读取了server.key 和 server.pem证书凭证文件
// tls.X509KeyPair：   从一对PEM编码的数据中解析公钥/私钥对
// tls.Certificate:    返回一个或者多个证书
// http2.NextProtoTLS：用户HTTP/2的TLS配置
func GetTLSConfig(certPemPath, certKeyPath string) *tls.Config {
	cert, _ := ioutil.ReadFile(certPemPath)
	key, _ := ioutil.ReadFile(certKeyPath)

	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Printf("TLS KeyPair err: %v\n\n", err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{pair},
		NextProtos:   []string{http2.NextProtoTLS},
	}
}
