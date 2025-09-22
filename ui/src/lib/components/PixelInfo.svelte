<script lang="ts">
	export let dimensions: { width: number; height: number } | null = null;
	export let selectedPixel: {
		x: number;
		y: number;
		hex: string;
		rgb: { r: number; g: number; b: number };
	} | null = null;
</script>

<div class="info-panel">
	<div class="stats-section">
		<h4>Canvas Stats</h4>
		{#if dimensions}
			<p><strong>Dimensions:</strong> {dimensions.width} x {dimensions.height} px</p>
			<p><strong>Total Pixels:</strong> {(dimensions.width * dimensions.height).toLocaleString()}</p>
		{:else}
			<p>Loading stats...</p>
		{/if}
	</div>

	<div class="inspector-section">
		<h4>Pixel Inspector</h4>
		{#if selectedPixel}
			<div class="pixel-details">
				<div class="color-swatch" style="background-color: {selectedPixel.hex};"></div>
				<dl>
					<div>
						<dt>Coordinates:</dt>
						<dd>({selectedPixel.x}, {selectedPixel.y})</dd>
					</div>
					<div>
						<dt>HEX:</dt>
						<dd>{selectedPixel.hex}</dd>
					</div>
					<div>
						<dt>RGB:</dt>
						<dd>
							rgb({selectedPixel.rgb.r}, {selectedPixel.rgb.g}, {selectedPixel.rgb.b})
						</dd>
					</div>
				</dl>
			</div>
		{:else}
			<p class="prompt">Click the canvas to inspect a pixel.</p>
		{/if}
	</div>
</div>

<style>
	.info-panel {
		border: 1px solid #ccc;
		border-radius: 8px;
		padding: 1rem;
		width: 400px;
		max-width: 90vw;
		background-color: #f9f9f9;
		display: flex;
		flex-direction: column;
		gap: 1.5rem;
	}
	h4 {
		margin-top: 0;
		margin-bottom: 0.5rem;
		border-bottom: 1px solid #eee;
		padding-bottom: 0.5rem;
	}
	.pixel-details {
		display: flex;
		align-items: center;
		gap: 1rem;
	}
	.color-swatch {
		width: 50px;
		height: 50px;
		border-radius: 4px;
		border: 1px solid #ccc;
		flex-shrink: 0;
	}
	dl {
		margin: 0;
		font-family: monospace;
		font-size: 0.9em;
	}
	dl div {
		display: flex;
		gap: 0.5rem;
	}
	dt {
		font-weight: bold;
		color: #555;
	}
	dd {
		margin: 0;
		color: #333;
	}
	.prompt {
		color: #777;
	}
</style>
