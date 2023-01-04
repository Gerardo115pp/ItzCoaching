<div id="itz-our-services-wrapper">
    <h2 id="iosw-title">Nuestros servicios</h2>
    <div id="iosw-services-container">
        <div class="itz-service">
            <img src="/resources/images/service_pit_stop-original.webp" alt="pit stop service" class="itz-service-icon">
            <h3 class="itz-service-name">
                Pit Stop<br/><i>Consultoría</i>
            </h3>
            <p class="itz-service-description">
                Identifica las necesidades de tu empresa y recibe soluciones y recomendaciones para alcanzar tus metas
            </p>
        </div>
        <div class="itz-service">
            <img src="/resources/images/service_mentorship-original.webp" alt="mentorship service" class="itz-service-icon">
            <h3 class="itz-service-name">
                Mentor meet-up<br/><i>Mentoría</i>
            </h3>
            <p class="itz-service-description">
                Conversa con una experta y aprende de su experiencia, recibe consejos y guía para tu negocio
            </p>
        </div>
        <div class="itz-service">
            <img src="/resources/images/service_group-M.webp" alt="learning group service" class="itz-service-icon">
            <h3 class="itz-service-name">
                Better together<br/><i>Aprendizaje grupal</i>
            </h3>
            <p class="itz-service-description">
                Aprovecha la oportunidad de aprender de la mano de otras emprendedoras
            </p>
        </div>
        <div class="itz-service">
            <img src="/resources/images/service_comunity-M.webp" alt="community service" class="itz-service-icon">
            <h3 class="itz-service-name">
                Workshop<br/><i>Capacitación</i>
            </h3>
            <p class="itz-service-description">
                Conecta con otras emprendedoras y desarrolla alianzas y colaboraciones
            </p>
        </div>
    </div>
</div>

<style>
    #itz-our-services-wrapper {
        display: flex;
        height: 100vh;
        flex-direction: column;
        justify-content: center;
        padding: var(--spacing-4) var(--spacing-h2);
        gap: var(--spacing-h3);
    }

    #iosw-title {
        font-size: var(--font-size-h2);
        font-weight: 600;
    }

    #iosw-services-container {
        display: flex;
        justify-content: space-between;
        margin-top: var(--spacing-4);
        padding: 0 var(--spacing-3);

    }

    .itz-service {
        max-width: 24%;
        display: flex;
        flex-direction: column;
        gap: var(--spacing-2);
        text-align: center;
    }
    
    .itz-service img {
        width: 90%;
        object-fit: cover;
    }

    .itz-service h3 {
        font-family: var(--font-titles);
        font-size: var(--font-size-4);
        color: var(--theme-red);
        text-transform: none;
    }

    .itz-service h3 i {
        font-family: var(--font-text);
        font-size: var(--font-size-4);
    }

    .itz-service p {
        font-family: var(--font-text);
        font-size: var(--font-size-3);
    }


</style>