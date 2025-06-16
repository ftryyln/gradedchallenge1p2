package handler

import (
	"encoding/json"
	"errors"
	"graded-challenge/model"
	"graded-challenge/service"
	"net/http"
)

// EmployeesHandler bertanggung jawab menangani HTTP request terkait data karyawan
type EmployeesHandler struct {
	Service *service.EmployeeService
}

// Response adalah struct generic untuk standar format response API
type Response[T any] struct {
	// Kode status (misal: "200", "404")
	ResponseCode    string `json:"responseCode"`
	// Pesan status (misal: "Success", "Not Found")
	ResponseMessage string `json:"responseMessage"`
	// Data hasil response (jika ada)
	Data            T      `json:"data,omitempty"`
}

// GetAllEmployees menangani request untuk mengambil semua data karyawan
func (h *EmployeesHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.GetAllEmployees()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response[any]{ResponseCode: "500", ResponseMessage: "Error fetching data"})
		return
	}
	json.NewEncoder(w).Encode(Response[any]{ResponseCode: "200", ResponseMessage: "Success", Data: data})
}

// GetEmployeeByID menangani request untuk mengambil data karyawan berdasarkan ID
func (h *EmployeesHandler) GetEmployeeByID(w http.ResponseWriter, r *http.Request, id int) {
	data, err := h.Service.GetEmployeeByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response[any]{ResponseCode: "404", ResponseMessage: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(Response[any]{ResponseCode: "200", ResponseMessage: "Success", Data: data})
}

// CreateEmployee menangani request untuk menambahkan data karyawan baru
func (h *EmployeesHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var p model.Employees
	json.NewDecoder(r.Body).Decode(&p)

	created, err := h.Service.CreateEmployee(p)
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response[any]{ResponseCode: "400", ResponseMessage: err.Error()})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response[any]{ResponseCode: "500", ResponseMessage: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response[any]{ResponseCode: "201", ResponseMessage: "Created", Data: created})
}

// UpdateEmployee menangani request untuk memperbarui data karyawan berdasarkan ID
func (h *EmployeesHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request, id int) {
	var p model.Employees
	json.NewDecoder(r.Body).Decode(&p)

	updated, err := h.Service.UpdateEmployee(id, p)
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response[any]{ResponseCode: "400", ResponseMessage: err.Error()})
			return
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response[any]{ResponseCode: "404", ResponseMessage: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(Response[any]{ResponseCode: "200", ResponseMessage: "Updated", Data: updated})
}

// DeleteEmployee menangani request untuk menghapus data karyawan berdasarkan ID
func (h *EmployeesHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request, id int) {
	err := h.Service.DeleteEmployee(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response[any]{ResponseCode: "404", ResponseMessage: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(Response[any]{ResponseCode: "200", ResponseMessage: "Deleted"})
}
