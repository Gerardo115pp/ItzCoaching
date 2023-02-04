<script>
    import { ExpertPublicData, expert_types } from '../../libs/itz-utils';
    import { GetActiveExpertsRequest } from '../../libs/HttpRequests';
    import ExpertCard from './page-components/ExpertCard.svelte';
    import createColorSchema, {
        supported_components
    } from '../../libs/ColorSchema';
    import { onMount, onDestroy } from 'svelte';


    let experts = [];

    const our_experts_color_schema = createColorSchema({
        color: "var(--dark-color)",
        contrast: "var(--clear-transparent-color)",
    }, supported_components.NAVBAR);

    onMount(() => {
        our_experts_color_schema.define();

        requestExpertsProfiles();
    });

    function requestExpertsProfiles() {
        const active_experts_request = new GetActiveExpertsRequest();

        const on_success = profiles => {
            experts = profiles.map(profile => {
                return new ExpertPublicData(profile.expert_id, profile.public_name, profile.brief, profile.professional_title, profile.description, profile.image_url, profile.expert_type);
            });
        };

        const on_error = error => {
            console.log(error);
        };

        active_experts_request.do(on_success, on_error);
    }

</script>

<main id="itz-experts-page">
    <div id="top-page-label">
        <h2 id="iep-title">Nuestras expertas</h2>
        <h3 id="iep-subtitle">Acceso a las mejores</h3>
    </div>
    <section id="experts-container">
        {#each experts as expert}
            <ExpertCard expert_data={expert}/>
        {/each}
    </section>
</main>

<style>
    #itz-experts-page {
        padding: var(--page-padding-top);
    }

    /* #top-page-label {
    } */
    
    #iep-title {
        font-size: var(--font-size-h2);
    }
    
    #iep-subtitle {
        font-size: var(--font-size-4);
    }

    #experts-container {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        justify-items: center;
        align-items: center;
        padding: var(--page-padding-2);
    }
</style>