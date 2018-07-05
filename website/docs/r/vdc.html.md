---
layout: "sewan"
page_title: "Sewan: vdc"
sidebar_current: "docs-sewan-resource-vdc"
description: |-
  Manages Sewan clouddc vdc.
---

**< WARNING : page in construction>**

# sewan\_vdc

Provides a virtual data center (VDC).

## Example Usage

It creates 10 instance of the described vdc with a dynamic name field.

```hcl
resource "sewan_clouddc_vdc" "vdc-example" {
  name = "vdc example"
  enterprise = "your-company"
  datacenter = "a datacenter"
  vdc_resources=[
  {
    resource="your-company-mono-ram"
    total=10
  },
  {
    resource="your-company-mono-cpu"
    total=10
  },
  {
    resource="your-company-mono-storage_enterprise"
    total=80
  },
  {
    resource="your-company-mono-storage_performance"
    total=20
  },
  {
    resource="your-company-mono-storage_high_performance"
    total=10
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
