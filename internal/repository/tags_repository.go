package repository

import "github.com/Jhnvlglmlbrt/crud-gin/model"

type TagsRepository interface {
	Save(tags model.Tags) error
	Update(tags model.Tags) error
	Delete(tagsID int) error
	FindByID(tagsID int) (tags model.Tags, err error)
	FindAll() ([]model.Tags, error)
}
