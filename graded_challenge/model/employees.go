package model

// Employees merepresentasikan data entitas karyawan
type Employees struct {
	ID    int    `json:"id"`              // ID unik karyawan
	Name  string `json:"name"`            // Nama lengkap karyawan
	Email string `json:"email"`           // Alamat email karyawan
	Phone string `json:"phone,omitempty"` // Nomor telepon karyawan (bisa kosong saat dikembalikan dalam JSON)
}
