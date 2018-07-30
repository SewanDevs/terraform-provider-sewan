provider "sewan" {
  api_token = "<your-company token>"
  api_url = "https://cloud-datacenter.fr/api/clouddc/"
}

resource "sewan_clouddc_vdc" "vdc-example" {
  name = "vdc example"
  instance_number = "${count.index}"
  enterprise = "your-company"
  datacenter = "a datacenter"
  vdc_resources=[
  {
    resource="ram"
    total=10
  },
  {
    resource="cpu"
    total=10
  },
  {
    resource="storage_enterprise"
    total=80
  },
  {
    resource="storage_performance"
    total=20
  },
  {
    resource="storage_high_performance"
    total=10
  },
  ]
}

//
// RESOURCES CREATED FROM A SEWAN'S CLOUD DATA CENTER TEMPLATE
//
resource "sewan_clouddc_vm" "template-created-vm" {
  depends_on = ["sewan_clouddc_vdc.vdc-example"]
  count = 2 //instance of this resource to create
  name = "template-created-vm"
  instance_number = "${count.index}"
  enterprise = "${sewan_clouddc_vdc.vdc-example.enterprise}"
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
  count = 5
  ram = 1 //GiB
  cpu = 1
  os = "CentOS"
  name = "vm"
  instance_number = "${count.index}"
  enterprise = "${sewan_clouddc_vdc.vdc-example.enterprise}"
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
      vlan="internal-2410"
      connected=true
    },
  ]
  vdc = "${sewan_clouddc_vdc.vdc-example.slug}"
  backup = "backup-no-backup"
  boot = "on disk"
}
