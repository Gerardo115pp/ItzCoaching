<script>
    import { onMount, onDestroy } from "svelte";
    import { getDayHours } from './time_utils';

    // STYLE CUSTOMIZATION
    export let component_id = "libery-timeline";

    // CUSTOMIZATION
    export let available_time_slots = []; 
    export let occupied_time_slots = [];
    export let day; // MonthDay
    export let first_hour_in_view = "11:00";
    const day_hours = getDayHours(day);
    const available_time_slot_label = "Available";
    const unavailable_time_slot_label = "Unavailable";

    // REFERENCES
    let timeline_component;
    let available_hours = new Set();
    $: available_hours = filterAvailableHours(day_hours, available_time_slots, occupied_time_slots);
    $: window.available_hours = available_hours;
    
    onMount(() => {
        setTimelineStyles();
    });

    const filterAvailableHours = day_hours => {
        let new_available_hours = new Set();
        day_hours.forEach(hour => {
            if (isHourSchedulable(hour)) {
                new_available_hours.add(hour.toString());
            }
        });

        return new_available_hours;
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

    const isHourSchedulable = hours_ts => {
        return isHourAvailable(hours_ts) && !isHourOccupied(hours_ts);
    }

    const isHourAvailable = hours_ts => {
        return available_time_slots.some(ts => {
            let is_available = ts.inTimeframe(hours_ts);
            return ts.inTimeframe(hours_ts);
        });
    }

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
    <ul id="day-hours">
        {#each day_hours as hour}
            <li id="timeline-slot-{hour.HourLocaleStart}" class="time-slot {available_hours.has(hour.toString()) ? 'schedule-available-hour' : 'schedule-unavailable-hour'}">
                <span class="time-slot-dot"></span>
                <div class="starting-hour-wrapper">{hour.HourLocaleStart}</div>
                <div class="time-slot-desciption">
                    {#if available_hours.has(hour.toString())}
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
        height: var(--libery-timeline-width);
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