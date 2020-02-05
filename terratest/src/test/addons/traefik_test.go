package addons

import (
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

func TestShouldNotContainTraefikIfDisabled(t *testing.T) {
	t.Parallel()
	helmChartPath, err := filepath.Abs(helmChartRelativePath)

	require.NoError(t, err)
	options := &helm.Options{
		SetValues: map[string]string{
			"traefik.enabled": "false",
		},
	}

	helmChart := NewHelmConfigParser(t, options, helmChartPath)

	for _, i := range traefikResources {
		require.False(t, helmChart.contains(SearchResourceOption{
			name: i.name,
			kind: i.kind,
		}))
	}
}

func TestTraefikShouldContainAllResources(t *testing.T) {
	t.Parallel()
	helmChartPath, err := filepath.Abs(helmChartRelativePath)

	require.NoError(t, err)
	options := &helm.Options{
		SetValues: map[string]string{
			"traefik.enabled": "true",
		},
	}

	helmChart := NewHelmConfigParser(t, options, helmChartPath)

	for _, i := range traefikResources {
		require.True(t, helmChart.contains(SearchResourceOption{
			name: i.name,
			kind: i.kind,
		}))
	}
}

var traefikResources = []SearchResourceOption{
	{
		name: "release-name-traefik",
		kind: "ConfigMap",
	},
	{
		name: "release-name-traefik",
		kind: "ServiceAccount",
	},
	{
		name: "release-name-traefik",
		kind: "ClusterRole",
	},
	{
		name: "release-name-traefik",
		kind: "Deployment",
	},
	{
		name: "release-name-traefik",
		kind: "ClusterRoleBinding",
	},
	{
		name: "release-name-traefik",
		kind: "Service",
	},
	{
		name: "release-name-traefik-test",
		kind: "Pod",
	},
	{
		name: "release-name-traefik-test",
		kind: "ConfigMap",
	},
}
