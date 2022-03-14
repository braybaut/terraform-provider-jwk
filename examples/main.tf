terraform {
  required_providers {
    jwk = {
      version = "0.1.0"
      source  = "github.com/braybaut/jwk"
    }
    tls = {
      source  = "hashicorp/tls"
      version = "3.1.0"
    }
  }
}

provider "tls" {
  # Configuration options
}

provider "jwk" {}

resource "tls_private_key" "example" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "jwk_document" "new_jwk" {
  public_key_pem = tls_private_key.example.public_key_pem
}

output "jwk_document" {
  value = jwk_document.new_jwk.jwk_document
}


