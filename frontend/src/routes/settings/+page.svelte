<script lang="ts">
    import type { PageProps } from './$types';
    import { FolderOpenIcon, SaveIcon } from '@lucide/svelte';
    import { SelectAnnoModsFolder, SaveConfigData } from '$lib/wailsjs/go/config/AMMConfig';
    let { data }: PageProps = $props();

    let modfolder = $derived(data.appConfig.modfolder);
    let apikey = $derived(data.appConfig.apikey);
    let apiendpoint = $derived(data.appConfig.apiendpoint);

    const openAnnoModFolder = async () => {
        modfolder = await SelectAnnoModsFolder();
    }

    const saveConfig = async () => {
        await SaveConfigData({ modfolder, apikey, apiendpoint });
    }
</script>

<h1>Settings</h1>
<form class="mx-auto w-full max-w-xl space-y-4">
    <fieldset class="space-y-4">
        <label class="label">
            <span class="label-text">Mod Download Folder</span>
            <div class="input-group grid-cols-[1fr_auto]">
                <input class="input" type="text" placeholder="Mod Download Folder" bind:value={modfolder} />
                <button class="ig-btn preset-filled" title="Open Folder." onclick={openAnnoModFolder}>
                    <FolderOpenIcon size={16} />
                </button>
            </div>
        </label>
        <label class="label">
            <span class="label-text">ModIO User API Key</span>
            <input class="input" type="text" placeholder="API Key" bind:value={apikey} />
        </label>
        <label class="label">
            <span class="label-text">ModIO User API Endpoint</span>
            <input class="input" type="text" placeholder="API Endpoint" bind:value={apiendpoint} />
        </label>                
    </fieldset>
	<fieldset class="flex justify-end">
		<button type="button" class="btn preset-outlined-surface-300-700" onclick={saveConfig}>
            <span>Save</span>
            <SaveIcon size={16} />
        </button>
	</fieldset>    
</form>