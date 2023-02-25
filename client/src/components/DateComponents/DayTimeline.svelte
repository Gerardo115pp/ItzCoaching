<script>
    import { newNotification } from '../notifications/events';
    import { getDayHours, TimeSlot } from './time_utils';
    import { onMount, onDestroy } from "svelte";

    // STYLE CUSTOMIZATION
    export let component_id = "libery-timeline";
    export let font_titles = "Roboto";
    export let font_body = "sans-serif";
    export let dark_color_1 = "#3c3c3c";
    export let dark_color_2 = "#7d7d7d";
    export let dark_color_3 = "#b7b7b7";
    export let light_color_1 = "#ffffff";
    export let light_color_2 = "#f2f2f2";
    export let primary_color = "#ff5a5f";
    export let primary_color_light = "#ff5a5f";
    export let secondary_color = "#54cae8";
    export let available_highlight_color = "#6fca13";
    export let unavailable_highlight_color = "#ff5a5f";
    export let font_size_1 = "8px";
    export let font_size_2 = "14px";
    export let font_size_3 = "18px";
    export let font_size_h3 = "22px";
    export let font_size_h2 = "28px";
    export let font_size_h1 = "34px";
    export let boxShadow = "0px 0px 10px 0px rgba(0,0,0,0.75)";


    // CUSTOMIZATION
    export let available_time_slots = []; // the hours in which the schedule owner is available to schedule appointments 
    export let occupied_time_slots = []; // appointments already scheduled
    export let available_durations = []; // the possible durations for an appointment
    export let day; // CalendarDay
    export let first_hour_in_view = "11:00"; // the timeline component will scroll to this hour, which is assumed to be in 24 hour format.
    export let onAppointmentSelected = () => {}; // callback function that will be called when an appointment is selected and the user clicks on the "schedule" button
    export let available_time_slot_label = "Available"; // the label that will be displayed on a available time slot
    export let unavailable_time_slot_label = "Unavailable"; // the label that will be displayed on a unavailable time slot
    
    // REFERENCES
    let day_hours = getDayHours(day); // a list of TimeSlots representing the hours of the day
    let timeline_component; //the main component is binded to this variable
    let appointment_start_time; // the start time of the appointment that is being scheduled
    let desired_appointment; // TimeSlot representing the desired appointment
    $: day_hours = filterAvailableHours(day_hours, available_time_slots, occupied_time_slots);

    onMount(() => {
        setTimelineStyles();
    });

    // Returns a set of the available hours in the day in string format so they can be quickly checked
    const filterAvailableHours = day_hours => {
        let new_day_hours = []

        day_hours.forEach(hour => {

            //Check if the hour is in the available time slots if not but it overlaps with one of the available time slots constraint it and
            // then check if the duration is greater than 1 minute
            let is_available = hour.IsAvailable;
            if (is_available) {
                for(let h = 0; h < available_time_slots.length; h++) {
                    let ts = available_time_slots[h];
                    if (hour.overlaps(ts)) {
                        hour.constraintInTs(ts);
                        break;
                    } else {
                        hour.IsAvailable = false;
                    }
                }
            }
            is_available = hour.IsAvailable && hour.getDurationinMinutes() > 1;

            // let is_not_occupied = !isHourOccupied(hour);
            let is_not_occupied = true;
            
            hour.IsAvailable = is_available && is_not_occupied;

            new_day_hours.push(hour);

        });

        return day_hours;
    }

    function setTimelineStyles() {
        const timeline_width = timeline_component.offsetWidth;

        timeline_component.style.setProperty('--libery-timeline-width', `${timeline_width}px`);
        timeline_component.style.setProperty('--libery-font-family-titles', `${font_titles}`);
        timeline_component.style.setProperty('--libery-font-family-paragraphs', `${font_body}`);
        timeline_component.style.setProperty('--libery-timeline-dark-color-1', `${dark_color_1}`);
        timeline_component.style.setProperty('--libery-timeline-dark-color-2', `${dark_color_2}`);
        timeline_component.style.setProperty('--libery-timeline-dark-color-3', `${dark_color_3}`);
        timeline_component.style.setProperty('--libery-timeline-light-color-1', `${light_color_1}`);
        timeline_component.style.setProperty('--libery-timeline-light-color-2', `${light_color_2}`);
        timeline_component.style.setProperty('--libery-timeline-primary-color', `${primary_color}`);
        timeline_component.style.setProperty('--libery-timeline-primary-color-light', `${primary_color_light}`);
        timeline_component.style.setProperty('--libery-timeline-secondary-color', `${secondary_color}`);
        timeline_component.style.setProperty('--libery-timeline-available-color', `${available_highlight_color}`);
        timeline_component.style.setProperty('--libery-timeline-unavailable-color', `${unavailable_highlight_color}`);
        timeline_component.style.setProperty('--libery-timeline-font-size-1', `${font_size_1}`);
        timeline_component.style.setProperty('--libery-timeline-font-size-2', `${font_size_2}`);
        timeline_component.style.setProperty('--libery-timeline-font-size-3', `${font_size_3}`);
        timeline_component.style.setProperty('--libery-timeline-font-size-h3', `${font_size_h3}`);
        timeline_component.style.setProperty('--libery-timeline-font-size-h2', `${font_size_h2}`);
        timeline_component.style.setProperty('--libery-timeline-font-size-h1', `${font_size_h1}`);
        timeline_component.style.setProperty('--boxes-shadow', `${boxShadow}`);




        const firstin_view_element = document.getElementById(`timeline-slot-${first_hour_in_view}`);
        firstin_view_element.scrollIntoView(
            {
                behavior: 'auto',
                block: 'start',
                inline: 'nearest'
            }
        );
    }

    // Returns true if the given hour is available to schedule an appointment
    const isHourAvailable = hours_ts => {
        return available_time_slots.some(ts => {
            let is_available = ts.inTimeframe(hours_ts);
            return is_available;
        });
    }

    // Returns true if the given hour is occupied by an appointment
    const isHourOccupied = hours_ts => {
        return occupied_time_slots.some(ts => {
            return ts.inTimeframe(hours_ts);
        });
    }

    // Returns true if the given hour is available to schedule an appointment
    const isHourSchudulable = hours_ts => {
        return isHourAvailable(hours_ts) && !isHourOccupied(hours_ts);
    }

    const defineNewAppointment = duration => {
        if (appointment_start_time === undefined) {
            newNotification("Debe seleccionar una hora de inicio para la cita");
            return;
        }

        const current_day = day.toDate();
        const [start_hour_str, start_minutes_str] = appointment_start_time.split(":");
        const start_time = new Date(current_day.getFullYear(), current_day.getMonth(), current_day.getDate(), start_hour_str, start_minutes_str);
        const end_time = new Date(start_time.getTime() + duration);

        const appointment = new TimeSlot(start_time, end_time);
        if (!isHourSchudulable(appointment)) {
            newNotification("La hora seleccionada no está disponible para agendar una cita");
            return;
        }

        // TODO: finish the scheduling process

        desired_appointment = appointment;
    }

