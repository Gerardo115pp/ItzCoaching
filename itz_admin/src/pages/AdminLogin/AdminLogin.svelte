<script>
    import { onMount } from "svelte";
    import { push, link } from "svelte-spa-router";
    import Input from "../../components/Input/Input.svelte";
    import bonhart_storage from "../../libs/bonhart-storage";
    import { AdminLoginRequest } from '../../libs/HttpRequests';
    import FieldData, { verifyFormFields } from "../../libs/FieldData";
    import { newNotification } from '../../components/notifications/events' ;

    const login_request = new AdminLoginRequest();
    let is_form_ready = false;
    let lock_form = false;



    let form_data = [
        new FieldData("admin_email", /[^\n\s;\'\"\`]{2,60}/, "email", "email"),
        new FieldData("admin_password", /[^\n\s]{8,60}/, "password", "password")
    ]

    // Link FormData to the request_data
        $: login_request.email = form_data[0].getFieldValue();
        $: login_request.password = form_data[1].getFieldValue();
    //

    onMount(() => {
        if (bonhart_storage.Token !== "") {
            push("/admin-panel");
        }
    });

    const verifyRegistrationForm = () => {
        if (lock_form) {
            return;
        }
        let is_valid = verifyFormFields(form_data);
        is_form_ready = is_valid;
        form_data = [...form_data]; // trigger svelte update
    }
    
    const loginUser = () => {
        if(is_form_ready) {
            const on_success = data => {
                if ("token" in data) {
                    bonhart_storage.Token = data.token;
                    push("/admin-panel");
                }
            }
            
            const on_error = error_code => {
                console.log(error);
                if (error_code !== 404) {
                    newNotification(`Error ${error_code}: No se pudo iniciar sesión. please report.`);
                } else {
                    newNotification(`Usuario o contraseña incorrectos.`);
                }
            }

            login_request.do(on_success, on_error);
        }
    }

</script>



<main id="itz-admin-login-page">
    <div id="ial-login-wrapper" class="round-box">
        <div id="itz-admin-login-title-wrapper">
            <h1>Itz Coaching</h1>
            <h2>para administradores</h2>
        </div>
        <div id="ial-login-form-container">
            <form on:submit={() => {}} id="form-fields">
                {#each form_data as field}
                    <div class="form-field-wrapper">
                        <Input
                            field_data={field}
                            isClear={true}
                            isSquared={true}
                            input_label={field.name}
                            onEnterPressed={() => {
                                verifyFormFields(form_data);
                                window.queueMicrotask(() => {
                                    if (is_form_ready) {
                                        loginUser();
                                    }
                                });
                            }}
                            onBlur={verifyRegistrationForm}
                        />
                    </div>
                {/each}
                <button class="full-btn" type="button" on:click={loginUser}>Entrar</button>
            </form>
        </div>
    </div>
</main>

<style>

    #itz-admin-login-page {
        --ial-height: 100vh;
        
        box-sizing: border-box;
        display: grid;
        min-height: var(--ial-height);
        place-items: center;
        background: var(--theme-red);
        z-index: 1;
    }

    :global(#content-wrap:has(#itz-admin-login-page)) {
        margin: 0;
    }

    #ial-login-wrapper {
        --login-form-width: 40%;
        --login-form-height: calc(var(--ial-height) * .6);

        box-sizing: border-box;
        display: flex;
        color: var(--theme-dark-red);
        width: var(--login-form-width);
        min-height: var(--login-form-height);
        flex-direction: column;
        background: var(--clear-color);
        align-items: center;
        padding: var(--spacing-4) var(--spacing-4);
        border: 1px solid #ccc;
    }

    #itz-admin-login-title-wrapper {
        text-align: center;
    }

    #itz-admin-login-title-wrapper h1 {
        font-size: var(--font-size-h2);
        height: var(--font-size-h1);
        text-transform: capitalize;
        margin: var(--spacing-1) auto;
        padding: 0 var(--spacing-3);
        border-bottom: 2px solid var(--theme-red);
        width: max-content;
    }

    #itz-admin-login-title-wrapper h2 {
        font-size: var(--font-size-2);
        font-family: var(--font-text);
        margin: var(--spacing-1) 0;
    }

    #ial-login-form-container {
        margin: var(--spacing-3) 0;
        width: 100%;
    }

    #form-fields {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }

    #form-fields .form-field-wrapper {
        margin: var(--spacing-1) 0;
        width: 60%;
    }

    #form-fields > button {
        margin: var(--spacing-3) auto;
    }


</style>