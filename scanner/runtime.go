package main

import (
	"context"
	"dagger.io/dagger"
	"encoding/json"
	"fmt"
	"os"
)

type Runtime struct {
	Ctx          context.Context
	DaggerClient *dagger.Client
	WorkDir      string
	DockerConfig *dockerConfig
	RegistryInfo *registryInfo
}

func NewRuntime(ctx context.Context) (*Runtime, error) {
	c, err := getDefaultDaggerClient(ctx)
	if err != nil {
		return nil, err
	}
	return &Runtime{
		Ctx:          ctx,
		DaggerClient: c,
	}, nil
}

type dockerConfig struct {
	Auths map[string]authConfig `json:"auths"`
}

type authConfig struct {
	Auth  string `json:"auth"`
	Email string `json:"email"`
}

type registryInfo struct {
	RegistryServer   string
	RegistryUsername string
	RegistryPassword string
	RegistryEmail    string
}

// GenerateSBOM generates a software bill of materials for the container image
func (r *Runtime) GenerateSBOM(targetImage string) (string, error) {
	const (
		sbomFileName = "sbom.json"
		syftImg      = "anchore/syft:latest"
	)
	var (
		syftCmd = []string{targetImage, "--scope", "all-layers", "-v", "-o", "spdx-json", "--file", "sbom.json"}
	)

	bom := r.DaggerClient.Container().From(syftImg)

	bom = bom.WithWorkdir(r.WorkDir).
		WithExec(syftCmd)

	fileID, err := bom.File(sbomFileName).ID(r.Ctx)
	if err != nil {
		return "", err
	}

	return r.ReadFile(fileID)
}

func (r *Runtime) ScanSBOM(sbomJson string) (string, error) {
	var (
		grypeImg = "anchore/grype:latest"
		grypeCmd = []string{"sbom:/sbom.json", "-v", "-o", "json", "--file", "vuln.json"}
	)

	grype := r.DaggerClient.Container().
		From(grypeImg).
		WithNewFile("/sbom.json",
			dagger.ContainerWithNewFileOpts{
				Contents: sbomJson,
			})

	grype = grype.WithExec(grypeCmd)

	fileID, err := grype.File("vuln.json").ID(r.Ctx)
	if err != nil {
		return "", err
	}

	return r.ReadFile(fileID)
}

func (r *Runtime) ParseVulnerabilityReport(report string) ([]Vulnerability, error) {
	vulns := make([]Vulnerability, 0)
	doc := &Document{}
	err := json.Unmarshal([]byte(report), &doc)
	if err != nil {
		return nil, err
	}

	for _, match := range doc.Matches {
		if match.Vulnerability.ID != "" {
			vulns = append(vulns, match.Vulnerability)
		}
	}

	return vulns, nil
}

func (r *Runtime) ParseVulnerabilityForSeverityLevels(vulns []Vulnerability) (map[string]int, int) {
	levels := make(map[string]int, 0)
	fixes := 0
	for _, vuln := range vulns {
		if levels[vuln.Severity] == 0 {
			levels[vuln.Severity] = 1
		} else {
			levels[vuln.Severity]++
		}
		if len(vuln.Fix.Versions) > 0 {
			fixes++
		}
	}

	return levels, fixes
}

func (r *Runtime) ReadFile(fileID dagger.FileID) (string, error) {
	return r.DaggerClient.File(fileID).Contents(r.Ctx)
}

func (r *Runtime) ExportFile(fileID dagger.FileID) error {
	_, err := r.DaggerClient.File(fileID).Export(r.Ctx, fmt.Sprintf("%s/%s", r.WorkDir, "sbom.json"))
	if err != nil {
		return err
	}
	return nil
}

func getDefaultDaggerClient(ctx context.Context) (*dagger.Client, error) {
	c, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return nil, err
	}
	return c, nil
}

type Vulnerability struct {
	VulnerabilityMetadata
	Fix        Fix        `json:"fix"`
	Advisories []Advisory `json:"advisories"`
}

type Fix struct {
	Versions []string `json:"versions"`
	State    string   `json:"state"`
}

