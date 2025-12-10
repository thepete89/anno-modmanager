export namespace config {
	
	export class AMMConfigData {
	    modfolder: string;
	    apikey: string;
	    apiendpoint: string;
	
	    static createFrom(source: any = {}) {
	        return new AMMConfigData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.modfolder = source["modfolder"];
	        this.apikey = source["apikey"];
	        this.apiendpoint = source["apiendpoint"];
	    }
	}

}

export namespace events {
	
	export enum AMMEvent {
	    REFRESH_CONFIG = "refresh_config",
	}

}

