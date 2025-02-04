package usecases

import (
	"errors"
	"fmt"
	"service-account/internal/entities"
	"service-account/internal/repositories"
)

type NasabahUseCase struct {
	nasabahRepo *repositories.NasabahRepository
	saldoRepo   *repositories.SaldoRepository
}

func NewNasabahUseCase(nasabahRepo *repositories.NasabahRepository, saldoRepo *repositories.SaldoRepository) *NasabahUseCase {
	return &NasabahUseCase{nasabahRepo: nasabahRepo, saldoRepo: saldoRepo}
}

func (uc *NasabahUseCase) DaftarNasabah(nasabah *entities.Nasabah) (string, error) {
	exists, err := uc.nasabahRepo.FindByNIKOrNoHP(nasabah.NIK, nasabah.NoHP)
	if err != nil {
		return "", fmt.Errorf("Error Find By NIK : %w", err )
	}
	if exists {
		return "", nil // NIK atau No HP sudah ada
	}

	noRekening, err := uc.nasabahRepo.Create(nasabah)
	if err != nil {
		return "", fmt.Errorf("Error Create Nasabah : %w", err )
	}

	// **Tambahkan Insert ke tabel saldo terlebih dahulu**
	err = uc.saldoRepo.InsertSaldo(noRekening, 0)
	if err != nil {
		return "", fmt.Errorf("Error Insert Saldo : %w", err)
	}

	// Inisialisasi saldo
	_, err = uc.saldoRepo.UpdateSaldo(noRekening, 0)
	if err != nil {
		return "", fmt.Errorf("Error Update Saldo : %w", err)
	}

	return noRekening, nil
}



// Tabung - Deposit funds into the account
func (uc *NasabahUseCase) Tabung(noRekening string, jumlah float64) error {
    if jumlah <= 0 {
        return errors.New("deposit amount must be greater than zero")
    }

    _, err := uc.saldoRepo.UpdateSaldo(noRekening, jumlah)
    return err
}

func (uc *NasabahUseCase) Tarik(noRekening string, jumlah float64) error {
    if jumlah <= 0 {
        return errors.New("withdraw amount must be greater than zero")
    }

    saldo, err := uc.saldoRepo.GetSaldo(noRekening)
    if err != nil {
        return err
    }

    if saldo < jumlah {
        return errors.New("insufficient balance")
    }

    _, err = uc.saldoRepo.UpdateSaldo(noRekening, -jumlah)
    return err
}



// CekSaldo - Check the current balance of the account
func (uc *NasabahUseCase) CekSaldo(noRekening string) (float64, error) {
    saldo, err := uc.saldoRepo.GetSaldo(noRekening)
    if err != nil {
        return 0, err
    }

    return saldo, nil
}