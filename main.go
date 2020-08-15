package main

import (
	"log"
	"net/http"
	"os"

	//"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/kben19/CleoApp/pkg/server"
)

const (
	InstanceConnectionName = "heroic-oarlock-286113:asia-southeast2:database-prod"
	DatabaseName           = "database-prod"
	DatabaseUser           = "postgres"
	DatabasePass           = "D3vSc@rl3t"
)

func main() {
	//dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
	//	InstanceConnectionName,
	//	DatabaseName,
	//	DatabaseUser,
	//	DatabasePass)

	// Create DB pool
	//db, err := sql.Open("cloudsqlpostgres", dsn)
	//if err != nil {
	//	log.Fatal("Failed to open a DB connection: ", err)
	//}
	//defer db.Close()

	usecases := server.Init(nil)

	server.RegisterHandler(usecases)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

