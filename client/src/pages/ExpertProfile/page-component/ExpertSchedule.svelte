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
    
    let current_schedule_mode = schdule_modes.CALENDAR_VIEW;
    let selected_day;
    const available_durations = [
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
        selected_day = date;
        console.log(selected_day);
        current_schedule_mode = schdule_modes.TIMELINE_VIEW;
    }

    const appointmentTimeSelected = appointment_time_slot => {
        newNotification(`Appointment time selected: ${appointment_time_slot.toString()}`);
    }
</script>

<section id="expert-schdule-appointer">
    <h3 id="esa-header">
        Agenda tu cita
    </h3>
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
        {:else if current_schedule_mode === schdule_modes.TIMELINE_VIEW && selected_day !== undefined}
            <div class="scheduling-component" id="esa-scheduling-timeline">
                <LiberyTimeline 
                    {available_durations}  
                    day={selected_day} 
                    available_time_slots={available_schedule_timeslots} 
                    onAppointmentSelected={appointmentTimeSelected}
                    primary_color="var(--theme-red)"
                    font_titles="var(--font-titles)"
                    font_body="var(--font-text)"
                    font_size_1="var(--font-size-2)"
                    font_size_2="var(--font-size-1)"
                    font_size_3="var(--font-size-2)"
                    available_highlight_color="var(--ready)"
                    unavailable_highlight_color="var(--danger)"
                    boxShadow="var(--boxes-shadow-2)"
                    available_time_slot_label="Disponible"
                    unavailable_time_slot_label="No disponible"
                />
            </div>
        {/if}
    </div>
</section>

<style>
    /* .esa-section-label {
        font-family: var(--font-text);
        font-size: var(--font-size-2);
        text-transform: none;
    }

    #esa-ad-durations {
        display: flex;
        flex-wrap: wrap;
        gap: var(--spacing-1);
        padding: var(--spacing-3);
    } */

    #esa-header {
        font-family: var(--font-titles);
        font-size: var(--font-size-h3);
        text-transform: none;
        margin-bottom: var(--spacing-2);
    }

    .scheduling-component {
        width: 90%;
    }
</style>