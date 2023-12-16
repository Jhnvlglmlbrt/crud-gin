package service

import (
	"github.com/Jhnvlglmlbrt/crud-gin/data/request"
	"github.com/Jhnvlglmlbrt/crud-gin/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest) error
	Update(tags request.UpdateTagsRequest) error
	Delete(tagsID int) error
	FindByID(tagsID int) (response.TagsResponse, error)
	FindAll() ([]response.TagsResponse, error)
}
