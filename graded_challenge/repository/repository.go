package repository

import (
	"database/sql"
	"errors"
	"graded-challenge/model"
)

// EmployeesRepository bertanggung jawab melakukan operasi database terkait tabel employees
type EmployeesRepository struct {
	DB *sql.DB // Koneksi ke database
}

// FindAll mengambil semua data karyawan (ID, nama, dan email)
func (r *EmployeesRepository) FindAll() ([]model.Employees, error) {
	rows, err := r.DB.Query("SELECT id, name, email FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employees
	for rows.Next() {
		var e model.Employees
		if err := rows.Scan(&e.ID, &e.Name, &e.Email); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, nil
}

// FindByID mengambil data karyawan berdasarkan ID (termasuk phone)
func (r *EmployeesRepository) FindByID(id int) (*model.Employees, error) {
	var e model.Employees
	err := r.DB.QueryRow("SELECT id, name, email, phone FROM employees WHERE id=?", id).
		Scan(&e.ID, &e.Name, &e.Email, &e.Phone)
	return &e, err
}

// Create menyimpan data karyawan baru ke database dan mengembalikan ID-nya
func (r *EmployeesRepository) Create(e model.Employees) (int64, error) {
	stmt, err := r.DB.Prepare("INSERT INTO employees (name, email, phone) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Email, e.Phone)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// Update memperbarui data karyawan berdasarkan ID
func (r *EmployeesRepository) Update(id int, e model.Employees) error {
	stmt, err := r.DB.Prepare("UPDATE employees SET name=?, email=?, phone=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Email, e.Phone, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

// Delete menghapus data karyawan berdasarkan ID
func (r *EmployeesRepository) Delete(id int) error {
	stmt, err := r.DB.Prepare("DELETE FROM employees WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("no rows deleted")
	}
	return nil
}
