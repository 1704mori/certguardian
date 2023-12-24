package certificates

type Info struct {
	Directories  []string `json:"directories" binding:"required"`
	Certificates []string `json:"certificates"`
}
