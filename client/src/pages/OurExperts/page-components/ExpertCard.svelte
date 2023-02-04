<script>
    import { users_server } from "../../../libs/HttpRequests";
    import { expert_types } from "../../../libs/itz-utils";
    import itz_logo from '../../../svg/MainLogo.svg'
    export let expert_data;
</script>


<div id="expert-data-card" class="{expert_data.expert_type === expert_types.MENTOR ? 'expert-mentor' : 'expert-consultant'}">
    <div id="expert-pic-wrapper">
        <figure id="expert-profile-picture">
            {#if expert_data.image !== undefined && expert_data.image !== ""}
                <img src="{users_server}/profile_pictures/{expert_data.image}" alt="profile">
            {:else}
                <div id="empty-image-placeholder">
                    <span id="itz-icon">
                        {@html itz_logo}
                    </span>
                </div>
            {/if}
        </figure>
        <span class="expert-type">
            {expert_data.expert_type === expert_types.MENTOR ? "Mentora" : "Consultora"}
        </span>
    </div>
    <div id="expert-info-wrapper">
        <h3 class="expert-name">
            {expert_data.name}
        </h3>
        <h4 class="expert-title">
            {expert_data.title}
        </h4>
        <p class="expert-description">
            {expert_data.subtitle}
        </p>
    </div>
</div>

<style>
    #expert-data-card {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: flex-start;
        width: 14vw;
        height: 40vh;
        padding: var(--spacing-2);
        border-radius: var(--boxes-roundness);
        gap: var(--spacing-3);
    }

    #expert-pic-wrapper {
        position: relative;
        width: 100%;
        height: 65%;
    }

    #expert-profile-picture {
        position: relative;
        border-radius: calc(var(--boxes-roundness) * 2);
        overflow: hidden;
        width: 100%;
        height: calc(2.2 * var(--spacing-h1));
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
        background: var(--theme-red);
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

    .expert-type {
        position: absolute;
        top:95%;
        right: 10%;
        font-weight: 600;
        padding: calc(var(--spacing-1) / 2);
        font-size: calc(.7 * var(--font-size-2));
        border-radius: var(--boxes-roundness);
    }

    .expert-mentor .expert-type {
        background: var(--theme-purple);
        color: var(--theme-pearl);
    }

    .expert-consultant .expert-type {
        background: var(--theme-red);
        color: var(--clear-color);
    }

    .expert-name {
        text-transform: capitalize;
        font-size: var(--font-size-3);
    }

    .expert-title {
        font-family: var(--font-text);
        font-size: calc(.8 * var(--font-size-2));
        text-transform: capitalize;
    }

    .expert-mentor .expert-title {
        color: var(--theme-purple);
    }

    .expert-consultant .expert-title {
        color: var(--theme-red);
    }

    .expert-description {
        font-family: var(--font-text);
        font-size: calc(.8 * var(--font-size-2));
    }
</style>