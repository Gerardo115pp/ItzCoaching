<script>
    import Carousel from 'svelte-carousel';

    const experts_data = [
        {
            name: "Itziar Olóriz",
            title: "Founder & CEO",
            profile_picture: "/resources/images/itz_placeholder-original.webp",
            session_count: 30,
            description: "Finanzas, estrategia, emprendimiento",
        },
        {
            name: "Ainhoa Olóriz",
            title: "Founder & COO",
            profile_picture: "/resources/images/itz_placeholder-original.webp",
            session_count: 30,
            description: "Finanzas, estrategia, emprendimiento",
        },
        {
            name: "Ane Olóriz",
            title: "Founder & CMO",
            profile_picture: "/resources/images/itz_placeholder-original.webp",
            session_count: 30,
            description: "Finanzas, estrategia, emprendimiento",
        },
        {
            name: "Ane Olóriz",
            title: "Founder & CMO",
            profile_picture: "/resources/images/itz_placeholder-original.webp",
            session_count: 30,
            description: "Finanzas, estrategia, emprendimiento",
        },
        {
            name: "Ane Olóriz",
            title: "Founder & CMO",
            profile_picture: "/resources/images/itz_placeholder-original.webp",
            session_count: 30,
            description: "Finanzas, estrategia, emprendimiento",
        }
    ]
</script>

<div id="itz-our-experts-wrapper">
    <h2 id="ioew-title">Nuestras expertas</h2>
    <div id="ioew-experts-corusel-wrapper">
        <Carousel
            autoplayDuration={200}
            duration={25000}
            timingFunction="linear"
            dots={false}
            arrows={false}
            particlesToShow={4}
            particlesToScroll={1}
            infinite
            pauseOnFocus
            autoplay
        >
            {#each experts_data as expert}
                <div class="itz-expert-card">
                    <div class="iec-profile-picture-wrapper">
                        <img src={expert.profile_picture} alt={expert.name} />
                        <span class="session-counter">
                            {expert.session_count} Sesiones
                        </span>
                    </div>
                    <div class="iec-expert-info-wrapper">
                        <h3 class="iec-expert-name">{expert.name}</h3>
                        <h4 class="iec-expert-title"><i>{expert.title}</i></h4>
                        <p class="iec-expert-tags">{expert.description}</p>
                    </div>
                </div>
            {/each}
        </Carousel>
    </div>
</div>

<style>
    #itz-our-experts-wrapper {
        height: 100vh;
        padding: var(--page-padding-1);
    }

    #ioew-title {
        font-size: var(--font-size-h2);
        font-weight: 600;
    }

    #ioew-experts-corusel-wrapper {
        width: 80%;
        padding: var(--spacing-h1);
        margin: 0 auto;
    }

    .itz-expert-card {
        width: 10vw !important;
        margin: var(--spacing-2) !important;
    }

    .iec-profile-picture-wrapper {
        display: grid;
        width: 80%;
        grid-template: repeat(16, 1fr) / repeat(10, 1fr);
        margin: 0 auto;
    }

    .iec-profile-picture-wrapper img {
        max-width: 100%;
        grid-area: 1 / 1 / 17 / 11;
        z-index: 1;
        border-radius: var(--button-roundness);
    }

    .session-counter {
        display: grid;
        grid-area: 15 / 7 / 16 / 10;
        z-index: 2;
        font-size: var(--font-size-1);
        place-items: center;
        background-color: var(--theme-pearl);
        color: var(--dark-color);
        border-radius: var(--button-roundness);
    }

    .iec-expert-info-wrapper {
        width: 60%;
        display: flex;
        text-align: center;
        flex-direction: column;
        justify-content: center;
        padding: var(--spacing-2);
        gap: var(--spacing-1);
        margin: 0 auto;
    }

    /* .iec-expert-info-wrapper * {
        border: 1px solid red;
    } */

    .iec-expert-name {
        font-family: var(--font-titles);
        font-size: var(--font-size-3);
        text-transform: none;
    }

    .iec-expert-title {
        font-family: var(--font-text);
        font-size: calc(.8 * var(--font-size-2));
        text-transform: none;
    }

    .iec-expert-tags {
        font-family: var(--font-text);
        font-size: calc(.8 * var(--font-size-2));
        text-transform: none;
        margin: 0;
    }
</style>