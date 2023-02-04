<script>
    import bonhart_storage from "../../libs/bonhart-storage";
    import createColorSchema, {
        supported_components,
    } from '../../libs/ColorSchema';
    import ExpertAdministrationSection from "./page-components/ExpertAdministrationSection.svelte";
    import { onMount } from 'svelte';

    onMount(() => {
        const color_schema = createColorSchema(
            {
                color: 'var(--theme-red)',
                contrast: '',
            },
            supported_components.NAVBAR
        );
        color_schema.define();
    });

</script>

<main id="itz-administrator-panel">
    <div class="hub-section-wrapper">
        <ExpertAdministrationSection />
    </div>
</main>

<style>
    #itz-administrator-panel {
        padding: var(--page-padding-top);
    }
</style>