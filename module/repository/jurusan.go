package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/paradewisudaitb/Backend/module/entity"
	uuid "github.com/satori/go.uuid"
)

type JurusanRepository struct {
	db *gorm.DB
}

func NewJurusanRepository(db *gorm.DB) JurusanRepository {
	return JurusanRepository{db: db}
}

func (repo JurusanRepository) GetOne(id uuid.UUID) (entity.Jurusan, error) {
	var jurusan entity.Jurusan
	repo.db.First(&jurusan, "id = ?", id)
	if jurusan.ID == "" {
		return jurusan, errors.New("Id jurusan not found")
	}
	return jurusan, nil
}

func (repo JurusanRepository) AddOne(jurusan, fakultas, fakultas_short, jurusan_short string) {
	jurusans := entity.Jurusan{Jurusan: jurusan, Fakultas: fakultas, FakultasShort: fakultas_short, JurusanShort: jurusan_short}
	repo.db.Create(&jurusans)
}

func (repo JurusanRepository) UpdateOne(id uuid.UUID, jurusan, fakultas, fakultas_short, jurusan_short string) error {
	var jurusans entity.Jurusan
	jurusan_update := map[string]interface{}{}
	if jurusan != "" {
		jurusan_update["jurusan"] = jurusan
	}
	if fakultas != "" {
		jurusan_update["fakultas"] = fakultas
	}
	if jurusan_short != "" {
		jurusan_update["jurusan_short"] = jurusan_short
	}
	if fakultas_short != "" {
		jurusan_update["fakultas_short"] = fakultas_short
	}
	repo.db.First(&jurusans, "id = ?", id.String())
	if (jurusans == entity.Jurusan{}) {
		return errors.New("Jurusan not found.")
	}
	repo.db.Model(&jurusans).Updates(jurusan_update)
	return nil
}

func (repo JurusanRepository) DeleteOne(id uuid.UUID) error {
	var jurusans entity.Jurusan
	repo.db.First(&jurusans, "id = ?", id)
	if jurusans.ID == "" {
		return errors.New("Id jurusan not found")
	}
	repo.db.Delete(&jurusans)
	return nil
}
