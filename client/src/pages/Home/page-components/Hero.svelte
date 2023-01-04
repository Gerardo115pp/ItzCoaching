<div id="imp-top-page-hero">
    <div id="imp-tph-color-filter-container">
        <div id="imp-tph-center-content">
            <h3 id="imp-tph-hero-title">
                <strong>Emprender</strong> tiene sus retos,<br/>Te <i>ayudamos</i> a resolverlos.
            </h3>
            <div id="imp-tph-hero-call-to-action">
                <button class="clear-btn">¡Hablemos!</button>
            </div>
        </div>
    </div>
</div>

<style>
    #imp-top-page-hero {
        box-sizing: border-box;
        background: url("/resources/images/main_itz_background-XL.webp") ;
        background-size: cover;
        background-repeat: no-repeat;
        height: 100vh;
        transition: all .2s ease-in-out;
    }
    
    #imp-tph-color-filter-container {
        display: grid;
        background: var(--theme-transparent-red);
        place-items: center;
        height: 100vh;
    }

    #imp-tph-center-content {
        display: flex;
        flex-direction: column;
        align-items: flex-end;
        /* justify-content: center; */
    }

    #imp-tph-hero-title {
        font-family: var(--font-text);
        color: var(--clear-color);
        font-size: var(--font-size-h2);
    }

    #imp-tph-hero-title strong {
        letter-spacing: var(--spacing-1);
    }
    
    #imp-tph-hero-title i {
        font-family: var(--font-moontime);
        font-weight: 500;
        text-transform: lowercase;
        font-size: calc(2.5 * var(--font-size-h1));
    }
</style>