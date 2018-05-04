provider "sewan" {
  api_token     = "securetoken=="
  endpoint    = "https://next.cloud-datacenter.fr/api/clouddc"
  timeout     = 60
  max_retries = 5
}

resource "sewan_serverVM" "skeleton-server" {
  name = "skeleton-server"
  cpus = 2
  ram  = 8192
  address = "127.0.0.2"
}

resource "sewan_clientVM" "skeleton-client1" {
  name = "client1"
  cpus = 1
  ram  = 2048
  address = "127.0.0.3"
}

resource "sewan_clientVM" "skeleton-client2" {
  name = "client2"
  cpus = 1
  ram  = 2048
  address = "127.0.0.3"
}
