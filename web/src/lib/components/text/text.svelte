<script lang="ts">
	import { clsx } from 'clsx';
	import type { SvelteHTMLElements } from 'svelte/elements';

	const colors = {
		default: 'text-on-surface-variant',
		emphasis: 'text-on-surface',
		error: 'text-red-500 dark:text-red-400',
		success: 'text-green-500 dark:text-green-400',
		warning: 'text-yellow-500 dark:text-yellow-400'
	} as const;

	type TextProps = SvelteHTMLElements['p'] & {
		variant?: keyof typeof colors;
		size?: 'lg' | 'base' | 'sm';
	};

	const {
		children,
		class: className,
		size = 'base',
		variant = 'default',
		...props
	}: TextProps = $props();
</script>

<p
	data-slot="text"
	{...props}
	class={clsx(
		className,
		colors[variant],
		size === 'lg' && 'text-lg/6 sm:text-base/6',
		size === 'base' && 'text-base/6 sm:text-sm/6',
		size === 'sm' && 'text-sm/6 sm:text-xs/6'
	)}
>
	{#if children}
		{@render children()}
	{/if}
</p>
