// Note: This is normally "jsr:@pandaci/workflow", but we're just using the local source here.
import { $, docker, cache } from "@pandaci/workflow";
import { env, initPnpm } from "../utils.ts";

export function testFrontendTask() {
  return docker(
    "node:22",
    { name: "Linting and testing frontend" },
    async () => {

      await cache({
        keys: ["web/package.json", "web/pnpm-lock.yaml"],
      })


      await initPnpm();
      await $`pnpm -F web lint`;
      await $`pnpm -F web sync`; // TODO - we should run svelte-kit check here too
    },
  );
}

export function deployDevFrontendTask() {
  return docker("node:22", { name: "Deploying frontend" }, async () => {
    await initPnpm();

    await $`pnpm -F web build`.env({
      PUBLIC_GITHUB_APP_NAME: "pandaci-com-dev",
      PUBLIC_API_URL: `https://${env.PANDACI_BRANCH}.api.dev.pandaci.com`,
      PUBLIC_KRATOS_URL: "https://auth.dev.pandaci.com",
      PUBLIC_STAGE: "dev",
    });

    const res =
      await $`pnpx wrangler pages deploy build --branch=${env.PANDACI_BRANCH} --commit-hash=${env.PANDACI_COMMIT_SHA} --project-name=pandaci-app-dev`
        .env({
          CLOUDFLARE_API_TOKEN: env.CLOUDFLARE_API_TOKEN,
          CLOUDFLARE_ACCOUNT_ID: env.CLOUDFLARE_ACCOUNT_ID,
        })
        .cwd("web")
        .text();

    // We have a worker that rewrites our dev URLs to the correct deployment
    // we store the latest deployment for each sha in a KV store
    const deploymentId =
      /https:\/\/([a-zA-Z0-9-]+)\.pandaci-app-dev\.pages\.dev/.exec(res)?.[1];

    await $`pnpx wrangler kv key put ${env.PANDACI_BRANCH} ${deploymentId} --remote --namespace-id=${env.CLOUDFLARE_KV_NAMESPACE_ID}`
      .env({
        CLOUDFLARE_API_TOKEN: env.CLOUDFLARE_API_TOKEN,
        CLOUDFLARE_ACCOUNT_ID: env.CLOUDFLARE_ACCOUNT_ID,
      });
  });
}

export function deployProdFrontendTask() {
  return docker("node:22", { name: "Deploying frontend" }, async () => {
    await initPnpm();

    await $`pnpm -F web build`.env({
      PUBLIC_GITHUB_APP_NAME: "pandaci-com",
      PUBLIC_API_URL: "https://api.pandaci.com",
      PUBLIC_KRATOS_URL: "https://auth.pandaci.com",
      PUBLIC_STAGE: "prod",
    });

    await $`pnpx wrangler pages deploy build --branch=${env.PANDACI_BRANCH} --commit-hash=${env.PANDACI_COMMIT_SHA} --project-name=pandaci-app-prod`
      .env({
        CLOUDFLARE_API_TOKEN: env.CLOUDFLARE_API_TOKEN,
        CLOUDFLARE_ACCOUNT_ID: env.CLOUDFLARE_ACCOUNT_ID,
      })
      .cwd("web");
  });
}
