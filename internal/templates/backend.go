package templates

// BunBackendServer returns a minimal Bun API server for Vite projects.
func BunBackendServer() string {
	return `const server = Bun.serve({
  port: 3001,
  fetch(req) {
    const url = new URL(req.url);

    if (url.pathname === "/api/hello") {
      return Response.json({ message: "Hello from Bun!" });
    }

    return new Response("Not Found", { status: 404 });
  },
});

console.log("Bun API server running at " + server.url);
`
}

// BunBackendPackageJSON returns a package.json for the Bun backend.
func BunBackendPackageJSON() string {
	return `{
  "name": "sparky-backend",
  "private": true,
  "type": "module",
  "scripts": {
    "dev": "bun --hot index.ts",
    "start": "bun index.ts"
  }
}
`
}
