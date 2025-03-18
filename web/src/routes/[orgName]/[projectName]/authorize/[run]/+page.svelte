<script lang="ts">
	import { page } from '$app/state';
	import Button from '$lib/components/button.svelte';
	import Heading from '$lib/components/heading.svelte';
	import Text from '$lib/components/text/text.svelte';
	import TextLink from '$lib/components/text/textLink.svelte';
	import { queries } from '$lib/queries';
	import { createQuery } from '@tanstack/svelte-query';

	const run = createQuery(() =>
		queries.runs.projectRuns(page.params.orgName, page.params.projectName)._ctx.get(page.params.run)
	);

	// SHow commit, show branch, show pr. Maybe do this with a details component
</script>

<!-- <Title title="Authorize Fork Deployment"></Title> -->

<div class="mt-16 flex flex-col space-y-4">
	<Heading size="lg" class="text-center">Authorize Fork Deployment</Heading>
	<Text size="lg" class="text-center">
		The Deployment for commit <TextLink class="text-on-surface font-semibold"
			>{run.data.gitSha}</TextLink
		> on pull request AlfieJones/alfiejones.dev#1 is awaiting authorization.
	</Text>

	<div class="flex flex-col justify-center">
		<Heading level={2} size="sm" class="text-center"></Heading>
		<Button color="dark/white" class="mx-auto mt-4">Approve</Button>
		<!-- TODO: add a review button once we have a link for the repo -->
	</div>
</div>
