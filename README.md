tfdoc
====

[![Coverage Status](https://coveralls.io/repos/github/chroju/tfdoc/badge.svg?branch=master)](https://coveralls.io/github/chroju/tfdoc?branch=test-ci) [![Go Report Card](https://goreportcard.com/badge/github.com/chroju/tfdoc)](https://goreportcard.com/report/github.com/chroju/tfdoc) [![CircleCI](https://circleci.com/gh/chroju/tfdoc/tree/master.svg?style=shield)](circleci.com/gh/chroju/tfdoc/tree/master)

tfdoc will help you write [Terraform](https://github.com/hashicorp/terraform) files (.tf) .


Description
----

tfdoc provides the Terraform documents about each resources on your terminal, like `ansible-doc` command. You don't need to check the documents with your web browser any more.


Install
----

Using `go get` or download binaries from [Releases](https://github.com/chroju/tfdoc/releases).

```bash
$ go get github.com/chroju/tfdoc
```


Usage
----

Output Terraform documents like this.

```
$ tfdoc aws_instance
aws_instance

Provides an EC2 instance resource. This allows instances to be created, updated,and deleted. Instances also support provisioning.


Argument Reference (= is mandatory):

= ami
    (Required) The AMI to use for the instance.

- availability_zone
    (Optional) The AZ to start the instance in.

- placement_group
    (Optional) The Placement Group to start the instance in.

...
```

There are some options to change output format.

### --url, -u

Output only Terraform document URL.

```
$ tfdoc --url aws_instance
https://www.terraform.io/docs/providers/aws/r/instance.html
```

### --snippet, -s

Output in the snippet format like this.

```
$ tfdoc --snippet aws_instance
resource "aws_instance" "sample" {

  // (Required) The AMI to use for the instance.
  ami = ""

  // (Optional) The AZ to start the instance in.
  availability_zone = ""

  // (Optional) The Placement Group to start the instance in.
  placement_group = ""

...
}
```

When using `--snippet` option, you can control output format more finely with some options. Using `--only-required` make a snippet with only required arguments, and `--no-comments` option eliminate all comments.

### --list, -l

List available resources with given provider.

```
$ tfdoc -l azurerm
azurerm_resource_group
azurerm_app_service
azurerm_app_service_plan
azurerm_app_service_active_slot
azurerm_app_service_custom_hostname_binding
azurerm_app_service_slot
azurerm_function_app
azurerm_role_assignment
```


License
----

MIT


Author
------

[chroju](https://chroju.net)
