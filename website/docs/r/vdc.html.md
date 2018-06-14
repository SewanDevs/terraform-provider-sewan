---
layout: "sewan"
page_title: "Sewan: vdc"
sidebar_current: "docs-sewan-resource-vdc"
description: |-
  Manages Sewan clouddc vdc.
---

# sewan\_vdc

Provides a virtual data center (VDC).

## Example Usage

It creates 10 instance of the described vdc with a dynamic name field.

```hcl
resource "sewan_clouddc_vdc" "terraform-built-vdc" {
  count = 10
  name = "terraform-vdc-charge-test${count.index}"
  enterprise = "sewan-rd-cloud-beta"
  datacenter = "dc1"
  vdc_resources=[
  {
    resource="sewan-rd-cloud-beta-mono-ram"
    total=1
  },
  {
    resource="sewan-rd-cloud-beta-mono-cpu"
    total=1
  },
  {
    resource="sewan-rd-cloud-beta-mono-storage_enterprise"
    total=1
  },
  {
    resource="sewan-rd-cloud-beta-mono-storage_performance"
    total=1
  },
  {
    resource="sewan-rd-cloud-beta-mono-storage_high_performance"
    total=1
  },
  ]
}

```

## Argument Reference

The following vdc creation arguments are supported :

* `name` - (Required, string) vdc name
* `enterprise` - (Required, string) name of the enterprise using sewan data center
* `datacenter` - (Required, string) name of the datacenter
* `resources` - (Required, listof maps) disks allocated to the vdc, minimum of 1 is required
  * `resource` - (Required, string) resource name
  * `total` - (Required, int) size of the resource (GiB for RAM or storage, number for CPU etc.)

## Attributes Reference

The following attributes are exported :

* `id` - ID of the new resource

## Import

Instance import is not supported.
