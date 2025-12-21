package templates

import "github.com/hotslug/go-sparky/internal/plan"

const styledMantineApp = `import sparky from './assets/sparky.png';

export default function App() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 to-slate-800 text-slate-100 overflow-hidden">
      <div className="absolute inset-0 bg-[url('data:image/svg+xml,%3Csvg%20width%3D%2260%22%20height%3D%2260%22%20viewBox%3D%220%200%2060%2060%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Cg%20fill%3D%22none%22%20fill-rule%3D%22evenodd%22%3E%3Cg%20fill%3D%22%23ffffff%22%20fill-opacity%3D%220.02%22%3E%3Ccircle%20cx%3D%2230%22%20cy%3D%2230%22%20r%3D%221%22%2F%3E%3C%2Fg%3E%3C%2Fg%3E%3C%2Fsvg%3E')] animate-pulse" />
      <div className="mx-auto flex min-h-screen max-w-5xl flex-col items-center justify-center gap-10 px-6 py-16 md:flex-row md:gap-12 relative z-10">
        <div className="relative group">
          <div className="absolute -inset-2 rounded-[28px] bg-gradient-to-r from-blue-500/20 to-purple-500/20 blur-xl group-hover:from-blue-500/30 group-hover:to-purple-500/30 transition-all duration-500" />
          <div className="flex flex-row items-center justify-center gap-6 h-full w-full animate-fade-in">
            <img
              src={sparky}
              alt="Go Sparky mascot"
              className="h-auto w-[500px] relative object-cover hover:scale-105 transition-transform duration-300"
            />
            <div className="flex flex-col items-start justify-center space-y-2">
              <div className="font-sparky text-[2.5rem] font-semibold tracking-tight text-white sm:text-[4rem] bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent">
                Go-Sparky
              </div>
              <div className="text-[1.5rem] text-slate-300 animate-fade-in-delay">Good boy.</div>
              <div className="inline-flex items-center uppercase gap-2 text-[10px] text-slate-400 bg-slate-800/50 px-2 py-1 rounded-full hover:bg-slate-700/50 transition-colors cursor-default">
                <span className="w-2 h-2 bg-green-400 rounded-full animate-pulse"></span>
                Vite + React + TypeScript
              </div>
              <div className="flex-shrink-0 text-start md:text-left" style={{ width: '400px' }}>
                <p className="mt-4 text-base leading-relaxed text-slate-300">
                  Go-Sparky is a CLI scaffolder that spins up a fast, opinionated
                  React stack with TypeScript, Tailwind, and optional add-ons like
                  Mantine, React Query, ESLint, Prettier, and Husky.
                </p>
                <p className="mt-4 text-base leading-relaxed text-slate-300">
                  It's a great way to get started with a new project.
                </p>
                <button className="mt-6 px-6 py-2 bg-gradient-to-r from-blue-500 to-purple-500 text-white rounded-lg hover:from-blue-600 hover:to-purple-600 transition-all duration-200 shadow-lg hover:shadow-xl hover:shadow-blue-500/25 transform hover:-translate-y-0.5 active:translate-y-0">
                  Get Started
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
`

