Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)


<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Maintainers
-----------

This provider plugin is maintained by the Sewan's team.

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

* Set up [Go](http://www.golang.org) your dev environment (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

* Clone Sewan's sdk and terraform plugin repositories to: `$GOPATH/src/github.com/SewanDevs/`

```sh
$ mkdir -p $GOPATH/src/github.com/SewanDevs/
$ cd $GOPATH/src/github.com/SewanDevs/
$ git clone git@github.com:SewanDevs/sewan_go_sdk.git
$ git clone git@github.com:SewanDevs/terraform-provider-sewan.git
```
* **Optional steps :**

To run unit tests of the provider ans sdk, run `make test`.

```sh
$ cd $GOPATH/src/github.com/SewanDevs/terraform-provider-sewan.git
$ make test
```

* Provider compilation

This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ cd $GOPATH/src/github.com/SewanDevs/terraform-provider-sewan
$ make build
```

Use the docker image for Sewan's plugin
---------------------------
*< available soon>*
