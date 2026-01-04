package templates

import "strings"

// BunServer returns a Bun server entry file.
func BunServer(includeAPI bool) string {
	apiBlock := ""
	if includeAPI {
		apiBlock = `
    if (url.pathname === "/api/hello") {
      return Response.json({ message: "Hello from Bun!" });
    }
`
	}

	content := `const server = Bun.serve({
  port: 3000,
  async fetch(req) {
    const url = new URL(req.url);
` + apiBlock + `
    const filePath = url.pathname === "/" ? "index.html" : url.pathname.slice(1);
    const file = Bun.file(new URL(filePath, import.meta.url));
    if (await file.exists()) {
      return new Response(file);
    }

    return new Response("Not Found", { status: 404 });
  },
});

console.log("Bun server running at " + server.url);
`

	return strings.TrimLeft(content, "\n")
}
