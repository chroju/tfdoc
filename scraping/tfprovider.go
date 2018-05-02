package scraping

type TfProvider struct {
	Name         string
	ResourceList []string
}

func (t *TfProvider) Doc(dummy ...bool) []string {
	return t.ResourceList
}
