# Terraform-provider-jwk


A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) datastructure that represents a cryptographic key, these public keys can be used to make integrations of services such as Auth and OIDC authentication. 

Generating this public key can be a manual process with some bash commands, but creating the public key directly with terraform will help prevent using bash commands and then passing the call as an input to terraform.


Supported resource:
- `jwk_document`

## Installation

OS X & Linux:

```sh
make build install
```

## Example Usage

this public_key can be generated using ```terraform-tls```  provider or manually using the ```file()```  function

```terraform
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
```

## Development Setup

```sh
make install
```

## Running Tests

How to run tests against this repository, and optionally generate a coverage report.

```sh
ginkgo -v -r -compilers 0 -cover -covermode count -coverprofile ./coverage.out
```

## Contributing

1. Fork it (<https://github.com/yourname/yourproject/fork>)
2. Create your feature branch (`git checkout -b foobar_stuff`)
3. Commit your changes (`git commit -am 'feat: add some foobar'`)
4. Push to the branch (`git push origin foobar_stuff`)
5. Create a new upstream pull request (`hub pull-request -b braybaut/terraform-provider-jwk:main`)

## Credits 

Made with ❤  By [Braybaut](https://twitter.com/braybaut)