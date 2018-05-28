provider "sewan" {
  api_token = "b9ede87382b298b0b664830b00decc19edfc81e3"
  api_url = "https://next.cloud-datacenter.fr/api/clouddc/"
}

resource "sewan_clouddc_vm" "skeleton-server1" {
  name = "skeleton-server1"
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

resource "sewan_clouddc_vm" "skeleton-client1" {
  name = "skeleton-client1"
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

resource "sewan_clouddc_vm" "skeleton-client2" {
  name = "skeleton-client2"
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
