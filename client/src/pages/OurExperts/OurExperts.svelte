<script>
    import { ExpertPublicData, expert_types } from '../../libs/itz-utils';
    import ExpertCard from './page-components/ExpertCard.svelte';
    import createColorSchema, {
        supported_components
    } from '../../libs/ColorSchema';
    import { onMount, onDestroy } from 'svelte';


    let experts = [
        new ExpertPublicData(
            "mock-id-01", 
            "Itziar Olóriz", 
            "Founder & CEO en Itz Coaching", 
            "Finanzas, estrategia, emprendimiento",
            "lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl,",
            "/resources/mock_experts/a.webp",
            expert_types.MENTOR
        ),
        new ExpertPublicData(
            "mock-id-02", 
            "Ana Sofía Sada", 
            "Account Associate de thinkPARALLAX", 
            "Impacto social, marca personal",
            "lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl,",
            "/resources/mock_experts/s.webp",
            expert_types.CONSULTANT
        ),
        new ExpertPublicData(
            "mock-id-03", 
            "Andrea Nacif", 
            "Team Lead Data Analyst en Nowports", 
            "Business Intelligence, Design Thinking",
            "lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl,",
            "/resources/mock_experts/b.jpg",
            expert_types.CONSULTANT
        ),
        new ExpertPublicData(
            "mock-id-04", 
            "Camila Zamora", 
            "Innovation Lead en Sigma Ventures", 
            "Operación, estrategia, gestión de proyectos",
            "lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl, vitae aliquam nisl nisl euismod nisl. Sed euismod, nisl vitae ultricies lacinia, nisl nisl aliquam nisl,",
            "/resources/mock_experts/c.jpg",
            expert_types.CONSULTANT
        ),

    ];

    const our_experts_color_schema = createColorSchema({
        color: "var(--dark-color)",
        contrast: "var(--clear-transparent-color)",
    }, supported_components.NAVBAR);

    onMount(() => {
        our_experts_color_schema.define();
    })

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