export const weekDays = {
    0: 'Sunday',
    1: 'Monday',
    2: 'Tuesday',
    3: 'Wednesday',
    4: 'Thursday',
    5: 'Friday',
    6: 'Saturday'
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

export const getCalendarMonth = (month, year) => {
    /* 
        The month parameter should be an integer representing the month number, with January being 0 and December being 11. 
        The year parameter should be the full year, for example 1993 or 2023, as a four-digit integer.
    */

    const first_day_month = new Date(year, month, 1);
    const first_week_day = first_day_month.getDay();
    const calendar_month = {};
    
}

class CalendarDay {
    #calendar_order;
    constructor(day_num, week_day, calendar_order, month) {
        this.#calendar_order = calendar_order;
        this.day_num = day_num;
        this.week_day = week_day;
        this.week_day_name = weekDays[week_day];
        this.month = month;
    }

    get CalendarOrder() {
        return this.#calendar_order;
    }
}

export class CalendarMonth {
    #name;
    #month;
    #year;
    #day_count;
    #days;
    #internal_date;
   
    constructor(year, month) {
        this.#internal_date = new Date(year, month, 1);
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
            current_day = new CalendarDay(previous_month_day_count - ((first_week_day-1) - h), h, first_week_day - (missing_days.length+1), this.#month - 1);
            missing_days.push(current_day);
        }

        return missing_days;
    }

    #getLastWeekMissingDays(last_day, calendar_order) {
        const last_week_day = last_day.week_day;
        const missing_days = [];
        let current_day;

        for (let h = last_week_day + 1; h <= 6; h++) {
            current_day = new CalendarDay((h - last_day.week_day), h, calendar_order, this.#month+1);
            missing_days.push(current_day);
            calendar_order++;
        }

        return missing_days;
    }

    #fillDaysInfo() {

        let new_month_days = this.#getFirstWeekMissingDays();
        let current_day;
        
        for (let h = 0; h < this.#day_count; h++) {
            current_day = new CalendarDay(h + 1, (new_month_days.length) % 7, new_month_days.length, this.#month);
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