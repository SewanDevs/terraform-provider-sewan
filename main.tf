provider "sewan" {
  api_token = "2fe8f62cfc506e87ece81ba41f0e9cd5528d009f"
  api_url = "https://next.cloud-datacenter.fr/api/clouddc/"
}

resource "sewan_clouddc_vdc" "terraform-built-vdc-charge-test" {
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


//resource "sewan_clouddc_vdc" "terraform-built-vdc" {
//  name = "terraform-built-vdc"
//  enterprise = "sewan-rd-cloud-beta"
//  datacenter = "dc1"
//  vdc_resources=[
//  {
//    resource="sewan-rd-cloud-beta-mono-ram"
//    total=8
//  },
//  {
//    resource="sewan-rd-cloud-beta-mono-cpu"
//    total=8
//  },
//  {
//    resource="sewan-rd-cloud-beta-mono-storage_enterprise"
//    total=10
//  },
//  {
//    resource="sewan-rd-cloud-beta-mono-storage_performance"
//    total=1
//  },
//  {
//    resource="sewan-rd-cloud-beta-mono-storage_high_performance"
//    total=1
//  },
//  ]
//}
//
//resource "sewan_clouddc_vm" "server" {
//  depends_on = ["sewan_clouddc_vdc.terraform-built-vdc"]
//  count = 3
//  name = "server${count.index}"
//  vdc = "${sewan_clouddc_vdc.terraform-built-vdc.slug}"
//  os = "CentOS"
//  ram  = 1
//  cpu = 1
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
//      size=1
//      v_disk="${sewan_clouddc_vdc.terraform-built-vdc.slug}-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//    {
//      name="disk-centos7-rd-DC1-2"
//      size=1
//      v_disk="${sewan_clouddc_vdc.terraform-built-vdc.slug}-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "${sewan_clouddc_vdc.terraform-built-vdc.slug}-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}

//
//resource "sewan_clouddc_vdc" "terraform-built-vdc-bis" {
//  name = "terraform-built-vdc-bis"
//  enterprise = "sewan-rd-cloud-beta"
//  datacenter = "dc1"
//  vdc_resources=[
//  {
//    resource="sewan-rd-cloud-beta-mono-ram"
//    total=8
//  },
//  {
//    resource="sewan-rd-cloud-beta-mono-cpu"
//    total=8
//  },
//  {
//    resource="sewan-rd-cloud-beta-mono-storage_enterprise"
//    total=10
//  },
//  {
//    resource="sewan-rd-cloud-beta-mono-storage_performance"
//    total=1
//  },
//  {
//    resource="sewan-rd-cloud-beta-mono-storage_high_performance"
//    total=1
//  },
//  ]
//}
//
//resource "sewan_clouddc_vm" "server-bis" {
//  depends_on = ["sewan_clouddc_vdc.terraform-built-vdc-bis"]
//  count = 5
//  name = "server-bis${count.index}"
//  vdc = "${sewan_clouddc_vdc.terraform-built-vdc-bis.slug}"
//  os = "CentOS"
//  ram  = 1
//  cpu = 1
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
//      size=1
//      v_disk="${sewan_clouddc_vdc.terraform-built-vdc-bis.slug}-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//    {
//      name="disk-centos7-rd-DC1-2"
//      size=1
//      v_disk="${sewan_clouddc_vdc.terraform-built-vdc-bis.slug}-sewan-rd-cloud-beta-mono-storage_enterprise"
//    },
//  ]
//  boot = "on disk"
//  vdc_resource_disk = "${sewan_clouddc_vdc.terraform-built-vdc-bis.slug}-sewan-rd-cloud-beta-mono-storage_enterprise"
//  backup = "backup-no-backup"
//}
