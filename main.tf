provider "sewan" {
  api_token = "6a54ebf17b29d12b3449aefc8a43c24f99e34bb4"
  api_url = "https://next.cloud-datacenter.fr/api/clouddc/"
}

//resource "sewan_clouddc_vdc" "terraform-vdc-charge-test" {
//  count = 10
//  name = "terraform-vdc-charge-test${count.index}"
//  enterprise = "sewan-rd-cloud-beta"
//  datacenter = "dc1"
//  vdc_resources=[
//  {
//    resource="sewan-rd-cloud-beta-mono-ram"
//    total=1
//  },
//  {
//    resource="sewan-rd-cloud-beta-mono-cpu"
//    total=1
//  },
//  ]
//}


resource "sewan_clouddc_vdc" "terraform-vdc" {
  name = "terraform-vdc"
  enterprise = "sewan-rd-cloud-beta"
  datacenter = "dc1"
  vdc_resources=[
  {
    resource="sewan-rd-cloud-beta-mono-ram"
    total=8
  },
  {
    resource="sewan-rd-cloud-beta-mono-cpu"
    total=8
  },
  {
    resource="sewan-rd-cloud-beta-mono-storage_enterprise"
    total=50
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

resource "sewan_clouddc_vm" "server" {
  depends_on = ["sewan_clouddc_vdc.terraform-vdc"]
  count = 10
  name = "server${count.index}"
  vdc = "${sewan_clouddc_vdc.terraform-vdc.slug}"
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
      storage_class="storage_enterprise"
    },
    {
      name="disk-centos7-rd-DC1-2"
      size=1
      storage_class="storage_enterprise"
    },
  ]
  boot = "on disk"
  storage_class = "storage_enterprise"
  backup = "backup-no-backup"
  comment = "update test"
}

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
