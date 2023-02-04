import { hasJwtExpired } from "./bonhart-utils";

export const standard_generic_stores = {
    PRODUCTS: 1,
    POSITIONING: 2,
}

const STORE_NAME = "ITZ-EXPERT-STORE"

class BonhartStorage {
    constructor() {
        this._token = "";
        this.state = {};
        this.generic = {}; // anything can be read/written here
        /* 

            main use case for this is to store data between pages. generic strcuture is:

            generic = {
                "store_name": {
                    "key": "value"
                }
            }
        */

        this.loadData();
    }

    
    /*=============================================
    =            Propertys            =
    =============================================*/
    
    get Token() {
        return this._token;
    }

    set Token(token) {
        this._token = token;
        this.saveData();
    }
    
    
    /*=====  End of Propertys  ======*/
    
    loadData = () => {
        let data = localStorage.getItem(STORE_NAME);
        if (data === null) {
            return;
        }
        
        data = JSON.parse(data);

        if (data.state) {
            this.state = data.state;
        }
        if (data.token) {
            if (!hasJwtExpired(data.token)) {
                this._token = data.token;
            } else {
                console.log("Token expired");
                this.saveData();
            }
        }
        if (data.generic) {
            this.generic = data.generic;
        }
    }

    readGeneric = (store_name, key) => {
        if (this.generic[store_name] && this.generic[store_name][key]) {
            return this.generic[store_name][key];
        }
        return null;
    }

    saveData = () => {
        let data = {
            state: this.state,
            token: this._token,
            generic: this.generic
        };

        localStorage.setItem(STORE_NAME, JSON.stringify(data));
        console.log(`${STORE_NAME} saved`);
    }   

    writeGeneric = (store, key, value) => {
        if (!this.generic[store]) {
            this.generic[store] = {};
        }
        this.generic[store][key] = value;
        this.saveData();
    }
}

const bonhart_storage = new BonhartStorage();
window.bonhart_storage = bonhart_storage;

export default bonhart_storage;