const zustandApp = `import sparky from './assets/sparky.png';
import { useSparkyStore } from './stores/useSparkyStore';

export default function App() {
  const { treats, hype, mood, addTreat, reset, toggleHype } = useSparkyStore();

  return (
    <div className="min-h-screen bg-linear-gradient-to-br from-slate-900 to-slate-800 text-slate-100 overflow-hidden">
      <div className="absolute inset-0 bg-[url('data:image/svg+xml,%3Csvg%20width%3D%2260%22%20height%3D%2260%22%20viewBox%3D%220%200%2060%2060%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Cg%20fill%3D%22none%22%20fill-rule%3D%22evenodd%22%3E%3Cg%20fill%3D%22%23ffffff%22%20fill-opacity%3D%220.02%22%3E%3Ccircle%20cx%3D%2230%22%20cy%3D%2230%22%20r%3D%221%22%2F%3E%3C%2Fg%3E%3C%2Fg%3E%3C%2Fsvg%3E')] animate-pulse" />
      <div className="mx-auto flex min-h-screen max-w-5xl flex-col items-center justify-center gap-10 px-6 py-16 md:flex-row md:gap-12 relative z-10">
        <div className="relative group">
          <div className="absolute -inset-2 rounded-[28px] bg-linear-to-r from-blue-500/20 to-purple-500/20 blur-xl group-hover:from-blue-500/30 group-hover:to-purple-500/30 transition-all duration-500" />
          <div className="flex flex-row items-center justify-center gap-6 h-full w-full animate-fade-in">
            <img
              src={sparky}
              alt="Go Sparky mascot"
              className="h-auto w-[500px] relative object-cover hover:scale-105 transition-transform duration-300"
            />
            <div className="flex w-full flex-col items-start justify-center space-y-2">
              <div className="font-sparky text-[2.5rem] font-semibold tracking-tight text-white sm:text-[4rem] bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent">
                Go-Sparky
              </div>
              <div className="text-[1.5rem] text-slate-300 animate-fade-in-delay">Good boy.</div>
              <div className="inline-flex items-center uppercase gap-2 text-[10px] text-slate-400 bg-slate-800/50 px-2 py-1 rounded-full hover:bg-slate-700/50 transition-colors cursor-default">
                <span className="w-2 h-2 bg-green-400 rounded-full animate-pulse"></span>
                Vite + React + TypeScript
              </div>
              <div className="flex-shrink-0 text-start md:text-left" style={{ width: '400px' }}>
                <p className="mt-4 text-base leading-relaxed text-slate-300">
                  Go-Sparky is a CLI scaffolder that spins up a fast, opinionated
                  React stack with TypeScript, Tailwind, and optional add-ons like
                  Mantine, React Query, ESLint, Prettier, and Husky.
                </p>
                <p className="mt-4 text-base leading-relaxed text-slate-300">
                  It's a great way to get started with a new project.
                </p>
              </div>
              <div className="mt-6 grid w-full gap-4">
                <div className="rounded-2xl border border-slate-700/60 bg-slate-800/40 px-5 py-4 shadow-xl backdrop-blur">
                  <div className="text-xs font-semibold tracking-[0.12em] text-slate-400 uppercase">Zustand demo</div>
                  <p className="mt-2 text-sm text-slate-300">
                    useSparkyStore keeps a tiny slice of UI state so your UI stays in sync.
                  </p>
                  <div className="mt-4 flex items-center justify-between rounded-xl bg-slate-900/60 px-4 py-3">
                    <div>
                      <div className="text-xs tracking-wide text-slate-400 uppercase">Treats</div>
                      <div className="text-3xl font-semibold text-white">{treats}</div>
                    </div>
                    <div className="flex gap-2">
                      <button
                        className="rounded-lg bg-slate-700 px-3 py-2 text-sm text-slate-100 transition hover:bg-slate-600"
                        type="button"
                        onClick={addTreat}
                      >
                        +1 treat
                      </button>
                      <button
                        className="rounded-lg border border-slate-600 px-3 py-2 text-sm text-slate-200 transition hover:border-slate-400"
                        type="button"
                        onClick={reset}
                      >
                        Reset
                      </button>
                    </div>
                  </div>
                  <div className="mt-3 flex items-center justify-between rounded-xl bg-slate-900/60 px-4 py-3">
                    <div>
                      <div className="text-xs tracking-wide text-slate-400 uppercase">Mood</div>
                      <div className="text-lg font-semibold text-white">{mood}</div>
                      <p className="text-xs text-slate-400">Hype is {hype ? "on" : "off"}.</p>
                    </div>
                    <button
                      className="rounded-lg bg-gradient-to-r from-blue-500 to-purple-500 px-3 py-2 text-sm font-semibold text-white shadow-lg transition hover:from-blue-600 hover:to-purple-600"
                      type="button"
                      onClick={toggleHype}
                    >
                      {hype ? "Calm down" : "Hype up"}
                    </button>
                  </div>
                </div>
                <div className="rounded-2xl border border-blue-500/30 bg-blue-500/5 px-5 py-4 shadow-xl">
                  <div className="text-xs font-semibold tracking-[0.12em] text-blue-200 uppercase">Default stack</div>
                  <ul className="mt-3 space-y-2 text-sm text-slate-200">
                    <li className="flex items-center gap-2">
                      <span className="h-2 w-2 rounded-full bg-emerald-400" />
                      Vite + React + TypeScript
                    </li>
                    <li className="flex items-center gap-2">
                      <span className="h-2 w-2 rounded-full bg-emerald-400" />
                      Tailwind + Prettier + ESLint
                    </li>
                    <li className="flex items-center gap-2">
                      <span className="h-2 w-2 rounded-full bg-emerald-400" />
                      Zustand store + React Query providers
                    </li>
                    <li className="flex items-center gap-2">
                      <span className="h-2 w-2 rounded-full bg-emerald-400" />
                      Husky pre-commit + Framer Motion
                    </li>
                  </ul>
                  <div className="mt-4 rounded-xl bg-blue-500/10 px-3 py-2 text-xs text-blue-100">
                    Edit src/stores/useSparkyStore.ts to shape your own state slices.
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
`

