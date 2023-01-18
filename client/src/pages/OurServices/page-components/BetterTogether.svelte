<div id="better-together-wrapper">
    <div id="btw-center-wrapper">
        <div class="service-title-card">
            <div class="service-icon-wrapper">
                <img src="/resources/images/service_group-original.webp" alt="Pit stop service">
            </div>
            <h3 class="service-name-red">Better together</h3>
            <h4 class="service-type">Aprendizaje grupal</h4>
        </div>
        <p class="emotional-speach">
            Si caminas sola llegas rápido, si lo haces acompañada llegas lejos. Aprende y<br/> emprende de la mano con otras mujeres como tú, que están en camino de desarrollo<br/> personal y profesional
        </p>
        <div class="service-step-wrapper">
            <div class="service-step-num">*</div>
            <div class="service-step-data">
                <p class="service-step-desc">Programa activo<br/><span>Empodera y emprende <i>Vol. II</i></span></p>
            </div>
        </div>
    </div>
</div>

<style>
    #better-together-wrapper {
        box-sizing: border-box;
        display: flex;
        height: 100vh;
        justify-content: center;
        align-items: center;
        /* padding: var(--page-padding-1); */
    }

    #btw-center-wrapper {
        display: grid;
        width: 100%;
        grid-template-columns: repeat(3, 1fr);
        padding: 0 var(--spacing-h1);
        gap: var(--spacing-1);
        place-items: center;
    }

    p.emotional-speach {
        box-sizing: border-box;
        grid-area: 1 / 2 / 2 / 4;
        font-size: var(--font-size-2);
        padding: var(--spacing-h3);
    }

    .service-title-card {
        width: 34%;
        text-align: center;
    }

    .service-step-desc span {
        color: var(--theme-red);
        font-weight: 600;
    }

    .service-step-desc i {
        color: var(--theme-red);
        font-weight: 500;
    }

    .service-step-wrapper {
        grid-area: 2 / 2 / 3 / 4;
        font-size: var(--font-size-2);
    }
</style>