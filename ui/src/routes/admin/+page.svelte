<script lang="ts">
	import { onMount } from 'svelte';
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

	$: selectionCanvasCoords = (() => {
		if (!canvasElement) {
			return { x1: 0, y1: 0, x2: 0, y2: 0 };
		}
		const rect = canvasElement.getBoundingClientRect();

		if (rect.width === 0 || rect.height === 0) {
			return { x1: 0, y1: 0, x2: 0, y2: 0 };
		}

		const scaleX = canvasElement.width / rect.width;
		const scaleY = canvasElement.height / rect.height;

		return {
			x1: Math.floor(startX * scaleX),
			y1: Math.floor(startY * scaleY),
			x2: Math.floor(endX * scaleX),
			y2: Math.floor(endY * scaleY)
		};
	})();


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

	<AdminPanel selection={selectionCanvasCoords} on:deleted={onSectionDeleted} />
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

		width: 40vw; 
        max-width: 1000px; 
	}

    .canvas-wrapper :global(canvas) {
        display: block; 
        width: 100%; 
        height: auto;
    }

	.selection-box {
		position: absolute;
		background-color: rgba(255, 0, 0, 0.3);
		border: 1px dashed red;
		pointer-events: none;
	}
</style>
