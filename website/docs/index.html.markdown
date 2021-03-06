---
layout: "sewan"
page_title: "Provider: Sewan"
sidebar_current: "docs-sewan-index"
description: |-
  The Sewan provider is used to interact with Sewan "AirDrum" API to provide vdc and vms.
---

# Sewan Provider

The Sewan provider is used to interact with [Sewan's cloud data center](https://www.sewan.fr/cloud-data-center/) (clouddc) API to provide virtual data centers (vdc) and virtual machines (vm).

Use the navigation to the left to read about the available data sources, pay attention to vm created from sewan's clouddc template that does not have a standard terraform resource behavior.

## Get an api token for your company

Contact the support.

## Provider arguments Reference

* `api_url` - *(Required, string)* clouddc's url
* `api_token` - *(Required, string)* session token
* `enterprise` - *(Required, string)* your enterprise name available in the
clouddc's url after login : "cloud-datacenter.fr/fr/enterprise/<enterprise name>/dashboard"

## Example Usage

```hcl
provider "sewan" {
  api_token = "your-company token"
  api_url = "https://cloud-datacenter.fr/api/clouddc/"
  enterprise = "your enterprise account's name"
}

resource "sewan_clouddc_vdc" "vdc-example" {
  name = "vdc example"
  datacenter = "a datacenter"
  vdc_resources=[
  {
    resource="ram"
    total=10
  },
  {
    resource="cpu"
    total=10
  },
  {
    resource="storage_enterprise"
    total=80
  },
  {
    resource="storage_performance"
    total=20
  },
  {
    resource="storage_high_performance"
    total=10
  },
  ]
}

//
// RESOURCES CREATED FROM A SEWAN'S CLOUD DATA CENTER TEMPLATE
//
resource "sewan_clouddc_vm" "template-created-vm" {
  depends_on = ["sewan_clouddc_vdc.vdc-example"]
  count = 10
  name = "template-created-vm"
  instance_number = "${count.index + 1}"
  nics=[
    {
      vlan="vlan name"
      connected=false
    },
  ]
  template = "a template"
  vdc = "${sewan_clouddc_vdc.vdc-example.slug}"
  backup = "backup-no-backup"
  storage_class = "storage_enterprise"
  boot = "on disk"
}

//
// TEMPLATE-LESS RESOURCES
//
resource "sewan_clouddc_vm" "vm" {
  depends_on = ["sewan_clouddc_vdc.vdc-example"]
  count = 10
  ram = 1
  cpu = 1
  os = "CentOS"
  name = "vm-${count.index}"
  disks=[
    {
      name="disk 1"
      size=1
      storage_class="storage_performance"
    },
    {
      name="disk 2"
      size=2
      storage_class="storage_enterprise"
    },
  ]
  nics=[
    {
      vlan="vlan 1"
      connected=false
    },
    {
      vlan="vlan 2"
      connected=true
    },
  ]
  vdc = "${sewan_clouddc_vdc.vdc-example.slug}"
  backup = "backup-no-backup"
  boot = "on disk"
}
```
