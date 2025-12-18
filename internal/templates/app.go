package templates

import "github.com/hotslug/go-sparky/internal/plan"

const styledMantineApp = `import '@mantine/core/styles.css';
import { Container, Title, Text, Button, Paper, Box, Stack } from '@mantine/core';

export default function App() {
  return (
    <Box className="min-h-screen bg-gradient-to-br from-gray-900 via-gray-800 to-black flex items-center justify-center p-6">
      <Container size="sm">
        <Paper
          shadow="xl"
          radius="lg"
          p="xl"
          className="backdrop-blur-xl border border-white/10 bg-white/10"
        >
          <Stack align="center" gap="md">
            <Title
              order={1}
              className="text-center text-3xl font-bold tracking-tight text-white drop-shadow"
            >
              Welcome to Your New App ⚡️
            </Title>

            <Text className="text-center text-gray-300 leading-relaxed max-w-md">
              This starter is powered by React, TypeScript, Mantine, and TailwindCSS —
              perfectly configured and beautifully styled. Let’s build something amazing.
            </Text>

            <Button
              size="md"
              radius="md"
              className="mt-4 font-medium"
              variant="white"
              color="dark"
            >
              Get Started
            </Button>
          </Stack>
        </Paper>
      </Container>
    </Box>
  );
}
`

const mantineApp = `import '@mantine/core/styles.css';
import { Button, Container, Stack, Text, Title } from '@mantine/core';

export default function App() {
  return (
    <Container size="sm" py="xl">
      <Stack gap="md">
        <Title order={1}>Welcome to your new app ⚡️</Title>
        <Text>React, TypeScript, and Mantine are ready to go.</Text>
        <Button>Get Started</Button>
      </Stack>
    </Container>
  );
}
`

const basicApp = `import sparky from './assets/sparky.png';

export default function App() {
  return (
    <div className="min-h-screen bg-[#1b120a] text-slate-100">
      <div className="mx-auto flex min-h-screen max-w-5xl flex-col items-center justify-center gap-10 px-6 py-16 md:flex-row md:gap-12">
        <div className="relative">
          <div className="absolute -inset-2 rounded-[28px] bg-gradient-to-br from-amber-400/50 via-orange-500/30 to-transparent blur-xl" />
          <div className="flex flex-row items-center justify-center gap-6 h-full w-full">
            <img
              src={sparky}
              width={288}
              height={288}
              alt="Go Sparky mascot"
              className="relative h-60 w-60 object-cover shadow-2xl sm:h-72 sm:w-72"
            />
            <div className="flex flex-col items-start justify-center">
              <div className="font-sparky text-[4rem] font-semibold tracking-tight text-white sm:text-[8rem]">
                Go-Sparky
              </div>
              <div className="text-[1.5rem] text-white">Good boy.</div>
              <div className="inline-flex items-center uppercase gap-2 text-[10px] text-gray-800">
                Vite + React + TypeScript
              </div>
            </div>
          </div>
        </div>

        <div className="flex-shrink-0 text-start md:text-left" style={{ width: '500px' }}>
          <p className="mt-0 text-base/7 leading-relaxed text-white">
            Go-Sparky is a CLI scaffolder that spins up a fast, opinionated
            React stack with TypeScript, Tailwind, and optional add-ons like
            Mantine, React Query, ESLint, Prettier, and Husky.
            <br />
            <br />
            It's a great way to get started with a new project.
          </p>
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

	if p.Mantine {
		return mantineApp
	}

	return basicApp
}
