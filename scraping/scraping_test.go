package scraping

import (
	"strings"
	"testing"
)

func TestTfResourceConvertURL(t *testing.T) {
	var cases = []struct {
		resourceName string
		expectedURL  string
	}{
		{"aws_instance", "https://www.terraform.io/docs/providers/aws/r/instance.html"},
		{"aws_lb_listener", "https://www.terraform.io/docs/providers/aws/r/lb_listener.html"},
		{"azurerm_container_group", "https://www.terraform.io/docs/providers/azurerm/r/container_group.html"},
		{"grafana_dashboard", "https://www.terraform.io/docs/providers/grafana/r/dashboard.html"},
	}

	for _, c := range cases {
		scraper, _ := NewScraper("resource", c.resourceName)
		if !strings.Contains(scraper.Url, c.expectedURL) {
			t.Error()
		}
	}
}

func TestTfProviderConvertURL(t *testing.T) {
	var cases = []struct {
		providerName string
		expectedURL  string
	}{
		{"aws", "https://www.terraform.io/docs/providers/aws/index.html"},
		{"grafana", "https://www.terraform.io/docs/providers/grafana/index.html"},
		{"terraform-enterprise", "https://www.terraform.io/docs/providers/terraform-enterprise/index.html"},
		{"nsxt", "https://www.terraform.io/docs/providers/nsxt/index.html"},
	}

	for _, c := range cases {
		scraper, _ := NewScraper("provider", c.providerName)
		if !strings.Contains(scraper.Url, c.expectedURL) {
			t.Errorf("expect: %s \n result: %s", c.expectedURL, scraper.Url)
		}
	}
}
