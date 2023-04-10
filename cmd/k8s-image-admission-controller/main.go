package main

import (
	"os"

	"k8s-image-admission-controller/pkg/config"
)

func main() {
	cmd := config.NewRootCommand()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// func runWebhookServer(certFile, keyFile string) {
// 	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Starting webhook server")
// 	http.HandleFunc("/validate", k8simageadmissioncontroller.ValidatePod)
// 	server := http.Server{
// 		Addr: fmt.Sprintf(":%d", 8080),
// 		TLSConfig: &tls.Config{
// 			Certificates: []tls.Certificate{cert},
// 		},
// 	}

// 	if err := server.ListenAndServeTLS("", ""); err != nil {
// 		panic(err)
// 	}
// }
