<script>
    import { GetExpertProfileRequest, GetExpertRequest,  users_server, GetExpertAppointments } from "../../libs/HttpRequests";
    import { Appointment } from "../../components/DateComponents/date_utils";
    import ExpertSchedule from "./page-component/ExpertSchedule.svelte";
    import itz_logo from "../../svg/MainLogo.svg";
    import { onMount, onDestroy } from "svelte";
    import createColorSchema, {
        supported_components
    } from '../../libs/ColorSchema';

    const expert_profile_page_schema = createColorSchema({
        color: "var(--dark-color)",
        contrast: "var(--clear-transparent-color)",
    }, supported_components.NAVBAR);

    export let params = {}
    
    let expert_id = params.id;
    let experts_appointments = [];
    let profile_data = {};
    let expert_data = {};

    onMount(() => {
        expert_profile_page_schema.define();

        getExpertProfile();
        getExpertData();

    });

    const getExpertProfile = () => {
        const profile_request = new GetExpertProfileRequest(expert_id);

        const on_success = profile => {
            profile_data = profile;
        };

        const on_error = error => {
            console.log(error);
        };

        profile_request.do(on_success, on_error);
    }

    const getExpertData = () => {
        const expert_request = new GetExpertRequest(expert_id);

        const on_success = expert => {
            expert_data = expert;
            getExpertAppointments();
            setStyles();
        };

        const on_error = error => {
            console.log(error);
        };

        expert_request.do(on_success, on_error);
    }

    const getExpertAppointments = () => {
        const appointments_request = new GetExpertAppointments(expert_id);

        const on_success = appointments => {
            let new_appointments = appointments.map(appointment => new Appointment(appointment));
                new_appointments.sort((a, b) => {
                return a.Start.getTime() - b.Start.getTime();
            });

            experts_appointments = new_appointments;
        };

        const on_error = error => {
            console.log(error);
        };

        appointments_request.do(on_success, on_error);
    }

    const setStyles = () => {
        const page_element = document.getElementById("itz-expert-profile-page");
        const expert_type_color = expert_data.expert_type === "mentor" ? "var(--theme-red)" : "var(--theme-purple)";

        page_element.style.setProperty("--expert-type-color", expert_type_color);
    }
</script>

<main id="itz-expert-profile-page">
    <div id="image-profile-wrapper">
        <figure id="expert-profile-picture">
            {#if profile_data.image_url !== undefined && profile_data.image_url !== ""}
                <img src="{users_server}/profile_pictures/{profile_data.image_url}" alt="profile">
            {:else}
                <div id="empty-image-placeholder">
                    <span id="itz-icon">
                        {@html itz_logo}
                    </span>
                </div>
            {/if}
        </figure>
        <span id="experts-achievements">
            {profile_data.brief}
        </span>
        <span id="expert-sold-meetings-counter">
            0 sesiones
            <!-- TODO: JD service required -->
        </span>
    </div>
    <div id="itz-profile-content">
        <hgroup id="profile-content-top">
            <h1 id="experts-name">
                {profile_data.public_name}
            </h1>
            <h3 id="expert-call-to-action">
                Agenda una sesión de <span class="ecta-expert-type-color">{expert_data.expert_type === "mentor" ? "mentoría" : "consultoría"}</span>
            </h3>
            <h2 id="experts-profession">
                {profile_data.professional_title}
            </h2>
        </hgroup>
        <p id="profile-description">
            {profile_data.description}
        </p>
    </div>
    <div id="itz-schedule-wrapper">
        <ExpertSchedule expert_id="{expert_id}" expert_appointments={experts_appointments}/>
    </div>
</main>

<style>

    /* DEBUG: BOXES */
    /* #itz-expert-profile-page * {
        border: 1px solid blue;
    } */


    /*=============================================
    =            top output panel            =
    =============================================*/
    



    #itz-expert-profile-page {
        --expert-type-color: var(--theme-red);

        box-sizing: border-box;
        width: 100%;
        display: grid;
        grid-template-columns: repeat(7, 1fr);
        gap: var(--spacing-4);
        padding: var(--page-padding-top);
    }

    #image-profile-wrapper {
        grid-column: 1 / 3;
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        gap: var(--spacing-3) 0;
        padding: 0 var(--spacing-4);
    }    

    #expert-profile-picture {
        position: relative;
        border-radius: calc(var(--boxes-roundness) * 2);
        overflow: hidden;
        width: 100%;
        height: calc(4.7 * var(--spacing-h1));
        margin: 0;
    }

    #expert-profile-picture img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        z-index: 1;
    }

    #empty-image-placeholder {
        display: grid;
        width: 100%;
        background: var(--expert-type-color);
        height: 100%;
        place-items: center;
        z-index: 1;
    }

    #itz-icon {
        width: 80%;
        background-size: cover;
    }
    
    :global(#itz-icon svg) {
        width: 100%;
        height: 100%;
        fill: var(--clear-color);
    }


    #experts-achievements {
        color: var(--expert-type-color);
        font-style: italic;
    }

    #expert-sold-meetings-counter {
        color: var(--expert-type-color);
        background: var(--theme-pearl);
    }

    
    /*=============================================
    =            Middle content            =
    =============================================*/

    #itz-profile-content {
        grid-column: 3 / 6;
    }

    #profile-content-top {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-2) 0;
    }

    #experts-name {
        font-size: var(--font-size-h2);
        font-weight: bold;
        text-transform: capitalize;
    }

    #experts-profession {
        font-size: var(--font-size-2);
        font-weight: 500;
        text-transform: none;
    }

    #expert-call-to-action {
        font-size: var(--font-size-4);
        text-transform: none;
        
    }
    
    .ecta-expert-type-color {
        background: var(--expert-type-color);
        color: var(--clear-color);
        padding: var(--spacing-1);
        border-radius: var(--boxes-roundness);
    }
    
    #profile-content-top :last-child {
        margin: var(--spacing-4) 0 0 0;
    }

    #profile-description {
        width: 100%;
        height: 50vh;
        resize: none;
        border: none;
    }

    
    /*=============================================
    =            Schedule section            =
    =============================================*/
    
    #itz-schedule-wrapper {
        grid-column: 6 / 8;
    }
    
    
    /*=============================================
    =            Editor Controls            =
    =============================================*/
    
    :global(#editor-controls .success) {
        background: var(--ready);
        color: var(--clear-color);
    }
    
    


</style>