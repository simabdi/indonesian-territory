package service

import (
	"context"
	"github.com/simabdi/indonesian-territory/entities"
	"gorm.io/gorm"
)

type AreaService interface {
	GetAllProvince(ctx context.Context) ([]*entities.Province, error)
	GetRegencyByProvinceID(ctx context.Context, provinceID int) ([]*entities.Regency, error)
	GetDistrictByRegencyID(ctx context.Context, regencyID int) ([]*entities.District, error)
	GetSubDistrictByDistrictID(ctx context.Context, districtID int) ([]*entities.SubDistrict, error)
}

type areaService struct {
	db *gorm.DB
}

func NewAreaService(db *gorm.DB) AreaService {
	return &areaService{db: db}
}

func (s *areaService) GetAllProvince(ctx context.Context) ([]*entities.Province, error) {
	var province []*entities.Province
	err := s.db.WithContext(ctx).Find(&province).Error
	if err != nil {
		return nil, err
	}
	return province, nil
}

func (s *areaService) GetRegencyByProvinceID(ctx context.Context, provinceID int) ([]*entities.Regency, error) {
	var regency []*entities.Regency
	err := s.db.WithContext(ctx).Where("province_id = ?", provinceID).Find(&regency).Error
	if err != nil {
		return nil, err
	}
	return regency, nil
}

func (s *areaService) GetDistrictByRegencyID(ctx context.Context, regencyID int) ([]*entities.District, error) {
	var district []*entities.District
	err := s.db.WithContext(ctx).Where("regency_id = ?", regencyID).Find(&district).Error
	if err != nil {
		return nil, err
	}
	return district, nil
}

func (s *areaService) GetSubDistrictByDistrictID(ctx context.Context, districtID int) ([]*entities.SubDistrict, error) {
	var subDistrict []*entities.SubDistrict
	err := s.db.WithContext(ctx).Where("district_id = ?", districtID).Find(&subDistrict).Error
	if err != nil {
		return nil, err
	}
	return subDistrict, nil
}
