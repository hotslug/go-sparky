package installer

import "os"

// DeleteDockerArtifacts deletes Dockerfile and docker-compose.yml if they match generated content.
func DeleteDockerArtifacts() error {
	_ = deleteFileIfContentMatches("Dockerfile", dockerfileContents)
	_ = deleteFileIfContentMatches("docker-compose.yml", dockerComposeContents)
	return nil
}

// DeleteVercelConfig deletes vercel.json if it matches generated content.
func DeleteVercelConfig() error {
	return deleteFileIfContentMatches("vercel.json", vercelConfig)
}

// DeleteNetlifyConfig deletes netlify.toml if it matches generated content.
func DeleteNetlifyConfig() error {
	return deleteFileIfContentMatches("netlify.toml", netlifyConfig)
}

func deleteFileIfContentMatches(path, expected string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if string(data) != expected {
		return nil
	}

	return os.Remove(path)
}
