import { TimeSlot } from "./time_utils"

export const weekDays = {
    0: 'Sunday',
    1: 'Monday',
    2: 'Tuesday',
    3: 'Wednesday',
    4: 'Thursday',
    5: 'Friday',
    6: 'Saturday'
}

export const weekDaysShort = {
    0: 'Sun',
    1: 'Mon',
    2: 'Tue',
    3: 'Wed',
    4: 'Thu',
    5: 'Fri',
    6: 'Sat'
}

export const spanishWeekDays = {
    0: 'Domingo',
    1: 'Lunes',
    2: 'Martes',
    3: 'Miércoles',
    4: 'Jueves',
    5: 'Viernes',
    6: 'Sábado'
}

export const spanishWeekDaysShort = {
    0: 'Dom',
    1: 'Lun',
    2: 'Mar',
    3: 'Mié',
    4: 'Jue',
    5: 'Vie',
    6: 'Sáb'
}

const weekDayToNumber = {
    'sunday': 0,
    'monday': 1,
    'tuesday': 2,
    'wednesday': 3,
    'thursday': 4,
    'friday': 5,
    'saturday': 6
}

export class Appointment {
    #utc_start;
    #duration;
    #ts
    constructor(data, lang="en") {
        this.id = data.id;
        this.expert_id = data.expert;
        this.customer_email = data.customer_email;
        this.customer_name = data.customer_name;
        this.#utc_start = parseUTCDate(data.utc_start);
        this.#duration = data.duration; // milliseconds
        this.status = data.status;
        this.created_at = parseUTCDate(data.created_at);
        this.lang = lang;
        this.#ts = this.#toTimeSlot();
    }

