<script>
    import instagram_logo from '../../../svg/instagram.svg';
    import whatsapp_logo from '../../../svg/whatsapp.svg';
    import linkedin_logo from '../../../svg/linkedin.svg';
    import facebook_logo from '../../../svg/facebook.svg';

</script>

<div id="itz-contact-wrapper">
    <div id="icw-color-background">
        <h3 id="icw-title">
            Trabajemos <mark>juntas</mark>
        </h3>
        <article id="itz-contact-details">
            <section class="itz-contact-field">
                <h4>Correo</h4>
                <a href="mailto:hello@itzcoaching.com">
                    hello@itzcoaching.com
                </a>
            </section>
            <section class="itz-contact-field">
                <h4>Whatsapp</h4>
                <a href="https://wa.me/message/BZBSTMU6MACIH1">
                    +52 22 1258 3648
                </a>
            </section>
            <section class="itz-contact-field">
                <h4>
                    Redes sociales
                </h4>
                <ul id="itz-social-networks">
                    <li class="itz-social-link">
                        <a href="https://www.instagram.com/itz.coaching/">
                            {@html instagram_logo}
                        </a>
                    </li>
                    <li class="itz-social-link">
                        <a href="https://www.facebook.com/itzcoachingg">
                            {@html facebook_logo}
                        </a>
                    </li>
                    <li class="itz-social-link">
                        <a href="https://wa.me/message/BZBSTMU6MACIH1">
                            {@html whatsapp_logo}
                        </a>
                    </li>
                    <li class="itz-social-link">
                        <a href="https://www.linkedin.com/company/itzcoaching">
                            {@html linkedin_logo}
                        </a>
                    </li>
                </ul>
            </section>
        </article>
    </div>
</div>

<style>
    #itz-contact-wrapper {
        box-sizing: border-box;
        height: 100vh;
        color: var(--clear-color);
        background: url('/resources/images/laptop_background-original.webp');
        background-size: cover;
        background-position: center;
        background-repeat: no-repeat;
    }

    #icw-color-background {
        box-sizing: border-box;
        width: 100%;
        height: 100%;
        background: var(--theme-transparent-red);
        padding: var(--navbar-height) var(--spacing-4) var(--spacing-4);
    }

    #icw-title {
        text-transform: none;
        font-size: var(--font-size-h2);
        font-weight: 600;
    }

    #icw-title mark {
        background: transparent;
        color: inherit;
        font-style: italic;
        font-weight: 500;
    }

    #itz-contact-details {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-h3);
        padding: var(--spacing-4) var(--spacing-h2);
    }
    
    .itz-contact-field {
        width: 50%;
    }

    .itz-contact-field h4 {
        text-transform: none;
        font-size: var(--font-size-4);
        font-weight: 600;
    }

    .itz-contact-field a {
        font-size: var(--font-size-3);
        font-weight: 500;
        color: var(--clear-color);
        line-height: 2em;
    }

    #itz-social-networks {
        display: flex;
        width: 100%;
        gap: var(--spacing-2);
    }

    .itz-social-link {
        width: 3%;
        display: inline-block;
    }

    :global(.itz-social-link svg) {
        display: block;
        width: 100%;
        height: 100%;
    }

</style>