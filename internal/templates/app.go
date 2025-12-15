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

const basicApp = `export default function App() {
  return (
    <div className="app">
      <h1>Hello from React + Vite + TypeScript ⚡️</h1>
    </div>
  );
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
