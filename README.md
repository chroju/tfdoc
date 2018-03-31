**NOTE: tfdoc is under construction. The contents of the under consideration in this document.**

tfdoc
====

tfdoc is a [Terraform](https://github.com/hashicorp/terraform) helper tool.


Description
----

tfdoc provides the Terraform documents about each resources on your terminal, like `ansible-doc` command. You don't need to check the documents with your web browser any more.


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

Output in the snippet format.

```
$ tfdoc --snippet <resource>
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

The below options are under construction.

```
$ tfdoc --import <resource>
$ tfdoc --list <provider>
$ tfdoc --snippet --only-required <resource>
$ tfdoc --snippet --no-comments <resource>
```


License
----

MIT


Author
------

[chroju](https://chroju.net)
