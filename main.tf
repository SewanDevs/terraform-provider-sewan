provider "sewan" {
  api_token   = "e9ffa8f6bcbaab2079dc675724cef504a405a24b"
  endpoint    = "https://next.cloud-datacenter.fr/api/clouddc"
  timeout     = 20
  max_retries = 3
}

resource "sewan_serverVM" "skeleton-server" {
  name = "skeleton-server"
  vdc = "sewan-rd-cloud-beta-dc1-terraf"
  ram  = 2
  cpu = 2
  disk_image = "nil"
  boot = "on disk"
  template = "centos7-rd-dc1"
  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
  backup = "backup-no-backup"
}

resource "sewan_clientVM" "skeleton-client1" {
  name = "client1"
  vdc = "sewan-rd-cloud-beta-dc1-terraf"
  ram  = 1
  cpu = 1
  disk_image = "nil"
  boot = "on disk"
  template = "centos7-rd-dc1"
  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
  backup = "backup-no-backup"
}

resource "sewan_clientVM" "skeleton-client2" {
  name = "client2"
  vdc = "sewan-rd-cloud-beta-dc1-terraf"
  ram  = 1
  cpu = 1
  disk_image = "nil"
  boot = "on disk"
  template = "centos7-rd-dc1"
  vdc_resource_disk = "sewan-rd-cloud-beta-dc1-terraf-sewan-rd-cloud-beta-mono-storage_enterprise"
  backup = "backup-no-backup"
}
