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

export class GetActiveExpertsRequest {
    toJson = attributesToJson.bind(this);

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${users_server}/public_profiles?active=1`, {
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

export class GetExpertProfileRequest {
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

export class GetExpertRequest {
    constructor(expert_id) {
        this.id = expert_id;
    }

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${users_server}/experts?id=${this.id}`, {
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

export class GetExpertSchedule {
    constructor(expert_id) {
        this.expert_id = expert_id;
    }

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${users_server}/schedules?id=${this.expert_id}`, {
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