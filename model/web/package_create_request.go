package web

type PackageCreateRequest struct {
	Name string `binding:"required" json:"name"`
}
