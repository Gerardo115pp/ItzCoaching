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

export class AdminLoginRequest {
    constructor() {
        this.email = "";
        this.password = "";
    }

    toJson = attributesToJson.bind(this);

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${auth_server}/tokens/admins`, {
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
    constructor(token) {
        this._token = token;
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

export class GetAllExpertsRequest {
    constructor(token) {
        this._token = token;
    }

    toJson = attributesToJson.bind(this);

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Authorization', `Bearer ${this._token}`);

        const request = new Request(`${users_server}/experts?all=1`, {
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

export class DeleteExpertRequest {
    constructor(token, id) {
        this._token = token;
        this.id = id;
    }

    toJson = attributesToJson.bind(this);

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Authorization', `Bearer ${this._token}`);

        const request = new Request(`${users_server}/experts`, {
            method: 'DELETE',
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

export class UpdateExpertActiveStatusRequest {
    constructor(token, id, new_status) {
        this._token = token;
        this.id = id;
        this.is_active = new_status;
    }

    toJson = attributesToJson.bind(this);

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Authorization', `Bearer ${this._token}`);

        const request = new Request(`${users_server}/experts`, {
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