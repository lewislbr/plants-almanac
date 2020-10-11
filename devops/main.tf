terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
      version = "~> 1.22.1"
    }
  }
}

variable "do_api_token" {}
variable "do_ssh_key_fingerprint" {}

provider "digitalocean" {
  token = var.do_api_token
}

data "template_file" "cloud-init" {
  template = file("${path.module}/cloud-init.yml")
}

resource "digitalocean_droplet" "plantdex" {
  image = "docker-20-04"
  name = "plantdex"
  region = "fra1"
  size = "s-1vcpu-1gb"
  monitoring = true
  ipv6 = true
  private_networking = false
  ssh_keys = [
    var.do_ssh_key_fingerprint
  ]
  user_data = data.template_file.cloud-init.rendered
}

output "public_ipv4" {
  value = digitalocean_droplet.plantdex.ipv4_address
}
