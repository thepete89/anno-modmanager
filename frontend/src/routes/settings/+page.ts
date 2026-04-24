import type { PageLoad } from './$types';
import { AMMConfig } from "$lib/bindings/anno-modmanager/core/config";

export const load: PageLoad = async () => {
    const appConfig = await AMMConfig.GetConfigData();
    
    return { appConfig }
};