# CDN Resources and Data Sources

## BYOD (Bring Your Own Delivery) CDN Resource
resource "ultradns_cdn" "byod_example" {
  account_name = var.account_name
  fqdn         = "www.${ultradns_zone.primary.id}"
  type         = "BYOD"
  name         = "my-byod-cdn"
  description  = "Example BYOD CDN"
  ttl          = 300
  content_type = "static"

  cdn_providers {
    client_cdn_id = "my-cdn-provider"
    cdn_name      = "My CDN Provider"
    description   = "Primary CDN provider"
    fqdn          = "cdn-provider.example.com."
  }

  config_properties = {
    cdnEnablementMap = jsonencode({
      worldDefault = ["my-cdn-provider"]
      asnOverrides = {}
      continents   = {}
    })
    trafficDistribution = jsonencode({
      worldDefault = {
        options = [{
          name            = "equal_distribution"
          description     = "Equal distribution across available CDNs"
          equalWeight     = true
          distribution    = [{ id = "my-cdn-provider" }]
        }]
      }
    })
  }

  preference_properties = {
    availabilityThresholds = jsonencode({
      world = 90
      continents = {
        NA = {
          countries = {
            US = 70
            CA = 70
            MX = 70
          }
          default = 79
        }
      }
    })
    performanceFiltering = jsonencode({
      world = {
        mode              = "relative"
        relativeThreshold = 0.2
      }
      continents = {
        NA = {
          mode              = "relative"
          relativeThreshold = 0.3
        }
      }
    })
    enabledSubdivisionCountries = jsonencode({
      continents = {
        NA = {
          countries = ["US", "CA", "MX"]
        }
      }
    })
  }
}

## SYNTHETIC CDN Resource
resource "ultradns_cdn" "synthetic_example" {
  account_name = var.account_name
  fqdn         = "api.${ultradns_zone.primary.id}"
  type         = "SYNTHETIC"
  name         = "my-synthetic-cdn"

  cdn_providers {
    client_cdn_id = "synthetic-provider"
  }

  config_properties = {
    checksConfig = jsonencode({
      protocol = "HTTP"
      path     = "/"
    })
  }

  preference_properties = {
    availabilityThresholds = jsonencode({
      world = 92
    })
    performanceFiltering = jsonencode({
      world = {
        mode              = "relative"
        relativeThreshold = 0.2
      }
    })
    enabledSubdivisionCountries = jsonencode({
      continents = {}
    })
  }
}

## Data Source: Read a Single CDN by FQDN
data "ultradns_cdn" "example" {
  account_name = var.account_name
  fqdn         = "www.${ultradns_zone.primary.id}"

  depends_on = [ultradns_cdn.byod_example]
}

## Data Source: List All CDNs for an Account
data "ultradns_cdn_list" "all" {
  account_name = var.account_name
  page         = 1
  size         = 100
}

## Output Examples
output "byod_cdn_resource_id" {
  value = ultradns_cdn.byod_example.resource_id
}

output "synthetic_cdn_name" {
  value = ultradns_cdn.synthetic_example.name
}

output "single_cdn_details" {
  value = data.ultradns_cdn.example
}

output "all_cdns_count" {
  value = data.ultradns_cdn_list.all.total_elements
}
