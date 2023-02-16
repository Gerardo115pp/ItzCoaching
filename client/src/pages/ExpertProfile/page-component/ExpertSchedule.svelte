<script>
    import { GetExpertSchedule } from "../../../libs/HttpRequests";
    import LiberyTimeline from "../../../components/DateComponents/DayTimeline.svelte";
    import LiberyCalendar  from "../../../components/DateComponents/Calendar.svelte";
    import { onMount, onDestroy } from "svelte";
    import { newNotification } from "../../../components/notifications/events";
    import { CalendarMonth } from "../../../components/DateComponents/date_utils"
    import { TimeSlot, convertUTCtoLocalTime } from "../../../components/DateComponents/time_utils";

    export let expert_id;

    let expert_schedule = {};
    let available_schedule_timeslots = [];

    // DELETE THIS
    let now = new Date();
    let current_month_days = new CalendarMonth(now.getFullYear(), now.getMonth());


    const schdule_modes = {
        CALENDAR_VIEW: "calendar_view",
        TIMELINE_VIEW: "timeline_view"
    }
    let current_schedule_mode = schdule_modes.TIMELINE_VIEW;
    const avaiable_durations = [
        {
            human_readable: "30 mins",
            value: 30
        },
        {
            human_readable: "45 mins",
            value: 45
        },
        {
            human_readable: "1 hora",
            value: 60
        }
    ];


    onMount(() => {
        getExpertSchedule();
    });

    const getExpertSchedule = () => {
        const schedule_request = new GetExpertSchedule(expert_id);

        const on_success = schedule => {
            expert_schedule = schedule;
            console.log(`time_availability: ${expert_schedule.time_availability.length}`)
            let new_available_timeslots = expert_schedule.time_availability.map(ts => {
                return new TimeSlot(convertUTCtoLocalTime(ts.start_time), convertUTCtoLocalTime(ts.end_time));
            });
            available_schedule_timeslots = new_available_timeslots;

            console.log(available_schedule_timeslots);
        };

        const on_error = error => {
            console.log(error);
        };

        schedule_request.do(on_success, on_error);
    };

    const handleAvailableDayClick = date => {
        newNotification(`Clicked on ${date.toLocaleDateString()}`)
    }

    const handleMarkedDayClick = date => {
        newNotification(`Show day schedule for ${date.toLocaleDateString()}`)
    }
</script>

<section id="expert-schdule-appointer">
    <div id="esa-available-durations">
        <h3 class="esa-section-label">Duración</h3>
        <div id="esa-ad-durations">
            {#each avaiable_durations as appointment_duration}
                <button class="clear-two-btn-thin" on:click={() => {}}>
                    {appointment_duration.human_readable}
                </button>
            {/each}
        </div>
    </div>
    <div id="esa-scheduling-component">
        {#if current_schedule_mode === schdule_modes.CALENDAR_VIEW && expert_schedule.week_availability !== undefined}
            <div class="scheduling-component" id="esa-scheduling-calendar">
                <LiberyCalendar 
                    active_day_color="var(--theme-red)"
                    inactive_day_color="var(--theme-light-red)"
                    today_color="var(--theme-purple)"
                    clear_color_1="var(--clear-color)"
                    marked_color="var(--ready)"
                    marked_weekdays={new Set(Object.keys(expert_schedule.week_availability).filter(day => expert_schedule.week_availability[day]))}
                    active_day_callback={handleAvailableDayClick}
                    marked_day_callback={handleMarkedDayClick}
                />
            </div>
        {:else}
            <div class="scheduling-component" id="esa-scheduling-timeline">
                <LiberyTimeline day={current_month_days.Days[0]} available_time_slots={available_schedule_timeslots}/>
            </div>
        {/if}
    </div>
</section>

<style>
    .esa-section-label {
        font-family: var(--font-text);
        font-size: var(--font-size-2);
        text-transform: none;
    }

    #esa-ad-durations {
        display: flex;
        flex-wrap: wrap;
        gap: var(--spacing-1);
        padding: var(--spacing-3);
    }

    .scheduling-component {
        width: 90%;
    }
</style>