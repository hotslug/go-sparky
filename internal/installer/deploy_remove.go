package installer

import (
	"os"

	"github.com/hotslug/go-sparky/internal/plan"
)

// DeleteDockerArtifacts deletes Dockerfile and docker-compose.yml if they match generated content.
func DeleteDockerArtifacts() error {
	_ = deleteFileIfContentMatches("Dockerfile", dockerfileViteContents, dockerfileBunContents)
	_ = deleteFileIfContentMatches("docker-compose.yml", dockerComposeViteContents, dockerComposeBunContents)
	return nil
}

// DeleteVercelConfig deletes vercel.json if it matches generated content.
func DeleteVercelConfig() error {
	return deleteFileIfContentMatches(
		"vercel.json",
		VercelConfig(plan.Plan{Bundler: plan.BundlerVite}),
		VercelConfig(plan.Plan{Bundler: plan.BundlerBun}),
	)
}

// DeleteNetlifyConfig deletes netlify.toml if it matches generated content.
func DeleteNetlifyConfig() error {
	return deleteFileIfContentMatches(
		"netlify.toml",
		NetlifyConfig(plan.Plan{Bundler: plan.BundlerVite}),
		NetlifyConfig(plan.Plan{Bundler: plan.BundlerBun}),
	)
}

func deleteFileIfContentMatches(path string, expected ...string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	for _, content := range expected {
		if string(data) == content {
			return os.Remove(path)
		}
	}

	return nil
}
