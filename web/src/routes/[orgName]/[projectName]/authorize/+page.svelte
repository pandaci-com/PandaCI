<script lang="ts">
	import Button from '$lib/components/button.svelte';
	import Card from '$lib/components/card.svelte';
	import Divider from '$lib/components/divider.svelte';
	import Heading from '$lib/components/heading.svelte';
	import Text from '$lib/components/text/text.svelte';
	import Title from '$lib/components/title.svelte';

	// Get commit url
	// Get pull request url
</script>

<!-- <Title title="Authorize Fork Deployment"></Title> -->

<div class="mt-16 flex flex-col space-y-4">
	<Heading size="lg" class="text-center">Authorize Fork Deployment</Heading>
	<Text size="lg" class="text-center">
		The Deployment for commit <span class="text-on-surface font-semibold">5c47dd2</span> on pull request
		AlfieJones/alfiejones.dev#1 is awaiting authorization.
	</Text>

	<div class="flex flex-col justify-center">
		<Heading level={2} size="sm" class="text-center"></Heading>
		<Button color="dark/white" class="mx-auto mt-4">Approve</Button>
		<!-- TODO: add a review button once we have a link for the repo -->
	</div>
</div>
