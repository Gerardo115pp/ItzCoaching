<script>
    import Router from "svelte-spa-router";
    import { routes } from "./routes.js";
    import Navbar from "./components/Navbar/Navbar.svelte";
    import NotificationsPopup from "./components/notifications/NotificationsPopup.svelte";
</script>

<style global>
    :root {
        /* Colors */
        --theme-red: #e84c4c;
        --theme-light-red: #ee7171;
        --theme-dark-red: #c63939;
        --theme-transparent-red: #e84c4c80;
        --theme-pearl: #FCEEDE;
        --theme-light-pearl: #F3F4F1;
        --theme-dark-pearl: #dfd6d0;
        --theme-transparent-pearl: #fceedee2;
        --theme-purple: #A475AF;
        --theme-light-purple: #b384bd;
        --theme-dark-purple: #8e5f8a;
        --theme-transparent-purple: #a475af80;
        --dark-color: #260B2D;
        --dark-transparent-color: #260b2dc3;
        --dark-light-color: #c5b5cb;
        --placeholder-color: #00000025;
        --clear-color: white;
        --clear-transparent-color: #ffffff80;
        --clear-background: white;
        --theme-gradient: linear-gradient(270deg, #F3F4F1 0%, #E6E2DF 100%);
        --danger: rgb(206, 18, 18);
        --ready: #62C77A;
        --trasparent-theme-color: #ccff6c88;
        --dark-tranparent-background: rgba(0, 0, 0, 0.3);
        


        /* Fonts */
        --font-titles: 'Libre Baskerville', cursive;
        --font-text: 'open-sans-v', sans-serif;
        --font-moontime: 'moontime', cursive;
        --font-subtitles: 'DM Sans', cursive;

        --font-size-1: 1.2257405515832482vh;
        --font-size-2: calc(var(--font-size-1) + calc(var(--font-size-1) * 0.55));
        --font-size-3: calc(var(--font-size-1) + calc(var(--font-size-1) * 0.7));
        --font-size-4: calc(var(--font-size-1) + calc(var(--font-size-1) * 1.1));
        --font-size-h1: calc(var(--font-size-1) * 5.5);
        --font-size-h2: calc(var(--font-size-h1) * 0.80);
        --font-size-h3: calc(var(--font-size-h1) * 0.6);
        --font-size-h4: calc(var(--font-size-h1) * 0.4);
        --font-size-small: calc(var(--font-size-1) * 0.8);
        
        font-size: var(--font-size-1);

        
        /* Spacing */
        --spacing-1: 0.6rem;
        --spacing-2: calc(var(--spacing-1) * 2);
        --spacing-3: calc(var(--spacing-1) * 3);
        --spacing-4: calc(var(--spacing-1) * 4);
        --spacing-h1: calc(var(--spacing-4) * 4);
        --spacing-h2: calc(var(--spacing-h1) * 0.75);
        --spacing-h3: calc(var(--spacing-h1) * 0.5);
        --spacing-h4: calc(var(--spacing-h1) * 0.25);

        --page-padding-1: var(--spacing-4) var(--spacing-h2);
        --page-padding-2: var(--spacing-h3);
        --page-padding-top: calc(var(--spacing-4) + var(--navbar-height)) var(--spacing-h2) var(--spacing-4);

        
        /* Misc */
        --navbar-height: 15vh; /* Not used yet */
        --footer-height: 56.72vh; /* Not used yet */
        --boxes-roundness: 6px;
        --letter-spacing: 0.39936px;
        --boxes-shadow: 0px 5px 4px 1px rgba(0, 0, 0, 0.1);
        letter-spacing: var(--letter-spacing);
        
        /* buttons */
        --buttons-padding: var(--spacing-2) var(--spacing-h3);
        --buttons-padding-thin: var(--spacing-1) var(--spacing-3);
        --buttons-text-transform: capitalize;
        --btn-font-size: var(--font-size-2);
        --button-roundness: var(--boxes-roundness);
        --button-shadow: var(--boxes-shadow);
        --button-font-weight: 500;
        
        
     
    }


    
/*=============================================
=            Normalize            =
=============================================*/

    
        * {
            box-sizing: border-box;
        }
    
        :global(body) {
            margin: 0;
            padding: 0;
            font-family: var(--font-text);
        }
    
        
        
/*=====  End of Normalize  ======*/
        
        
        
    
/*=============================================
=            Theme styles            =
=============================================*/

        :global(h1, h2, h3, h4, h5, h6) {
            font-family: var(--font-titles);
            font-weight: 400;
            text-transform: uppercase;
            margin: 0;
        }

        :global(.itz-subtitle) {
            font-family: var(--font-subtitles);
            font-weight: lighter;
            color: var(--theme-red);
        }

        :global(a) {
            color: var(--theme-color);
            text-decoration: none;
        }
    
        :global(.page-container) {
            width: 100%;
            min-height: calc(100vh - var(--navbar-height));
        }


        :global(.page-title) {
            text-transform: uppercase;
            font-size: var(--font-size-h1);
            font-weight: 300;
        }

        :global(.full-btn) {
            font-family: var(--font-text);
            box-sizing: border-box;
            padding: var(--buttons-padding);
            color: var(--theme-pearl);
            background-color: var(--theme-red);
            border-radius: var(--button-roundness);
            font-weight: var(--button-font-weight);
            font-size: var(--btn-font-size);
            border: none;
            box-shadow: var(--button-shadow);
            transition: all 0.3s ease-in-out;
            text-transform: var(--buttons-text-transform);
        }

        :global(.full-btn-thin) {
            font-family: var(--font-text);
            box-sizing: border-box;
            padding: var(--buttons-padding-thin);
            color: var(--theme-pearl);
            background-color: var(--theme-red);
            border-radius: var(--button-roundness);
            font-weight: var(--button-font-weight);
            font-size: var(--btn-font-size);
            border: none;
            box-shadow: var(--button-shadow);
            transition: all 0.3s ease-in-out;
            text-transform: var(--buttons-text-transform);
        }
        
        :global(.clear-btn) {
            font-family: var(--font-text);
            box-sizing: border-box;
            padding: var(--buttons-padding);
            color: var(--theme-red);
            background-color: var(--theme-pearl);
            border-radius: var(--button-roundness);
            font-weight: var(--button-font-weight);
            font-size: var(--btn-font-size);
            border: none;
            box-shadow: var(--button-shadow);
            transition: all 0.3s ease-in-out;
            text-transform: var(--buttons-text-transform);

        }

        :global(.full-two-btn) {
            font-family: var(--font-text);
            box-sizing: border-box;
            padding: var(--buttons-padding);
            color: var(--theme-pearl);
            background-color: var(--theme-purple);
            border-radius: var(--button-roundness);
            font-weight: var(--button-font-weight);
            font-size: var(--btn-font-size);
            border: none;
            box-shadow: var(--button-shadow);
            transition: all 0.3s ease-in-out;
            text-transform: var(--buttons-text-transform);
        }

        :global(.full-two-btn-thin) {
            font-family: var(--font-text);
            box-sizing: border-box;
            padding: var(--buttons-padding-thin);
            color: var(--theme-pearl);
            background-color: var(--theme-purple);
            border-radius: var(--button-roundness);
            font-weight: var(--button-font-weight);
            font-size: var(--btn-font-size);
            border: none;
            box-shadow: var(--button-shadow);
            transition: all 0.3s ease-in-out;
            text-transform: var(--buttons-text-transform);
        }

        :global(.clear-two-btn) {
            font-family: var(--font-text);
            box-sizing: border-box;
            padding: var(--buttons-padding);
            color: var(--theme-purple);
            background-color: var(--theme-pearl);
            border-radius: var(--button-roundness);
            font-weight: var(--button-font-weight);
            font-size: var(--btn-font-size);
            border: none;
            box-shadow: var(--button-shadow);
            transition: all 0.3s ease-in-out;
            text-transform: var(--buttons-text-transform);
        }

        :global(.icon-wrapper) {
            display: flex;
            align-items: center;
            justify-content: center;
            width: var(--spacing-2);
            margin: 0;
        }

        :global(.glassmorphism) {
            background: rgba(255, 171, 171, 0.25);
            box-shadow: 0 8px 32px 0 rgba( 31, 38, 135, 0.37 );
            border: 1px solid rgba(59, 58, 58, 0.18);
            backdrop-filter: blur(2px);
        }

        @media (pointer: fine) {
            
            :global(.full-btn:hover, .full-btn-thin:hover) {
                background-color: var(--theme-light-red);
                transform: scale(1.05);
                box-shadow: 0px 5px 8px 3px rgba(0, 0, 0, 0.18);
            }
            
            :global(.clear-btn:hover) {
                background-color: var(--theme-dark-pearl);
                transform: scale(1.05);
                box-shadow: 0px 5px 8px 3px rgba(0, 0, 0, 0.18);
            }

            :global(.full-two-btn:hover) {
                background-color: var(--theme-light-purple);
                transform: scale(1.05);
                box-shadow: 0px 5px 8px 3px rgba(0, 0, 0, 0.18);
            }

            :global(.clear-two-btn:hover) {
                background-color: var(--theme-dark-pearl);
                transform: scale(1.05);
                box-shadow: 0px 5px 8px 3px rgba(0, 0, 0, 0.18);
            }
        }

        :global(.round-image) {
            border-radius: var(--boxes-roundness);
            box-shadow: 0px 5px 17px 4px rgb(0 0 0 / 26%);
        }

        :global(.round-box) {
            border-radius: var(--boxes-roundness);
            box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.1);
        }

        :global(.divider) {
            border:1px solid var(--theme-one-color) !important;
            background-color: var(--clear-one-color);
            margin: 0;
        }


    
    
