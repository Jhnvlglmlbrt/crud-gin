package request

type UpdateTagsRequest struct {
	ID   int    `validate:"required"`
	Name string `validate:"required,min=1,max=200" json:"name"`
}
