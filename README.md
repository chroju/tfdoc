**NOTE: tfdoc is under construction. The contents of the under consideration in this document.**

tfdoc
====

tfdoc is a [Terraform](https://github.com/hashicorp/terraform) helper tool.


Description
----

tfdoc provides the Terraform documents about each resources on your terminal, like `ansible-doc` command. You don't need to check the documents with your web browser any more.


Usage
----

```
$ tfdoc <resource>
$ tfdoc --url <resource>
$ tfdoc --import <resource>
$ tfdoc --list <provider>
```

### tfdoc snippet

```
$ tfdoc snippet <resource>
$ tfdoc snippet --only-required <resource>
$ tfdoc snippet --no-comments <resource>
```

License
----

MIT
