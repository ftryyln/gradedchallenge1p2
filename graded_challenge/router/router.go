package router

import (
	"graded-challenge/handler"
	"net/http"
	"strconv"
	"strings"
)

// SetupRouter mengatur routing HTTP dan menghubungkan endpoint ke handler
func SetupRouter(ctrl *handler.EmployeesHandler) http.Handler {
	// Membuat instance multiplexer (router bawaan net/http)
	mux := http.NewServeMux()

	// Handler untuk endpoint "/employees"
	mux.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Ambil semua data karyawan
			ctrl.GetAllEmployees(w, r)
		case http.MethodPost:
			// Buat karyawan baru
			ctrl.CreateEmployee(w, r)
		default:
			// Method selain GET dan POST tidak diizinkan
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Handler untuk endpoint "/employees/{id}"
	mux.HandleFunc("/employees/", func(w http.ResponseWriter, r *http.Request) {
		// Mengambil ID dari URL path
		idStr := strings.TrimPrefix(r.URL.Path, "/employees/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// Routing berdasarkan HTTP method
		switch r.Method {
		case http.MethodGet:
			// Ambil data karyawan berdasarkan ID
			ctrl.GetEmployeeByID(w, r, id)
		case http.MethodPut:
			// Update data karyawan berdasarkan ID
			ctrl.UpdateEmployee(w, r, id)
		case http.MethodDelete:
			// Hapus data karyawan berdasarkan ID
			ctrl.DeleteEmployee(w, r, id)
		default:
			// Method selain GET, PUT, DELETE tidak diizinkan
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Mengembalikan router yang sudah dikonfigurasi
	return mux
}
