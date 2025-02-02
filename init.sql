CREATE TABLE nasabah (
    no_rekening VARCHAR(20) PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    nik VARCHAR(16) UNIQUE NOT NULL,
    no_hp VARCHAR(15) UNIQUE NOT NULL
);

CREATE TABLE saldo (
    no_rekening VARCHAR(20) PRIMARY KEY REFERENCES nasabah(no_rekening),
    saldo BIGINT NOT NULL DEFAULT 0
);