---
layout: "sewan"
page_title: "Sewan: vm"
sidebar_current: "docs-sewan-resource-vm"
description: |-
  Manages Sewan clouddc vm.
---

**< WARNING : page in construction>**

# sewan\_vm

Provides a clouddc virtual machine in a specific virtual data center (VDC).

## Example Usage

```hcl
//
// RESOURCES CREATED FROM A SEWAN'S CLOUD DATA CENTER TEMPLATE
//
resource "sewan_clouddc_vm" "template-created-vm" {
  depends_on = ["sewan_clouddc_vdc.vdc-example"]
  count = 10
  name = "template-created-vm${count.index}"
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
  ram = 1 //GiB
  cpu = 1
  os = "CentOS"
  name = "vm${count.index}"
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

## Argument Reference

### VM created from Sewan's cloud data center own templates

The following vm creation arguments are supported :

* `name` - (Required, string) vm name
* `os` - (Required, string) OS of the vm
* `ram` - (Required, string) ram size in GiB
* `cpu` - (Required, string) number of allocated cpus
* `disks` - (Required, listof maps) disks allocated to the vm, minimum of 1 is required
  * `name` - (Required, string) disk name
  * `size` - (Required, int) disk size in GiB
  * `storage_class` - (Required, string) type of virtual disks (accepted values : "Enterprise Storage", "Performance Storage", "High Performance Storage")
* `nics` - (Required, list of maps) disks allocated to the vm, minimum of 1 is required
  * `vlan` - (Required, string) nic vlan name
  * `connected` - (Optional, boolean) nic status
* `vdc` - (Required, string) vm's virtual data center name
* `boot` - (Optional, string) boot mode (accepted values : "", "on disk")
* `storage_class` - (Required, string) vdc resource disk
* `backup` - (Required, string) backup mode (accepted values : "backup-no-backup", "backup-7-days", "backup-31-days")
* `disk_image` - (Optional, string)

### VM fully configured by terraform plan file (.tf file)

The following vm creation arguments are supported :

* `name` - (Required, string) vm name
* `os` - (Required, string) OS of the vm
* `ram` - (Required, string) ram size in GiB
* `cpu` - (Required, string) number of allocated cpus
* `disks` - (Required, listof maps) disks allocated to the vm, minimum of 1 is required
  * `name` - (Required, string) disk name
  * `size` - (Required, int) disk size in GiB
  * `storage_class` - (Required, string) type of virtual disks (accepted values : "Enterprise Storage", "Performance Storage", "High Performance Storage")
* `nics` - (Required, list of maps) disks allocated to the vm, minimum of 1 is required
  * `vlan` - (Required, string) nic vlan name
  * `connected` - (Optional, boolean) nic status
* `vdc` - (Required, string) vm's virtual data center name
* `boot` - (Optional, string) boot mode (accepted values : "", "on disk")
* `backup` - (Required, string) backup mode (accepted values : "backup-no-backup", "backup-7-days", "backup-31-days")
* `disk_image` - (Optional, string)

## Attributes Reference

The following attributes are exported :

* `id` - ID of the new resource

## Import

Instance import is not yet supported, coming soon.
