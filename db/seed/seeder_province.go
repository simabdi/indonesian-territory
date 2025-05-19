package seed

import (
	"github.com/simabdi/indonesian-territory/entities"
	"gorm.io/gorm"
)

func ProvinceSeed(DB *gorm.DB) error {
	province := []*entities.Province{
		{ID: 1, Name: "BALI"},
		{ID: 2, Name: "BANGKA BELITUNG"},
		{ID: 3, Name: "BANTEN"},
		{ID: 4, Name: "BENGKULU"},
		{ID: 5, Name: "D.I. YOGYAKARTA"},
		{ID: 6, Name: "DKI JAKARTA"},
		{ID: 7, Name: "GORONTALO"},
		{ID: 8, Name: "JAMBI"},
		{ID: 9, Name: "JAWA BARAT"},
		{ID: 10, Name: "JAWA TENGAH"},
		{ID: 11, Name: "JAWA TIMUR"},
		{ID: 12, Name: "KALIMANTAN BARAT"},
		{ID: 13, Name: "KALIMANTAN SELATAN"},
		{ID: 14, Name: "KALIMANTAN TENGAH"},
		{ID: 15, Name: "KALIMANTAN TIMUR"},
		{ID: 16, Name: "KEPULAUAN RIAU"},
		{ID: 17, Name: "LAMPUNG"},
		{ID: 18, Name: "MALUKU"},
		{ID: 19, Name: "MALUKU UTARA"},
		{ID: 20, Name: "NANGGROE ACEH DARUSSALAM"},
		{ID: 21, Name: "NUSA TENGGARA BARAT"},
		{ID: 22, Name: "NUSA TENGGARA TIMUR"},
		{ID: 23, Name: "PAPUA"},
		{ID: 24, Name: "RIAU"},
		{ID: 25, Name: "SULAWESI SELATAN"},
		{ID: 26, Name: "SULAWESI TENGAH"},
		{ID: 27, Name: "SULAWESI TENGGARA"},
		{ID: 28, Name: "SULAWESI UTARA"},
		{ID: 29, Name: "SUMATERA BARAT"},
		{ID: 30, Name: "SUMATERA SELATAN"},
		{ID: 31, Name: "SUMATERA UTARA"},
		{ID: 32, Name: "PAPUA BARAT"},
		{ID: 33, Name: "SULAWESI BARAT"},
		{ID: 34, Name: "KALIMANTAN UTARA"},
		{ID: 35, Name: "PAPUA BARAT DAYA"},
		{ID: 36, Name: "PAPUA TENGAH"},
		{ID: 37, Name: "PAPUA PEGUNUNGAN"},
		{ID: 38, Name: "PAPUA SELATAN"},
	}

	return DB.Create(&province).Error
}
