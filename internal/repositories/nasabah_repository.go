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
	rand.Seed(time.Now().UnixNano()) // pastikan menggunakan metode terbaru, sesuai versi Go
	prefix := "1"
	noUrut := rand.Int63n(1000000000000) // Angka acak 12 digit
	return fmt.Sprintf("%s%013d", prefix, noUrut) // Gabungkan prefix dan nomor acak menjadi satu nomor rekening
}


// **Daftar Nasabah (dengan generate nomor rekening)**
func (r *NasabahRepository) Create(nasabah *entities.Nasabah) (string, error) {
    var noRekening string
    query := `INSERT INTO nasabah (nama, nik, no_hp, no_rekening) 
              VALUES ($1, $2, $3, $4) RETURNING no_rekening`
    
    noRekening = generateNoRekening() // Memanggil fungsi untuk menghasilkan nomor rekening
    
    err := r.db.QueryRow(query, nasabah.Nama, nasabah.NIK, nasabah.NoHP, noRekening).Scan(&noRekening)
    
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
