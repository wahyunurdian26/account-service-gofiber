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
    tx, err := r.db.Begin()
    if err != nil {
        return 0, fmt.Errorf("failed to begin transaction: %w", err)
    }
    defer tx.Rollback() // Jika terjadi error, rollback otomatis

    query := `UPDATE saldo SET saldo = saldo + $1 WHERE no_rekening = $2 RETURNING saldo`
    var saldo float64
    err = tx.QueryRow(query, nominal, noRekening).Scan(&saldo)
    if err != nil {
        return 0, fmt.Errorf("error updating saldo: %w", err)
    }

    if err := tx.Commit(); err != nil {
        return 0, fmt.Errorf("failed to commit transaction: %w", err)
    }

    return saldo, nil
}

func (r *SaldoRepository) InsertSaldo(noRekening string, nominal float64) error {
    query := `INSERT INTO saldo (no_rekening, saldo) VALUES ($1, $2)`
    _, err := r.db.Exec(query, noRekening, nominal)
    if err != nil {
        return fmt.Errorf("Error Insert Saldo: %w", err)
    }
    return nil
}

func (r *SaldoRepository) GetSaldo(noRekening string) (float64, error) { // Menggunakan string dan float64
    var saldo float64
    query := `SELECT saldo FROM saldo WHERE no_rekening = $1`
    err := r.db.QueryRow(query, noRekening).Scan(&saldo)
    return saldo, err
}
