package version

const versionPrefix = "terraform-provider-ultradns-"

//go:generate go run gen.go
func GetProviderVersion() string {
	return versionPrefix + version
}
