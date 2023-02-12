<script>
    import PublicProfileEditor from './page-components/PublicProfileEditor.svelte';
    import ExpertsSchedule from './page-components/ExpertsSchedule.svelte';
    import bonhart_storage from "../../libs/bonhart-storage";
    import { parseJwt } from '../../libs/bonhart-utils';
    import { push } from 'svelte-spa-router';
    import { onMount } from 'svelte';
    import createColorSchema, {
        supported_components,
    } from '../../libs/ColorSchema';

    let expert_data = {};

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
    });

</script>

<main id="itz-administrator-panel">
    <section class="admin-panel-section pp-editor-section-wrapper">
        <PublicProfileEditor {expert_data} />
    </section>
    <section class="admin-panel-section">
        <ExpertsSchedule />
    </section>
    
</main>

<style>
    #itz-administrator-panel {
        padding: var(--page-padding-top);
    }
</style>