/*=====  End of Theme styles  ======*/
    

/*=============================================
=            Fonts            =
=============================================*/

    @font-face {
        font-family: 'open-sans-v';
        src: url('/resources/fonts/OpenSansV.ttf') format('truetype');
        font-weight: 400;
        font-style: normal;
        font-variation-settings: 'wght' 400 'wdth' 80;
    }

    @font-face {
        font-family: 'moontime';
        src: url('/resources/fonts/moontime.ttf') format('truetype');
        font-weight: 500;
        font-style: normal;
    }

/*=====  End of Fonts  ======*/


    
/*=============================================
=            mobile            =
=============================================*/
    
    @media only screen and (max-width: 768px) {
        :global(:root) {
            --font-size-1: 1.3278855975485189vh;
            --font-size-h1: calc(var(--font-size-1) * 3.5);

            --spacing-1: 2vw;
            --navbar-height: 12.172949002217294vh; /* Not used yet */
            --footer-height: 89.72vh; /* Not used yet */
            --boxes-roundness: 6px;
            --letter-spacing: 0.39936px;

            --btn-font-size: var(--font-size-2);

            --buttons-padding: var(--spacing-2) var(--spacing-2);
        }

        :global(.icon-wrapper) {
                display: flex;
                align-items: center;
                justify-content: center;
                width: var(--spacing-2);
                margin: 0;
        }
        
    }

