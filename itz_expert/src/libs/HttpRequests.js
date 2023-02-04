import { createUnsecureJWT } from "./bonhart-utils";

export const auth_server = AUTH_SERVER;
export const users_server = USERS_SERVER;

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

export class ExpertLoginRequest {
    constructor() {
        this.email = "";
        this.password = "";
    }

    toJson = attributesToJson.bind(this);

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${auth_server}/tokens/experts`, {
            method: 'POST',
            headers: headers,
            body: this.toJson()
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                promise.json().then(data => {
                    on_success(data);
                });
            } else {
                on_error(promise.status);
            }
        });
    }
}

export class RegisterExpertRequest {
    constructor() {
        this._token = "";
        this.name = "";
        this.username = "";
        this.email = "";
        this.password = "";
        this.expert_type = "";
    }

    toJson = attributesToJson.bind(this);

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Authorization', `Bearer ${this._token}`);

        const request = new Request(`${users_server}/experts`, {
            method: 'POST',
            headers: headers,
            body: this.toJson()
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                on_success();
            } else {
                on_error(promise.status);
            }
        });
    }
}

export class ExpertProfileRequest {
    constructor(expert_id) {
        this.id = expert_id;
    }

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${users_server}/public_profiles?id=${this.id}`, {
            method: 'GET',
            headers: headers
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                promise.json().then(data => {
                    on_success(data);
                });
            } else {
                on_error(promise.status);
            }
        });
    }
}

export class ExpertProfileUpdateRequest {
    constructor(public_profile, token) {
        this._token = token;
        this.public_name = public_profile.public_name;
        this.description = public_profile.description;
        this.professional_title = public_profile.professional_title;
        this.brief = public_profile.brief;
        this.image_url = public_profile.image_url;
        this.expert_id = public_profile.expert_id;
    }

    toJson = attributesToJson.bind(this);

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Authorization', `Bearer ${this._token}`);

        const request = new Request(`${users_server}/public_profiles`, {
            method: 'PATCH',
            headers: headers,
            body: this.toJson()
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                on_success();
            } else {
                on_error(promise.status);
            }
        });
    }
}
