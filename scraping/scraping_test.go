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
		if !strings.Contains(scraper.URL, c.expectedURL) {
			t.Error()
		}
	}
}

func TestConvertURLError(t *testing.T) {
	var cases = []struct {
		docType      string
		resourceName string
	}{
		{"resource", "awsinstance"},
		{"errorDocType", "aws_lb_listener"},
		{"resource", "aws_lb_listener_error"},
	}

	for _, c := range cases {
		scraper, err := NewScraper(c.docType, c.resourceName)
		if err == nil {
			t.Errorf("url: %s, docType: %s, resourceName: %s", scraper.URL, c.docType, c.resourceName)
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
		if !strings.Contains(scraper.URL, c.expectedURL) {
			t.Errorf("expect: %s \n result: %s", c.expectedURL, scraper.URL)
		}
	}
}

func TestGetTfResourceDoc(t *testing.T) {
	var cases = []struct {
		resourceName   string
		expectedString string
	}{
		{"aws_lb_listener", "Provides a Load Balancer Listener resource."},
		{"azurerm_container_group", "Create as an Azure Container Group instance."},
		{"grafana_dashboard", "The dashboard resource allows a dashboard to be created on a Grafana server."},
		{"bitbucket_hook", "Provides a Bitbucket hook resource."},
		{"pagerduty_service_integration", "The ID of the service the integration should belong to."},
		{"vault_okta_auth_backend_group", "(Optional) Vault policies to associate with this group"},
		{"kubernetes_resource_quota", "(Required) Standard resource quota's metadata."},
	}

	for _, c := range cases {
		scraper, _ := NewScraper("resource", c.resourceName)
		resource, _ := scraper.Scrape()
		if !strings.Contains(strings.Join(resource.Doc(false, false, false), " "), c.expectedString) {
			t.Errorf("expect: %s \n result: %s", c.expectedString, scraper.Name)
		}
	}
}

func TestGetTfResourceDocWithOptions(t *testing.T) {
	noCommentSnippet := `
  load_balancer_arn                          = ""
  port                                       = ""
  protocol                                   = ""
  ssl_policy                                 = ""`
	requiredOnlySnippet := `
  // (Required) The port on which the load balancer is listening.
  port = ""

  // (Required) An Action block. Action blocks are documented below.
  default_action {`
	noCommentAndRequiredOnlySnippet := `
  load_balancer_arn = ""
  port              = ""

  default_action {`

	var cases = []struct {
		resourceName    string
		expectedString  string
		isSnippet       bool
		needlessComment bool
		requiredOnly    bool
	}{
		{"aws_lb_listener", "resource \"aws_lb_listener\"", true, false, false},
		{"aws_lb_listener", noCommentSnippet, true, true, false},
		{"aws_lb_listener", requiredOnlySnippet, true, false, true},
		{"aws_lb_listener", noCommentAndRequiredOnlySnippet, true, true, true},
	}

	for _, c := range cases {
		scraper, _ := NewScraper("resource", c.resourceName)
		resource, _ := scraper.Scrape()
		snippet := strings.Join(resource.Doc(c.isSnippet, c.needlessComment, c.requiredOnly), "\n")
		if !strings.Contains(snippet, c.expectedString) {
			t.Errorf("expect: %s \n result: %s", c.expectedString, snippet)
		}
	}
}

func TestGetTfProviderDoc(t *testing.T) {
	var cases = []struct {
		providerName   string
		expectedString string
	}{
		{"aws", "aws_main_route_table_association"},
		{"grafana", "grafana_data_source"},
		// TODO: fix newxt raws test bug
		// {"profitbricks", "profitbricks_ipfailover"},
		{"nsxt", "nsxt_logical_router_link_port_on_tier1"},
		{"datadog", "datadog_monitor"},
	}

	for _, c := range cases {
		scraper, _ := NewScraper("provider", c.providerName)
		provider, _ := scraper.Scrape()
		if !strings.Contains(strings.Join(provider.Doc(), " "), c.expectedString) {
			t.Errorf("expect: %s \n result: %s", c.expectedString, scraper.Name)
		}
	}
}
