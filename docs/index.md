---
layout: "ultradns"
page_title: "Provider: UltraDNS"
description: |-
  The Terraform UltraDNS provider is a plugin for Terraform used to manage UltraDNS resources using Terraform codes. You must configure the provider with the proper credentials before you can use it.
---

# UltraDNS Provider

The Terraform UltraDNS provider is a plugin for Terraform used to manage UltraDNS resources using Terraform codes. You must configure the provider with the proper credentials before you can use it.

## Example Usage

```terraform
terraform {
  required_providers {
    ultradns = {
      source = "ultradns/ultradns"
      version = "1.0.0"
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

Static credentials can be provided by adding an `username`, `password`, `hosturl` 
in-line in the ultradns provider block:

Usage:

```terraform
provider "ultradns" {
 username = "username"
 password = "password"
 hosturl = "https://api.test.ultradns.net/"
}
```

### Environment Variables

You can provide your credentials via the `ULTRADNS_USERNAME`,`ULTRADNS_PASSWORD` ,`ULTRADNS_HOST_URL`  environment variables, representing your username, password, hosturl respectively.

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

 The following arguments are supported in the UltraDNS
 `provider` block:

* `username` - This is the username for ultradns rest api. It must be provided, but
  it can also be sourced from the `ULTRADNS_USERNAME` environment variable.

* `password` - This is the password for ultradns rest api. It must be provided, but
  it can also be sourced from the `ULTRADNS_PASSWORD` environment variable.

* `hosturl` - This is the url for ultradns rest api. It must be provided, but
  it can also be sourced from the `ULTRADNS_HOST_URL` environment variables.
