<script>
    import PublicProfileEditor from './page-components/PublicProfileEditor.svelte';
    import ExpertsSchedule from './page-components/ExpertsSchedule.svelte';
    import { GetExpertRequest } from '../../libs/HttpRequests';
    import bonhart_storage from "../../libs/bonhart-storage";
    import { parseJwt } from '../../libs/bonhart-utils';
    import { expert_panel_events } from './events';
    import { push } from 'svelte-spa-router';
    import { onMount, onDestroy } from 'svelte';
    import createColorSchema, {
        supported_components,
    } from '../../libs/ColorSchema';

    let expert_data = {};
    let updated_expert_data = {}; // this has reliable is_available and is_active values. 

    // PARSE EXPERT DATA FROM TOKEN
    if (bonhart_storage.Token === "") {
        push("/")
    } else {
        expert_data = parseJwt(bonhart_storage.Token);
        console.log(expert_data);
    }

    onMount(() => {
        const color_schema = createColorSchema(
            {
                color: 'var(--theme-red)',
                contrast: 'var(--clear-color)',
            },
            supported_components.NAVBAR
        );
        color_schema.define();

        requestExpertData();

        document.addEventListener(expert_panel_events.EXPERT_DATA_CHANGED, requestExpertData);
    });

    onDestroy(() => {
        document.removeEventListener(expert_panel_events.EXPERT_DATA_CHANGED, requestExpertData);
    });

    const requestExpertData = () => {
        const expert_request = new GetExpertRequest(expert_data.id);

        const on_success = (response_data) => {
            updated_expert_data = response_data;
        };

        const on_error = (error) => {
            console.log(error);
        };

        expert_request.do(on_success, on_error);
    }

</script>

<main id="itz-administrator-panel">
    <section class="admin-panel-section">
        {#if expert_data.id !== undefined}
            <ExpertsSchedule expert_id={expert_data.id} is_schedule_active={updated_expert_data.is_available}/>
        {/if}
    </section>
    <section class="admin-panel-section pp-editor-section-wrapper">
        <PublicProfileEditor {expert_data} />
    </section>
</main>

<style>
    #itz-administrator-panel {
        padding: var(--page-padding-top);
    }

    #itz-administrator-panel section:not(:last-child) {
        margin-bottom: var(--spacing-h2);
    }
</style>