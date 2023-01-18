<script>
    import PitStop from './page-components/PitStopSection.svelte';
    import MentorMeet from './page-components/MentorMeet.svelte';
    import BetterTogether from './page-components/BetterTogether.svelte';
    import BusinessMoman from './page-components/BusinessMoman.svelte';
    import createColorSchema, {
        supported_components,
        checkLastComponentCollition
    } from '../../libs/ColorSchema';
    import { onMount, onDestroy } from 'svelte';

    window.scrollTo(0, 0);

    let section_collition_listener = undefined;
    const sections_color_schemas = {
        PIT_STOP: {
            ord: 0,
            section_id: "itz-meet-section",
            ref: undefined,
            color_schema: createColorSchema({
                color: "var(--dark-color)",
                contrast: "var(--clear-color)",
            }, supported_components.NAVBAR)
        },
        MENTOR_MEET: {
            ord: 1,
            section_id: "itz-mentor-meet",
            ref: undefined,
            color_schema: createColorSchema({
                color: "var(--dark-color)",
                contrast: "var(--clear-color)",
            }, supported_components.NAVBAR)
        },
        BETTER_TOGETHER: {
            ord: 2,
            section_id: "itz-better-together",
            ref: undefined,
            color_schema: createColorSchema({
                color: "var(--dark-color)",
                contrast: "var(--clear-color)",
            }, supported_components.NAVBAR)
        },
        BUSINESS_MOMAN: {
            ord: 3,
            section_id: "itz-business-moman",
            ref: undefined,
            color_schema: createColorSchema({
                color: "var(--dark-color)",
                contrast: "var(--clear-color)",
            }, supported_components.NAVBAR)
        },
    }

    
    onMount(() => {
        // Change the navbar color schema
        sections_color_schemas.PIT_STOP.color_schema.define();


        section_collition_listener = checkLastComponentCollition.bind({}, sections_color_schemas);

        window.addEventListener('scroll', section_collition_listener);
    });

    
    onDestroy(() => {
        window.removeEventListener('scroll', section_collition_listener);
    });
</script>

<main id="experts-catalog-page">
    <section bind:this={sections_color_schemas.PIT_STOP.ref} id={sections_color_schemas.PIT_STOP.section_id}>
        <PitStop />
    </section>
    <section bind:this={sections_color_schemas.MENTOR_MEET.ref} id={sections_color_schemas.MENTOR_MEET.section_id}>
        <MentorMeet />
    </section>
    <section bind:this={sections_color_schemas.BETTER_TOGETHER.ref} id={sections_color_schemas.BETTER_TOGETHER.section_id}>
        <BetterTogether />
    </section>
    <section bind:this={sections_color_schemas.BUSINESS_MOMAN.ref} id={sections_color_schemas.BUSINESS_MOMAN.section_id}>
        <BusinessMoman />
    </section>
</main>


<style>
    :global(.service-title-card) {
        width: 45%;
        display: flex;
        flex-direction: column;
        align-items: center;
        grid-area: 1 / 1 / 3 / 2;
    }

    :global(.service-icon-wrapper) {
        width: 100%;
    }

    :global(.service-icon-wrapper img) {
        width: 110%;
    }

    :global(.service-name-red) {
        color: var(--theme-red);
        text-transform: none;
        font-size: var(--font-size-h2);
        font-weight: bold;
        margin: 0;
    }

    :global(.service-type) {
        text-transform: none;
        font-size: var(--font-size-4);
        margin: 0;
    }
    
    :global(.service-step-wrapper) {
        --service-step-num-font-size: calc(3 * var(--font-size-h1));
        position: relative;
        height: (var(--service-step-num-font-size) + var(--spacing-h4));
    }
    
    :global(.service-step-num) {
        font-family: var(--font-titles);
        position: absolute;
        font-size: var(--service-step-num-font-size);
        color: var(--theme-pearl);
        top: 0;
        left: 0;
        z-index: -1;
        margin: 0;
        line-height: .8;
    }
    
    :global(.service-step-data) {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-1);
        padding: var(--spacing-2) 0 0 var(--spacing-h3);
    }

    :global(.service-step-name) {
        font-family: var(--font-titles);
        color: var(--theme-red);
        font-size: var(--font-size-3);
        margin: 0;
    }

    :global(.service-step-desc) {
        font-family: var(--font-text);
        font-size: var(--font-size-2);
        margin: 0;
    }

    :global(.service-step-type) {
        font-family: var(--font-text);
        font-size: calc(.8 * var(--font-size-2));
        color: var(--theme-red);
        font-style: italic;
        margin: 0;
    }
    
    :global(.service-characteristics) {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-1);
    }
    
    :global(.characteristics-label) {
        font-family: var(--font-text);
        font-size: calc(1.2 * var(--font-size-2));
        color: var(--theme-red);
        font-style: italic;
        margin: 0;
    }

    :global(.characteristics-list) {
        font-size: var(--font-size-2);
    }
</style>