const mantineApp = `import '@mantine/core/styles.css';
import { Badge, Box, Button, Container, Group, Image, Stack, Text, Title } from '@mantine/core';
import sparky from './assets/sparky.png';

export default function App() {
  return (
    <Box
      style={{
        minHeight: '100vh',
        background: 'linear-gradient(135deg, #0f172a 0%, #1e293b 100%)',
        color: '#f8fafc',
        position: 'relative',
        overflow: 'hidden',
      }}
    >
      <Box
        style={{
          position: 'absolute',
          inset: 0,
          backgroundImage:
            'radial-gradient(circle at 20% 20%, rgba(255, 255, 255, 0.06) 0px, transparent 120px), radial-gradient(circle at 80% 20%, rgba(59, 130, 246, 0.14) 0px, transparent 240px)',
          opacity: 0.6,
        }}
      />
      <Container size="lg" py={80} style={{ position: 'relative', zIndex: 1 }}>
        <Group align="center" justify="center" gap="xl" wrap="wrap">
          <Box
            style={{
              position: 'relative',
              padding: '12px',
              borderRadius: '28px',
              background:
                'linear-gradient(135deg, rgba(59, 130, 246, 0.12), rgba(139, 92, 246, 0.12))',
            }}
          >
            <Image
              src={sparky}
              alt="Go Sparky mascot"
              fit="contain"
              style={{
                width: '420px',
                maxWidth: '100%',
                borderRadius: '20px',
              }}
            />
          </Box>
          <Stack gap="xs" style={{ maxWidth: '420px' }}>
            <Title
              order={1}
              className="font-sparky"
              style={{
                fontSize: 'clamp(2.5rem, 4vw, 4rem)',
                letterSpacing: '-0.02em',
                background: 'linear-gradient(90deg, #60a5fa, #a78bfa)',
                WebkitBackgroundClip: 'text',
                color: 'transparent',
              }}
            >
              Go-Sparky
            </Title>
            <Text style={{ color: '#cbd5f5', fontSize: '1.2rem' }}>Good boy.</Text>
            <Badge
              size="xs"
              tt="uppercase"
              style={{
                backgroundColor: 'rgba(30, 41, 59, 0.7)',
                color: '#94a3b8',
                letterSpacing: '0.08em',
              }}
            >
              Vite + React + TypeScript
            </Badge>
            <Text style={{ color: '#cbd5f5' }}>
              Go-Sparky is a CLI scaffolder that spins up a fast, opinionated React stack with
              TypeScript, Tailwind, and optional add-ons like Mantine, React Query, ESLint,
              Prettier, and Husky.
            </Text>
            <Text style={{ color: '#cbd5f5' }}>It's a great way to get started with a new project.</Text>
            <Button
              size="md"
              radius="md"
              variant="gradient"
              gradient={{ from: 'blue', to: 'violet', deg: 90 }}
            >
              Get Started
            </Button>
          </Stack>
        </Group>
      </Container>
    </Box>
  );
}
`

