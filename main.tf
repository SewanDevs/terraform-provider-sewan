provider "sewan" {
  api_token = "da4c32a1b5403c9f1a88f984bc94bd802800120e"
  api_url = "https://next.cloud-datacenter.fr/api/clouddc/"
}

resource "sewan_clouddc_vdc" "terraform-built-vdc1" {
  name = "terraform-built-vdc1"
  enterprise = "sewan-rd-cloud-betan"
  datacenter = "dc1"
  vdc_resources = [
    {
      "resource"="sewan-rd-cloud-beta-mono-ram"
      "total"="1"
    },
    {
      "resource"="sewan-rd-cloud-beta-mono-cpu"
      "total"="1"
    },
    {
      "resource"="sewan-rd-cloud-beta-mono-storage_enterprise"
      "total"="10"
    },
    {
      "resource"="sewan-rd-cloud-beta-mono-storage_performance"
      "total"="10"
    },
    {
      "resource"="sewan-rd-cloud-beta-mono-storage_high_performance"
      "total"="10"
    },
  ]
}

//resource "sewan_clouddc_vm" "skeleton-server1" {
//  name = "skeleton-server1"
//  vdc = "sewan-rd-cloud-beta-dc1-terraf"
//  os = "CentOS"
//  ram  = 8
//  cpu = 4
//  disk_image = ""
//  nics=[
//  {
//    vlan="internal-2412"
//    connected=false
//  },
//  {
//    vlan="internal-2410"
//    connected=true
//  },
//  ]
//  disks=[
//    {
//      name="disk-centos7-rd-DC1-1"
//      size=40
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//    {
//      name="disk-centos7-rd-DC1-2"
//      size=20
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    }
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "skeleton-client1" {
//  name = "skeleton-client1"
//  vdc = "sewan-rd-cloud-beta-dc1-terraf"
//  os = "Debian"
//  ram  = 1
//  cpu = 1
//  disk_image = ""
//  nics=[
//  {
//    vlan="internal-2404"
//    connected=true
//  },
//  ]
//  disks=[
//    {
//      name="disk-centos7-rd-DC1-1"
//      size=20
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "skeleton-client2" {
//  name = "skeleton-client2"
//  vdc = "sewan-rd-cloud-beta-dc1-terraf"
//  os = "CentOS"
//  ram  = 1
//  cpu = 1
//  disk_image = ""
//  nics=[
//  {
//    vlan="internal-2404"
//    connected=true
//  },
//  ]
//  disks=[
//    {
//      name="disk-centos7-rd-DC1-1"
//      size=20
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "skeleton-client3" {
//  name = "skeleton-client3"
//  vdc = "sewan-rd-cloud-beta-dc1-terraf"
//  os = "CentOS"
//  ram  = 1
//  cpu = 1
//  disk_image = ""
//  nics=[
//  {
//    vlan="internal-2404"
//    connected=true
//  },
//  ]
//  disks=[
//    {
//      name="disk-centos7-rd-DC1-1"
//      size=20
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "skeleton-client4" {
//  name = "skeleton-client4"
//  vdc = "sewan-rd-cloud-beta-dc1-terraf"
//  os = "CentOS"
//  ram  = 1
//  cpu = 1
//  disk_image = ""
//  nics=[
//  {
//    vlan="internal-2404"
//    connected=true
//  },
//  ]
//  disks=[
//    {
//      name="disk-centos7-rd-DC1-1"
//      size=20
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "skeleton-client5" {
//  name = "skeleton-client5"
//  vdc = "sewan-rd-cloud-beta-dc1-terraf"
//  os = "CentOS"
//  ram  = 1
//  cpu = 1
//  disk_image = ""
//  nics=[
//  {
//    vlan="internal-2404"
//    connected=true
//  },
//  ]
//  disks=[
//    {
//      name="disk-centos7-rd-DC1-1"
//      size=20
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "skeleton-client6" {
//  name = "skeleton-client6"
//  vdc = "sewan-rd-cloud-beta-dc1-terraf"
//  os = "CentOS"
//  ram  = 1
//  cpu = 1
//  disk_image = ""
//  nics=[
//  {
//    vlan="internal-2404"
//    connected=true
//  },
//  ]
//  disks=[
//    {
//      name="disk-centos7-rd-DC1-1"
//      size=20
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "skeleton-client7" {
//  name = "skeleton-client7"
//  vdc = "sewan-rd-cloud-beta-dc1-terraf"
//  os = "CentOS"
//  ram  = 1
//  cpu = 1
//  disk_image = ""
//  nics=[
//  {
//    vlan="internal-2404"
//    connected=true
//  },
//  ]
//  disks=[
//    {
//      name="disk-centos7-rd-DC1-1"
//      size=20
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "skeleton-client8" {
//  name = "skeleton-client8"
//  vdc = "sewan-rd-cloud-beta-dc1-terraf"
//  os = "CentOS"
//  ram  = 1
//  cpu = 1
//  disk_image = ""
//  nics=[
//  {
//    vlan="internal-2404"
//    connected=true
//  },
//  ]
//  disks=[
//    {
//      name="disk-centos7-rd-DC1-1"
//      size=20
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "skeleton-client9" {
//  name = "skeleton-client9"
//  vdc = "sewan-rd-cloud-beta-dc1-terraf"
//  os = "CentOS"
//  ram  = 1
//  cpu = 1
//  disk_image = ""
//  nics=[
//  {
//    vlan="internal-2404"
//    connected=true
//  },
//  ]
//  disks=[
//    {
//      name="disk-centos7-rd-DC1-1"
//      size=20
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "skeleton-client10" {
//  name = "skeleton-client10"
//  vdc = "sewan-rd-cloud-beta-dc1-terraf"
//  os = "CentOS"
//  ram  = 1
//  cpu = 1
//  disk_image = ""
//  nics=[
//  {
//    vlan="internal-2404"
//    connected=true
//  },
//  ]
//  disks=[
//    {
//      name="disk-centos7-rd-DC1-1"
//      size=20
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "skeleton-client11" {
//  name = "skeleton-client11"
//  vdc = "sewan-rd-cloud-beta-dc1-terraf"
//  os = "CentOS"
//  ram  = 1
//  cpu = 1
//  disk_image = ""
//  nics=[
//  {
//    vlan="internal-2404"
//    connected=true
//  },
//  ]
//  disks=[
//    {
//      name="disk-centos7-rd-DC1-1"
//      size=20
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "skeleton-client12" {
//  name = "skeleton-client12"
//  vdc = "sewan-rd-cloud-beta-dc1-terraf"
//  os = "CentOS"
//  ram  = 1
//  cpu = 1
//  disk_image = ""
//  nics=[
//  {
//    vlan="internal-2404"
//    connected=true
//  },
//  ]
//  disks=[
//    {
//      name="disk-centos7-rd-DC1-1"
//      size=20
//      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
