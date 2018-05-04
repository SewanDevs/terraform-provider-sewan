# terraform Sewan's airdrum API connector : A step by step guide

## Set up your environment
* install go : golang.org
* install terraform : terraform.io
* install an IDE : Sewan provide pro license of pycharm, ask you sysadmin for one

## Sources
This guide is build with the help of the following refs :
* www.terraform.io/guides/terraform-provider-development-program.html
* Hashicorp connector tutorial : www.youtube.com/watch?v=2BvpqmFpchI
* www.terraform.io/docs/extend/writing-custom-providers.html

## Setup
Go packages url to get with "go get <package url>" :
* github.com/hashicorp/terraform/ : terraform go package and plugin

## GLOSSARY
* provider
* resource

## Sewan skeleton provider creation steps
### i/ provide Sewan's provider schema
### ii/ Build Sewan's plugin
go build -o terraform-provider-<provider name=sewan>
### iii/ Define a Sewan's resources
redo : go build -o terraform-provider-<provider name=sewan>
### iv/ Invoke Sewan's provider
terraform init
terraform plan
### v/ Implements REST calls
go build -o terraform-provider-<provider name=sewan>
terraform init
terraform plan
terraform apply
