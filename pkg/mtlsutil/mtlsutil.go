package mtlsutil

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc/credentials"
)

// Creates mTLS creds for the client side
// Makes a copy of the system trust store and appends the passed cert
func NewMutualTLSClientCreds(certFile, keyFile, serverName string) (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	certPool := NewCustomCertPool(certFile)

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   serverName,
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	})

	return creds, nil
}

// Creates mTLS creds for the server side
// Makes a copy of the system trust store and appends the passed cert
func NewMutualTLSServerCreds(certFile, keyFile string) (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	certPool := NewCustomCertPool(certFile)

	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certPool,
	})

	return creds, nil
}

// Creates an in-mem copy of the system trust store (if possible) and appends the custom cert
// If using a valid CA cert, just use the system default trust store
func NewCustomCertPool(certFile string) *x509.CertPool {
	// Get the SystemCertPool, continue with an empty pool on error
	certPool, _ := x509.SystemCertPool()
	if certPool == nil {
		certPool = x509.NewCertPool()
	}

	// Append cert to the system pool to "trust it"
	certPEMBlock, _ := ioutil.ReadFile(certFile)
	if ok := certPool.AppendCertsFromPEM(certPEMBlock); !ok {
		log.Println("No certs appended, using system certs only")
	}

	return certPool
}
