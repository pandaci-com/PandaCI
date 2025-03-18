<script lang="ts">
	import {
		Button,
		Combobox,
		Heading,
		Field,
		Label,
		Input,
		Card,
		Divider,
		Title,
		Text
	} from '$lib/components';
	import ComboboxInput from '$lib/components/combobox/comboboxInput.svelte';
	import Listbox from '$lib/components/listbox/listbox.svelte';
	import { queries } from '$lib/queries';
	import { createQuery } from '@tanstack/svelte-query';
	import { GithubLogo, GitlabLogo } from 'phosphor-svelte';
	import { PUBLIC_GITHUB_APP_NAME } from '$env/static/public';
	import ListGithubRepos from './listGithubRepos.svelte';
	import Skeleton from '$lib/components/skeleton.svelte';

	let search = $state('');

	const installations = createQuery(() => queries.github.listInstallation({ perPage: 100 }));

	let selectedInstallation = $state<{ label: string; value: string }>();

	$effect(() => {
		if (!selectedInstallation && installations.data?.installations.length) {
			const selectedInstallationId = localStorage.getItem('new-github-selected-install');

			const install =
				installations.data?.installations.find(({ id }) => id === selectedInstallationId) ??
				installations.data?.installations[0];

			selectedInstallation = { value: install.id, label: install.name };
		}
	});

	$effect(() => {
		if (selectedInstallation) {
			localStorage.setItem('new-github-selected-install', selectedInstallation.value);
		}
	});

	const selectedInstallationObj = $derived(
		installations.data?.installations.find(({ id }) => id === selectedInstallation?.value)
	);

	const comboboxData = $derived([
		...(installations.data?.installations.map(({ name, id }) => ({
			label: name,
			value: id
		})) ?? []),
		{
			label: 'Add github account',
			value: 'Add github account',
			onClick: () =>
				window.location.assign(
					`https://github.com/apps/${PUBLIC_GITHUB_APP_NAME}/installations/new`
				)
		}
	]);
</script>

<Title title="New project">
	<Listbox
		selected={{ label: 'Github', value: 'github', icon: GithubLogo }}
		items={[
			{ label: 'Github', value: 'github', icon: GithubLogo },
			{ label: 'GitLab (comming soon)', value: 'gitlab', icon: GitlabLogo, disabled: true }
		]}
	>
		{#snippet item(item)}
			<item.icon class="mr-1" data-slot="icon" />
			<!-- <svelte:element this={item?.icon} class="mr-1" data-slot="icon" /> -->
			<span>{item.label}</span>
		{/snippet}

		{#snippet button(btnProps, item, { icon })}
			<Button {...btnProps} outline class="flex w-full items-center justify-between">
				{#if item}
					<item.icon class="mr-1" data-slot="icon" />
				{/if}
				<span class="pr-6"> {item?.label}</span>
				{@render icon()}
			</Button>
		{/snippet}
	</Listbox>
</Title>

<Card class="mx-auto mt-16 flex w-full max-w-3xl flex-col space-y-10">
	<div class="grid w-full grid-cols-1 gap-8 sm:grid-cols-3 sm:gap-4">
		<Field>
			<Label>Git namespace</Label>
			{#if installations.isLoading || !selectedInstallation}
				<div
					data-slot="control"
					class="h-9 w-full rounded-lg bg-zinc-500/10 dark:bg-white/10"
				></div>
			{:else}
				<Combobox bind:selected={selectedInstallation} items={comboboxData}>
					<ComboboxInput />
				</Combobox>
			{/if}
		</Field>

		<Field class="sm:col-span-2">
			<Label>Search repositories</Label>
			<Input autofocus bind:value={search} />
		</Field>
	</div>

	<Divider />

	{#if selectedInstallationObj !== undefined}
		<ListGithubRepos installation={selectedInstallationObj} {search} />
	{/if}

	{#if !installations.data?.installations.length && !installations.isLoading}
		<div class="flex flex-col justify-center">
			<Heading level={2} class="text-center">No installations found</Heading>
			<Text class="text-center">Please add a github account to continue.</Text>
			<Button
				color="dark/white"
				href={`https://github.com/apps/${PUBLIC_GITHUB_APP_NAME}/installations/new`}
				class="mx-auto mt-4">Add github account</Button
			>
		</div>
	{/if}

	{#if installations.isLoading}
		<div class="flex flex-col justify-center">
			{#each Array.from({ length: 5 }) as _, i (i)}
				<Skeleton class="my-4 h-16 w-full" />
			{/each}
		</div>
	{/if}
</Card>
