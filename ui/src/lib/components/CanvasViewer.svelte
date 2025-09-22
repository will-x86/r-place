<script lang="ts">
	import { onMount, createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';

	/**
	 * When true, enables click-to-inspect functionality.
	 * @default false
	 */
	export let inspectorEnabled = false;
	export let canvasElement: HTMLCanvasElement;

	const dispatch = createEventDispatcher();
	let intervalId: ReturnType<typeof setInterval>;

	function rgbToHex(r: number, g: number, b: number): string {
		const toHex = (c: number) => ('0' + c.toString(16)).slice(-2);
		return `#${toHex(r)}${toHex(g)}${toHex(b)}`;
	}

	function handleCanvasClick(event: MouseEvent) {
		const ctx = canvasElement.getContext('2d', { willReadFrequently: true });
		if (!ctx) return;

		const rect = canvasElement.getBoundingClientRect();
		
		const clickX = event.clientX - rect.left;
		const clickY = event.clientY - rect.top;
		
		const scaleX = canvasElement.width / rect.width;
		const scaleY = canvasElement.height / rect.height;
		
		const x = Math.floor(clickX * scaleX);
		const y = Math.floor(clickY * scaleY);

		try {
			const pixel = ctx.getImageData(x, y, 1, 1).data;
			const r = pixel[0];
			const g = pixel[1];
			const b = pixel[2];

			dispatch('inspect', {
				x,
				y,
				hex: rgbToHex(r, g, b),
				rgb: { r, g, b }
			});
		} catch (e) {
			console.error('Could not get pixel data. Probably a CORS issue.', e);
		}
	}

	onMount(() => {
		const ctx = canvasElement.getContext('2d');
		if (!ctx) return;

		const refreshCanvas = () => {
			const img = new Image();
			img.crossOrigin = 'Anonymous';
			img.src = (import.meta.env.VITE_API_URL || '') + `/api/canvas?t=${Date.now()}`;

			img.onload = () => {
				const hasSizeChanged =
					canvasElement.width !== img.naturalWidth || canvasElement.height !== img.naturalHeight;

				if (hasSizeChanged) {
					canvasElement.width = img.naturalWidth;
					canvasElement.height = img.naturalHeight;
					dispatch('dimensions', { width: img.naturalWidth, height: img.naturalHeight });
				}

				ctx.drawImage(img, 0, 0);
			};
		};

		refreshCanvas();
		intervalId = setInterval(refreshCanvas, 2000);

		if (browser && inspectorEnabled) {
			canvasElement.addEventListener('click', handleCanvasClick);
		}

		return () => {
			clearInterval(intervalId);
			if (browser && inspectorEnabled) {
				canvasElement.removeEventListener('click', handleCanvasClick);
			}
		};
	});
</script>

<canvas bind:this={canvasElement} class:inspector-enabled={inspectorEnabled}></canvas>

<style>

canvas {
		border: 1px solid #333;
		background-color: #f0f0f0;
		display: block;

		width: 100%;
		height: auto; 

		image-rendering: pixelated;
		image-rendering: -moz-crisp-edges;
		image-rendering: crisp-edges;
	}
	.inspector-enabled {
		cursor: pointer;
	}
</style>
