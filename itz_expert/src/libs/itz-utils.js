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

export const convertTimeToUTC = time_string => {
    // time_string is a string representing a time in the format "HH:MM" in the local timezone
    // returns a string representing the same time in UTC with the format "HH:MM"
    const time = new Date();
    console.log("Time string:", time_string);
    if (time_string.search(/\s?[APapmM]{2}/) >= 0) {
        console.log("Time string is in the format 'HH:MM AM/PM'");
        // time_string is in the format "HH:MM AM/PM"
        const [hours, minutes_am_pm] = time_string.split(":");
        const [minutes, am_pm] = minutes_am_pm.split(" ");
        time.setHours(parseInt(hours) + (am_pm === "PM" ? 12 : 0));
        time.setMinutes(minutes);
    } else {
        console.log("Time string is in the format 'HH:MM'");
        const [hours, minutes] = time_string.split(":");
        time.setHours(hours);
        time.setMinutes(minutes);
        time.setSeconds(0);
        time.setMilliseconds(0);
    }

    return `${time.getUTCHours().toString().padStart(2, '0')}:${time.getUTCMinutes().toString().padStart(2, '0')}`;
}

export const convertUTCtoLocalTime = utc_time => {
    // utc_time is a string representing a time in the format "HH:MM" in UTC
    // returns a string representing the same time in the local timezone with the format "HH:MM"

    const time = new Date();
    const [hours, minutes] = utc_time.split(":");
    time.setUTCHours(hours);
    time.setUTCMinutes(minutes);
    time.setUTCSeconds(0);

    return `${time.getHours().toString().padStart(2, '0')}:${time.getMinutes().toString().padStart(2, '0')}`;
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

export const expert_types = {
    CONSULTANT: "consultant",
    MENTOR: "mentor",
}
