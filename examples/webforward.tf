resource "ultradns_webforward" "www" {
  zone_name             = "example.com."
  request_to            = "www.example.com"
  default_redirect_to   = "https://example.com/"
  default_forward_type  = "HTTP_301_REDIRECT"
  relative_forward_type = "PARAMETER_AND_PATH"
}
