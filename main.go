package main

import (
	"graded-challenge/config"
	"graded-challenge/handler"
	"graded-challenge/repository"
	"graded-challenge/router"
	"graded-challenge/service"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load environment variables dari file .env jika ada
	config.LoadEnv()

	// Inisialisasi koneksi database
	db := config.InitDB()
	if db == nil {
		log.Fatal("Database not initialized.")
	}
	defer db.Close()

	// Inisialisasi repository, service, dan handler
	repo := &repository.EmployeesRepository{DB: db}
	svc := service.NewEmployeeService(repo)
	h := &handler.EmployeesHandler{Service: svc}

	// Setup router dan jalankan HTTP server
	r := router.SetupRouter(h)

	// Menjalankan server pada railway
	var PORT = os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(": "+PORT, r))
}
