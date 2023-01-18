<div id="meet-up-wrapper">
    <div id="muw-center-wrapper">
        <div class="service-title-card">
            <div class="service-icon-wrapper">
                <img src="/resources/images/service_mentorship-original.webp" alt="Pit stop service">
            </div>
            <h3 class="service-name-red">Mentor meet-up</h3>
            <h4 class="service-type">Mentoría</h4>
        </div>
        <div class="service-step-wrapper">
            <div class="service-step-num">1</div>
            <div class="service-step-data">
                <p class="service-step-desc">Elige la <span>experta</span> de tu interés</p>
            </div>
        </div>
        <div class="service-step-wrapper">
            <div class="service-step-num">2</div>
            <div class="service-step-data">
                <p class="service-step-desc">Agenda tu <span>sesión</span> con ella</p>
            </div>
        </div>
        <div class="service-step-wrapper">
            <div class="service-step-num">3</div>
            <div class="service-step-data">
                <p class="service-step-desc">Disfruta de una <span>conversación</span> en la que podrás recibir consejos para tu negocio</p>
                <!-- <h5 class="service-step-type"></h5> -->
            </div>
        </div>
        <div class="service-characteristics">
            <span class="characteristics-label">Características</span>
            <ul class="characteristics-list">
                <li class="characteristic-item">Acceso a una platica virtual con una experta</li>
                <li class="characteristic-item">Sesión individual</li>
                <li class="characteristic-item">Duración 30 min, 45 min o 1hr</li>
            </ul>
        </div>
    </div>
</div>

<style>

    /* #meet-up-wrapper * {
        border: 1px solid red;
    } */

    #meet-up-wrapper {
        box-sizing: border-box;
        display: flex;
        height: 100vh;
        justify-content: center;
        align-items: center;
        padding: var(--page-padding-1);

    }

    #muw-center-wrapper {
        display: grid;
        width: 100%;
        grid-template-columns: repeat(3, 1fr);
        padding: 0 var(--spacing-h2);
        gap: var(--spacing-3);
    }   

    .service-title-card {
        width: 54%;
        text-align: center;
    }

    .service-step-desc span {
        color: var(--theme-red);
        font-weight: 600;
    }
</style>