<script>
    import { CalendarMonth, getWeekDay, weekDays, convertWeekDayToNumber } from "./date_utils";
    import { onMount } from "svelte";


    // STYLE CUSTOMIZATION
    export let component_id = "calendar-component";
    export let active_day_color = "#e97f6a";
    export let inactive_day_color = "#da9f93";
    export let today_color = "#54cae8";
    export let marked_color = "#6fca13"
    export let clear_color_1 = "white";
    export let clear_color_2 = "#f2f2f2";
    export let dark_color_1 = "#3c3c3c";
    export let dark_color_2 = "#7d7d7d";
    export let dark_color_3 = "#b7b7b7";

    // COSTUMIZATION
    export let marked_weekdays = new Set(); // Set of week days to mark
    $: marked_weekdays = convertWeekDaysSet(marked_weekdays); // Convert the weekdays to numbers

    
    // BEHAVIOR
    export let is_current_date_active = false; // If the current date is active or not
    
    // INTERNAL STYLE VARIABLES
    let small_mode_on = false;


    // REFERENCES
    let calendar_component;

    // CALLBACKS
    export let active_day_callback = date => {};
    export let marked_day_callback = date => {};

    export let current_displayed_date = new Date(); // Current date
    let calendar_month_data = new CalendarMonth(current_displayed_date.getFullYear(), current_displayed_date.getMonth());

    onMount(() => {
        setCalendarStyles();
    });



    function changeMonth(increment) {
        let new_month = calendar_month_data.Month;
        let new_year = calendar_month_data.Year;
        
        new_month += increment;

        if (new_month < 0 || new_month > 11) {
            new_year += increment;
        }

        new_month = new_month < 0 ? 11 : new_month%12;
        current_displayed_date = new Date(new_year, new_month, 1);
        calendar_month_data = new CalendarMonth(new_year, new_month);
    }

    const convertWeekDaysSet = weekdays => {
        let converted_weekdays = new Set();

        weekdays.forEach(weekday => {
            converted_weekdays.add(convertWeekDayToNumber(weekday));
        });

        return converted_weekdays;
    }

    const isDayMarked = month_day => {
        let is_marked = false;

        is_marked = marked_weekdays.has(month_day.week_day);

        return is_marked;
    }

    const monthDayClicked = month_day => {
        //IDEA: Just add the avaiable week days logic in the ExpertSchedule.svelte not here
        //TODO: Remove this, because the idea is related to itz coaching project and this should be a generic component

        if (isDayMarked(month_day)  && !month_day.isPastDate()) {
            if (!month_day.isCurrentDate() || is_current_date_active) {
                marked_day_callback(month_day.toDate());
            }
        } else if (month_day.inCurrentMonth() && !month_day.isPastDate()) {
            if (!month_day.isCurrentDate() || is_current_date_active) {
                active_day_callback(month_day.toDate());
            }
        } 
    }

    function setCalendarStyles() {
        const component_width = calendar_component.offsetWidth;
        small_mode_on = component_width <= 500;

        calendar_component.style.setProperty("--libery-calendar-width", `${component_width}px`);
        calendar_component.style.setProperty("--libery-calendar-primary-color", `${active_day_color}`);
        calendar_component.style.setProperty("--libery-calendar-primary-color-light", `${inactive_day_color}`);
        calendar_component.style.setProperty("--libery-calendar-marked-color", `${marked_color}`);
        calendar_component.style.setProperty("--libery-calendar-secondary-color", `${today_color}`);
        calendar_component.style.setProperty("--libery-calendar-light-color-1", `${clear_color_1}`);
        calendar_component.style.setProperty("--libery-calendar-light-color-2", `${clear_color_2}`);
        calendar_component.style.setProperty("--libery-calendar-dark-color-1", `${dark_color_1}`);
        calendar_component.style.setProperty("--libery-calendar-dark-color-2", `${dark_color_2}`);
        calendar_component.style.setProperty("--libery-calendar-dark-color-3", `${dark_color_3}`);

    }

</script>

