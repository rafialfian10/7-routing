package model

import (
	"log"
	"strconv"
	"time"
)

// Create struct, struct berfungsi untuk membuat struktur dari tipe data
type Project struct {
	ProjectName string
	StartDate   string
	EndDate     string
	Desc        string
	Tech        []string
	Image       string
}

// Buat array of object sebagai local storage
var DataProjects = []Project{
	{
		ProjectName: "Dumbways 2022",
		StartDate:   "2022-11-24",
		EndDate:     "2022-12-24",
		Desc:        "Halo Dumbways",
		Tech:        []string{"node", "angular", "react", "typescript"},
		Image:       "public/assets/img/saitama.png",
	},
}

// Function render time
func (p Project) RenderTime(date string) string { // parameter didapatkan dari pemanggilan funtion di file project-detail.html
	// method time.Parse() berfungsi untuk memparsing date. par1: layout format dari waktu yang diparsing, par2: string yang ingin diparsing
	time, err := time.Parse("2006-01-02", date)

	if err != nil {
		log.Fatal(err)
	}

	// Buat slice yang akan digunakan untuk format date yang akan diparsing
	Months := [...]string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Des"}

	// strconv.Itoa() berfungsi untuk mengkonversi int menjadi string
	return strconv.Itoa(time.Day()) + " " + Months[time.Month()-1] + " " + strconv.Itoa(time.Year())
}

func (p Project) DurationTime(startDate string, endDate string) string {
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	duration := end.Sub(start).Hours() // Selisih waktu akan dikonversi menjadi jam
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
