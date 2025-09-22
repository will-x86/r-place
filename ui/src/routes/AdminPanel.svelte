<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let selection: { x1: number; y1: number; x2: number; y2: number };

	const dispatch = createEventDispatcher();

	let apiKey = '';
	let statusMessage = 'Enter API key and drag on canvas to select a section.';

	$: hasSelection = selection.x1 !== selection.x2 || selection.y1 !== selection.y2;

	async function deleteSection() {
		if (!hasSelection) {
			statusMessage = 'Error: No section selected.';
			return;
		}
		if (!apiKey) {
			statusMessage = 'Error: API Key is required.';
			return;
		}

		statusMessage = 'Deleting section...';
		try {
			const res = await fetch(import.meta.env.VITE_API_URL + '/api/admin/section', {
				method: 'DELETE',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${apiKey}`
				},
				body: JSON.stringify({
					x1: Math.min(selection.x1, selection.x2),
					y1: Math.min(selection.y1, selection.y2),
					x2: Math.max(selection.x1, selection.x2),
					y2: Math.max(selection.y1, selection.y2)
				})
			});

			if (res.status === 401) throw new Error('Unauthorized. Check API Key.');
			if (!res.ok) throw new Error(`Server error: ${res.statusText || res.status}`);

			statusMessage = 'Section cleared successfully.';
			dispatch('deleted');
		} catch (error) {
			statusMessage = `Error: ${error instanceof Error ? error.message : 'Unknown error'}`;
		}
	}
</script>

<div class="admin-panel">
	<h3>Admin Panel</h3>
	<div class="controls">
		<input type="password" placeholder="API Key" bind:value={apiKey} />
		<button on:click={deleteSection} disabled={!hasSelection}>Delete Section</button>
	</div>
	<p>{statusMessage}</p>
</div>

<style>
	.admin-panel {
		border: 2px solid red;
		padding: 1rem;
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
		width: 500px;
	}
	.controls {
		display: flex;
		gap: 0.5rem;
	}
</style>
