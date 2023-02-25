<script>
    import { onMount, onDestroy } from "svelte";
    import { getDayHours } from './time_utils';

    // STYLE CUSTOMIZATION
    export let component_id = "libery-timeline";

    // CUSTOMIZATION
    export let available_time_slots = []; // the hours in which the schedule owner is available to schedule appointments 
    export let occupied_time_slots = []; // appointments already scheduled
    export let available_durations = []; // the possible durations for an appointment
    export let day; // MonthDay
    export let first_hour_in_view = "11:00"; // the timeline component will scroll to this hour, which is assumed to be in 24 hour format.
    const available_time_slot_label = "Available"; // the label that will be displayed on a available time slot
    const unavailable_time_slot_label = "Unavailable"; // the label that will be displayed on a unavailable time slot
    
    // REFERENCES
    let day_hours = getDayHours(day); // a list of TimeSlots representing the hours of the day
    let timeline_component; //the main component is binded to this variable
    let available_hours = new Set(); // the hours than can be scheduled are stored in string format here, e.g new Set(["11:12 - 12:30",...]). this is done to improve performance when mounting the dom elements and checking for correct css classes.
    let appointment_start_time; // the start time of the appointment that is being scheduled
    $: day_hours = filterAvailableHours(day_hours, available_time_slots, occupied_time_slots);
    $: window.available_hours = available_hours; // debug,DELETE THIS
    $: window.day_hours = day_hours; // debug,DELETE THIS

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


            console.log(`hour ${hour.toString()} is available: ${is_available}`)

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

</script>

<div id={component_id} class="timeline-component" bind:this={timeline_component}>
    <header id="day-banner">
        <h3>{day.toString()}</h3>
    </header>
    <div class="appointment-definer">
        <input type="time" id="appointment-start-time" bind:value={appointment_start_time}>
        <div class="available-scheduling-durations">
            {#each available_durations as duration}
                <button class="available-duration">+{duration.value} mins</button>
            {/each}
        </div>
        <div id="appointment-preview">
            
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
                <div class="time-slot-controls">
                    {#if available_hours.has(hour.toString())}
                         <button class="schedule-btn material-symbols-outlined">arrow_forward_ios</button>
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
        box-shadow: 0 0 10px 2px rgba(255, 255, 255, 0.8), 0 0 48px 28px rgba(0, 0, 0, 0.09) ;
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
    
    .appointment-definer * {
        border: 1px solid yellowgreen;
    }

    .appointment-definer {
        display: grid;
        align-items: center;
        justify-items: center;
        grid-template-columns: repeat(3, 1fr);
        gap: calc(var(--libery-timeline-width) * 0.02);
        padding: max(15px, calc(var(--libery-timeline-width) * 0.02)) max(15px, calc(var(--libery-timeline-width) * 0.1));
    }

    #appointment-start-time {
        grid-column: 1 / 2;
        font-family: var(--libery-font-family-titles);
        font-size: var(--libery-timeline-font-size-2);
        color: var(--libery-timeline-dark-color-1);
        border: none;
        border-bottom: 2px solid var(--libery-timeline-dark-color-2);
        background-color: transparent;
        outline: none;
    }

    .available-scheduling-durations {
        grid-column: 2 / 4;
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
        /* height: calc(var(--libery-timeline-width) * 0.2); */
    }
    
    /*=====  End of Schedule Appointer  ======*/
    
    
    
    
    /*=============================================
    =            time slot            =
    =============================================*/
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

    .time-slot-controls .schedule-btn {
        border: none;
        background-color: transparent;
        color: var(--libery-timeline-dark-color-2);
    }
    
    /*=====  End of time slot  ======*/
    
    
</style>