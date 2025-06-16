package service

import (
	"errors"
	"graded-challenge/model"
	"graded-challenge/repository"
)

// Custom error untuk validasi input
var ErrValidation = errors.New("name & email wajib diisi")

// EmployeeService menyediakan logika bisnis terkait entity karyawan
type EmployeeService struct {
	// Repository untuk akses data karyawan
	repo *repository.EmployeesRepository
}

// NewEmployeeService mengembalikan instance baru dari EmployeeService
func NewEmployeeService(repo *repository.EmployeesRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

// GetAllEmployees mengambil semua data karyawan dari repository
func (s *EmployeeService) GetAllEmployees() ([]model.Employees, error) {
	return s.repo.FindAll()
}

// GetEmployeeByID mengambil satu data karyawan berdasarkan ID
func (s *EmployeeService) GetEmployeeByID(id int) (*model.Employees, error) {
	return s.repo.FindByID(id)
}

// CreateEmployee membuat data karyawan baru setelah validasi input
func (s *EmployeeService) CreateEmployee(e model.Employees) (model.Employees, error) {
	if e.Name == "" || e.Email == "" || e.Phone == "" {
		return model.Employees{}, ErrValidation
	}

	id, err := s.repo.Create(e)
	if err != nil {
		return model.Employees{}, err
	}
	e.ID = int(id)
	return e, nil
}

// UpdateEmployee memperbarui data karyawan berdasarkan ID setelah validasi
func (s *EmployeeService) UpdateEmployee(id int, e model.Employees) (model.Employees, error) {
	if e.Name == "" || e.Email == "" || e.Phone == "" {
		return model.Employees{}, ErrValidation
	}

	err := s.repo.Update(id, e)
	if err != nil {
		return model.Employees{}, err
	}
	e.ID = id
	return e, nil
}

// DeleteEmployee menghapus data karyawan berdasarkan ID
func (s *EmployeeService) DeleteEmployee(id int) error {
	return s.repo.Delete(id)
}
