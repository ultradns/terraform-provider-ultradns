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