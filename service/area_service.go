package service

import (
	"context"
	"github.com/simabdi/indonesian-territory/entities"
	"gorm.io/gorm"
)

type AreaService interface {
	GetProvinceByID(ctx context.Context, id int) (*entities.Province, error)
	GetRegencyByID(ctx context.Context, id int) (*entities.Regency, error)
	GetDistrictByID(ctx context.Context, id int) (*entities.District, error)
	GetSubDistrictByID(ctx context.Context, id int) (*entities.SubDistrict, error)
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

func (s *areaService) GetProvinceByID(ctx context.Context, id int) (*entities.Province, error) {
	var province *entities.Province
	err := s.db.WithContext(ctx).First(&province, id).Error
	if err != nil {
		return nil, err
	}
	return province, nil
}

func (s *areaService) GetRegencyByID(ctx context.Context, id int) (*entities.Regency, error) {
	var regency *entities.Regency
	err := s.db.WithContext(ctx).First(&regency, id).Error
	if err != nil {
		return nil, err
	}
	return regency, nil
}

func (s *areaService) GetDistrictByID(ctx context.Context, id int) (*entities.District, error) {
	var district *entities.District
	err := s.db.WithContext(ctx).First(&district, id).Error
	if err != nil {
		return nil, err
	}
	return district, nil
}

func (s *areaService) GetSubDistrictByID(ctx context.Context, id int) (*entities.SubDistrict, error) {
	var subDistrict *entities.SubDistrict
	err := s.db.WithContext(ctx).First(&subDistrict, id).Error
	if err != nil {
		return nil, err
	}
	return subDistrict, nil
}
