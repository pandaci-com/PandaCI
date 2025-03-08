<div align="center">
<p>No YAML required!</p>
<h1>PandaCI - Code CI/CD your in TypeScript</h1>
<p>PandaCI is a modern CI/CD platform with a simple but extremely powerful styntax learnable in minutes. Avoid hours of documentation and just use a language your team already knows.</p>
</p>
<div align="center">
  <a href="https://pandaci.com/home">Homepage</a>
  .
  <a href="https://pandaci.com/docs">Documentation</a>
</div>
</div>
<br />

![Example PandaCI workflow](https://raw.githubusercontent.com/pandaci-com/pandaci/main/.assets/example-run.png)

## Project structure

Aside from the TypeScript sdk, the project is written in Go with Svelte for the
frontend. The highlights of the project structure are:

- `web` - The Svelte frontend
- `schema` - Our sql schema
- `proto` - The protobuf files and sdks for communication between jobs
- `sdk/ts` - The TypeScript sdk
- `cmd/core` - The main entrypoint for the backend
- `cmd/workflow` - The entrypoint for the workflow sidecar
- `cmd/job` - The entrypoint for the job sidecar

## Contributing

We welcome all contributions! Be it fixing a bug, helping with a feature,
suggesting ideas or even writing about us on your blog.

If you're looking to contribute to a more significant feature, please reach out
to us first and we'll help you get started.

## Disclosing security issues

If you think you've found a security vulnerability, please refrain from posting
it publicly on the forums, the chat, or GitHub. You can find all the info for
responsible disclosure in our
[SECURITY.md](https://github.com/pandaci-com/pandaci/blob/main/SECURITY.md)

tes
