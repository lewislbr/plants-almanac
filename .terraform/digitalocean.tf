variable "do_api_token" {}
variable "do_ssh_key_fingerprint" {}

provider "digitalocean" {
  token = var.do_api_token
}

resource "digitalocean_droplet" "plantdex" {
  image              = "docker-20-04"
  ipv6               = true
  monitoring         = true
  name               = "plantdex"
  private_networking = false
  region             = "fra1"
  size               = "s-1vcpu-1gb"
  ssh_keys = [
    var.do_ssh_key_fingerprint
  ]
  user_data = file("${path.module}/cloud-init.yaml")
}

output "public_ipv4" {
  value = digitalocean_droplet.plantdex.ipv4_address
}