type Advisory struct {
	ID   string `json:"id"`
	Link string `json:"link"`
}

type Coordinates struct {
	RealPath     string `json:"path" cyclonedx:"path"`                 // The path where all path ancestors have no hardlinks / symlinks
	FileSystemID string `json:"layerID,omitempty" cyclonedx:"layerID"` // An ID representing the filesystem. For container images, this is a layer digest. For directories or a root filesystem, this is blank.
}

type VulnerabilityMetadata struct {
	ID          string   `json:"id"`
	DataSource  string   `json:"dataSource"`
	Namespace   string   `json:"namespace,omitempty"`
	Severity    string   `json:"severity,omitempty"`
	URLs        []string `json:"urls"`
	Description string   `json:"description,omitempty"`
	Cvss        []Cvss   `json:"cvss"`
}

type Cvss struct {
	Version        string      `json:"version"`
	Vector         string      `json:"vector"`
	Metrics        CvssMetrics `json:"metrics"`
	VendorMetadata interface{} `json:"vendorMetadata"`
}

type CvssMetrics struct {
	BaseScore           float64  `json:"baseScore"`
	ExploitabilityScore *float64 `json:"exploitabilityScore,omitempty"`
	ImpactScore         *float64 `json:"impactScore,omitempty"`
}

type Document struct {
	Matches        []Match        `json:"matches"`
	IgnoredMatches []IgnoredMatch `json:"ignoredMatches,omitempty"`
	Source         *source        `json:"source"`
	Distro         distribution   `json:"distro"`
	Descriptor     descriptor     `json:"descriptor"`
}

// Match is a single item for the JSON array reported
type Match struct {
	Vulnerability          Vulnerability           `json:"vulnerability"`
	RelatedVulnerabilities []VulnerabilityMetadata `json:"relatedVulnerabilities"`
	MatchDetails           []MatchDetails          `json:"matchDetails"`
	Artifact               Package                 `json:"artifact"`
}

// MatchDetails contains all data that indicates how the result match was found
type MatchDetails struct {
	Type       string      `json:"type"`
	Matcher    string      `json:"matcher"`
	SearchedBy interface{} `json:"searchedBy"`
	Found      interface{} `json:"found"`
}

type Package struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Type         Type              `json:"type"`
	Locations    []Coordinates     `json:"locations"`
	Language     Language          `json:"language"`
	Licenses     []string          `json:"licenses"`
	CPEs         []string          `json:"cpes"`
	PURL         string            `json:"purl"`
	Upstreams    []UpstreamPackage `json:"upstreams"`
	MetadataType MetadataType      `json:"metadataType,omitempty"`
	Metadata     interface{}       `json:"metadata,omitempty"`
}

type UpstreamPackage struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

type IgnoredMatch struct {
	Match
	AppliedIgnoreRules []IgnoreRule `json:"appliedIgnoreRules"`
}

type IgnoreRule struct {
	Vulnerability string             `json:"vulnerability,omitempty"`
	FixState      string             `json:"fix-state,omitempty"`
	Package       *IgnoreRulePackage `json:"package,omitempty"`
}

type IgnoreRulePackage struct {
	Name     string `json:"name,omitempty"`
	Version  string `json:"version,omitempty"`
	Type     string `json:"type,omitempty"`
	Location string `json:"location,omitempty"`
}

type source struct {
	Type   string      `json:"type"`
	Target interface{} `json:"target"`
}

// distribution provides information about a detected Linux distribution.
type distribution struct {
	Name    string   `json:"name"`    // Name of the Linux distribution
	Version string   `json:"version"` // Version of the Linux distribution (major or major.minor version)
	IDLike  []string `json:"idLike"`  // the ID_LIKE field found within the /etc/os-release file
}

type descriptor struct {
	Name                  string      `json:"name"`
	Version               string      `json:"version"`
	Configuration         interface{} `json:"configuration,omitempty"`
	VulnerabilityDBStatus interface{} `json:"db,omitempty"`
}

type Type string

type Language string

type MetadataType string
