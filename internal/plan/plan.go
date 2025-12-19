package plan

// Plan captures the requested project configuration derived from CLI flags.
type Plan struct {
	Name       string
	Mantine    bool
	Tailwind   bool
	ReactQuery bool
	Eslint     bool
	Prettier   bool
	Husky      bool
	StyledApp  bool
	Framer     bool
	Docker     bool
	Vercel     bool
	Netlify    bool
}
