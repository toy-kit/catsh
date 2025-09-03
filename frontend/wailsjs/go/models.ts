export namespace types {
	
	export class Config {
	    theme: string;
	    locale: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.theme = source["theme"];
	        this.locale = source["locale"];
	    }
	}
	export class WailsConfigInfo {
	    companyName: string;
	    productName: string;
	    productVersion: string;
	    copyright: string;
	    comments: string;
	
	    static createFrom(source: any = {}) {
	        return new WailsConfigInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.companyName = source["companyName"];
	        this.productName = source["productName"];
	        this.productVersion = source["productVersion"];
	        this.copyright = source["copyright"];
	        this.comments = source["comments"];
	    }
	}
	export class WailsConfigAuthor {
	    name: string;
	    email: string;
	
	    static createFrom(source: any = {}) {
	        return new WailsConfigAuthor(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.email = source["email"];
	    }
	}
	export class WailsConfig {
	    name: string;
	    outputfilename: string;
	    repository: string;
	    homepage: string;
	    author: WailsConfigAuthor;
	    info: WailsConfigInfo;
	
	    static createFrom(source: any = {}) {
	        return new WailsConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.outputfilename = source["outputfilename"];
	        this.repository = source["repository"];
	        this.homepage = source["homepage"];
	        this.author = this.convertValues(source["author"], WailsConfigAuthor);
	        this.info = this.convertValues(source["info"], WailsConfigInfo);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class LoadData {
	    wails_config: WailsConfig;
	    config: Config;
	    locales: Record<string, any>;
	
	    static createFrom(source: any = {}) {
	        return new LoadData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.wails_config = this.convertValues(source["wails_config"], WailsConfig);
	        this.config = this.convertValues(source["config"], Config);
	        this.locales = source["locales"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	

}

