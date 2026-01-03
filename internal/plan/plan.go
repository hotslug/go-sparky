package plan

// BundlerType tracks the chosen frontend bundler.
type BundlerType string

const (
	BundlerVite BundlerType = "vite"
	BundlerBun  BundlerType = "bun"
)

// Plan captures the requested project configuration derived from CLI flags.
type Plan struct {
	Name       string
	Bundler    BundlerType
	Mantine    bool
	Tailwind   bool
	ReactQuery bool
	Zustand    bool
	Eslint     bool
	Prettier   bool
	Husky      bool
	StyledApp  bool
	Framer     bool
	Docker     bool
	Vercel     bool
	Netlify    bool
	Storybook  bool
}

// IsVite returns true when the plan targets Vite.
func (p Plan) IsVite() bool { return p.Bundler == BundlerVite }

// IsBun returns true when the plan targets Bun.
func (p Plan) IsBun() bool { return p.Bundler == BundlerBun }

// PackageManager returns the package manager for the bundler.
func (p Plan) PackageManager() string {
	if p.IsBun() {
		return "bun"
	}
	return "pnpm"
}
