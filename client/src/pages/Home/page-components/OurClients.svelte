<script>
    const clients_response = [
        {
            name: "Arte Espacio",
            physical_client_name: "Fer Gianacópulos",
            logo: "/resources/client_logos/arte_espacio.png",
            website: "https://www.arte-espacio.com/",
            sessions: 30
        },
        {
            name: "atelina",
            physical_client_name: "Cecy Rodríguez",
            logo: "/resources/client_logos/atelina.png",
            website: "https://atelina.com/",
            sessions: 22
        },
        {
            name: "Cake a wish",
            physical_client_name: "María José",
            logo: "/resources/client_logos/cake_wish.png",
            website: "https://www.cakeawish.com.ar/",
            sessions: 26
        },
        {
            name: "Ellucero",
            physical_client_name: "Lucero",
            logo: "/resources/client_logos/ellucero.png",
            website: "https://www.ellucero.com/",
            sessions: 26
        },
        {
            name: "Recode media",
            physical_client_name: "Ana de Armas",
            logo: "/resources/client_logos/recode.png",
            website: "https://www.recode.media/",
            sessions: 26
        },
        {
            name: "Giana",
            physical_client_name: "Martha Giana",
            logo: "/resources/client_logos/giana.png",
            website: "https://www.giana.com.ar/",
            sessions: 2
        },
        {
            name: "glam on the beach",
            physical_client_name: "Tatiana",
            logo: "/resources/client_logos/glam_on_the_beach.png",
            website: "https://www.glamonthebeach.com/",
            sessions: 2
        },
        {
            name: "Movimiento camino a la salud",
            physical_client_name: "Natalia",
            logo: "/resources/client_logos/camino_salud.png",
            website: "https://www.movimientocaminoalasalud.com/",
            sessions: 2
        },
        {
            name: "Vef",
            physical_client_name: "Verónica",
            logo: "/resources/client_logos/vef.png",
            website: "https://www.vef.com.ar/",
            sessions: 2
        },
        {
            name: "Print your life",
            physical_client_name: "Gisela",
            logo: "/resources/client_logos/print_life.png",
            website: "https://www.printyourlife.com.ar/",
            sessions: 2
        },
        {
            name: "Skinni Tahini",
            physical_client_name: "Yael",
            logo: "/resources/client_logos/skinni_tree.png",
            website: "https://www.skinnitahini.com/",
            sessions: 2
        },
        {
            name: "Gaby Flo",
            physical_client_name: "Gabriela Flores",
            logo: "/resources/client_logos/gaby_flo.png",
            website: "https://www.gabyflo.com/",
            sessions: 2
        },
        {
            name: "Centro Baniano",
            physical_client_name: "Sofía",
            logo: "/resources/client_logos/centro_baniano.png",
            website: "https://www.centrobaniano.com/",
            sessions: 22
        },
        {
            name: "Paulina Cervantes Denimart",
            physical_client_name: "Paulina Cervantes",
            logo: "/resources/client_logos/denimart.png",
            website: "https://www.paulinacervantes.com/",
            sessions: 12
        },
        {
            name: "Holy Butter",
            physical_client_name: "Eliana",
            logo: "/resources/client_logos/holy_butter.png",
            website: "https://www.holybutter.com/",
            sessions: 12
        },
        {
            name: "CAPI, Centro de Atención Psicológica Integral",
            physical_client_name: "Cecilia gomez",
            logo: "/resources/client_logos/capi.png",
            website: "https://www.capi.com.ar/",
            sessions: 12
        },
        {
            name: "Fincoach",
            physical_client_name: "Ursula",
            logo: "/resources/client_logos/fincoach.png",
            website: "https://www.fincoach.com.ar/",
            sessions: 3
        },
        {
            name: "Detalles Dulces",
            physical_client_name: "Zulma",
            logo: "/resources/client_logos/detalles_dulces.png",
            website: "https://www.detallesdulces.com.ar/",
            sessions: 2
        },
        {
            name: "Un cariñito",
            physical_client_name: "Oriana",
            logo: "/resources/client_logos/un_carinito.png",
            website: "https://www.uncarinito.com.ar/",
            sessions: 1
        },
        {
            name: "Pabela",
            physical_client_name: "Paula",
            logo: "/resources/client_logos/pabela.png",
            website: "https://www.pabela.com.es/",
            sessions: 2
        },
        {
            name: "Medical Meet",
            physical_client_name: "Marcela",
            logo: "/resources/client_logos/medical_meet.png",
            website: "https://www.medicalmeet.com.ar/",
            sessions: 3
        },
        {
            name: "Insight Evolution",
            physical_client_name: "Jasmine",
            logo: "/resources/client_logos/insight_evolution.png",
            website: "https://www.insightevolution.com/",
            sessions: 4
        },
        {
            name: "Annie de Hoyos",
            physical_client_name: "Annie de Hoyos",
            logo: "/resources/client_logos/annie_de_hoyos.png",
            website: "https://www.anniehoyos.com/",
            sessions: 5
        },
        {
            name: "La trividad",
            physical_client_name: "Karina",
            logo: "/resources/client_logos/la_trividad.png",
            website: "https://www.latrividaddelosnumeros.com/",
            sessions: 6
        }

    ];
