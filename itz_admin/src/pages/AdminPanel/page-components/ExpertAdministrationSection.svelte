<script>
    import RegisterExpert from "./RegisterExpert.svelte";
    import ExpertsList from "./ExpertsList.svelte";

    const section_options = {
        REGISTER_EXPERT: "register-expert",
        LIST_EXPERTS: "list-experts",
    };

    let selected_option = section_options.LIST_EXPERTS;

    const setContentOption = content_option => {
        selected_option = content_option;
    }
</script>
<article id="register-expert-admin-component">
    <!-- IMPORTED BY THE PAGE COMPONENT -->
    <header id="reac-header">
        <h4>administracion de expertas</h4>
    </header>
    <section id="reac-action-options">
        <button on:click={() => setContentOption(section_options.REGISTER_EXPERT)} class="reac-option full-btn-thin">Registrar experta</button>
        <button on:click={() => setContentOption(section_options.LIST_EXPERTS)} class="reac-option full-btn-thin">Listado</button>
    </section>
    <section id="selected-option-content">
        <div class="content-option-wrapper {selected_option === section_options.REGISTER_EXPERT ? 'selected-content-option' : '' }">
            <RegisterExpert/>
        </div>
        <div class="content-option-wrapper {selected_option === section_options.LIST_EXPERTS ? 'selected-content-option' : '' }">
            <ExpertsList/>
        </div>
    </section>
</article>

<style>
    #register-expert-admin-component {
        box-sizing: border-box;
        border: 1px solid var(--dark-light-color);
        padding: var(--spacing-4);
        border-radius: var(--boxes-roundness);
        box-shadow: var(--boxes-shadow);
    }

    #register-expert-admin-component header h4 {
        text-transform: none;
        font-size: var(--font-size-3);
        font-family: var(--font-text);
        font-weight: 600;
        color: var(--dark-color);
    }

    #reac-action-options {
        display: flex;
        margin: var(--spacing-4);
        gap: var(--spacing-2);
    }

    .reac-option {
        text-transform: lowercase;
    }

    
    /*=============================================
    =            selected content            =
    =============================================*/
    
    .content-option-wrapper {
        display: none;
    }

    .content-option-wrapper.selected-content-option {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
    }
    
    
    /*=====  End of selected content  ======*/
    
    

</style>