/*=====  End of mobile  ======*/
    
    
    
/*=============================================
=            Animations            =
=============================================*/

    @keyframes slide-fwd-bottom {
        0% {
            -webkit-transform: translateZ(160px) translateY(-100px);
                    transform: translateZ(160px) translateY(-100px);
        }
        100% {
            -webkit-transform: translateZ(0) translateY(0);
                    transform: translateZ(0) translateY(0);
        }
    }

    :global(.slide-fwd-bottom) {
        -webkit-animation: slide-fwd-bottom 0.45s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
        animation: slide-fwd-bottom 0.45s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
    }
    
    @keyframes scale-in-ver-top {
        0% {
            -webkit-transform: scaleY(0);
                    transform: scaleY(0);
            -webkit-transform-origin: 100% 0%;
                    transform-origin: 100% 0%;
            opacity: 1;
        }
        100% {
            -webkit-transform: scaleY(1);
                    transform: scaleY(1);
            -webkit-transform-origin: 100% 0%;
                    transform-origin: 100% 0%;
            opacity: 1;
        }
    }
    
    :global(.scale-in-ver-top) {
        -webkit-animation: scale-in-ver-top 0.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
                animation: scale-in-ver-top 0.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
    }

    @keyframes -global-slit-in-diagonal-1 {
        0% {
            -webkit-transform: translateZ(-800px) rotate3d(1, 1, 0, 90deg);
                    transform: translateZ(-800px) rotate3d(1, 1, 0, 90deg);
            -webkit-animation-timing-function: ease-in;
                    animation-timing-function: ease-in;
            opacity: 0;
        }
        54% {
            -webkit-transform: translateZ(-160px) rotate3d(1, 1, 0, 87deg);
                    transform: translateZ(-160px) rotate3d(1, 1, 0, 87deg);
            -webkit-animation-timing-function: ease-in-out;
                    animation-timing-function: ease-in-out;
            opacity: 1;
        }
        100% {
            -webkit-transform: translateZ(0) rotate3d(1, 1, 0, 0);
                    transform: translateZ(0) rotate3d(1, 1, 0, 0);
            -webkit-animation-timing-function: ease-out;
                    animation-timing-function: ease-out;
        }
    }

    :global(.slit-in-diagonal-1) {
        -webkit-animation: slit-in-diagonal-1 0.45s ease-out both;
                animation: slit-in-diagonal-1 0.45s ease-out both;
    }

/*=====  End of Animations  ======*/
    
    
    
/*=============================================
=            DEBUG            =
=============================================*/
    
        :global(.debug) {
            border: 1px solid red !important;
        }

        :global(.debug-green) {
            border: 1px solid red;
            background-color: greenyellow;
        }
        
        :global(.debug-blue) {
            border: 1px solid red;
            background-color: blue;
        }

        :global(.debug-yellow) {
            border: 1px solid red;
            background-color: yellow;
        }

/*=====  End of DEBUG  ======*/

</style>

<svelte:head>
    <link rel="preconnect" href="https://fonts.googleapis.com"><link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=DM+Sans&family=Libre+Baskerville:ital,wght@0,400;0,700;1,400&display=swap" rel="stylesheet">
    <link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200" />
</svelte:head>
<div id="page-content">
    <NotificationsPopup />
    <Navbar/>
    <div id="content-wrap">
        <Router {routes}/>
    </div>
</div>
