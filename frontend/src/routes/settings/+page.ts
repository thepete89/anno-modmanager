import type { PageLoad } from './$types';
import { GetConfigData } from "$lib/wailsjs/go/config/AMMConfig";

export const load: PageLoad = async () => {
    const appConfig = await GetConfigData();
    
    return { appConfig }
};