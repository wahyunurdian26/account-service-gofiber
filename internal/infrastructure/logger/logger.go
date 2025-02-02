package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log adalah instance dari logrus.Logger yang akan digunakan di seluruh aplikasi.
var Log *logrus.Logger

// InitLogger menginisialisasi logger dengan konfigurasi default.
func InitLogger() {
	Log = logrus.New()
	Log.SetOutput(os.Stdout) // Output log ke console
	Log.SetFormatter(&logrus.JSONFormatter{}) // Format log sebagai JSON
	Log.SetLevel(logrus.InfoLevel) // Set level log ke INFO
}