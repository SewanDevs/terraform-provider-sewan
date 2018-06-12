provider "sewan" {
  api_token = "51f27cf40ee99d60db93d75e196ec70bc4872ea7"
  api_url = "https://next.cloud-datacenter.fr/api/clouddc/"
}

resource "sewan_clouddc_vdc" "terraform-built-vdc1" {
  name = "terraform-built-vdc1"
  enterprise = "sewan-rd-cloud-beta"
  datacenter = "dc1"
  vdc_resources=[
  {
    resource="sewan-rd-cloud-beta-mono-ram"
    total=5
  },
  {
    resource="sewan-rd-cloud-beta-mono-cpu"
    total=5
  },
  {
    resource="sewan-rd-cloud-beta-mono-storage_enterprise"
    total=5
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

resource "sewan_clouddc_vm" "server1" {
  depends_on = ["sewan_clouddc_vdc.terraform-built-vdc1"]
  name = "server1"
  vdc = "sewan-rd-cloud-beta-dc1-terraf"
  os = "CentOS"
  ram  = 1
  cpu = 1
  disk_image = ""
  nics=[
  {
    vlan="internal-2412"
    connected=false
  },
  {
    vlan="internal-2410"
    connected=true
  },
  ]
  disks=[
    {
      name="disk-centos7-rd-DC1-1"
      size=1
      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
    },
    {
      name="disk-centos7-rd-DC1-2"
      size=1
      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
    },
  ]
  boot = "on disk"
  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
  backup = "backup-no-backup"
}

resource "sewan_clouddc_vm" "client1" {
  depends_on = ["sewan_clouddc_vdc.terraform-built-vdc1"]
  name = "client1"
  vdc = "sewan-rd-cloud-beta-dc1-terraf"
  os = "Debian"
  ram  = 1
  cpu = 1
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
      size=1
      v_disk="sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
    },
  ]
  boot = "on disk"
  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
  backup = "backup-no-backup"
}
//
//resource "sewan_clouddc_vm" "client2" {
//  name = "client2"
//  vdc = "terraform-built-vdc1"
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
//      v_disk="sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "client3" {
//  name = "client3"
//  vdc = "terraform-built-vdc1"
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
//      v_disk="sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "client4" {
//  name = "client4"
//  vdc = "terraform-built-vdc1"
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
//      v_disk="sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "client5" {
//  name = "client5"
//  vdc = "terraform-built-vdc1"
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
//      v_disk="sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "client6" {
//  name = "client6"
//  vdc = "terraform-built-vdc1"
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
//      v_disk="sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "client7" {
//  name = "client7"
//  vdc = "terraform-built-vdc1"
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
//      v_disk="sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "client8" {
//  name = "client8"
//  vdc = "terraform-built-vdc1"
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
//      v_disk="sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "client9" {
//  name = "client9"
//  vdc = "terraform-built-vdc1"
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
//      v_disk="sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "client10" {
//  name = "client10"
//  vdc = "terraform-built-vdc1"
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
//      v_disk="sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "client11" {
//  name = "client11"
//  vdc = "terraform-built-vdc1"
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
//      v_disk="sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
//resource "sewan_clouddc_vm" "client12" {
//  name = "client12"
//  vdc = "terraform-built-vdc1"
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
//      v_disk="sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
//