const basicApp = `import sparky from './assets/sparky.png';

export default function App() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 to-slate-800 text-slate-100 overflow-hidden">
      <div className="absolute inset-0 bg-[url('data:image/svg+xml,%3Csvg%20width%3D%2260%22%20height%3D%2260%22%20viewBox%3D%220%200%2060%2060%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Cg%20fill%3D%22none%22%20fill-rule%3D%22evenodd%22%3E%3Cg%20fill%3D%22%23ffffff%22%20fill-opacity%3D%220.02%22%3E%3Ccircle%20cx%3D%2230%22%20cy%3D%2230%22%20r%3D%221%22%2F%3E%3C%2Fg%3E%3C%2Fg%3E%3C%2Fsvg%3E')] animate-pulse" />
      <div className="mx-auto flex min-h-screen max-w-5xl flex-col items-center justify-center gap-10 px-6 py-16 md:flex-row md:gap-12 relative z-10">
        <div className="relative group">
          <div className="absolute -inset-2 rounded-[28px] bg-gradient-to-r from-blue-500/20 to-purple-500/20 blur-xl group-hover:from-blue-500/30 group-hover:to-purple-500/30 transition-all duration-500" />
          <div className="flex flex-row items-center justify-center gap-6 h-full w-full animate-fade-in">
            <img
              src={sparky}
              alt="Go Sparky mascot"
              className="h-auto w-[500px] relative object-cover hover:scale-105 transition-transform duration-300"
            />
            <div className="flex flex-col items-start justify-center space-y-2">
              <div className="font-sparky text-[2.5rem] font-semibold tracking-tight text-white sm:text-[4rem] bg-linear-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent">
                Go-Sparky
              </div>
              <div className="text-[1.5rem] text-slate-300 animate-fade-in-delay">Good boy.</div>
              <div className="inline-flex items-center uppercase gap-2 text-[10px] text-slate-400 bg-slate-800/50 px-2 py-1 rounded-full hover:bg-slate-700/50 transition-colors cursor-default">
                <span className="w-2 h-2 bg-green-400 rounded-full animate-pulse"></span>
                Vite + React + TypeScript
              </div>
              <div className="flex-shrink-0 text-start md:text-left" style={{ width: '400px' }}>
                <p className="mt-4 text-base leading-relaxed text-slate-300">
                  Go-Sparky is a CLI scaffolder that spins up a fast, opinionated
                  React stack with TypeScript, Tailwind, and optional add-ons like
                  Mantine, React Query, ESLint, Prettier, and Husky.
                </p>
                <p className="mt-4 text-base leading-relaxed text-slate-300">
                  It's a great way to get started with a new project.
                </p>
                <button className="mt-6 px-6 py-2 bg-linear-to-r from-blue-500 to-purple-500 text-white rounded-lg hover:from-blue-600 hover:to-purple-600 transition-all duration-200 shadow-lg hover:shadow-xl hover:shadow-blue-500/25 transform hover:-translate-y-0.5 active:translate-y-0">
                  Get Started
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
`

// AppTemplate selects the correct App.tsx template based on the plan.
func AppTemplate(p plan.Plan) string {
	if p.StyledApp && p.Mantine {
		return styledMantineApp
	}

	if p.Zustand {
		return zustandApp
	}

	return basicApp
}
