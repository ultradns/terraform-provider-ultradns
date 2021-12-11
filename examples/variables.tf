variable "ultradns_username" {
  type        = string
  description = "UltraDNS username"
}

variable "ultradns_password" {
  type        = string
  sensitive   = true
  description = "UltraDNS password"
}

variable "ultradns_host_url" {
  type        = string
  description = "UltraDNS hosturl"
}

variable "ultradns_primary_zone_name" {
  type        = string
  description = "zone name for testing the provider"
}