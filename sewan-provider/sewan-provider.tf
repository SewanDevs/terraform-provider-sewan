provider "sewan" {}

resource "sewan_test" "sewan_ressource_test1" {
	name = "pof_walkthrough1"
	paused = false
	resolution = 1
	url = "www.sewan.fr"
	port = 80
}
