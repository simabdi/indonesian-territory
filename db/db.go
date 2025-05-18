package db

import (
	"gorm.io/gorm"
	"master-indonesian-territory/db/seed"
	"master-indonesian-territory/entities"
)

func InitDB(db *gorm.DB) error {
	status := db.Migrator().HasTable(&entities.Province{})
	if status == false {
		db.AutoMigrate(&entities.Province{}, &entities.Regency{}, &entities.District{}, &entities.SubDistrict{})

		if err := seed.ProvinceSeed(db); err != nil {
			return err
		}
		if err := seed.RegencySeed(db); err != nil {
			return err
		}
		if err := seed.DistrictsSeed(db); err != nil {
			return err
		}
		if err := seed.SubDistrictsSeedBatch1(db); err != nil {
			return err
		}
		if err := seed.SubDistrictsSeedBatch2(db); err != nil {
			return err
		}
		if err := seed.SubDistrictsSeedBatch3(db); err != nil {
			return err
		}
		if err := seed.SubDistrictsSeedBatch4(db); err != nil {
			return err
		}
		if err := seed.SubDistrictsSeedBatch5(db); err != nil {
			return err
		}
		if err := seed.SubDistrictsSeedBatch6(db); err != nil {
			return err
		}
		if err := seed.SubDistrictsSeedBatch7(db); err != nil {
			return err
		}
	}
	return nil
}
