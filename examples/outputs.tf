output "zone" {
  value = data.ultradns_zone.zone
}

output "record" {
  value = data.ultradns_record.record
}

output "rdpool" {
  value = data.ultradns_rdpool.rdpool
}

output "sfpool" {
  value = data.ultradns_sfpool.sfpool
}

output "slbpool" {
  value = data.ultradns_slbpool.slbpool
}

output "sbpool" {
  value = data.ultradns_sbpool.sbpool
}

output "tcpool" {
  value = data.ultradns_tcpool.tcpool
}

# output "dirpool" {
#   value = data.ultradns_dirpool.dirpool
# }

output "probehttp" {
  value = data.ultradns_probe_http.probehttp
}

output "probeping" {
  value = data.ultradns_probe_ping.probeping
}

output "probedns" {
  value = data.ultradns_probe_dns.probedns
}

output "probetcp" {
  value = data.ultradns_probe_tcp.probetcp
}