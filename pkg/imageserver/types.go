package imageserver

import (
	iiapi "github.com/openshift/image-inspector/pkg/api"
)

// ImageServer abstracts the serving of image information.
type ImageServer interface {
	// ServeImage Serves the image
	ServeImage(meta *iiapi.InspectorMetadata,
		scanReport []byte,
		htmlScanReport []byte) error
}

// ImageServerOptions is used to configure an image server.
type ImageServerOptions struct {
	// ServePath is the root path/port of serving. ex 0.0.0.0:8080
	ServePath string
	// HealthzURL is the relative url of the health check. ex /healthz
	HealthzURL string
	// APIURL is the relative url where the api will be served.  ex /api
	APIURL string
	// APIVersions are the supported API versions.
	APIVersions iiapi.APIVersions
	// MetadataURL is the relative url of the metadata content.  ex /api/v1/metadata
	MetadataURL string
	// ContentURL is the relative url of the content.  ex /api/v1/content/
	ContentURL string
	// ImageServeURL is the location that the image is being served from.
	// NOTE: if the image server supports a chroot the server implementation will perform
	// the chroot based on this URL.
	ImageServeURL string
	// ScanType is the type of the scan that was done on the inspected image
	ScanType string
	// ScanReportURL is the url to publish the scan report
	ScanReportURL string
	// HTMLScanReport wether or not to publish an HTML scan report
	HTMLScanReport bool
	// HTMLScanReportURL url for the scan html report
	HTMLScanReportURL string
}
