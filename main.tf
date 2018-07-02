provider "sewan" {
  api_token = "bb5bb501371090a55e7a0f8ca03a8361232e8111"
  api_url = "https://next.cloud-datacenter.fr/api/clouddc/"
}

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
    total=80
  },
  {
    resource="sewan-rd-cloud-beta-mono-storage_performance"
    total=20
  },
  {
    resource="sewan-rd-cloud-beta-mono-storage_high_performance"
    total=10
  },
  ]
}

resource "sewan_clouddc_vm" "template-server" {
  depends_on = ["sewan_clouddc_vdc.terraform-vdc"]
  count = 2
  name = "template-server${count.index}"
  ram = 1
  cpu = 2
  //disks=[
  //  {
  //    name="add disk test"
  //    size=16
  //    storage_class="storage_performance"
  //  }
  //]
  nics=[
    {
      vlan="internal-2412"
      connected=false
    },
  ]
  //template = "Debian 7"
  template = "centos7-rd-DC1"
  vdc = "${sewan_clouddc_vdc.terraform-vdc.slug}"
  backup = "backup-no-backup"
  storage_class = "storage_enterprise"
  boot = "on disk"
  //dynamic_field = "test_dynamic_field"
}

//resource "sewan_clouddc_vm" "template-server" {
//  depends_on = ["sewan_clouddc_vdc.terraform-vdc"]
//  count = 1
//  name = "template-server${count.index}"
//  template = "centos7-rd-DC1"
//  disks=[
//    {
//      name="centos dc1"
//      size=15
//      storage_class="storage_enterprise"
//      delete = true
//    }
//  ]
//  nics=[
//    {
//      vlan="internal-2410"
//      connected=true
//    },
//    {
//      vlan="internal-2410"
//      connected=true
//    },
//  ]
//  vdc = "${sewan_clouddc_vdc.terraform-vdc.slug}"
//  backup = "backup-no-backup"
//  storage_class = "storage_enterprise"
//}