</script>

<div id="itz-our-clients-wrapper">
    <h2 id="iocw-title">Nuestros clientes</h2>
    <div id="itz-our-clients-container">
        {#each clients_response as client}
            <div class="itz-client">
                <div class="ic-logo-wrapper">
                    <img src="{client.logo}" alt="cliente {client.name}"/>
                </div>
                <div class="ic-data-wrapper glassmorphism">
                    <span class="ic-dw-client-name">
                        {client.physical_client_name}
                    </span>
                    <span class="ic-dw-session-count">
                        {client.sessions} sesiones
                    </span>
                </div>
            </div>
        {/each}
    </div>
</div>

<style>
    #itz-our-clients-wrapper {
        min-height: 100vh;
        padding: var(--page-padding-2);
    }

    #iocw-title {
        text-transform: none;
        font-size: var(--font-size-h2);
        font-weight: 600;
    }

    #itz-our-clients-container {
        --client-container-width: 75vw;
        --client-row-ratio: 6;
        --client-gap: var(--spacing-1);
        --client-width-height: calc(calc(var(--client-container-width) / var(--client-row-ratio)) - calc(var(--client-gap) * 2));
        
        box-sizing: border-box;
        width: var(--client-container-width);
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        gap: var(--client-gap);
        margin: var(--spacing-3) auto;
    }

    .itz-client {
        box-sizing: border-box;
        width: var(--client-width-height);
        height: var(--client-width-height);
        display: grid;
        grid-template: repeat(2, 1fr) / repat(2, 1fr);
        margin-bottom: var(--spacing-3);
    }
    
    .ic-logo-wrapper {
        display: flex;
        width: 100%;
        height: 100%;
        grid-area: 1 / 1 / 3 / 3;
        justify-content: center;
        align-items: center;
    }
    
    .ic-logo-wrapper img {
        width: calc(var(--client-width-height) * .8);
        max-height: calc(var(--client-width-height) * .8);
        object-fit: contain;
    }
    
    .ic-data-wrapper {
        grid-area: 1 / 1 / 3 / 3;
        display: flex;
        justify-content: center;
        flex-direction: column;
        align-items: center;
        color: var(--dark-color);
        font-size: var(--font-size-3);
        border-radius: var(--button-roundness);
        opacity: 0;
        transition: all .4s ease-in;

    }

    @media(pointer: fine) {
        .itz-client:hover .ic-data-wrapper {
            opacity: 1;
            transform: scale(1.1);
            animation: slit-in-diagonal-1 0.45s ease-out both;
        }
    }
    
    .ic-dw-client-name {
        font-weight: 600;
    }

    .ic-dw-session-count {
        font-family: var(--font-titles);
    }
</style>