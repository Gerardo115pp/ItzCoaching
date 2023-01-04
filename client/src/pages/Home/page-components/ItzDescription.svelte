<script>

</script>

<div id="itz-desc-content-wrapper">
    <div id="idcw-media-wrapper">
        <video controls=true src="/resources/videos/home_video.mp4">
            <track kind="captions"/>
        </video>
    </div>
    <div id="idcw-description">
        <p id="itz-description">
            <span class='itz-highlight'>Empoderamos</span> a <i>mujeres</i> fomentando su<br/><span class='itz-highlight'>colaboración</span> y desarrollo profesional para que alcancen el <strong>éxito</strong> en<br/>el mundo empresarial.
        </p>
    </div>

</div>

<style>
    #itz-desc-content-wrapper {
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 100%;
        max-width: 100vw;
        padding: calc(var(--navbar-height) + var(--spacing-h2)) 0 0 0;
    }

    #idcw-media-wrapper {
        width: 70%;
    }

    #idcw-media-wrapper video {
        width: 100%
    }

    #idcw-description {
        display: grid;
        width: 100%;
        height: calc(100vh - var(--navbar-height));
        margin-top: var(--spacing-h3);
        place-items: center;
    }

    #itz-description {
        font-size: var(--font-size-h3);
        text-align: center;
        line-height: 1.4;
    }
    
    #itz-description .itz-highlight {
        color: var(--theme-purple);
    }
    
    #itz-description i {
        font-family: var(--font-moontime);
        font-size: calc(1.5 * var(--font-size-h1));
        line-height:.5;
    }

    #itz-description strong {
        font-weight: 600;
    }
</style>