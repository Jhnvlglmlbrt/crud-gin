package service

import (
	"log"

	"github.com/Jhnvlglmlbrt/crud-gin/data/request"
	"github.com/Jhnvlglmlbrt/crud-gin/data/response"
	"github.com/Jhnvlglmlbrt/crud-gin/model"
	"github.com/Jhnvlglmlbrt/crud-gin/repository"
	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) *TagsServiceImpl {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}

// Create implements TagsService.
func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) error {
	if err := t.Validate.Struct(tags); err != nil {
		return err
	}

	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)

	return nil
}

// Update implements TagsService.
func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) error {
	if err := t.Validate.Struct(tags); err != nil {
		return err
	}

	tagData, err := t.TagsRepository.FindByID(tags.ID)
	if err != nil {
		return err
	}

	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
	return nil
}

// Delete implements TagsService.
func (t *TagsServiceImpl) Delete(tagsID int) error {
	err := t.TagsRepository.Delete(tagsID)
	if err != nil {
		log.Printf("Ошибка при удалении тега: %v", err)
		return err
	}
	return nil
}

// FindAll implements TagsService.
func (t *TagsServiceImpl) FindAll() ([]response.TagsResponse, error) {
	result, err := t.TagsRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			ID:   value.ID,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

// FindByID implements TagsService.
func (t *TagsServiceImpl) FindByID(tagsID int) (response.TagsResponse, error) {
	tagData, err := t.TagsRepository.FindByID(tagsID)
	if err != nil {
		return response.TagsResponse{}, err
	}

	tagResponse := response.TagsResponse{
		ID:   tagData.ID,
		Name: tagData.Name,
	}

	return tagResponse, nil
}
