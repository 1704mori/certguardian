package certificates

import "github.com/1704mori/certguardian/internal/domain"

// type DirectoryInfo map[string][]string
type DirectoryInfo map[string]map[string]domain.Info

type Info struct {
	Directories DirectoryInfo `json:"directories" binding:"required"`
}
