---
layout: "sewan"
page_title: "Sewan: vm"
sidebar_current: "docs-sewan-resource-vm"
description: |-
  Manages Sewan clouddc vm.
---

# sewan\_vm

Provides a clouddc virtual machine in a specific virtual data center (VDC).

## Example Usage

```hcl
resource "sewan_clouddc_vm" "vm_name" {
  name = "<vm_name>"
  vdc = "<vm's vdc name>"
  os = "Debian"
  ram  = "4"
  cpu = "2"
  disk_image = ""
  nics=[
  {
    vlan="vlan1"
    connected=true
  },
  {
    vlan="vlan2"
    connected=false
  },
  ]
  disks=[
    {
      name="disk1"
      size=20
      v_disk="<Type of used virtual disks>"
    },
    {
      name="disk2"
      size=42
      v_disk="<Type of used virtual disks>"
    },
  ]
  boot = "on disk"
  backup = "backup-no-backup"

```

## Argument Reference

The following vm creation arguments are supported :

* `name` - (Required, string) vm name
* `os` - (Required, string) OS of the vm
* `ram` - (Required, string) ram size in GiB
* `cpu` - (Required, string) number of allocated cpus
* `disks` - (Required, listof maps) disks allocated to the vm, minimum of 1 is required
  * `name` - (Required, string) disk name
  * `size` - (Required, int) disk size in GiB
  * `v_disk` - (Required, string) type of virtual disks (accepted values : "Enterprise Storage", "Performance Storage", "High Performance Storage")
* `nics` - (Required, list of maps) disks allocated to the vm, minimum of 1 is required
  * `vlan` - (Required, string) nic vlan name
  * `connected` - (Optional, boolean) nic status
* `vdc` - (Required, string) vm's virtual data center name
* `boot` - (Optional, string) boot mode (accepted values : "", "on disk")
* `vdc_resource_disk` - (Required, string) vdc resource disk
* `backup` - (Required, string) backup mode (accepted values : "backup-no-backup", "backup-7-days", "backup-31-days")
* `disk_image` - (Optional, string)
* `comment` - (Optional, string)
* `dynamic_field` - (Optional, string)

## Attributes Reference

The following attributes are exported :

* `id` - ID of the new resource

## Import

Instance import is not supported.
