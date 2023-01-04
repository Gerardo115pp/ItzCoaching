<script>
    import Hero from "./page-components/Hero.svelte";
    import ItzDescription from "./page-components/ItzDescription.svelte";
    import ItzPurpleDivider from "./page-components/PurpleDivider.svelte";
    import WhyItz from "./page-components/WhyItz.svelte";
    import OurServices from "./page-components/OurServices.svelte";
    import OurExperts from "./page-components/OurExperts.svelte";
    import ItzStats from "./page-components/ItzStats.svelte";
    import OurClients from "./page-components/OurClients.svelte";
    import createColorSchema, {
        supported_components,
    } from '../../libs/ColorSchema';
    import { onMount, onDestroy } from 'svelte';

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
    };
    

    // SECTION POSITION CHECKING
    const check_every = 50;
    let skiped_sections = check_every; // first check is always performed

    onMount(() => {
        window.addEventListener('scroll', checkLastComponentCollition);
    });

    onDestroy(() => {
        window.removeEventListener('scroll', checkLastComponentCollition);
    });

    function sectionReachTopScreen(dom_element) {
        if (dom_element === undefined) {
            return false;
        }

        const element_rect = dom_element.getBoundingClientRect();

        const element_top = element_rect.top;

        const viewport_height = Math.max(document.documentElement.clientHeight, window.innerHeight || 0);

        return element_top <= viewport_height * .14;
    }
    
    function checkLastComponentCollition() {
        if (skiped_sections < check_every) {
            skiped_sections++;
            return;
        }

        const sections = Object.values(sections_color_schemas);

        if (sections.length === 0) {
            return;
        }

        let last_collided_section = sections[0];
        
        for (let section of sections) {
            if (sectionReachTopScreen(section.ref)) {
                if (section.ord > last_collided_section.ord) {
                    last_collided_section = section;
                }
            }
        }

        last_collided_section.color_schema.define();
    }
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
</main>

<style>

</style>