<div bind:this={calendar_component} id={component_id} class="calendar-component">
    <header class="calendar-month-descriptor">
        <button on:click={() => changeMonth(-1)} class="previous-month material-symbols-outlined">
            arrow_back_ios
        </button>
        <div class="month-descriptor-wrapper">
            <h3 class="week-day">{getWeekDay(current_displayed_date)}</h3>
            <h3 class="month-data">{calendar_month_data.Name} {current_displayed_date.getDate()}, {calendar_month_data.Year}</h3>
        </div>
        <button on:click={() => changeMonth(1)} class="next-month material-symbols-outlined">
            arrow_forward_ios
        </button>
    </header>
    <ul class="calendar-week-days">
        {#each Object.values(weekDays) as week_day}
            <li class="calendar-week-day">
                <span class="week-day-wrapper">{!small_mode_on ? week_day : week_day[0].toUpperCase()}</span>
            </li>
        {/each}
    </ul>
    <ul class="calendar-days">
        {#each calendar_month_data.Days as month_day, h}
             <li
                on:click={() => monthDayClicked(month_day)}
                class="calendar-day {(h+1)%7 === 0 ? 'no-border-right' : ''} {h >= (calendar_month_data.Days.length - 7) ? 'no-border-bottom' : ''} {month_day.month === calendar_month_data.Month ? 'current-month-day' : 'other-month-day'} {month_day.isPastDate() ? 'day-in-past' : 'day-in-future'} {month_day.isCurrentDate() ? 'today' : ''} {isDayMarked(month_day) ? 'marked-day' : ''}">
                <span class="month-day-wrapper {month_day}">{month_day.day_num}</span>
            </li>
        {/each}
    </ul>
</div>

<style>
    /* DEBUG: BOXES */
    /* #calendar-component * {
        border: 1px solid blue;
    } */

        .calendar-component {
            --libery-calendar-width: 100%;
            --libery-font-family-titles: 'Roboto', sans-serif;
            --libery-font-family-paragraphs: sans-serif;
            --libery-calendar-dark-color-1: #3c3c3c;
            --libery-calendar-dark-color-2: #7d7d7d;
            --libery-calendar-dark-color-3: #b7b7b7;
            --libery-calendar-light-color-1: white;
            --libery-calendar-light-color-2: #f2f2f2;
            --libery-calendar-primary-color: #e97f6a;
            --libery-calendar-primary-color-light: #da9f93;
            --libery-calendar-secondary-color: #54cae8;
            --libery-calendar-marked-color: #6fca13;
            --libery-calendar-font-size-1: 8px;
            --libery-calendar-font-size-2: 14px;
            --libery-calendar-font-size-3: 18px;
            --libery-calendar-font-size-h3: 22px;
            --libery-calendar-font-size-h2: 28px;
            --libery-calendar-font-size-h1: 34px;

            --libery-calendar-padding-2: 2vh;
            --libery-calendar-padding-1: 1vh;

            width: 100%;
            border: 1px solid var(--libery-calendar-dark-color-3);
            border-radius: 5px;
        }
        
        .calendar-component ul {
            list-style: none;
            padding: 0;
            margin: 0;
        }

    
    /*=============================================
    =            Header            =
    =============================================*/
        

        header.calendar-month-descriptor {
            background: var(--libery-calendar-light-color-1);
            font-family: var(--libery-font-family-titles);
            display: flex;
            justify-content: center;
            align-items: center;
            border-bottom: 1px solid var(--libery-calendar-dark-color-3);
        }

        header.calendar-month-descriptor button {
            background: none;
            border: none;
            outline: none;
            height: calc(var(--libery-calendar-font-size-h2) * 2.8);
            width: calc(var(--libery-calendar-font-size-h2) * 2.8);
            font-size: var(--libery-calendar-font-size-2);
            color: var(--libery-calendar-dark-color-2);
            transition: all 0.2s ease-in-out;
        }

        @media(pointer: fine) {
            header.calendar-month-descriptor button:hover {
                color: var(--libery-calendar-light-color-1);
                background: var(--libery-calendar-primary-color);
            }
        }

        .previous-month {
            margin-right: auto;
        }

        .week-day {
            font-family: var(--libery-font-family-paragraphs);
            font-size: var(--libery-calendar-font-size-3);
            font-weight: bold;
            color: var(--libery-calendar-dark-color-2);
        }

        .month-data {
            font-family: var(--libery-font-family-titles);
            font-size: var(--libery-calendar-font-size-h2);
            color: var(--libery-calendar-dark-color-2);

        }
        
        .next-month {
            margin-left: auto;
        }
    
    /*=====  End of Header  ======*/
    
    
    /*=============================================
    =            Week days list            =
    =============================================*/
    
        .calendar-week-days {
            display: flex;
            border-bottom: 1px solid var(--libery-calendar-dark-color-3);
        }
        
        .calendar-week-day {
            text-align: center;
            background: var(--libery-calendar-light-color-1);
            color: var(--libery-calendar-dark-color-2);
            width: calc(var(--libery-calendar-width) / 7);
            padding: 1em 0;
            border-right: 1px solid var(--libery-calendar-dark-color-3);
        }

        .calendar-week-day:last-child {
            border-right: none;
        }
    
    /*=====  End of Week days list  ======*/
        
        .calendar-days {
            display: grid;
            grid-template-columns: repeat(7, 1fr);
        }
        
        .calendar-day {
            box-sizing: border-box;
            font-family: var(--libery-font-family-paragraphs);
            height: calc(var(--libery-calendar-width) / 7);
            border-right: .5px solid var(--libery-calendar-dark-color-3);
            border-bottom: .5px solid var(--libery-calendar-dark-color-3);
            padding: var(--libery-calendar-padding-1);
            font-size: var(--libery-calendar-font-size-2);
        }
        
        .calendar-day.current-month-day {
            background: var(--libery-calendar-primary-color);
            color: var(--libery-calendar-light-color-1);
            transition: all 0.2s ease-in-out;
        }

        @media(pointer: fine) {
            .calendar-day.current-month-day:hover {
                transform: scale(1.1);
                border: none;
                box-shadow: 0 0 3px 0 var(--libery-calendar-dark-color-1);
            }
        }

        .calendar-day.other-month-day {
            color: var(--libery-calendar-light-color-2);
            background: var(--libery-calendar-dark-color-3);
        }

        .calendar-day.current-month-day.day-in-past {
            background: var(--libery-calendar-primary-color-light);
            color: var(--libery-calendar-light-color-2);
        }

        .calendar-day.calendar-day.current-month-day.today {
            background: var(--libery-calendar-secondary-color);
            color: var(--libery-calendar-light-color-1);
        }

        .calendar-day.calendar-day.current-month-day.marked-day.day-in-future:not(.today) {
            background: var(--libery-calendar-marked-color);
            color: var(--libery-calendar-light-color-1);
        }

        .calendar-day.no-border-right {
            border-right: none;
        }

        .calendar-day.no-border-bottom {
            border-bottom: none;
        }
</style>