    #toTimeSlot = () => {
        const end = new Date(this.#utc_start.getTime() + this.#duration);
        return new TimeSlot(this.#utc_start, end);
    }

    get Start() {
        return this.#utc_start;
    }

    get End() {
        return this.#ts.End;
    }

    get Duration() {
        return this.#duration;
    }

    get TimeSlot() {
        return this.#ts;
    }

    get MonthName() {
        const options = { month: 'long' };
        const month_name = this.#utc_start.toLocaleTimeString(this.lang, options);
        return month_name;
    }


    toString = () => {
        let date_string = "";
        let week_day_start = weekDaysShort[this.#utc_start.getDay()];

        if (this.lang === "es") {
            week_day_start = spanishWeekDaysShort[this.#utc_start.getDay()];
        }

        return `${week_day_start}. ${this.#utc_start.getDate()} ${this.lang === 'en' ? 'of' : 'de'} ${this.MonthName} at ${this.#utc_start.toLocaleTimeString('en-US', { hour: 'numeric', minute: 'numeric', hour12: true })}`
    }

    dateString = () => {
        return `${this.#utc_start.getDate()}-${this.#utc_start.getMonth()+1}-${this.#utc_start.getFullYear()}`
    }

    timeString = () => {    
        return `${this.#utc_start.toLocaleTimeString(this.lang, { hour: 'numeric', minute: 'numeric', hour12: true })} - ${this.Duration / 1000 / 60} mins`
    }
}

export const getDaysInMonth = date => {
    return new Date(date.getFullYear(), date.getMonth() + 1, 0).getDate();
}

export const getWeekDay = (date) => {
    return weekDays[date.getDay()];
}

export const getMonthName = (date) => {
    return date.toLocaleString('default', { month: 'long' });
}

export const convertWeekDayToNumber = (week_day) => {
    week_day = week_day.toLowerCase();
    return weekDayToNumber[week_day];
}

export const monthDayFromDate = (date) => {
    return new MonthDay(date.getFullYear(), date.getMonth(), date.getDate());
}

export const formatUTCDate = (date) => {
    const year = date.getUTCFullYear();
    const month = String(date.getUTCMonth() + 1).padStart(2, '0');
    const day = String(date.getUTCDate()).padStart(2, '0');
    const hours = String(date.getUTCHours()).padStart(2, '0');
    const minutes = String(date.getUTCMinutes()).padStart(2, '0');
    const seconds = String(date.getUTCSeconds()).padStart(2, '0');
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}

export const parseUTCDate = (date_string) => {
    const date_parts = date_string.split(/[^\d]/);
    const year = parseInt(date_parts[0]);
    const month = parseInt(date_parts[1]) - 1;
    const day = parseInt(date_parts[2]);
    const hours = parseInt(date_parts[3]);
    const minutes = parseInt(date_parts[4]);
    const seconds = parseInt(date_parts[5]);

    return new Date(Date.UTC(year, month, day, hours, minutes, seconds));
}


class CalendarDay {
    #calendar_order;
    #calendar;
    constructor(day_num, week_day, calendar_order, month, calendar) {
        this.#calendar_order = calendar_order;
        this.day_num = day_num;
        this.week_day = week_day;
        this.#calendar = calendar;
        this.week_day_name = weekDays[week_day];
        this.month = month;
        this.is_past_date = this.isPastDate();
    }

    get CalendarOrder() {
        return this.#calendar_order;
    }

    toString = () => {
        const current_date = this.toDate();
        return current_date.toLocaleDateString('default', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' });
    }

    toShortString = () => {
        const current_date = this.toDate();
        const day = current_date.getDate().toString().padStart(2, '0'); 
        const month = (current_date.getMonth() + 1).toString().padStart(2, '0');
        const year = current_date.getFullYear().toString().padStart(4, '0');

        return `${day}/${month}/${year}`;
    }

    isPastDate = () => {
        const input_date = new Date(this.#calendar.Year, this.month, this.day_num);
        
        input_date.setHours(0, 0, 0, 0);
        
        return input_date < this.#calendar.CurrentDate;
    }

    isCurrentDate = () => {
        const input_date = new Date(this.#calendar.Year, this.month, this.day_num);
        
        input_date.setHours(0, 0, 0, 0);
        
        return input_date.getTime() === this.#calendar.CurrentDate.getTime();
    }

    inCurrentMonth = () => {
        return this.month === this.#calendar.Month;
    }

    toDate = () => {
        return new Date(this.#calendar.Year, this.month, this.day_num);
    }
    

}

export class CalendarMonth {
    #name;
    #month;
    #year;
    #day_count;
    #days;
    #internal_date;
    #current_date;
   
    constructor(year, month) {
        this.#internal_date = new Date(year, month, 1);
        this.#current_date = new Date();
        this.#current_date.setHours(0, 0, 0, 0); 
        this.#month = month;
        this.#year = year;
        this.#days = [];
        this.#day_count = getDaysInMonth(this.#internal_date);
        this.#name = getMonthName(this.#internal_date);
        this.#fillDaysInfo(); // Fill the days array with the month days and the previous month days that belong to this month first and last weeks
    }

    get Month() {
        return this.#month;
    }

    set Month(month) {
        this.#month = month;
        this.#internal_date.setMonth(month);
        this.#name = getMonthName(this.#internal_date);
        this.#day_count = getDaysInMonth(this.#internal_date);
        this.#fillDaysInfo();
    }

    get Year() {
        return this.#year;
    }

    set Year(year) {
        this.#year = year;
        this.#internal_date.setFullYear(year);
        this.#day_count = getDaysInMonth(this.#internal_date);
        this.#fillDaysInfo();
    }

    get CurrentDate() {
        return this.#current_date;
    }

    get Name() {
        return this.#name;
    }

    get DayCount() {
        return this.#day_count;
    }

    get Days() {
        return this.#days;
    }

    toDate() {
        return this.#internal_date;
    }

    #getFirstWeekMissingDays() {
        const first_week_day = this.#internal_date.getDay();
        const missing_days = [];
        const previous_month_day_count = getDaysInMonth(new Date(this.#year, this.#month - 1, 1));
        let current_day;

        for (let h = first_week_day - 1; h >= 0; h--) {
            current_day = new CalendarDay(previous_month_day_count - ((first_week_day-1) - h), h, first_week_day - (missing_days.length+1), this.#month - 1, this);
            missing_days.push(current_day);
        }

        return missing_days;
    }

    #getLastWeekMissingDays(last_day, calendar_order) {
        const last_week_day = last_day.week_day;
        const missing_days = [];
        let current_day;

        for (let h = last_week_day + 1; h <= 6; h++) {
            current_day = new CalendarDay((h - last_day.week_day), h, calendar_order, this.#month+1, this);
            missing_days.push(current_day);
            calendar_order++;
        }

        return missing_days;
    }

    #fillDaysInfo() {

        let new_month_days = this.#getFirstWeekMissingDays();
        let current_day;
        
        for (let h = 0; h < this.#day_count; h++) {
            current_day = new CalendarDay(h + 1, (new_month_days.length) % 7, new_month_days.length, this.#month, this);
            new_month_days.push(current_day);
        }

        new_month_days = [...new_month_days, ...this.#getLastWeekMissingDays(current_day, new_month_days.length)];
        new_month_days = new_month_days.sort((a, b) => a.CalendarOrder - b.CalendarOrder);

        this.#days = new_month_days;
    }
}


// DEBUG
// window.CalendarMonth = CalendarMonth;
// window.getDaysInMonth = getDaysInMonth;
// window.getWeekDay = getWeekDay;
// window.getMonthName = getMonthName;

// Path: src/components/Calendar/date_utils.js