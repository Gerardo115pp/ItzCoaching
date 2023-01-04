<div id="itz-stats-wrapper">
    <h2 id="isw-title">ItzCo en números</h2>
    <div id="isw-stats-container">
        <div class="itz-stat">
            <h1 class="itz-stat-value">15</h1>
            <h3 class="itz-stat-name">expertas</h3>
        </div>
        <div class="itz-stat">
            <h1 class="itz-stat-value">+198</h1>
            <h3 class="itz-stat-name">sesiones impartidas</h3>
        </div>
        <div class="itz-stat">
            <h1 class="itz-stat-value">+64</h1>
            <h3 class="itz-stat-name">emprendedoras</h3>
        </div>
        <div class="itz-stat">
            <h1 class="itz-stat-value">10</h1>
            <h3 class="itz-stat-name">áreas de negocio</h3>
        </div>
        <div class="itz-stat">
            <h1 class="itz-stat-value">+83</h1>
            <h3 class="itz-stat-name">diagnósticos realizados</h3>
        </div>
        <div class="itz-stat">
            <h1 class="itz-stat-value">3</h1>
            <h3 class="itz-stat-name">países</h3>
        </div>
    </div>
</div>

<style>
    #itz-stats-wrapper {
        box-sizing: border-box;
        display: flex;
        flex-direction: column;
        height: 100vh;
        background: var(--theme-pearl);
        color: var(--dark-color);
        padding: var(--page-padding-1);
        justify-content: center;
    }

    #isw-title {
        color: var(--theme-red);
        font-size: var(--font-size-h2);
        font-weight: 600;
        text-transform: none;
    }

    #isw-stats-container {
        display: flex;
        width: 80%;
        flex-wrap: wrap;
        justify-content: space-around;
        margin: var(--spacing-4) auto;
        row-gap: var(--spacing-h1);
    }

    .itz-stat {
        display: grid;
        width: 30%;
        font-size: var(--font-size-h1);
        place-items: center;
    }

    .itz-stat-value {
        font-size: 1.5em;
        font-weight: 600;
    }

    .itz-stat-name {
        font-size: var(--font-size-h4);
        font-weight: 400;
    }
</style>