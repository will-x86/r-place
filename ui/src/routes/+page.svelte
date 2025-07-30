<script lang="ts">
	import { onMount } from 'svelte';

	/*interface PixelUpdate {
		x: number;
		y: number;
		hex: string;
	}
    */

	const CANVAS_WIDTH: number = 1000;
	const CANVAS_HEIGHT: number = 1000;

	let canvasElement: HTMLCanvasElement;

	onMount(() => {
		const ctx = canvasElement.getContext('2d');
		if (!ctx) return;

		const backgroundImage = new Image();
		backgroundImage.src = '/api/canvas';
		backgroundImage.onload = () => {
			ctx.drawImage(backgroundImage, 0, 0);
		};

		/*const ws = new WebSocket(`ws://${window.location.host}/ws`);

		ws.onmessage = (event: MessageEvent<string>) => {
			const data: PixelUpdate = JSON.parse(event.data);

			ctx.fillStyle = data.hex;
			ctx.fillRect(data.x, data.y, 1, 1);
		};

		ws.onerror = (error) => console.error('WebSocket Error:', error);
        */
	});
</script>

<main>
	<canvas bind:this={canvasElement} width={CANVAS_WIDTH} height={CANVAS_HEIGHT}></canvas>
</main>
