package user

type UriRequest struct {
	ID uint `uri:"id" binding:"required,numeric,gt=0"`
}
