package entities

type (
	Province struct {
		ID   uint   `gorm:"primaryKey"`
		Name string `gorm:"size:100;not null"`
	}

	Regency struct {
		ID         uint   `gorm:"primaryKey"`
		Name       string `gorm:"size:100;not null"`
		ProvinceID uint
	}

	District struct {
		ID        uint   `gorm:"primaryKey"`
		Name      string `gorm:"size:100;not null"`
		RegencyID uint
		Code      string `gorm:"size:20"`
		MapLat    string `gorm:"size:100"`
		MapLng    string `gorm:"size:100"`
	}

	SubDistrict struct {
		ID         uint   `gorm:"primaryKey"`
		Name       string `gorm:"size:100;not null"`
		DistrictID uint
		Code       string  `gorm:"size:30"`
		MapLat     *string `gorm:"size:50"`
		MapLng     *string `gorm:"size:50"`
	}
)
