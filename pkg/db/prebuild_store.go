// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package db

import (
	. "github.com/daytonaio/daytona/pkg/db/dto"
	"github.com/daytonaio/daytona/pkg/prebuild"
	"gorm.io/gorm"
)

type PrebuildStore struct {
	db *gorm.DB
}

func NewPrebuildStore(db *gorm.DB) (*PrebuildStore, error) {
	err := db.AutoMigrate(&PrebuildConfigDTO{})
	if err != nil {
		return nil, err
	}

	return &PrebuildStore{db: db}, nil
}

func (p *PrebuildStore) Find(id string) (*prebuild.PrebuildConfig, error) {
	prebuildConfigDTO := PrebuildConfigDTO{}
	tx := p.db.Where("id = ?", id).First(&prebuildConfigDTO)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, prebuild.ErrPrebuildNotFound
		}
		return nil, tx.Error
	}

	prebuildConfig := ToPrebuildConfig(prebuildConfigDTO)

	return prebuildConfig, nil
}

// TODO: Upsert should be implemented
func (p *PrebuildStore) Upsert(prebuildConfig *prebuild.PrebuildConfig) error {
	prebuildConfigDTO := ToPrebuildConfigDTO(prebuildConfig)
	tx := p.db.Save(&prebuildConfigDTO)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (p *PrebuildStore) Delete(prebuildConfig *prebuild.PrebuildConfig) error {
	tx := p.db.Where("id = ?", prebuildConfig.Id).Delete(&PrebuildConfigDTO{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return prebuild.ErrPrebuildNotFound
	}

	return nil
}
