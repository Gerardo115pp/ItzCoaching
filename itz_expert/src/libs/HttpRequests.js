import { createUnsecureJWT } from "./bonhart-utils";

export const auth_server = AUTH_SERVER;
export const users_server = USERS_SERVER;
export const jd_server = JD_SERVER;
export const payments_server = PAYMENTS_SERVER;

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

export class PatchExpertProfileRequest {
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

export class PatchExpertAvailabilityRequest {
    constructor(expert_id, availability, token) {
        this._token = token;
        this.id = expert_id;
        this.is_available = availability;
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

export class PostExpertProfilePicture {
    constructor(file_blob, token) {
        this._token = token;
        this.file = file_blob;
    }


    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Authorization', `Bearer ${this._token}`);

        const form = new FormData();

        form.append('profile_picture', this.file, this.file.name);


        const request = new Request(`${users_server}/profile_pictures`, {
            method: 'POST',
            headers: headers,
            body: form
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

export class PutExpertSchedule {
    constructor(expert_id, token) {
        this._token = token;
        this.owner = expert_id;
        this.week_availability = {};
        this.time_availability = [];
    }

    toJson = attributesToJson.bind(this);

    do = (on_success, on_error) => {
        if (this.week_availability.mondey === undefined) {
            return;
        }
        
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Authorization', `Bearer ${this._token}`);

        const request = new Request(`${users_server}/schedules`, {
            method: 'PUT',
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

export class GetExpertAppointments {
    constructor(expert_id) {
        this.expert_id = expert_id;
    }

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${jd_server}/appointments/expert?id=${this.expert_id}`, {
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