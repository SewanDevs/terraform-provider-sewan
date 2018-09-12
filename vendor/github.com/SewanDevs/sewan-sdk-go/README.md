Sewan's cloud data center go sdk
================================

- Website: https://www.sewan.fr/
- Travis build : [![Build Status](https://travis-ci.com/SewanDevs/sewan-sdk-go.svg?branch=github_release)](https://travis-ci.com/SewanDevs/sewan-sdk-go)
- SonarQube analysis : [![Sonar Status](https://sonarcloud.io/api/project_badges/measure?project=sewan-sdk-go&metric=alert_status)](https://sonarcloud.io/dashboard?id=sewan-sdk-go)
- [![Go Report Card](https://goreportcard.com/badge/github.com/SewanDevs/sewan-sdk-go)](https://goreportcard.com/report/github.com/SewanDevs/sewan-sdk-go)
- [Snyk security audit](https://app.snyk.io) : beta version, no badge available, No known vulnerabilities found.

<img src="http://entreprises.smallizbeautiful.fr/logo/Sewan-Communications.jpg" width="500px">

Maintainers
-----------

This sdk is maintained by the Sewan's team.

It is consumed by Sewan's terraform provider plugin (github.com:SewanDevs/terraform-provider-sewan.git) to communicate with [Sewan's cloud data center](https://www.sewan.fr/cloud-data-center/).

Requirements
------------

-	[Go](https://golang.org/doc/install) 1.10.x (to build the provider plugin)

Doc
--------------------

* Architecture doc : Available under doc folder, it contains sequence diagrams and a module diagram.
* Standard golang sdk doc :
> prerequisite : current repository have been cloned in GOPATH
> ```sh
> run godoc -http=:6060
> ```
> Open the webpage "http://localhost:6060/pkg/" and navigate through the tree to sewan-sdk-go package
