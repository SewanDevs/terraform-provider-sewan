provider "sewan" {
  api_token = "ad313c5dc3b067f61f73f3a74182f14eb1a81809"
  api_url = "http://airdrumnext-api-vma-1.mgt.sewan.fr:8081/api/clouddc/vm/"
}

resource "sewan_clouddc_vm" "skeleton-server1" {
  name = "skeleton-server1H"
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
  name = "skeleton-client1H"
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
  name = "skeleton-client2H"
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
