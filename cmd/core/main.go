package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/furee/backend/cmd/core/config"
	"github.com/furee/backend/cmd/core/routes"
	"github.com/gorilla/handlers"
)

func main() {
	id := time.Now().UTC().Add(time.Duration(525960 * time.Minute)).Unix()
	secret := "^qwertyuiop" //public secret key

	authCompareByte := sha256.Sum256([]byte(fmt.Sprintf("%s%d", secret, id)))

	fmt.Println(id)
	fmt.Printf("%x\n", authCompareByte)

	conf, err := config.GetCoreConfig()
	if err != nil {
		panic(err)
	}

	handler, log, err := config.NewRepoContext(conf)
	if err != nil {
		panic(err)
	}

	headers := handlers.AllowedHeaders(conf.Route.Headers)
	methods := handlers.AllowedMethods(conf.Route.Methods)
	origins := handlers.AllowedOrigins([]string{conf.Route.Origins.InternalTools})
	credentials := handlers.AllowCredentials()

	router := routes.GetCoreEndpoint(conf, handler, log)

	port := fmt.Sprintf(":%s", conf.App.Port)
	log.Info("server listen to port ", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(headers, methods, origins, credentials)(router)))
}
