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

It creates an instance of the described vdc with a dynamic name field.

*Warning : multiple vdc instances are not supported as vm dependence to vdc are unique.*

```hcl
resource "sewan_clouddc_vdc" "vdc-example" {
  name = "vdc example"
  enterprise = "your-company"
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
```

## Argument Reference

The following vdc creation arguments are supported :

* `name` - (Required, string) vdc name
* `enterprise` - (Required, string) name of your enterprise
* `datacenter` - (Required, string) name of the datacenter
* `vdc_resources` - (Required, list of maps) resources allocated to the vdc, minimum of 1 is required
  * `resource` - (Required, string) resource name (accepted resources for storage "storage_enterprise", "storage_performance", "storage_high_performance"), other resources : "ram" & "cpu"
  * `total` - (Required, int) size of the resource (GiB for RAM or storage class, number for CPU etc.)

## Attributes Reference

The following attributes are exported :

* `id` - ID of the new resource

## Import

Instance import is not yet supported.
