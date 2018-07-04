Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: *empty*

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Maintainers
-----------

This provider plugin is maintained by the Sewwan team at *<insert here a like to Sewan's maintenance team>*

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Usage
---------------------

```
# For example, restrict sewan version in 0.1.x
provider "sewan" {
  version = "~> 0.1"
}
```

Building The Provider
---------------------

*<We assume that gitlab.com/sewan/terraform-provider-sewan will be the public repo for sewan terraform provider*

Clone repository to: `$GOPATH/src/gitlab.com/sewan/terraform-provider-sewan`

```sh
$ mkdir -p $GOPATH/src/gitlab.com/terraform-providers; cd $GOPATH/src/gitlab.com/terraform-providers
$ git clone git@gitlab.com:sewan/terraform-provider-sewan
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/gitlab.com/terraform-providers
$ make build
```

Using the provider
----------------------
## Fill in for each provider

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-sewan
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
