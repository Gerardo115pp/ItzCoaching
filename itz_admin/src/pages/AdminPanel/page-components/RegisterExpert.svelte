<script>
    import FieldData, { verifyFormFields, resetFormFields } from "../../../libs/FieldData";
    import { newNotification } from '../../../components/notifications/events' ;
    import CTloader from "../../../components/Loaders/CTloader.svelte";
    import { RegisterExpertRequest } from '../../../libs/HttpRequests';
    import Input from "../../../components/Input/Input.svelte";
    import bonhart_storage from "../../../libs/bonhart-storage";

    const register_expert_request = new RegisterExpertRequest(bonhart_storage.Token);
    let is_form_ready = false;
    let lock_form = false;



    let form_data = [
        new FieldData("new_expert_name", /^[\u0041-\u005A\u0061-\u007A\u0080-\u00FF\s]+$/, "nombre"),
        new FieldData("new_expert_username", /[A-z0-9_]{3,20}/, "usuario"),
        new FieldData("new_expert_email", /[A-z0-9._%+-]+@[A-z0-9.-]+\.[A-z]{2,4}/, "correo", "email"),
        new FieldData("new_expert_password", /[^\n\s]{8,60}/, "contraseña","password")
    ]
    let expert_type_select;

    // Link FormData to the request_data
        $: register_expert_request.name = form_data[0].getFieldValue();
        $: register_expert_request.username = form_data[1].getFieldValue();
        $: register_expert_request.email = form_data[2].getFieldValue();
        $: register_expert_request.password = form_data[3].getFieldValue();
    //

    const registerExpert = e => {
        e.preventDefault();
        if (is_form_ready) {
            lock_form = true;
            register_expert_request.expert_type = expert_type_select.value === "1" ? "mentor" : "consultant";

            const on_success = () => {
                lock_form = false;
                resetFormFields(form_data);
                verifyRegistrationForm();
                expert_type_select.value = "";
                newNotification("Experta registrada con éxito.");
            }

            const on_error = error_code => {
                console.log(error);
                if (error_code !== 404) {
                    newNotification(`Error ${error_code}: No se pudo registrar a la experta. please report.`);
                } else {
                    newNotification(`Usuario o contraseña incorrectos.`);
                }
                lock_form = false;
            }

            register_expert_request.do(on_success, on_error);
        } else {
            newNotification("Formulario incompleto o equivocado.");
        }
    }

    const verifyRegistrationForm = () => {
        if (lock_form) {
            return;
        }
        let is_valid = verifyFormFields(form_data);
        if (expert_type_select === undefined || expert_type_select.value === "") {
            is_valid = false;
        }

        is_form_ready = is_valid;
        form_data = [...form_data]; // trigger svelte update
    }
</script>

<form action="?" method="post" id="expert-registration-form">
    <fieldset name="expert-basic-info">
        {#each form_data as field}
            <Input
                field_data={field}
                input_label={`${field.name}:`}
                input_padding="var(--spacing-1)"
                disabled={lock_form}
                autocomplete="off"
                onEnterPressed={() => {
                    verifyFormFields(form_data);
                    window.queueMicrotask(() => {
                        if (is_form_ready) {
                            // loginUser();
                        }
                    });
                }}
                onBlur={verifyRegistrationForm}
            />
        {/each}
    </fieldset>
    <div id="expert-type-selectfield">
        <label for="expert-type">tipo de experta</label>
        <select on:change={verifyRegistrationForm} on:blur={verifyRegistrationForm}  bind:this={expert_type_select} name="expert-type" id="expert-type">
            <option disabled selected hidden value=""></option>
            <option value="1">mentora</option>
            <option value="2">consultora</option>
        </select>
    </div>
    {#if lock_form}
        <CTloader />
    {:else}
        <button id="register-expert-btn" on:click={registerExpert} class="full-btn" disabled={!is_form_ready || lock_form}>registrar</button>
    {/if}
</form>

<style>

    /* DEBUG BOXES */
    /* #expert-registration-form, #expert-registration-form * {
        border: 1px solid blue;
    } */

    #expert-registration-form {
        display: flex;
        width: 100%;
        padding: var(--spacing-3);
        flex-direction: column;
        align-items: center;
        gap: var(--spacing-4);
    }

    fieldset[name="expert-basic-info"] {
        width: 80%;
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        border: none;
        gap: var(--spacing-4);
    }

    #expert-type-selectfield {
        width: 80%;
    }

    #expert-type-selectfield label {
        font-size: var(--font-size-2);
        color: var(--theme-red);
        font-weight: 600;
    }
    
    #expert-type-selectfield label::after {
        content: ":";
    }

    #expert-type-selectfield select {
        width: 40%;
        background: var(--theme-red);
        font-size: var(--font-size-2);
        border: none;
        padding: var(--spacing-1);
        color: var(--clear-color);
        font-weight: 600;
        border-radius: var(--boxes-roundness);
    }

    #expert-type-selectfield select:focus {
        outline: none;
    }

    #expert-type-selectfield select option {
        background: var(--clear-color);
        color: var(--theme-red);
    }

    #expert-registration-form button#register-expert-btn {
        text-transform: lowercase;
    }

    #expert-registration-form button#register-expert-btn:disabled {
        /* background: var(--theme-red);
        color: var(--clear-color); */
        filter: grayscale(1);
    }
</style>