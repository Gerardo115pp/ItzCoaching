import { push } from 'svelte-spa-router';   
import bonhart_storage from './bonhart-storage';


export const isLoggedIn = (redirect=true) => {
    let json_web_token = bonhart_storage.Token;
    console.log("isLoggedIn", json_web_token);

    let is_logged_in = json_web_token !== "";

    if (!is_logged_in && redirect) {
        alert("You must be logged in to access this page.");
        push("/login");
    }

    return is_logged_in;
}

export const logout = () => {
    bonhart_storage.Token = "";
    push("/login");
}


export const createUnsecureJWT = payload => {
    /* 
        Keep in mind that this method of creating a JWT is not secure, as the JWT is not signed and could be easily tampered with. It is only suitable for passing simple parameters that do not need to be secured.
    */

    const headers = {
        alg: "none",
        typ: "JWT"
    }

    const encoded_headers = window.btoa(JSON.stringify(headers)); // stupid vscode doesnt relize we are not working in node but in the browser

    const encoded_payload = window.btoa(JSON.stringify(payload));

    return `${encoded_headers}.${encoded_payload}.`;
}

function attributesToJson() {
    const json_data = {};
    console.log("AttributestoJson:" + this);
    Object.entries(this).forEach(([key, value]) => {
        if (!(this[key] instanceof Function) && key[0] !== '_') {
            json_data[key] = value;
        }
    });
    return JSON.stringify(json_data);
}

export class ExpertPublicData {
    constructor(id, name, subtitle, title, presentation, image, expert_type='consultant') {
        this.id = id;
        this.name = name;
        this.subtitle = subtitle;
        this.title = title;
        this.presentation = presentation;
        this.image = image;
        this.expert_type = expert_type;
    }

    toJson = attributesToJson.bind(this);
}

export const getUrlPARAM = key => {
    let url_string = window.location.href; 
    url_string = url_string.replace(/\/.{0,3}#/, ""); // remove #
    let url = new URL(url_string);
    return url.searchParams.get(key);
}

export const expert_types = {
    CONSULTANT: "consultant",
    MENTOR: "mentor",
}
