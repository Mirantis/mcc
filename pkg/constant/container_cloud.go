package constant

// Various constants for the Container Cloud bundle download.
const (
	InternalCdnBaseUrl         = "https://artifactory.mcp.mirantis.net/binary-dev-kaas-virtual"
	InternalEuCdnBaseUrl       = "https://artifactory.mirantis.com/binary-kaas-local"
	PublicCICdnBaseUrl         = "https://binary-dev-kaas-mirantis-com.s3.amazonaws.com"
	PublicCdnBaseUrl           = "https://binary.mirantis.com"
	DefaultReleasesPath        = "releases"
	LatestKaaSRelease          = "1.12.0"
	BootstrapEnvFile           = "bootstrap.env"
	DefaultKaaSReleasesPath    = "kaas"
	DefaultClusterReleasesPath = "cluster"
	DefaultCDNRegion           = "public"
	DefaultTargetDir           = "kaas-bootstrap"
	KaaSReleasesPath           = "releases/kaas"
	ClusterReleasesPath        = "releases/cluster"
)

// Environment variables designations for the Container Cloud bundle download
const (
	TargetDirEnvVar           = "TARGET_DIR"
	KaaSReleasesYamlEnvVar    = "KAAS_RELEASES_YAML"
	ClusterReleasesDirEnvVar  = "CLUSTER_RELEASES_DIR"
	KaaSCDNRegionEnvVar       = "KAAS_CDN_REGION"
	KaaSCDNBaseURLEnvVar      = "KAAS_CDN_BASE_URL"
	KaaSReleasesBaseURLEnvVar = "KAAS_RELEASES_BASE_URL"
)

var LatestClusterReleases = []string{"3.8.0", "4.5.0", "4.6.0", "5.5.0", "5.6.0"}