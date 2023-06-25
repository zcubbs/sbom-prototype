package config

import (
	"context"
	"github.com/uptrace/uptrace-go/uptrace"
	_ "go.opentelemetry.io/otel"
)

func ConfigTelemetry(ctx context.Context) func() {
	uptrace.ConfigureOpentelemetry(
		uptrace.WithDSN("http://project2_secret_token@localhost:14317/2"),
		uptrace.WithServiceName("sbom-scanner"),
		uptrace.WithServiceVersion("v1.0.0"),
		uptrace.WithDeploymentEnvironment("dev"),
	)

	return func() {
		uptrace.Shutdown(ctx)
	}
}
