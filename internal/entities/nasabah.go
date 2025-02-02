package entities

type Nasabah struct {
	NoRekening int    `json:"no_rekening" db:"no_rekening"`
	Nama       string `json:"nama" db:"nama"`
	NIK        string `json:"nik" db:"nik"`
	NoHP       string `json:"no_hp" db:"no_hp"`
}

type Saldo struct {
	NoRekening int `json:"no_rekening" db:"no_rekening"`
	Saldo      int `json:"saldo" db:"saldo"`
}