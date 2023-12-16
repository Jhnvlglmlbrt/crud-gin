package controller

import (
	"net/http"
	"strconv"

	"github.com/Jhnvlglmlbrt/crud-gin/data/request"
	"github.com/Jhnvlglmlbrt/crud-gin/service"
	"github.com/Jhnvlglmlbrt/crud-gin/utils"
	"github.com/gin-gonic/gin"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

// Create controller
func (controller *TagsController) Create(ctx *gin.Context) error {
	var createTagRequest request.CreateTagsRequest
	if err := ctx.ShouldBindJSON(&createTagRequest); err != nil {
		return err
	}

	controller.tagsService.Create(createTagRequest)
	utils.RespondWithJson(ctx, http.StatusOK, nil)
	return nil
}

func (controller *TagsController) Update(ctx *gin.Context) error {
	tagID, err := strconv.Atoi(ctx.Param("tagsID"))
	if err != nil {
		return err
	}

	var updateTagRequest request.UpdateTagsRequest
	if err := ctx.ShouldBindJSON(&updateTagRequest); err != nil {
		return err
	}

	updateTagRequest.ID = tagID
	controller.tagsService.Update(updateTagRequest)

	utils.RespondWithJson(ctx, http.StatusOK, nil)
	return nil
}

// Delete controller
func (controller *TagsController) Delete(ctx *gin.Context) error {
	tagID, err := strconv.Atoi(ctx.Param("tagsID"))
	if err != nil {
		return err
	}

	controller.tagsService.Delete(tagID)

	utils.RespondWithJson(ctx, http.StatusOK, nil)
	return nil
}

// FindByID controller
func (controller *TagsController) FindByID(ctx *gin.Context) error {
	tagID, err := strconv.Atoi(ctx.Param("tagsID"))
	if err != nil {
		return err
	}

	tagResponse, err := controller.tagsService.FindByID(tagID)
	if err != nil {
		return err
	}

	utils.RespondWithJson(ctx, http.StatusOK, tagResponse)
	return nil
}

// FindAll controller
func (controller *TagsController) FindAll(ctx *gin.Context) error {
	tagResponse, err := controller.tagsService.FindAll()
	if err != nil {
		return err
	}

	utils.RespondWithJson(ctx, http.StatusOK, tagResponse)

	return nil
}
