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

