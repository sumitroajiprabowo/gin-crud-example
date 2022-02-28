package web

type PackageUpdateRequest struct {
	Id   int64  `validate:"required" json:"id"`
	Name string `validate:"required,min=1,max=100" json:"name"`
}
