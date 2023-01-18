<script>
    import Hero from "./page-components/Hero.svelte";
    import ItzDescription from "./page-components/ItzDescription.svelte";
    import ItzPurpleDivider from "./page-components/PurpleDivider.svelte";
    import WhyItz from "./page-components/WhyItz.svelte";
    import OurServices from "./page-components/OurServices.svelte";
    import OurExperts from "./page-components/OurExperts.svelte";
    import ItzStats from "./page-components/ItzStats.svelte";
    import OurClients from "./page-components/OurClients.svelte";
    import Contact from "./page-components/Contact.svelte";
    import createColorSchema, {
        supported_components,
        checkLastComponentCollition
    } from '../../libs/ColorSchema';
    import { onMount, onDestroy } from 'svelte';

    let section_collition_listener = undefined;
    const sections_color_schemas = {
        HERO: {
            ord: 0,
            section_id: "itz-hero-section",
            ref: undefined,
            color_schema: createColorSchema({
                color: "var(--clear-color)",
                contrast: "transparent",
            }, supported_components.NAVBAR)
            
        },
        DESCRIPTION: {
            ord: 1,
            section_id: "itz-desc-section",
            ref: undefined,
            color_schema: createColorSchema({
                color: "var(--dark-color)",
                contrast: "white",
            }, supported_components.NAVBAR)
        },
        PURPLE_DIVIDER: {
            ord: 2,
            section_id: "itz-purple-divider",
            ref: undefined,
            color_schema: createColorSchema({
                color: "var(--theme-pearl)",
                contrast: "var(--theme-purple)",
            }, supported_components.NAVBAR)
        },
        WHY_ITZ: {
            ord: 3,
            section_id: "itz-why-itz",
            ref: undefined,
            color_schema: createColorSchema({
                color: "var(--dark-color)",
                contrast: "var(--clear-color)",
            }, supported_components.NAVBAR)
        },
        OUR_SERVICES: {
            ord: 4,
            section_id: "itz-our-services",
            ref: undefined,
            color_schema: createColorSchema({
                color: "var(--dark-color)",
                contrast: "var(--clear-transparent-color)",
            }, supported_components.NAVBAR)
        },
        OUR_EXPERTS: {
            ord: 5,
            section_id: "itz-our-experts",
            ref: undefined,
            color_schema: createColorSchema({
                color: "var(--dark-color)",
                contrast: "var(--clear-color)",
            }, supported_components.NAVBAR)
        },
        ITZ_STATS: {
            ord: 6,
            section_id: "itz-itz-stats",
            ref: undefined,
            color_schema: createColorSchema({
                color: "var(--dark-color)",
                contrast: "var(--theme-pearl)",
            }, supported_components.NAVBAR)
        },
        OUR_CLIENTS: {
            ord: 7,
            section_id: "itz-our-clients",
            ref: undefined,
            color_schema: createColorSchema({
                color: "var(--dark-color)",
                contrast: "var(--clear-transparent-color)",
            }, supported_components.NAVBAR)
        },
        CONTACT: {
            ord: 8,
            section_id: "itz-contact",
            ref: undefined,
            color_schema: createColorSchema({
                color: "var(--clear-color)",
                contrast: "transparent",
            }, supported_components.NAVBAR)
        }
    };

    window.scrollTo(0, 0);

    onMount(() => {
        section_collition_listener = checkLastComponentCollition.bind({}, sections_color_schemas);
        window.addEventListener('scroll', section_collition_listener);
    });

    onDestroy(() => {
        window.removeEventListener('scroll', section_collition_listener);
    });
</script>

<main id="itz-main-page">
    <section bind:this={sections_color_schemas.HERO.ref} id={sections_color_schemas.HERO.section_id}>
        <Hero/>
    </section>
    <section bind:this={sections_color_schemas.DESCRIPTION.ref} id={sections_color_schemas.DESCRIPTION.section_id}>
        <ItzDescription/>
    </section>
    <section bind:this={sections_color_schemas.PURPLE_DIVIDER.ref} id={sections_color_schemas.PURPLE_DIVIDER.section_id}>
        <ItzPurpleDivider/>
    </section>
    <section bind:this={sections_color_schemas.WHY_ITZ.ref} id={sections_color_schemas.WHY_ITZ.section_id}>
        <WhyItz/>
    </section>
    <section bind:this={sections_color_schemas.OUR_SERVICES.ref} id={sections_color_schemas.OUR_SERVICES.section_id}>
        <OurServices/>
    </section>
    <section bind:this={sections_color_schemas.OUR_EXPERTS.ref} id={sections_color_schemas.OUR_EXPERTS.section_id}>
        <OurExperts/>
    </section>
    <section bind:this={sections_color_schemas.ITZ_STATS.ref} id={sections_color_schemas.ITZ_STATS.section_id}>
        <ItzStats/>
    </section>
    <section bind:this={sections_color_schemas.OUR_CLIENTS.ref} id={sections_color_schemas.OUR_CLIENTS.section_id}>
        <OurClients/>
    </section>
    <section bind:this={sections_color_schemas.CONTACT.ref} id={sections_color_schemas.CONTACT.section_id}>
        <Contact/>
    </section>
</main>

<style>

</style>