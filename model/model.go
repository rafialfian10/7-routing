package model

import (
	"strconv"
	"time"
)

// Create struct, struct berfungsi untuk membuat struktur dari tipe data
type Project struct {
	ID          int
	ProjectName string
	StartDate   time.Time
	EndDate     time.Time
	Desc        string
	Tech        []string
	Image       string
}

// ----------------------------------------------------------------

// Function render time
func (p Project) RenderTime(date time.Time) string { // parameter didapatkan dari pemanggilan funtion di file project-detail.html
	// Buat slice yang akan digunakan untuk format date yang akan diparsing
	Months := [...]string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Des"}

	// strconv.Itoa() berfungsi untuk mengkonversi int menjadi string
	return strconv.Itoa(date.Day()) + " " + Months[date.Month()-1] + " " + strconv.Itoa(date.Year())
}

func (p Project) DurationTime(startDate time.Time, endDate time.Time) string {
	duration := endDate.Sub(startDate).Hours() // Selisih waktu akan dikonversi menjadi jam
	day := 0
	month := 0
	year := 0

	for duration >= 24 {
		day += 1
		duration -= 24
	}

	for day >= 30 {
		month += 1
		day -= 30
	}

	for month >= 12 {
		year += 1
		month -= 12
	}

	if year != 0 && month != 0 {
		return strconv.Itoa(year) + " year, " + strconv.Itoa(month) + " month, " + strconv.Itoa(day) + " day"
	} else if month != 0 {
		return strconv.Itoa(month) + " month, " + strconv.Itoa(day) + " day"
	} else {
		return strconv.Itoa(day) + " Day"
	}
}
