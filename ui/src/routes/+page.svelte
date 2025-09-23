<script lang="ts">
	import CanvasViewer from '$lib/components/CanvasViewer.svelte';
	import PixelInfo from '$lib/components/PixelInfo.svelte';

	let canvasDimensions: { width: number; height: number } | null = null;
	let inspectedPixel: {
		x: number;
		y: number;
		hex: string;
		rgb: { r: number; g: number; b: number };
	} | null = null;

	function handleDimensions(event: CustomEvent<{ width: number; height: number }>) {
		canvasDimensions = event.detail;
	}

	function handleInspect(event: CustomEvent) {
		inspectedPixel = event.detail;
	}
</script>

<main>
	<h1>pal/place Viewer</h1>

	<CanvasViewer
		inspectorEnabled={true}
		on:dimensions={handleDimensions}
		on:inspect={handleInspect}
	/>

	<PixelInfo dimensions={canvasDimensions} selectedPixel={inspectedPixel} />
</main>

<style>
main {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1.5rem;
		padding: 1rem;

		width: 40%;
		margin: 0 auto; 
	}
</style>
