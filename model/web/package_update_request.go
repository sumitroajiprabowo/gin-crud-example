package web

type PackageUpdateRequest struct {
	Id   int64  `validate:"required"`
	Name string `binding:"required" json:"name"`
}
