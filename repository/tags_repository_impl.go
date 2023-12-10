package repository

import (
	"errors"

	"github.com/Jhnvlglmlbrt/crud-gin/data/request"
	"github.com/Jhnvlglmlbrt/crud-gin/model"
	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	DB *gorm.DB
}

func NewTagsRepositoryImpl(db *gorm.DB) *TagsRepositoryImpl {
	return &TagsRepositoryImpl{DB: db}
}

// Delete implements TagsRepository.
func (t *TagsRepositoryImpl) Delete(tagsID int) error {
	var tags model.Tags
	result := t.DB.Where("id = ?", tagsID).Delete(&tags)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("tag is not found")
		}
		return result.Error
	}

	return nil
}

// FindAll implements TagsRepository.
func (t *TagsRepositoryImpl) FindAll() ([]model.Tags, error) {
	var tags []model.Tags
	result := t.DB.Find(&tags)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("tags not found")
		}
		return nil, result.Error
	}

	return tags, nil
}

// FindByID implements TagsRepository.
func (t *TagsRepositoryImpl) FindByID(tagsID int) (tags model.Tags, err error) {
	var tag model.Tags
	result := t.DB.Find(&tag, tagsID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return tag, errors.New("tag is not found")
		}
		return tag, result.Error
	}

	return tag, nil
}

// Save implements TagsRepository.
func (t *TagsRepositoryImpl) Save(tags model.Tags) error {
	if result := t.DB.Create(&tags); result.Error != nil {
		return result.Error
	}
	return nil
}

// Update implements TagsRepository.
func (t *TagsRepositoryImpl) Update(tags model.Tags) error {
	var updateTag = request.UpdateTagsRequest{
		ID:   tags.ID,
		Name: tags.Name,
	}

	if result := t.DB.Model(&tags).Updates(updateTag); result.Error != nil {
		return result.Error
	}
	return nil
}
