<script lang="ts">
    import type { PageProps } from './$types';
    let { data }: PageProps = $props();

    let configCorrect: boolean = $derived(data.configCorrect);
</script>

{#if configCorrect} 
    <h1>Welcome to SvelteKit</h1>
    <p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>
{:else}
    <div class="flex justify-center items-center">
        <p>Looks like your configuration is not set up correctly! Go to settings and configure the app.</p>
    </div>
{/if}
