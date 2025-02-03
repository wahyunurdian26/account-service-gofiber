package controllers

import (
	"fmt"
	"service-account/internal/entities"
	"service-account/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

type NasabahController struct {
	nasabahUseCase *usecases.NasabahUseCase
}

func NewNasabahController(nasabahUseCase *usecases.NasabahUseCase) *NasabahController {
	return &NasabahController{nasabahUseCase: nasabahUseCase}
}

// DaftarNasabah - Registrasi Nasabah Baru
func (c *NasabahController) DaftarNasabah(ctx *fiber.Ctx) error {
    var nasabah entities.Nasabah
    if err := ctx.BodyParser(&nasabah); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"remark": "Invalid request payload"})
    }

    noRekening, err := c.nasabahUseCase.DaftarNasabah(&nasabah)
    if err != nil {
        fmt.Println("Error saat mendaftar nasabah:", err)
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"remark": err.Error()})
    }

    // Pastikan noRekening bukan kosong
    if noRekening == "" {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"remark": "NIK or No HP already exists"})
    }

    return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"no_rekening": noRekening})
}



// **Tabung - Deposit ke Rekening**
func (c *NasabahController) Tabung(ctx *fiber.Ctx) error {
	var request struct {
		NoRekening string  `json:"no_rekening"`
		Jumlah     float64 `json:"jumlah"`
	}

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"remark": "Invalid request payload"})
	}

	if err := c.nasabahUseCase.Tabung(request.NoRekening, request.Jumlah); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"remark": "Internal server error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"remark": "Deposit successful"})
}

// **Tarik - Penarikan Dana dari Rekening**
func (c *NasabahController) Tarik(ctx *fiber.Ctx) error {
	var request struct {
		NoRekening string  `json:"no_rekening"`
		Jumlah     float64 `json:"jumlah"`
	}

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"remark": "Invalid request payload"})
	}

	if err := c.nasabahUseCase.Tarik(request.NoRekening, request.Jumlah); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"remark": "Internal server error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"remark": "Withdrawal successful"})
}

// **CekSaldo - Melihat Saldo Nasabah**
func (c *NasabahController) CekSaldo(ctx *fiber.Ctx) error {
	noRekening := ctx.Params("no_rekening")
	saldo, err := c.nasabahUseCase.CekSaldo(noRekening)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"remark": "Internal server error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"no_rekening": noRekening, "saldo": saldo})
}
