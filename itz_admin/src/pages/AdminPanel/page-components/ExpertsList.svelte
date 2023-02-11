<script>
    import { onMount } from 'svelte';
    import bonhart_storage from '../../../libs/bonhart-storage';
    import { GetAllExpertsRequest, DeleteExpertRequest, UpdateExpertActiveStatusRequest } from '../../../libs/HttpRequests';
    import Switch from '../../../components/Input/Switch.svelte';

    let experts_list = [];
    let inactive_experts_metric = 0; // experts that cannot access the platform because an admin has deactivated their account
    let unavailable_experts_metric = 0; // experts that have an active account but have set their schedule as unavailable

    onMount(() => {
        getExperts();
    });

    const deleteExpert = id => {
        const delete_request = new DeleteExpertRequest(bonhart_storage.Token, id);

        const on_success = (response) => {
            console.log(response);
            experts_list = experts_list.filter(expert => expert.id !== id);
        }

        const on_error = (error) => {
            console.log(error);
        }

        delete_request.do(on_success, on_error);
    }

    const toggleExpertActiveStatus = (id, is_active) => {
        is_active = !is_active;

        const update_request = new UpdateExpertActiveStatusRequest(bonhart_storage.Token, id, is_active);

        const on_success = (response) => {
            console.log(response);
            inactive_experts_metric += is_active ? -1 : 1;
        }

        const on_error = (error) => {
            console.log(error);
        }

        update_request.do(on_success, on_error);
    }

    function getExperts() {
        const expert_request = new GetAllExpertsRequest(bonhart_storage.Token);

        const on_success = (response) => {
            if (response !== null) {
                experts_list = response;
                experts_list.forEach(expert => {
                    inactive_experts_metric += expert.is_active ? 0 : 1;
                    unavailable_experts_metric += expert.is_available ? 0 : 1;
                });
            }
        }

        const on_error = (error) => {
            console.log(error);
        }

        expert_request.do(on_success, on_error);
    }
</script>

<div id="expert-list-wrapper">
    <header id="experts-metrics">
        <div class="expert-metric-wrapper">
            <span class="expert-metric-name">
                expertas
            </span>
            <span class="expert-metric-value">
                {experts_list.length}
            </span>
        </div>
        <div class="expert-metric-wrapper">
            <span class="expert-metric-name">
                disponibles
            </span><span class="expert-metric-value">
                {experts_list.length - unavailable_experts_metric}
            </span>
        </div>
        <div class="expert-metric-wrapper">
            <span class="expert-metric-name">
                inactivas
            </span>
            <span class="expert-metric-value">
                {inactive_experts_metric}
            </span>
        </div>
    </header>
    <!-- <hr/> -->
    <ul id="expert-list">
        {#each experts_list as expert}
            <li class="expert-item">
                <div class="ei-column expert-identifier">
                    <span class="expert-name">
                        {expert.name}
                    </span>
                    <span class="expert-email">
                        ({expert.email})
                    </span>
                </div>
                <div class="ei-column expert-data">
                    <span class="expert-availability {expert.is_available ? 'expert-available-status' : ''}">
                        {expert.is_available ? 'disponible' : 'no disponible'}
                    </span>
                </div>
                <div class="ei-column expert-controls">
                    <div class="expert-control">
                        <span class="ec-control-name">
                            activa
                        </span>
                        <span class="ec-control-input">
                            <Switch checked={expert.is_active} onChange={() => toggleExpertActiveStatus(expert.id, expert.is_active)} />
                        </span>
                    </div>
                    <div class="expert-control">
                        <button on:click={() => deleteExpert(expert.id)} class="ec-control-input full-btn-thin delete-expert-btn material-symbols-outlined">
                            close
                        </button>
                    </div>
                    
                </div>
            </li>
        {/each}
    </ul>
</div>

<style>
    #expert-list-wrapper {
        width: 60%;
    }

    #experts-metrics {
        display: flex;
        align-items: center;
        margin-bottom: 1rem;
    }

    .expert-metric-wrapper {
        width: max(14%, 8vw);
        font-size: var(--font-size-2);
    }

    .expert-metric-name {
        font-weight: 600;
    }

    .expert-metric-name::after {
        content: ': ';
    }

    .expert-metric-value {
        color: var(--theme-red);
    }

    
    /*----------  Expert list  ----------*/
    
    #expert-list {
        --expert-list-item-height: calc(1.6 * var(--spacing-h4));

        width: 100%;
        height: calc(6 * var(--expert-list-item-height));
        overflow-y: scroll;
        font-size: var(--font-size-2);
        padding: 0;
        margin: 0;
        list-style: none;
    }

    .expert-item {
        box-sizing: border-box;
        display: flex;
        background: var(--clear-color);
        height: var(--expert-list-item-height);
        color: var(--dark-color);
        align-items: center;
        border-bottom: 1px solid var(--theme-gray-2);
    }

    .ei-column {
        padding: 0 var(--spacing-3);
    }

    .expert-item:nth-child(odd) {
        background: var(--theme-light-red);
    }

    /* DEBUG: BOXES */
    /* .expert-item * {
        border: 1px solid blue;
    } */

    .expert-identifier {
        display: flex;
        width: 33%;
        background: var(--theme-red);
        height: 98%;
        color: white;
        align-items: center;
        border-bottom: 1px solid white;
    }

    .expert-available-status {
        color: var(--ready-strong);
    }

    .expert-data {
        width: 33%;
    }

    .expert-controls {
        width: 33%;
        display: flex;
        align-items: center;
        justify-content: flex-end;
        column-gap: var(--spacing-1);
    }

    .expert-control {
        display: flex;
        width: 25%;
        align-items: center;
        justify-content: center;
        
    }

    .ec-control-name {
        font-weight: 600;
    }

    .ec-control-name::after {
        content: ': ';
    }



</style>