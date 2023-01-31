<script>
    import { push, link } from "svelte-spa-router";
    import DropDown from "./sub-components/DropDown.svelte";
    import bonhart_storage from "../../libs/bonhart-storage";
    import createColorSchema, {
        supported_components,
    } from "../../libs/ColorSchema";
    import { onMount, onDestroy } from "svelte";
    import itz_logo from "../../svg/MainLogo.svg";

    let site_page = window.location.hash.slice(1);
    const dropdown_names = {
        INICIO: {
            name: "inicio",
            href: '/',
            options: [
                {
                    name: "inicio",
                    link: "/",
                },
                {
                    name: "nosotros",
                    link: "/nosotros",
                },
                {
                    name: "servicios",
                    link: "/servicios",
                },
                {
                    name: "contacto",
                    link: "/contacto",
                },
            ]
        },
        SERVICIOS: {
            name: "servicios",
            href: '/servicios',
            options: [
                {
                    name: "mentoría",
                    link: "/mentoring",
                },
                {
                    name: "consultoría",
                    link: "/consulting",
                },
                {
                    name: "Aprendizaje grupal",
                    link: "/learning",
                },
                {
                    name: "comunidad",
                    link: "/community",
                },
            ]
        },
        EXPERTAS: {
            name: "expertas",
            href: '/expertas',
            options: [
                {
                    name: "Mentoras",
                    link: "/mentors",
                },
                {
                    name: "Consultoras",
                    link: "/consultants",
                },
                {
                    name: "Para agendar",
                    link: "/schedule",
                },
                {
                    name: "Reviews",
                    link: "/Reviews",
                },
                {
                    name: "Equipo interno",
                    link: "/internal-team",
                },
            ]
        },
        COMUNIDAD: {
            name: "comunidad",
            href: null,
            options: [
                {
                    name: "Objetivo",
                    link: "/objective",
                },
                {
                    name: "Membresías",
                    link: "/memberships",
                },
                {
                    name: "Eventos",
                    link: "/events",
                },
                {
                    name: "Beneficios",
                    link: "/benefits",
                },
            ]
        },
        CONTACTO: {
            name: "contacto",
            href: null,
            options: [
                {
                    name: "Datos",
                    link: "/data",
                },
                {
                    name: "Forms",
                    link: "/forms",
                },
            ]
        }
    };

    onMount(() => {
        const color_schema = createColorSchema(
            {
                color: "var(--clear-color)",
                contrast: "",
            },
            supported_components.NAVBAR
        );
        color_schema.define();

        window.addEventListener("hashchange", setActivePage);
    });

    onDestroy(() => {
        window.removeEventListener("hashchange", setActivePage);
    });

    function setActivePage() {
        let current_hash = window.location.hash;
        current_hash = current_hash.replace("#", "");

        if (current_hash !== site_page) {
            site_page = current_hash;
        }
    }
</script>

<nav id="itz-main-navbar">
    <div id="im-navbar-container">
        <div id="im-navbar-logo">
            <div id="im-lg-logo">
                {@html itz_logo}
            </div>
        </div>
        <div id="itz-options">
            {#each Object.values(dropdown_names) as ddn}
                <span class="itz-navoption drop-down-parent">
                    {#if ddn.href !== null}
                        <a href={ddn.href} use:link>
                            <span id={site_page === ddn.href ? "itz-current-route" : ""}>
                                {ddn.name + (console.log(`${site_page} === ${ddn.href}: ${site_page === ddn.href}`) || "")}
                                <!-- ths console log is for Debug purposes -->
                            </span>
                        </a>
                    {:else}
                        <span>
                            {ddn.name}
                        </span>
                    {/if}
                    <DropDown
                        drop_down_name={ddn.name}
                        drop_down_options={ddn.options}
                    />
                </span>
            {/each}
        </div>
    </div>
</nav>

<style>
    /*=============================================
    =            Header            =
    =============================================*/

    #itz-main-navbar {
        box-sizing: border-box;
        position: fixed;
        width: 100vw;
        max-width: 100%;
        height: var(--navbar-height);
        display: flex;
        align-items: center;
        padding: 0 var(--spacing-h3);
        background: var(--navbar-schema-background);
        top: 0;
        left: 0;
        border-bottom: 1px solid var(--primary-color-midlight);
        z-index: 2;
        transition: all 0.2s ease-in-out;
    }

    #im-navbar-container {
        display: flex;
        justify-content: space-between;
        align-items: center;
        width: 100%;
    }

    #im-navbar-logo {
        display: flex;
        align-items: center;
        justify-content: flex-start;
    }

    #im-lg-logo {
        width: 9.4vw;
        margin-right: 0.3rem;
    }

    :global(#im-lg-logo svg) {
        width: 100%;
        fill: var(--navbar-schema-color);
        /* height: 100%; */
    }

    #itz-options {
        display: flex;
        align-items: center;
        justify-content: flex-end;
        gap: var(--spacing-2);
        width: 50%;
    }

    .itz-navoption {
        font-family: "Times New Roman", Times, serif;
        display: flex;
        align-items: center;
        font-size: calc(1.2 * var(--font-size-3));
        margin: 0 1ex;
        color: var(--navbar-schema-color);
        cursor: default;
        padding: var(--spacing-1);
        border-radius: var(--boxes-roundness);
        text-transform: capitalize;
        transition: all 0.2s ease-in-out;
    }

    :global(.itz-navoption a) {
        color: var(--primary-color-midlight);
        text-decoration: none;
        transition: all 0.2s ease-in-out;
    }

    .itz-navoption.drop-down-parent {
        border-radius: var(--boxes-roundness) var(--boxes-roundness) 0 0 !important;
    }

    #itz-current-route {
        border-bottom: 3px solid var(--navbar-schema-color);
    }

    @media (pointer: fine) {
        .itz-navoption:hover {
            color: var(--clear-color) !important;
            background: var(--primary-color-midlight);
            box-shadow: 0px 0px 13px 3px rgba(0, 0, 0, 0.1) !important;
            transform: scale(1.05);
        }

        /* :global(.itz-navoption:hover a) {
            color: var(--clear-color) !important;
        } */
    }
</style>
