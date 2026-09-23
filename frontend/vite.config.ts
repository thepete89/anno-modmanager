import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import wails from "@wailsio/runtime/plugins/vite";
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [wails("./src/lib/bindings"), tailwindcss(), sveltekit()],
	server: {
		host: "127.0.0.1",
		port: Number(process.env.WAILS_VITE_PORT) || 9245, 
		strictPort: true
	}
});
