<script>
    import { ExpertProfileRequest, ExpertProfileUpdateRequest } from '../../../libs/HttpRequests';
    import { onMount } from 'svelte';
    import itz_logo from '../../../svg/MainLogo.svg';
    import bonhart_storage from '../../../libs/bonhart-storage';
    import { newNotification } from '../../../components/notifications/events';
    import { logout } from '../../../libs/bonhart-utils';

    export let expert_data;

    let profile_data = {};

    let new_public_name;
    let new_professional_title;
    let new_description;
    let new_brief;
    let content_has_changed = false;

    onMount(() => {
        if (expert_data === undefined) {
            console.error("Expert data is undefined");
            return;
        }

        const profile_request = new ExpertProfileRequest(expert_data.id);

        const on_success = (response) => {
            profile_data = response;
            new_public_name = profile_data.public_name;
            new_professional_title = profile_data.professional_title;
            new_description = profile_data.description;
            new_brief = profile_data.brief;
        }

        const on_error = (error) => {
            console.error(error);
        }

        profile_request.do(on_success, on_error);
    })

    const contentChanged = () => {
        // this prevents svelte from doing large verifications in case the DOM has to be updated.
        if (!content_has_changed) {
            content_has_changed = true;
        }
    }

    const resetEditor = () => {
        const save_btn = document.querySelector("#editor-controls .full-btn");
        save_btn.classList.remove("success");

        content_has_changed = false;
    }

    const saveContent = () => {
        if (content_has_changed) {
            profile_data.public_name = new_public_name;
            profile_data.professional_title = new_professional_title;
            profile_data.description = new_description;
            profile_data.brief = new_brief;

            const profile_update_request = new ExpertProfileUpdateRequest(profile_data, bonhart_storage.Token);

            const on_success = () => {
                const save_btn = document.querySelector("#editor-controls .full-btn");
                save_btn.classList.add("success");
                window.setTimeout(resetEditor, 2000);
            }

            const on_error = (error) => {
                if (error === 401 || error === 403) {
                    logout();
                } else {
                    newNotification("Ah ocurrido un error al intentar actualizar tu perfil, por favor intenta de nuevo más tarde. Si el problema persiste, por favor contacta al administrador.");
                }
            }

            profile_update_request.do(on_success, on_error);
        }
    }

</script>

<section id="itz-expert-profile-editor">
    <div id="profile-editor-output-wrapper">
        <span class="peow-instruction">
            Has click en cualquier campo para editar, cuando termines haz click en el botón de guardar.
        </span>
        <span class="output"></span>
    </div>
    <div id="image-profile-wrapper">
        <figure id="expert-profile-picture">
            {#if profile_data.image_url !== ""}
                <img src="{profile_data.image_url}" alt="profile">
            {:else}
                <div id="empty-image-placeholder">
                    <span id="itz-icon">
                        {@html itz_logo}
                    </span>
                </div>
            {/if}
        </figure>
        <span id="experts-achievements">
            <input type="text" class="profile-attribute-holder" on:change={contentChanged} bind:value={new_brief}>
        </span>
        <span id="expert-sold-meetings-counter">
            0 sesiones
            <!-- TODO: JD service required -->
        </span>
    </div>
    <div id="itz-profile-content">
        <hgroup id="profile-content-top">
            <h1 id="experts-name">
                <input on:change={contentChanged} type="text" class="profile-attribute-holder" bind:value={new_public_name}>
            </h1>
            <h2 id="experts-profession">
                <input on:change={contentChanged} type="text" class="profile-attribute-holder" bind:value={new_professional_title}>
            </h2>
        </hgroup>
        <textarea on:change={contentChanged} name="description" id="profile-description" bind:value={new_description}></textarea>
    </div>
    <div id="itz-schedule-wrapper"></div>
    <div id="editor-controls">
        {#if content_has_changed}
             <button on:click={saveContent} class="full-btn">Guardar</button>
        {/if}
    </div>
</section>

<style>

    /* DEBUG: BOXES */
    /* #itz-expert-profile-editor * {
        border: 1px solid blue;
    } */

    .profile-attribute-holder {
        border: none;
        font-size: inherit;
        font-family: inherit;
        color: inherit;
    }

    /*=============================================
    =            top output panel            =
    =============================================*/
    
    #profile-editor-output-wrapper {
        grid-column: 1 / 8;
    }

    .peow-instruction {
        color: var(--theme-dark-red);
        font-size: var(--font-size-2);
        font-style: italic;
    }

    #itz-expert-profile-editor {
        width: 100%;
        display: grid;
        grid-template-columns: repeat(7, 1fr);
        gap: var(--spacing-4);
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
        border-radius: calc(var(--boxes-roundness) * 2);
        overflow: hidden;
        width: 100%;
        height: calc(4.7 * var(--spacing-h1));
        margin: 0;
    }

    #empty-image-placeholder {
        display: grid;
        width: 100%;
        background: var(--theme-red);
        height: 100%;
        place-items: center;
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
        color: var(--theme-red);
        font-style: italic;
    }

    #expert-sold-meetings-counter {
        color: var(--theme-purple);
        background: var(--theme-pearl);
    }

    
    /*=============================================
    =            Middle content            =
    =============================================*/

    #itz-profile-content {
        grid-column: 3 / 6;
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