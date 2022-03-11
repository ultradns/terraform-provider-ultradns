---
layout: "ultradns"
page_title: "Provider: UltraDNS"
description: |-
  The UltraDNS provider is a plugin for Terraform used to manage UltraDNS resources using HashiCorp Configuration Language (HCL). You must configure the provider with the proper credentials before you can use it.
---

# UltraDNS Provider

The UltraDNS provider is a plugin for <a href="https://www.terraform.io">Terraform</a> used to manage UltraDNS resources using HashiCorp Configuration Language (HCL). You must configure the provider with the proper credentials before you can use it.

## Example Usage

```terraform
terraform {
  required_providers {
    ultradns = {
      source = "ultradns/ultradns"
      version = ">= 1.3.0"
    }
  }
}

# Configure the UltraDNS Provider
provider "ultradns" {
  
}
```
## Authentication
The following methods are supported when providing the credential configuration for the UltraDNS provider:

- Static credentials
- Environment variables

### Static Credentials
Hard-coded credentials are not recommended in any Terraform configuration.

Static credentials are provided by adding `username`, `password`, `hosturl` 
in-line in the ultradns provider block:

Usage:

```terraform
provider "ultradns" {
 username = "username"
 password = "password"
 hosturl = "https://api.ultradns.net/"
}
```

### Environment Variables

If using Environment variables, your credentials are provided via the `ULTRADNS_USERNAME`, `ULTRADNS_PASSWORD`, and `ULTRADNS_HOST_URL` Environment variables, which represent your username, password, and hosturl respectively.

Usage:

```terraform
provider "ultradns" {}
```

```sh
$ export ULTRADNS_USERNAME="ULTRADNS_USERNAME"
$ export ULTRADNS_PASSWORD="ULTRADNS_PASSWORD"
$ export ULTRADNS_HOST_URL="ULTRADNS_HOST_URL"
$ terraform plan
```

## Argument Reference

 The following arguments are supported in the UltraDNS `provider` block:

* `username` - This is the username for UltraDNS REST API. It must be provided, but
  it can also be sourced from the `ULTRADNS_USERNAME` environment variable.

* `password` - This is the password for UltraDNS REST API. It must be provided, but
  it can also be sourced from the `ULTRADNS_PASSWORD` environment variable.

* `hosturl` - This is the url for UltraDNS REST API. It must be provided, but
  it can also be sourced from the `ULTRADNS_HOST_URL` environment variables.
