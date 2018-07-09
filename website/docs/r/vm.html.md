---
layout: "sewan"
page_title: "Sewan: vm"
sidebar_current: "docs-sewan-resource-vm"
description: |-
  Manages Sewan clouddc vm.
---

# sewan\_vm

Provides a [Sewan's cloud data center](https://www.sewan.fr/cloud-data-center/) (VDC : virtual data center) virtual machine (VM).

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
      vlan="vlan 1"
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

### VM creation with Sewan's cloud data center own templates

To consult the list of available templates for your company or create new ones, access your company account on [cloud-datacenter.fr](https://cloud-datacenter.fr).

* `name` - *(Required, string)* vm name
* `template` - *(Required, string)* optional field required for creating a vm from a template
* Arguments handled by the template
  * `os` - Can not be set as it is template provided
  * `ram` - *(Optional, string)* template provided value can be modified, value unit : GiB
  * `cpu` - *(Optional, string)* template provided value can be modified

* `disks` - *(Optional, list of maps)* disks allocated to the vm, minimum of 1 is required
  * `name` - *(Required, string)* disk name
  * `size` - *(Required, int)* disk size in GiB
  * `storage_class` - *(Required, string)* type of virtual disks (accepted values : "Enterprise Storage", "Performance Storage", "High Performance Storage")
  * `deletion` *(Optional, bool)* set to "true" to deleted a template created disk

  **Warning, please read this** : On creation, additional disk can not be created, only the template provided disks are created. Additional disks can be added once the vm is created. Moreover, a disks created by a template can be deleted by providing it's name and set deleted field to "true"

* `nics` - *(Optional, list of maps)* network interfaces allocated to the vm, can be nil
  * `vlan` - *(Required, string)* nic vlan name
  * `connected` - *(Optional, boolean)* nic status
* `vdc` - *(Required, string)* vm's virtual data center name
* `boot` - *(Optional, string)* boot mode (accepted values : "", "on disk")
* `storage_class` - *(Required, string)* type of template created disks storage type (accepted values : "storage_enterprise", "storage_performance", "storage_high_performance") [more infos](https://www.sewan.fr/cloud-data-center/)
* `backup` - *(Required, string)* backup mode (accepted values : "backup-no-backup", "backup-7-days", "backup-31-days")
* `disk_image` - *(Optional, string)* disk image to boot when `boot` is set to "on disk", it can be set only after a vm creation

### VM fully configured by terraform plan file (.tf file)

* `name` - *(Required, string)* vm name
* `os` - Can not be set as it is template provided
* `ram` - *(Optional, string)* template provided value can be modified, value unit : GiB
* `cpu` - *(Optional, string)* template provided value can be modified

* `disks` - *(Required, list of maps)* disks allocated to the vm, minimum of 1 is required
  * `name` - *(Required, string)* disk name
  * `size` - *(Required, int)* disk size in GiB
  * `storage_class` - *(Required, string)* disks storage type (accepted values : "storage_enterprise", "storage_performance", "storage_high_performance") [more infos](https://www.sewan.fr/cloud-data-center/)

* `nics` - *(Optional, list of maps)* network interfaces allocated to the vm, can be nil
  * `vlan` - *(Required, string)* nic vlan name
  * `connected` - *(Optional, boolean)* nic status
* `vdc` - *(Required, string)* vm's virtual data center name
* `boot` - *(Optional, string)* boot mode (accepted values : "", "on disk")
* `backup` - *(Required, string)* backup mode (accepted values : "backup-no-backup", "backup-7-days", "backup-31-days")
* `disk_image` - *(Optional, string)* disk image to boot when `boot` is set to "on disk", it can be set only after a vm creation

## Attributes Reference

The following attributes are exported :

* `id` - ID of the new resource

## Import

Instance import is not yet supported, coming soon.
