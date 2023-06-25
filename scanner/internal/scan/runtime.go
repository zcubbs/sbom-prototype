package scan

import (
	"context"
	"dagger.io/dagger"
	"encoding/json"
	"fmt"
	"os"
)

const (
	sbomFileName = "sbom.json"
)

type Runtime struct {
	Ctx          context.Context
	DaggerClient *dagger.Client
	WorkDir      string
	DockerConfig *dockerConfig
	RegistryInfo *registryInfo
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

// GenerateSBOM generates a software bill of materials for the container image
func (r *Runtime) GenerateSBOM(targetImage string) (string, error) {
	const (
		syftImg = "anchore/syft:latest"
	)
	var (
		syftCmd = []string{targetImage, "--scope",
			"all-layers", "-v", "-o",
			"spdx-json",
			"--file",
			fmt.Sprintf("/%s", sbomFileName),
		}
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
		grypeCmd = []string{fmt.Sprintf("sbom:/%s", sbomFileName), "-v", "-o", "json", "--file", "vuln.json"}
	)

	grype := r.DaggerClient.Container().
		From(grypeImg).
		WithNewFile(fmt.Sprintf("/%s", sbomFileName),
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
	_, err := r.DaggerClient.File(fileID).
		Export(r.Ctx,
			fmt.Sprintf(
				"%s/%s",
				r.WorkDir, sbomFileName),
		)
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
