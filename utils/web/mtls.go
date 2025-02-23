package main

// mtls.go creates an MTLS server or client.
// MTLS is mutual TLS (enabled with ClientAuth = tls.RequireAndVerifyClientCert)
//  - client authenticates the server as with TLS, but
//  - server also authenticates the client (using client's TLS certificate)
// To run a test first invoke the server using -server cmd-line flag then invoke client(s)

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	Address = "localhost:8043"

	ServerCertFile = `testdata/svr.crt` // test self-signed cert
	ServerKeyFile  = `testdata/svr.key`

	ClientCertFile = `testdata/confman.crt`
	ClientKeyFile  = `testdata/confman.key`
)

func main() {
	server := flag.Bool("server", false, "Run as server")
	flag.Parse()

	if *server {
		start()
	} else {
		fmt.Println(httpsRequest("https://" + Address + "/"))
	}
}

//////////// SERVER ////////////

func start() {
	// Load client's cert into new pool
	pool := x509.NewCertPool()
	pem, err2 := os.ReadFile(ClientCertFile)
	if err2 != nil {
		log.Fatalf("Error %v reading client certificate file %q", err2, ClientCertFile)
	}
	if !pool.AppendCertsFromPEM(pem) {
		log.Fatalf("Error in certificate file %q", ClientCertFile)
	}

	// Load our cert and private key (maybe encrypt private key on disk)
	cert, err := tls.LoadX509KeyPair(ServerCertFile, ServerKeyFile)
	if err != nil {
		log.Fatalf("Error %v loading certificate %q and key %q", err, ServerCertFile, ServerKeyFile)
	}

	// Start the server with TLS
	server := http.Server{
		Addr:    Address,
		Handler: http.HandlerFunc(httpRequestHandler),
		TLSConfig: &tls.Config{
			ClientAuth:   tls.RequireAndVerifyClientCert, // MTLS: client must provide a cert + server must verify it
			ClientCAs:    pool,                           // MTLS: we need a cert. to verify the cert. sent from client
			Certificates: []tls.Certificate{cert},
			ServerName:   "localhost",
		},
	}
	defer server.Close()

	log.Fatal(server.ListenAndServeTLS("", ""))
}

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello,World!\n"))
}

//////////// CLIENT ////////////

func httpsRequest(url string) string {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      mustGetPool(ServerCertFile),
				Certificates: mustGetCertificates(ClientCertFile, ClientKeyFile), // MTLS: required by server (ClientAuth=tls.RequireAndVerifyClientCert)
				ServerName:   "localhost",
			},
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("Response status:", resp.Status)
	}
	msg, _ := io.ReadAll(resp.Body)
	return string(msg)
}

func mustGetPool(files ...string) *x509.CertPool {
	pool := x509.NewCertPool()

	for _, f := range files {
		pem, err := os.ReadFile(f)
		if err != nil {
			log.Fatalf("mustGetPool: error %v reading certificate file %q", err, f)
		}
		if ok := pool.AppendCertsFromPEM(pem); !ok {
			log.Fatalf("mustGetPool: error in certificate file %q", f)
		}
	}
	return pool
}

func mustGetCertificates(files ...string) (r []tls.Certificate) {
	if len(files)%2 != 0 {
		log.Fatalf("mustGetCertificates: files (cert/key pairs) must have an even number of elements")
	}
	for i := 0; i < len(files); i += 2 {
		cert, err := tls.LoadX509KeyPair(files[i], files[i+1])
		if err != nil {
			log.Fatalf("mustGetCertificates: error %v loading x509 certificate %q, key %q", err, files[i], files[i+1])
		}
		r = append(r, cert)
	}
	return
}
