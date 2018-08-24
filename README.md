Sewan Terraform Provider
========================

- Webistes : https://www.sewan.fr, https://www.terraform.io
- Travis build : [![Build Status](https://travis-ci.com/SewanDevs/terraform-provider-sewan.svg?branch=github_release)](https://travis-ci.com/SewanDevs/terraform-provider-sewan)
- SonarQube analysis : ![Sonar Status](https://sonarcloud.io/api/project_badges/measure?project=terraform-provider-sewan-key&metric=alert_status)

<img src="http://entreprises.smallizbeautiful.fr/logo/Sewan-Communications.jpg" width="500px">

Maintainers
-----------

This provider plugin is maintained by the Sewan's team.

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.10.x

Usage
---------------------

Take a look in the website folder to get fully explained examples and documentation.

Use the docker image for Sewan's plugin
---------------------------

```sh
docker pull sewan/terraform-provider-sewan
```

[Docker image description and source](https://hub.docker.com/r/sewan/terraform-provider-sewan/)


Building The Provider
---------------------
* [Install terraform](https://www.terraform.io/intro/getting-started/install.html)

* Set up [Go](http://www.golang.org) your dev environment with version 1.10.x . You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

* Clone Sewan's sdk and terraform plugin repositories to: `$GOPATH/src/github.com/SewanDevs/` and terraform sources
```sh
git clone https://github.com/SewanDevs/sewan_go_sdk.git $GOPATH/src/github.com/SewanDevs/sewan_go_sdk
git clone https://github.com/SewanDevs/terraform-provider-sewan.git $GOPATH/src/github.com/SewanDevs/terraform-provider-sewan
git clone https://github.com/hashicorp/terraform.git $GOPATH/src/github.com/hashicorp/terraform
```

* Install additional library used in unit test:
```sh
go get -u github.com/google/go-cmp/cmp
```

* **Optional steps :**

  To run unit tests of the provider ans sdk, run `make test`.
```sh
cd $GOPATH/src/github.com/SewanDevs/terraform-provider-sewan
make test
```

* Provider compilation

  This will build the provider and put the provider binary in the `$GOPATH/bin` directory.
```sh
cd $GOPATH/src/github.com/SewanDevs/terraform-provider-sewan
make build
```
