// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlschema

type CertificateSigningRequestInfo struct {
	ManagementPlaneInfo *ManagementPlaneInfo `json:"managementPlaneInfo"`
	Subject             string               `json:"subject"`
	KeyAlgorithm        string               `json:"keyAlgorithm"`
}

type CertificationResult struct {
	Certificate       string `json:"certificate"`
	CaCertificate     string `json:"caCertificate"`
	ClientCertificate string `json:"clientCertificate"`
}

type ManagementPlaneInfo struct {
	DirectorURL string `json:"directorURL"`
}

type Token struct {
	Token string `json:"token"`
}