</script>

<div id={component_id} class="timeline-component" bind:this={timeline_component}>
    <header id="day-banner">
        <h3>{day.toString()}</h3>
    </header>
    <div class="appointment-definer">
        <div id="appointment-start-wrapper">
            <span class="start-time-label definer-label">
                Inicio
            </span>
            <input type="time" id="appointment-start-time" bind:value={appointment_start_time}>
        </div>
        <div id="available-scheduling-durations-wrapper">
            <span class="available-durations-label definer-label">
                Duración
            </span>
            <div class="available-scheduling-durations">
                {#each available_durations as duration}
                    <button on:click={() => defineNewAppointment((duration.value * 1000) * 60)} class="available-duration">+{duration.value} mins</button>
                {/each}
            </div>
        </div>
        <div id="appointment-preview">
            <div class="flip-card">
                {#if desired_appointment !== undefined}
                    <span class="appointment-preview-content appointment-preview-front">
                        {day.toShortString()}  {desired_appointment.toString()}
                    </span>
                    <button on:click={() => onAppointmentSelected(desired_appointment)}  class="appointment-preview-content schedule-button appointment-preview-back">
                        Agendar
                    </button>
                {/if}
            </div>
        </div>
    </div>
    <ul id="day-hours">
        {#each day_hours as hour}
            <li id="timeline-slot-{hour.HourLocaleStart}" class="time-slot {hour.IsAvailable ? 'schedule-available-hour' : 'schedule-unavailable-hour'}">
                <span class="time-slot-dot"></span>
                <div class="starting-hour-wrapper">{hour.toString()}</div>
                <div class="time-slot-desciption">
                    {#if hour.IsAvailable}
                         {available_time_slot_label}
                    {:else}
                        {unavailable_time_slot_label}
                    {/if}
                </div>
            </li>
        {/each}
    </ul>
</div>

<style>
    .timeline-component {
        --libery-timeline-width: 100%;
        --libery-font-family-titles: 'Roboto', sans-serif;
        --libery-font-family-paragraphs: sans-serif;
        --libery-timeline-dark-color-1: #3c3c3c;
        --libery-timeline-dark-color-2: #7d7d7d;
        --libery-timeline-dark-color-3: #b7b7b7;
        --libery-timeline-light-color-1: white;
        --libery-timeline-light-color-2: #f2f2f2;
        --libery-timeline-primary-color: #e97f6a;
        --libery-timeline-primary-color-light: #da9f93;
        --libery-timeline-secondary-color: #54cae8;
        --libery-timeline-available-color: #6fca13;
        --libery-timeline-unavailable-color: #e02f2f;
        --libery-timeline-font-size-1: 8px;
        --libery-timeline-font-size-2: 14px;
        --libery-timeline-font-size-3: 18px;
        --libery-timeline-font-size-h3: 22px;
        --libery-timeline-font-size-h2: 28px;
        --libery-timeline-font-size-h1: 34px;
        --boxes-shadow: 4px 0 8px 2px rgba(0, 0, 0, 0.2);

        display: flex;
        flex-direction: column;
        justify-content: center;
        height: calc(var(--libery-timeline-width) * 1.2);
        /* border: 1px solid var(--libery-timeline-dark-color-2); */
        border-radius: 3px;
    }

    #day-banner {
        padding: min(15px, calc(var(--libery-timeline-width) * 0.02));
    }

    #day-banner h3 {
        font-family: var(--libery-font-family-titles);
        font-size: var(--libery-timeline-font-size-h3);
        color: var(--libery-timeline-dark-color-1);
        text-transform: capitalize;
    }

    
    /*=============================================
    =            Schedule Appointer            =
    =============================================*/
    
    /* .appointment-definer * {
        border: 1px solid yellowgreen;
    } */

   
    @-webkit-keyframes blink-continuous {
        0%,
        25%
        {
            opacity: 1;
        }
        30%{
            opacity: 0;
        }
        35%{
            opacity: 1;
        }
        40%{
            opacity: 0;
        }
        45%,
        100% {
            opacity: 1;
        }
    }

    @keyframes blink-continuous {
        0%,
        25%
        {
            opacity: 1;
        }
        30%{
            opacity: 0;
        }
        35%{
            opacity: 1;
        }
        40%{
            opacity: 0;
        }
        45%,
        100% {
            opacity: 1;
        }
    }


    .appointment-definer {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        gap: calc(var(--libery-timeline-width) * 0.02) calc(var(--libery-timeline-width) * 0.1);
        padding: max(15px, calc(var(--libery-timeline-width) * 0.02)) max(15px, calc(var(--libery-timeline-width) * 0.1));
        margin-bottom: calc(var(--libery-timeline-width) * 0.1);
    }

    .definer-label {
        font-family: var(--libery-font-family-paragraphs);
        font-size: var(--libery-timeline-font-size-2);
        font-weight: 600;
        color: var(--libery-timeline-dark-color-1);
    }

    #appointment-start-wrapper {
        grid-column: 1 / 2;
        display: flex;
        flex-direction: column;
        row-gap: calc(var(--libery-timeline-width) * 0.02);
    }

    #appointment-start-time {
        font-family: var(--libery-font-family-titles);
        font-size: var(--libery-timeline-font-size-2);
        color: var(--libery-timeline-dark-color-1);
        border: none;
        border-bottom: 2px solid var(--libery-timeline-dark-color-2);
        background-color: transparent;
        outline: none;
    }

    #available-scheduling-durations-wrapper {
        display: flex;
        flex-direction: column;
        grid-column: 2 / 4;
    }

    .available-scheduling-durations {
        display: flex;
        gap: calc(var(--libery-timeline-width) * 0.01);
    }

    .available-duration {
        font-family: var(--libery-font-family-titles);
        font-size: var(--libery-timeline-font-size-2);
        color: var(--libery-timeline-light-color-1);
        border: none;
        border-radius: 3px;
        background-color: var(--libery-timeline-primary-color);
        outline: none;
        cursor: pointer;
        padding: 5px 10px;
    }

    #appointment-preview {
        grid-column: 1 / 4;
        width: 100%;
        /* perspective: 1000px; */
    }

    .flip-card {
        position: relative;
        width: 100%;
        height: 100%;
        transition: all 0.8s ease-in-out;
        transform-style: preserve-3d;
    }

    .appointment-preview-content {
        box-sizing: border-box;
        position: absolute;
        width: 100%;
        height: max-content;
        background: var(--libery-timeline-available-color);
        color: var(--libery-timeline-light-color-1);
        text-align: center;
        font-size: var(--libery-timeline-font-size-2);
        border-radius: 3px;
        padding: calc(var(--libery-timeline-width) * 0.01) calc(var(--libery-timeline-width) * 0.02);
        border: none;
        backface-visibility: hidden;
    }

    #appointment-preview span {
        animation-duration: 4s;
        animation-name: blink-continuous;
        animation-timing-function: ease;
        animation-iteration-count: infinite;
        animation-delay: 2s;
        animation-fill-mode: both;
    }

    .appointment-preview-back {
        transform: rotateY(180deg);
    }

    @media only screen and (max-width: 786px) {
        #appointment-preview span {
            animation-name: none;
        }

    }
    
    @media(pointer: fine) {
        #appointment-preview:hover .flip-card {
            transform: rotateY(180deg);
        }


        
        #appointment-preview:hover span {
            animation-name: none;
        }
    }
    
    /*=====  End of Schedule Appointer  ======*/
    
    
    
    
    /*=============================================
    =            time slot            =
    =============================================*/
    /* DEBUG */
    /* #day-hours * {
        border: 1px solid blue;
    } */

    #day-hours {
        display: grid;
        grid-template-columns: 1fr;
        overflow-y: auto;
        list-style: none;
        height: calc(var(--libery-timeline-width) * 0.8);
        gap: calc(var(--libery-timeline-width) * 0.02);
    }

    .time-slot {
        position: relative;
        box-sizing: border-box;
        display: flex;
        width: 90%;
        height: calc(var(--libery-timeline-width) * .14);
        align-items: center;
        justify-content: space-around;
        border-radius: 2px;
        box-shadow: var(--boxes-shadow);
        border-left: 3px solid var(--libery-timeline-primary-color);
        z-index: 1;
    }

    .time-slot.schedule-available-hour {
        border-color: var(--libery-timeline-available-color);
    }

    .time-slot.schedule-unavailable-hour {
        border-color: var(--libery-timeline-unavailable-color);
    }

    .time-slot .time-slot-dot {
        position: absolute;
        top: 50%;
        left: -10px;
        width: 5px;
        height: 5px;
        border-radius: 2px;
        border: 3px solid var(--libery-timeline-light-color-1);
        transform: translateX(30%) translateY(-50%);
        border-radius: 50%;
        z-index: 2;
    }

    .time-slot.schedule-available-hour .time-slot-dot {
        background-color: var(--libery-timeline-available-color);
    }

    .time-slot.schedule-unavailable-hour .time-slot-dot {
        background-color: var(--libery-timeline-unavailable-color);
    }


    
    /*=====  End of time slot  ======*/
    
    
</style>