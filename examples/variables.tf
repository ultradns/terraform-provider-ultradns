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
  description = "primary zone name for testing the provider"
}

variable "ultradns_secondary_zone_name" {
  type        = string
  description = "secondary zone name for testing the provider"
}

variable "ultradns_alias_zone_name" {
  type        = string
  description = "alias zone name for testing the provider"
}

variable "ultradns_primary_name_server" {
  type        = string
  description = "primary name server for testing the provider"
}

variable "ultradns_notification_email_address" {
  type        = string
  description = "notification email address for testing the provider"
}
