package scraping

// TfProvider is terraform provider object such as aws, azurerm, google
type TfProvider struct {
	Name         string
	ResourceList []string
}

// Doc return terraform document
func (t *TfProvider) Doc(dummy ...bool) []string {
	return t.ResourceList
}
