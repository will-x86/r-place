<script lang="ts">
	import { onMount} from 'svelte';
	import { browser } from '$app/environment'; 
	import CanvasViewer from '$lib/components/CanvasViewer.svelte';
	import AdminPanel from '$lib/components/AdminPanel.svelte';

	let canvasElement: HTMLCanvasElement;

	let isDragging = false;
	let startX = 0;
	let startY = 0;
	let endX = 0;
	let endY = 0;

	$: selectionRect = {
		left: Math.min(startX, endX),
		top: Math.min(startY, endY),
		width: Math.abs(endX - startX),
		height: Math.abs(endY - startY)
	};

	$: selectionCoords = {
		x1: Math.floor(startX),
		y1: Math.floor(startY),
		x2: Math.floor(endX),
		y2: Math.floor(endY)
	};

	const handleMouseDown = (e: MouseEvent) => {
		const rect = canvasElement.getBoundingClientRect();
		startX = e.clientX - rect.left;
		startY = e.clientY - rect.top;
		endX = startX;
		endY = startY;
		isDragging = true;
	};

	const handleMouseMove = (e: MouseEvent) => {
		if (!isDragging) return;
		const rect = canvasElement.getBoundingClientRect();
		endX = e.clientX - rect.left;
		endY = e.clientY - rect.top;
	};

	const handleMouseUp = () => {
		isDragging = false;
	};

	function onSectionDeleted() {
		startX = endX = startY = endY = 0;
	}

	onMount(() => {
		if (browser) {
			if (canvasElement) {
				canvasElement.addEventListener('mousedown', handleMouseDown);
			}
			window.addEventListener('mousemove', handleMouseMove);
			window.addEventListener('mouseup', handleMouseUp);
		}

		return () => {
			if (browser) {
				if (canvasElement) {
					canvasElement.removeEventListener('mousedown', handleMouseDown);
				}
				window.removeEventListener('mousemove', handleMouseMove);
				window.removeEventListener('mouseup', handleMouseUp);
			}
		};
	});


</script>

<main>
	<h1>r/place Admin</h1>

	<div class="canvas-wrapper">
		<CanvasViewer bind:canvasElement />

		{#if isDragging || selectionRect.width > 0}
			<div
				class="selection-box"
				style="left: {selectionRect.left}px; top: {selectionRect.top}px; width: {selectionRect.width}px; height: {selectionRect.height}px;"
			></div>
		{/if}
	</div>

	<AdminPanel selection={selectionCoords} on:deleted={onSectionDeleted} />
</main>

<style>
	main {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1rem;
		padding: 1rem;
	}

	.canvas-wrapper {
		position: relative;
		cursor: crosshair;
	}

	.selection-box {
		position: absolute;
		background-color: rgba(255, 0, 0, 0.3);
		border: 1px dashed red;
		pointer-events: none;
	}
</style>
