<script>
    import { Appointment } from '../../../components/DateComponents/date_utils'
    import { GetExpertAppointments } from '../../../libs/HttpRequests';
    import { onMount } from 'svelte';

    export let expert_id;

    let appointments = {};
    let appointment_count = 0;
    $: console.log(appointments);

    onMount(() => {
        requestAppointments();
    });

    const requestAppointments = () => {
        const appointments_request = new GetExpertAppointments(expert_id);

        const on_success = (response_data) => {
            appointment_count = response_data.length;
            mapDaysToAppointments(response_data);
        };

        const on_error = (error) => {
            console.log(error);
        };

        appointments_request.do(on_success, on_error);
    };

    const  mapDaysToAppointments = all_appointments => {
        let new_appointments = {};
        let expert_appointments = all_appointments.map(a => {
            const appointment = new Appointment(a, 'es');
            return appointment;
        });

        expert_appointments.sort((a, b) => {
            return a.Start.getTime() - b.Start.getTime();
        });

        for (let a of expert_appointments) {
            new_appointments[a.dateString()] = [] || new_appointments[a.dateString()];

            new_appointments[a.dateString()].push(a);
        }

        appointments = {...new_appointments};
    }

</script>

<article id="itz-expert-appointments">
    <headers id="iea-appointments-header">
        <h3>
            Tu calendario
        </h3>
    </headers>
    {#if appointment_count > 0}
        <div id="iea-appointments-days-container">
            {#each Object.keys(appointments) as appointment_day}
                 <div class="appointment-day-wrapper">
                    <h3>{appointment_day}</h3>
                    <hr>
                    <div class="appointments-wrapper">
                        {#each appointments[appointment_day] as appointment}
                            <div class="appointment">
                                <p class="appointment-email">{appointment.customer_email}</p>
                                <p class="appointment-time">{appointment.timeString()}</p>
                            </div>
                        {/each}
                    </div>
                </div>
            {/each}
        </div>
    {/if}
</article>

<style>
    /* #itz-expert-appointments * {
        border: 1px solid var(--theme-red);
    } */

    #itz-expert-appointments {
        box-sizing: border-box;
        border-radius: var(--boxes-roundness);
        border: 1px solid var(--dark-light-color);
        overflow: hidden;
    }
    
    #iea-appointments-header {
        display: block;
        width: 100%;
        box-sizing: border-box;
        font-size: var(--font-size-4);
        background: var(--theme-purple);
        color: var(--theme-pearl);
        padding: var(--spacing-3);
    }

    #iea-appointments-days-container {
        height: 440px;
        padding: var(--spacing-3);
        overflow-y: auto;
    }

    .appointment-day-wrapper {
        display: grid;
        grid-template-columns: repeat(6, 1fr);
        height: 80%;
    }

    .appointment-day-wrapper h3 {
        grid-column: 1 / 2;
        height: max-content;
        font-size: var(--font-size-3);
        color: var(--dark-color);
        text-align: center;
    }

    .appointment-day-wrapper hr {
        width: 100%;
        height: 2px;
        grid-column: 2 / 7;
        background-color: var(--dark-light-color);
        border: 1px solid var(--dark-light-color);
    }

    .appointments-wrapper {
        grid-column: 1 / 7;
        height: 300px
    }

    .appointment {
        width: max-content;
        color: var(--theme-pearl);
        background: var(--theme-purple);
        padding: var(--spacing-1);
        border-radius: var(--boxes-roundness);
    }

    .appointment-time {
        font-weight: bold;
    }
</style>