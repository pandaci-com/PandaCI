// Note: This is normally "jsr:@pandaci/workflow", but we're just using the local source here.
import { $, Config, docker } from "@pandaci/workflow";
import { env, initFly } from "./utils.ts";

export const config: Config = {
  name: "Cleanup",
  on: {
    push: {
      branches: [],
    },
    pr: {
      events: ["closed"],
      targetBranches: ["main"],
    },
  },
};

await docker("node:22", { name: "Cleanup fly" }, async () => {
  await initFly();

  const flyEnv = {
    FLY_ACCESS_TOKEN: env.FLY_ACCESS_TOKEN_DEV,
    FLY_APP: `pandaci-core-prev-${env.PANDACI_BRANCH}-${env.PANDACI_PR_NUMBER}`,
  };

  $`flyctl apps delete ${flyEnv.FLY_APP} --yes`.env(flyEnv);

  $`npx wrangler kv key delete ${env.PANDACI_BRANCH} --remote --namespace-id=${env.CLOUDFLARE_KV_NAMESPACE_ID}`
    .env({
      CLOUDFLARE_API_TOKEN: env.CLOUDFLARE_API_TOKEN,
      CLOUDFLARE_ACCOUNT_ID: env.CLOUDFLARE_ACCOUNT_ID,
    });

  $`npx wrangler kv key delete ${env.PANDACI_BRANCH} --remote --namespace-id=${env.CLOUDFLARE_API_KV_ID}`
    .env({
      CLOUDFLARE_API_TOKEN: env.CLOUDFLARE_API_TOKEN,
      CLOUDFLARE_ACCOUNT_ID: env.CLOUDFLARE_ACCOUNT_ID,
    });
});
