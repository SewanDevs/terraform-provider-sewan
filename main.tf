provider "sewan" {
  api_token = "45acaa30efc5d3f2c737d543c27f6ab890f822d5"
  api_url = "https://next.cloud-datacenter.fr/api/clouddc/"
}

resource "sewan_clouddc_vdc" "terraform-vdc" {
  name = "terraform-vdc"
  enterprise = "sewan-rd-cloud-beta"
  datacenter = "dc1"
  vdc_resources=[
  {
    resource="sewan-rd-cloud-beta-mono-ram"
    total=10
  },
  {
    resource="sewan-rd-cloud-beta-mono-cpu"
    total=10
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

//
// RESOURCES CREATED FROM TEMPLATE
//
resource "sewan_clouddc_vm" "template-server" {
  depends_on = ["sewan_clouddc_vdc.terraform-vdc"]
  count = 2
  name = "template-server${count.index}"
  ram = 1
  cpu = 2
  //disks=[
  //{
  //  name= "disk-centos7-rd-DC1-1"
  //  storage_class="storage_enterprise"
  //  size=16
  //  deletion= false
  //},
  //]
  //{
  //  name="add disk test"
  //  size=16
  //  storage_class="storage_enterprise"
  //},
  //{
  //  name="add disk test2"
  //  size=0
  //  storage_class="storage_enterprise"
  //},
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
}

//
// TEMPLATE-LESS RESOURCES
//
resource "sewan_clouddc_vm" "server" {
  depends_on = ["sewan_clouddc_vdc.terraform-vdc"]
  count = 5
  ram = 1
  cpu = 1
  os = "CentOS"
  name = "server${count.index}"
  disks=[
    {
      name="disk 1"
      size=1
      storage_class="storage_enterprise"
    },
    {
      name="disk 2"
      size=2
      storage_class="storage_enterprise"
    },
  ]
  nics=[
    {
      vlan="internal-2410"
      connected=true
    },
  ]
  vdc = "${sewan_clouddc_vdc.terraform-vdc.slug}"
  storage_class="storage_enterprise"
  backup = "backup-no-backup"
  boot = "on disk"
}
