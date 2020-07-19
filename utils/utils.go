package utils

// GetProviderList returns a list of cloud service providers
func GetProviderList() []string {

	providers := []string{"aws", "azure", "gcp"}
	return providers
}
