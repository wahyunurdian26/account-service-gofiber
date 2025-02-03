package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SaldoRepository struct {
    db *sqlx.DB
}

func NewSaldoRepository(db *sqlx.DB) *SaldoRepository {
    return &SaldoRepository{db: db}
}

func (r *SaldoRepository) UpdateSaldo(noRekening string, nominal float64) (float64, error) {
    query := `UPDATE saldo SET saldo = saldo + $1 WHERE no_rekening = $2`
    result, err := r.db.Exec(query, nominal, noRekening)
    if err != nil {
        return 0, err
    }

    // Pastikan ada baris yang diperbarui
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }
    if rowsAffected == 0 {
        return 0, fmt.Errorf("no rows affected, rekening not found")
    }

    // Ambil saldo terbaru setelah update
    var saldo float64
    err = r.db.QueryRow(`SELECT saldo FROM saldo WHERE no_rekening = $1`, noRekening).Scan(&saldo)
    if err != nil {
        return 0, err
    }

    return saldo, nil
}

func (r *SaldoRepository) GetSaldo(noRekening string) (float64, error) { // Menggunakan string dan float64
    var saldo float64
    query := `SELECT saldo FROM saldo WHERE no_rekening = $1`
    err := r.db.QueryRow(query, noRekening).Scan(&saldo)
    return saldo, err
}
