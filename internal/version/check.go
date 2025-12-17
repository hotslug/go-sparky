package version

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// NodeVersion represents a Node.js version
type NodeVersion struct {
	Major int
	Minor int
	Patch int
}

// npmPackageInfo represents the npm registry response
type npmPackageInfo struct {
	DistTags struct {
		Latest string `json:"latest"`
	} `json:"dist-tags"`
	Versions map[string]struct {
		Engines struct {
			Node string `json:"node"`
		} `json:"engines"`
	} `json:"versions"`
}

// CheckNodeVersion checks if the installed Node.js version meets Vite's requirements.
// It dynamically fetches the required version from npm registry.
func CheckNodeVersion() error {
	version, err := GetNodeVersion()
	if err != nil {
		return fmt.Errorf("failed to check Node.js version: %w", err)
	}

	requirement, err := GetViteNodeRequirement()
	if err != nil {
		// Continue with fallback silently
		requirement = ">=20.19.0 || >=22.12.0" // Fallback
	}

	if !IsVersionSupportedByRequirement(version, requirement) {
		return &NodeVersionError{
			Current:     version,
			Requirement: requirement,
		}
	}

	return nil
}

// NodeVersionError represents a Node.js version compatibility error
type NodeVersionError struct {
	Current     *NodeVersion
	Requirement string
}

// Error implements the error interface with a nicely formatted message
func (e *NodeVersionError) Error() string {
	var sb strings.Builder

	sb.WriteString("\n")
	sb.WriteString("╭─────────────────────────────────────────────────────────────╮\n")
	sb.WriteString("│  ⚠️  Node.js Version Incompatibility                        │\n")
	sb.WriteString("╰─────────────────────────────────────────────────────────────╯\n")
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("  Current version:  v%d.%d.%d\n", e.Current.Major, e.Current.Minor, e.Current.Patch))
	sb.WriteString(fmt.Sprintf("  Required:         %s\n", e.Requirement))
	sb.WriteString("\n")
	sb.WriteString("  Vite requires a newer Node.js version to run properly.\n")
	sb.WriteString("\n")
	sb.WriteString("  📦 Upgrade Node.js:\n")
	sb.WriteString("\n")

	// Provide specific upgrade instructions based on requirement
	if strings.Contains(e.Requirement, "20.19") {
		sb.WriteString("    # Using nvm (recommended)\n")
		sb.WriteString("    nvm install 20 && nvm use 20\n")
		sb.WriteString("\n")
		sb.WriteString("    # Using Homebrew\n")
		sb.WriteString("    brew upgrade node@20\n")
	} else {
		sb.WriteString("    # Using nvm (recommended)\n")
		sb.WriteString("    nvm install --lts && nvm use --lts\n")
		sb.WriteString("\n")
		sb.WriteString("    # Using Homebrew\n")
		sb.WriteString("    brew upgrade node\n")
	}

	sb.WriteString("\n")
	sb.WriteString("  After upgrading, verify with: node --version\n")
	sb.WriteString("\n")

	return sb.String()
}

// GetViteNodeRequirement fetches the Node.js version requirement from Vite's npm package
func GetViteNodeRequirement() (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("https://registry.npmjs.org/vite/latest")
	if err != nil {
		return "", fmt.Errorf("failed to fetch vite package info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("npm registry returned status %d", resp.StatusCode)
	}

	var pkgInfo struct {
		Engines struct {
			Node string `json:"node"`
		} `json:"engines"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&pkgInfo); err != nil {
		return "", fmt.Errorf("failed to decode package info: %w", err)
	}

	if pkgInfo.Engines.Node == "" {
		return "", fmt.Errorf("no node engine requirement found")
	}

	return pkgInfo.Engines.Node, nil
}

// GetNodeVersion returns the currently installed Node.js version
func GetNodeVersion() (*NodeVersion, error) {
	cmd := exec.Command("node", "--version")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("node command not found or failed to execute: %w", err)
	}

	versionStr := strings.TrimSpace(string(output))
	// Remove 'v' prefix (e.g., "v20.9.0" -> "20.9.0")
	versionStr = strings.TrimPrefix(versionStr, "v")

	// Parse version string
	re := regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)`)
	matches := re.FindStringSubmatch(versionStr)
	if len(matches) != 4 {
		return nil, fmt.Errorf("invalid version format: %s", versionStr)
	}

	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	patch, _ := strconv.Atoi(matches[3])

	return &NodeVersion{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}

// IsVersionSupportedByRequirement checks if a Node.js version satisfies a semver requirement string
// Supports formats like: "^18.0.0 || >=20", ">=20.19.0 || >=22.12.0"
func IsVersionSupportedByRequirement(v *NodeVersion, requirement string) bool {
	// Split by OR (||)
	parts := strings.Split(requirement, "||")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if matchesRequirement(v, part) {
			return true
		}
	}

	return false
}

// matchesRequirement checks if a version matches a single requirement
func matchesRequirement(v *NodeVersion, req string) bool {
	// Handle >= operator
	if strings.HasPrefix(req, ">=") {
		minVersion := parseVersionString(strings.TrimPrefix(req, ">="))
		if minVersion == nil {
			return false
		}
		return compareVersions(v, minVersion) >= 0
	}

	// Handle > operator
	if strings.HasPrefix(req, ">") {
		minVersion := parseVersionString(strings.TrimPrefix(req, ">"))
		if minVersion == nil {
			return false
		}
		return compareVersions(v, minVersion) > 0
	}

	// Handle ^ operator (compatible with)
	if strings.HasPrefix(req, "^") {
		baseVersion := parseVersionString(strings.TrimPrefix(req, "^"))
		if baseVersion == nil {
			return false
		}
		// ^X.Y.Z means >=X.Y.Z and <(X+1).0.0
		if v.Major != baseVersion.Major {
			return false
		}
		return compareVersions(v, baseVersion) >= 0
	}

	// Handle ~ operator (approximately equivalent)
	if strings.HasPrefix(req, "~") {
		baseVersion := parseVersionString(strings.TrimPrefix(req, "~"))
		if baseVersion == nil {
			return false
		}
		// ~X.Y.Z means >=X.Y.Z and <X.(Y+1).0
		if v.Major != baseVersion.Major || v.Minor != baseVersion.Minor {
			return false
		}
		return compareVersions(v, baseVersion) >= 0
	}

	// Handle exact version
	exactVersion := parseVersionString(req)
	if exactVersion == nil {
		return false
	}
	return compareVersions(v, exactVersion) == 0
}

// parseVersionString parses a version string like "20.19.0" into a NodeVersion
func parseVersionString(s string) *NodeVersion {
	s = strings.TrimSpace(s)
	parts := strings.Split(s, ".")

	if len(parts) < 1 {
		return nil
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil
	}

	minor := 0
	if len(parts) > 1 {
		minor, _ = strconv.Atoi(parts[1])
	}

	patch := 0
	if len(parts) > 2 {
		patch, _ = strconv.Atoi(parts[2])
	}

	return &NodeVersion{Major: major, Minor: minor, Patch: patch}
}

// compareVersions compares two versions
// Returns: -1 if v1 < v2, 0 if v1 == v2, 1 if v1 > v2
func compareVersions(v1, v2 *NodeVersion) int {
	if v1.Major != v2.Major {
		if v1.Major < v2.Major {
			return -1
		}
		return 1
	}

	if v1.Minor != v2.Minor {
		if v1.Minor < v2.Minor {
			return -1
		}
		return 1
	}

	if v1.Patch != v2.Patch {
		if v1.Patch < v2.Patch {
			return -1
		}
		return 1
	}

	return 0
}
