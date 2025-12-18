package installer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWritePostCSSConfig_MatchesDocs(t *testing.T) {
	wd := t.TempDir()
	orig, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	defer func() { _ = os.Chdir(orig) }()

	if err := os.Chdir(wd); err != nil {
		t.Fatalf("chdir temp dir: %v", err)
	}

	if err := WritePostCSSConfig(); err != nil {
		t.Fatalf("WritePostCSSConfig: %v", err)
	}

	got, err := os.ReadFile(filepath.Join(wd, "postcss.config.cjs"))
	if err != nil {
		t.Fatalf("read postcss.config.cjs: %v", err)
	}

	const want = `module.exports = {
  plugins: {
    'postcss-preset-mantine': {},
    'postcss-simple-vars': {
      variables: {
        'mantine-breakpoint-xs': '36em',
        'mantine-breakpoint-sm': '48em',
        'mantine-breakpoint-md': '62em',
        'mantine-breakpoint-lg': '75em',
        'mantine-breakpoint-xl': '88em',
      },
    },
  },
};
`

	if string(got) != want {
		t.Fatalf("postcss.config.cjs mismatch.\nGot:\n%s\n\nWant:\n%s", got, want)
	}
}
