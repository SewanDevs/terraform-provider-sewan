---
layout: "sewan"
page_title: "Provider: Sewan"
sidebar_current: "docs-sewan-index"
description: |-
  The Sewan provider is used to interact with Sewan "AirDrum" API to provide vdc and vms.
---

# Sewan Provider

The Sewan provider is used to interact with Sewan "AirDrum" API to provide vdc and vms.

Use the navigation to the left to read about the available data sources.

## Example Usage

```hcl
provider "sewan" {
  api_token = "111111111111111111"
  api_url = "https://next.cloud-datacenter.fr/api/clouddc/vm/"
}

resource "sewan_clouddc_vm" "server_resource_name" {
  name = "server_name"
  vdc = "sewan-rd-cloud-beta-dc1-terraf"
  os = "CentOS"
  ram  = "2"
  cpu = "2"
  disk_image = ""
  nics=[
  {
    vlan="internal-2412"
    connected=true
  },
  {
    vlan="internal-2410"
    connected=true
  },
  ]
  disks=[
    {
      name="disk-centos7-rd-DC1-1"
      size=20
      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
    },
    {
      name="disk-centos7-rd-DC1-2"
      size=20
      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
    }
  ]
  boot = "on disk"
  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
  backup = "backup-no-backup"
}

resource "sewan_clouddc_vm" "client_resource_name" {
  name = "client_name"
  vdc = "sewan-rd-cloud-beta-dc1-terraf"
  os = "CentOS"
  ram  = "1"
  cpu = "1"
  disk_image = ""
  nics=[
  {
    vlan="internal-2404"
    connected=true
  },
  ]
  disks=[
    {
      name="disk-centos7-rd-DC1-1"
      size=20
      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
    },
  ]
  boot = "on disk"
  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
  backup = "backup-no-backup"
}
```

NB 1 : add info about how to get a valid token

NB 2 : add info about how to choose api url