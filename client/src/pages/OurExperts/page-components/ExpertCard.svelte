<script>
    import { expert_types } from "../../../libs/itz-utils";
    export let expert_data;
</script>


<div id="expert-data-card" class="{expert_data.expert_type === expert_types.MENTOR ? 'expert-mentor' : 'expert-consultant'}">
    <div id="expert-pic-wrapper">
        <img src="{expert_data.image}" alt="{expert_data.name}" class="expert-pic">
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

    .expert-pic {
        width: 100%;
        height: 100%;
        object-fit: cover;
        border-radius: var(--boxes-roundness);
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