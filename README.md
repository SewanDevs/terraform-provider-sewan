Sewan Terraform Provider
========================

- Webistes : https://www.sewan.fr, https://www.terraform.io
- Travis build : [![Build Status](https://travis-ci.com/SewanDevs/terraform-provider-sewan.svg?branch=github_release)](https://travis-ci.com/SewanDevs/terraform-provider-sewan)
- SonarQube analysis : [![Sonar Status](https://sonarcloud.io/api/project_badges/measure?project=terraform-provider-sewan-key&metric=alert_status)](https://sonarcloud.io/dashboard?id=terraform-provider-sewan-key)
- [![Go Report Card](https://goreportcard.com/badge/github.com/SewanDevs/terraform-provider-sewan)](https://goreportcard.com/report/github.com/SewanDevs/terraform-provider-sewan)
- [Snyk security audit](https://app.snyk.io) : beta version, no badge available, No known vulnerabilities found.

<img src="http://entreprises.smallizbeautiful.fr/logo/Sewan-Communications.jpg" width="500px">

Maintainers
-----------

This provider plugin is maintained by the Sewan's team.

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.10.x
- [Go](https://golang.org/doc/install) 1.10.x
- [golint](https://github.com/golang/lint)

Usage
---------------------

Take a look in [the website folder](https://github.com/SewanDevs/terraform-provider-sewan/blob/github_release/website/docs) to get fully explained examples and documentation and [terraform guide](https://www.terraform.io/intro/getting-started):
- Infrastructure helloworld example : [helloworldInfrastructure folder](https://github.com/SewanDevs/terraform-provider-sewan/blob/github_release/helloworldInfrastructure)
- [Global Sewan's provider usage](https://github.com/SewanDevs/terraform-provider-sewan/blob/github_release/website/docs/index.html.markdown)
- [vm (virtual machine) configuration](https://github.com/SewanDevs/terraform-provider-sewan/blob/github_release/website/docs/r/vm.html.md)
- [vdc (virtual data center) configuration](https://github.com/SewanDevs/terraform-provider-sewan/blob/github_release/website/docs/r/vdc.html.md)

Use the docker image for Sewan's plugin
---------------------------

```sh
docker pull sewan/terraform-provider-sewan
```

[Docker image description and source](https://hub.docker.com/r/sewan/terraform-provider-sewan/), it contains a running environment containing terraform and current plugin, plus an example infrastructure configuration file.


Building The Provider
---------------------
* [Install terraform](https://www.terraform.io/intro/getting-started/install.html)

* Set up [Go](https://golang.org/doc/install) in your dev environment with version 1.10.x . You'll also need to correctly setup a [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) and add `export GOPATH=$HOME/go` and `PATH=$PATH:$GOPATH` to your `$HOME/.bashrc` file, then source it.

* Install golint
```sh
go get -u golang.org/x/lint/golint
```

* Clone Sewan's terraform plugin repositorie to: `$GOPATH/src/github.com/SewanDevs/`
```sh
git clone https://github.com/SewanDevs/terraform-provider-sewan.git $GOPATH/src/github.com/SewanDevs/terraform-provider-sewan
```

* **Optional steps :**

  To run unit tests, golint, gofmt and govet on the provider, run `make test`.
```sh
cd $GOPATH/src/github.com/SewanDevs/terraform-provider-sewan
make test
```

* Provider build

  This build the provider and put the provider binary in the `$GOPATH/bin` directory.

  NB : The build fails if any [golint](https://github.com/golang/lint),
  [gofmt](https://golang.org/cmd/gofmt/) or [govet](https://golang.org/cmd/vet/) issue is found.
```sh
cd $GOPATH/src/github.com/SewanDevs/terraform-provider-sewan
make build
```

* Provider install

  This moves the plugin binary from `$GOPATH/bin` to `$HOME/.terraform.d/plugins`, where terraform's plugins are stored.
```sh
make install
```

  NB : Notice this install is user limited, for system-wide installation,
  move the plugin binary to /opt/
