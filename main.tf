provider "sewan" {
  air_drum_token = "17f061821bac9e12f9a2ded3928e624ae7c28448"
  air_drum_url = "https://next.cloud-datacenter.fr/api/clouddc/vm/"
}

resource "sewan_serverVM" "skeleton-server1" {
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

resource "sewan_serverVM" "skeleton-client1" {
  name = "client1"
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

resource "sewan_serverVM" "skeleton-client2" {
  name = "client2"
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
