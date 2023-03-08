<script>
    import { TimeSlot, convertUTCtoLocalTime } from "../../../components/DateComponents/time_utils";
    import LiberyTimeline from "../../../components/DateComponents/DayTimeline.svelte";
    import LiberyCalendar  from "../../../components/DateComponents/Calendar.svelte";
    import { CalendarMonth, formatUTCDate } from "../../../components/DateComponents/date_utils"
    import { newNotification } from "../../../components/notifications/events";
    import { PostAppointmentRequest } from "../../../libs/HttpRequests";
    import LiberyInput from "../../../components/Input/Input.svelte";
    import { GetExpertSchedule } from "../../../libs/HttpRequests";
    import Input from "../../../components/Input/Input.svelte";
    import FieldData from "../../../libs/FieldData"
    import { onMount, onDestroy } from "svelte";

    export let expert_id;
    export let expert_appointments = [];
    $: window.expert_appointments = expert_appointments;

    let expert_schedule = {};
    let available_schedule_timeslots = [];


    const schdule_modes = {
        CALENDAR_VIEW: "calendar_view",
        TIMELINE_VIEW: "timeline_view"
    }
    
    let current_schedule_mode = schdule_modes.CALENDAR_VIEW;
    let selected_day; // type: CalendarDay
    let selected_time_slot; // type: TimeSlot
    let email_field = new FieldData("customer-email", /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/, "customer-email", 'email');

    let selected_day_appointments = []; // type: []TimeSlot

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

    const appointmentTimeSelected = appointment_time_slot => {
        selected_time_slot = appointment_time_slot;
        if (!isFormReady()) {
            if (selected_time_slot === undefined) {
                newNotification("Por favor selecciona un horario"); // currently this is impossible to reach, or atleast should be
            } else if (!email_field.isReady()) {
                newNotification("Por favor ingresa un correo electrónico");
            }
            return;
        }

        const start_time = selected_day.toDate();
        const end_time = selected_day.toDate();
        start_time.setHours(selected_time_slot.Start.getHours());
        start_time.setMinutes(selected_time_slot.Start.getMinutes());
        start_time.setSeconds(0);

        end_time.setHours(selected_time_slot.End.getHours());
        end_time.setMinutes(selected_time_slot.End.getMinutes());
        end_time.setSeconds(0);


        const appointment_request = new PostAppointmentRequest(expert_id, formatUTCDate(start_time), formatUTCDate(end_time), email_field.getFieldValue());

        const on_success = appointment => {
            let url = appointment.session_url
            window.location.href = url;
        };

        const on_error = error => {
            newNotification("Hubo un error al crear la cita");
        };

        appointment_request.do(on_success, on_error);
    }

    const getExpertSchedule = () => {
        const schedule_request = new GetExpertSchedule(expert_id);

        const on_success = schedule => {
            expert_schedule = schedule;
            let new_available_timeslots = expert_schedule.time_availability.map(ts => {
                return new TimeSlot(convertUTCtoLocalTime(ts.start_time), convertUTCtoLocalTime(ts.end_time));
            });
            available_schedule_timeslots = new_available_timeslots;

        };

        const on_error = error => {
            console.log(error);
        };

        schedule_request.do(on_success, on_error);
    };

    const getSelectedDayAppointments = new_selected_day => {
        let selected_day_appointments = [];

        if (new_selected_day !== undefined) {

            let new_selected_day_appointments = [];
            for(let a of expert_appointments) {
                if (new_selected_day.isSameDay(a.Start)) {
                    new_selected_day_appointments.push(a.TimeSlot);
                }
            }

            selected_day_appointments = new_selected_day_appointments;
        }
        return selected_day_appointments;
    }

    const handleAvailableDayClick = date => {
        // IGNORE THIS i dont remember whats the difference between a marked day and an available day
        // okey i remember, a marked day is a day that the expert is available, but an available day is just a day that is not in the past. what a mess
        newNotification(`Clicked on ${date.toLocaleDateString()}`)
    }

    const handleMarkedDayClick = date => {
        selected_day_appointments = getSelectedDayAppointments(date);
        selected_day = date;
        current_schedule_mode = schdule_modes.TIMELINE_VIEW;
    }

    const isFormReady = () => {
        return selected_day !== undefined && selected_time_slot !== undefined && email_field.isReady();
    }

    const validateEmail = () => {
        if (!email_field.verify()) {
            newNotification("El correo electrónico no es válido");
        }
    }
</script>

<section id="expert-schdule-appointer">
    <h3 id="esa-header">
        Agenda tu cita
    </h3>
    <div id="esa-customer-email-wrapper">
        <div class="esa-customer-email-collector">
            <Input 
                field_data={email_field}
                title_font="var(--font-titles)"
                font_size="var(--font-size-1)"
                input_padding="var(--spacing-1)"
                input_dark_color="var(--theme-red)"
                input_label="Tu correo electrónico"
                isClear={true}
                onBlur={validateEmail}
            />
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
        {:else if current_schedule_mode === schdule_modes.TIMELINE_VIEW && selected_day !== undefined}
            <div class="scheduling-component" id="esa-scheduling-timeline">
                <LiberyTimeline 
                    {available_durations}  
                    day={selected_day} 
                    available_time_slots={available_schedule_timeslots} 
                    occupied_time_slots={selected_day_appointments}
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


    #esa-header {
        font-family: var(--font-titles);
        font-size: var(--font-size-h3);
        text-transform: none;
        margin-bottom: var(--spacing-2);
    }

    #esa-customer-email-wrapper {
        box-sizing: border-box;
        display: flex;
        /* justify-content: center; */
        padding: 0 var(--spacing-3);
        margin-bottom: var(--spacing-2);
    }

    .esa-customer-email-collector {
        width: 80%;
    }

    .scheduling-component {
        width: 90%;
    }
</style>