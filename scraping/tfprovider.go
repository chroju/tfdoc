package scraping

import (
	"strings"
)

type TfProvider struct {
	Name         string
	ResourceList []string
}

func (t *TfProvider) Doc(dummy bool) string {
	return strings.Join(t.ResourceList, "\n")
}
