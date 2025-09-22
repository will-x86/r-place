<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import AdminPanel from './AdminPanel.svelte';

	let canvasElement: HTMLCanvasElement;
	let intervalId: ReturnType<typeof setInterval>;
	let isAdminPage = false;

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

	onMount(() => {
		// Simple client-side routing check
		if (window.location.pathname === '/admin') {
			isAdminPage = true;
		}

		const ctx = canvasElement.getContext('2d');
		if (!ctx) return;

		const refreshCanvas = () => {
			const img = new Image();
			img.src = import.meta.env.VITE_API_URL + `/api/canvas?t=${Date.now()}`;

			img.onload = () => {
				canvasElement.width = img.naturalWidth;
				canvasElement.height = img.naturalHeight;
				ctx.drawImage(img, 0, 0);
			};
		};

		refreshCanvas();
		intervalId = setInterval(refreshCanvas, 2000);

		if (isAdminPage) {
			canvasElement.addEventListener('mousedown', handleMouseDown);
			window.addEventListener('mousemove', handleMouseMove);
			window.addEventListener('mouseup', handleMouseUp);
		}
	});

	onDestroy(() => {
		clearInterval(intervalId);
		if (isAdminPage) {
			canvasElement.removeEventListener('mousedown', handleMouseDown);
			window.removeEventListener('mousemove', handleMouseMove);
			window.removeEventListener('mouseup', handleMouseUp);
		}
	});

	function onSectionDeleted() {
		startX = endX = startY = endY = 0;
	}
</script>

<main>
	<h1>{isAdminPage ? 'r/place Admin' : 'r/place Viewer'}</h1>

	<div class="canvas-wrapper">
		<canvas bind:this={canvasElement}></canvas>

		{#if isAdminPage && (isDragging || selectionRect.width > 0)}
			<div
				class="selection-box"
				style="left: {selectionRect.left}px; top: {selectionRect.top}px; width: {selectionRect.width}px; height: {selectionRect.height}px;"
			></div>
		{/if}
	</div>

	{#if isAdminPage}
		<AdminPanel selection={selectionCoords} on:deleted={onSectionDeleted} />
	{/if}
</main>

<style>
	main {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 1rem;
	}
	.canvas-wrapper {
		position: relative;
		line-height: 0; 
	}
	canvas {
		border: 1px solid #333;
		background-color: #f0f0f0;
	}
	.selection-box {
		position: absolute;
		background-color: rgba(255, 0, 0, 0.3);
		border: 1px dashed red;
		pointer-events: none;
	}
</style>
