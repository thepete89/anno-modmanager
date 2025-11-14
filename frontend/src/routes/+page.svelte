<script lang="ts">
import { IsConfigLoaded } from "$lib/wailsjs/go/main/App";

let configLoaded: boolean = $derived.by(() => {
    let value = false;
    IsConfigLoaded().then((v) => {
        value = v;
    });
    return value;
});
</script>

{#if configLoaded} 
    <h1>Welcome to SvelteKit</h1>
    <p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>
{:else}
    <p>Looks like your configuration is not set up correctly! Go to settings and configure the app.</p>
{/if}
