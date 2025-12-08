<script lang="ts">
    import type { PageProps } from './$types';
    let { data }: PageProps = $props();

    let modfolder = $derived(data.appConfig.modfolder);
    let apikey = $derived(data.appConfig.apikey);
    let apiendpoint = $derived(data.appConfig.apiendpoint);
</script>

<h1>Settings</h1>
<form class="mx-auto w-full max-w-md space-y-4">
    <label class="label">
        <span class="label-text">Mod Download Folder</span>
        <input class="input" type="text" placeholder="Mod Download Folder" bind:value={modfolder} />
    </label>
    <label class="label">
        <span class="label-text">ModIO User API Key</span>
        <input class="input" type="text" placeholder="API Key" bind:value={apikey} />
    </label>
    <label class="label">
        <span class="label-text">ModIO User API Endpoint</span>
        <input class="input" type="text" placeholder="API Endpoint" bind:value={apiendpoint} />
    </label>                
</form>