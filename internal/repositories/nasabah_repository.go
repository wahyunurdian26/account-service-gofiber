package repositories

import (
	"fmt"
	"math/rand"
	"service-account/internal/entities"
	"time"

	"github.com/jmoiron/sqlx"
)

type NasabahRepository struct {
	db *sqlx.DB
}

func NewNasabahRepository(db *sqlx.DB) *NasabahRepository {
	return &NasabahRepository{db: db}
}

// **Fungsi untuk generate nomor rekening**
func generateNoRekening() string {
    // Gunakan rand.New dengan sumber acak yang didasarkan pada waktu
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    // Prefix untuk jenis rekening, misalnya "1" untuk Tabungan
    prefix := "1"
    
    // Generate nomor urut acak (12 digit) antara 100000000000 dan 999999999999
    noUrut := r.Int63n(900000000000) + 100000000000 // Angka antara 100000000000 dan 999999999999
    
    // Gabungkan prefix dan nomor urut menjadi satu nomor rekening dengan format yang benar
    return fmt.Sprintf("%s%012d", prefix, noUrut) // %012d memastikan panjang nomor rekening 12 digit
}

// **Daftar Nasabah (dengan generate nomor rekening)**
func (r *NasabahRepository) Create(nasabah *entities.Nasabah) (string, error) {
    var noRekening string
    query := `INSERT INTO nasabah (nama, nik, no_hp) VALUES ($1, $2, $3) RETURNING no_rekening`
    err := r.db.QueryRow(query, nasabah.Nama, nasabah.NIK, nasabah.NoHP).Scan(&noRekening)
    
    if err != nil {
        fmt.Println("Error saat INSERT nasabah:", err)
        return "", fmt.Errorf("error saat insert nasabah: %w", err) 
    }
    return noRekening, nil
}




func (r *NasabahRepository) FindByNIKOrNoHP(nik, noHP string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM nasabah WHERE nik = $1 OR no_hp = $2`
	err := r.db.QueryRow(query, nik, noHP).Scan(&count)


    if err != nil {
        fmt.Println("Error saat mencari NIK atau No HP:", err)
        return false, fmt.Errorf("error saat mencari NIK atau No HP: %w", err)
    }
	return count > 0, nil
}
