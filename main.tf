provider "sewan" {
  air_drum_token = "17f061821bac9e12f9a2ded3928e624ae7c28448"
  air_drum_url = "https://next.cloud-datacenter.fr/api/clouddc/vm/"
}

resource "sewan_serverVM" "skeleton-server" {
  name = "skeleton-server"
  vdc = "sewan-rd-cloud-beta-dc1-terraf"
  ram  = "2"
  cpu = "2"
  disk_image = ""
  boot = "on disk"
  template = "centos7-rd-dc1"
  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
  backup = "backup-no-backup"
}

resource "sewan_clientVM" "skeleton-client1" {
  name = "client1"
  vdc = "sewan-rd-cloud-beta-dc1-terraf"
  ram  = "1"
  cpu = "1"
  disk_image = ""
  boot = "on disk"
  template = "centos7-rd-dc1"
  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
  backup = "backup-no-backup"
}

resource "sewan_clientVM" "skeleton-client2" {
  name = "client2"
  vdc = "sewan-rd-cloud-beta-dc1-terraf"
  ram  = "1"
  cpu = "1"
  disk_image = ""
  boot = "on disk"
  template = "centos7-rd-dc1"
  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
  backup = "backup-no-backup"
}
