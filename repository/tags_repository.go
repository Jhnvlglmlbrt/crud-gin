package repository

import "github.com/Jhnvlglmlbrt/crud-gin/model"

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(tagsID int)
	FindByID(tagsID int) (tags model.Tags, err error)
	FindAll() []model.Tags
}
