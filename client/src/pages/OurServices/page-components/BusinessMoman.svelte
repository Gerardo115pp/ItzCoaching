<!-- Need to write Moman because apparenlty github dosnt like wom.an word, so it will be moman -->
<div id="better-together-wrapper">
    <div id="btw-center-wrapper">
        <div class="service-title-card">
            <div class="service-icon-wrapper">
                <img src="/resources/images/service_comunity-original.webp" alt="Pit stop service">
            </div>
            <h3 class="service-name-red">Business<br/>Women<br/>Collective</h3>
            <h4 class="service-type">Comunidad</h4>
        </div>
        <p class="emotional-speach">
            Sé parte de nuestra comunidad que busca conectar a mujeres, compartir experiencias<br/>y dar un sentido de pertenencia a las emprendedoras
        </p>
        <div class="service-step-wrapper">
            <div class="service-step-num">*</div>
            <div class="service-step-data">
                <p class="service-step-desc">Próximamente<br/><span>Membresía</span></p>
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

    .service-step-wrapper {
        grid-area: 2 / 2 / 3 / 4;
        font-size: var(--font-size-2);
    }
</style>