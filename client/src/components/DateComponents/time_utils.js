import { MonthDay } from './date_utils';

export const convertLocalTimeToUTC = time_string => {
    // time_string is a string representing a time in the format "HH:MM" in the local timezone
    // returns a string representing the same time in UTC with the format "HH:MM"
    const time = new Date();
    console.log("Time string:", time_string);
    if (time_string.search(/\s?[APapmM]{2}/) >= 0) {
        console.log("Time string is in the format 'HH:MM AM/PM'");
        // time_string is in the format "HH:MM AM/PM"
        const [hours, minutes_am_pm] = time_string.split(":");
        const [minutes, am_pm] = minutes_am_pm.split(" ");
        time.setHours(parseInt(hours) + (am_pm === "PM" ? 12 : 0));
        time.setMinutes(minutes);
    } else {
        console.log("Time string is in the format 'HH:MM'");
        const [hours, minutes] = time_string.split(":");
        time.setHours(hours);
        time.setMinutes(minutes);
        time.setSeconds(0);
        time.setMilliseconds(0);
    }

    return time;
}

export const convertLocalTimeToUTCString = time_string => {
    const time = convertLocalTimeToUTC(time_string);
    return `${time.getUTCHours().toString().padStart(2, '0')}:${time.getUTCMinutes().toString().padStart(2, '0')}`
}

export const convertUTCtoLocalTime = utc_time => {
    // utc_time is a string representing a time in the format "HH:MM" in UTC
    // returns a string representing the same time in the local timezone with the format "HH:MM"

    const time = new Date();
    const [hours, minutes] = utc_time.split(":");
    time.setUTCHours(hours);
    time.setUTCMinutes(minutes);
    time.setUTCSeconds(0);

    return time;
}

export const convertUTCtoLocalTimeString = utc_time => {
    const time = convertUTCtoLocalTime(utc_time);
    return `${time.getHours().toString().padStart(2, '0')}:${time.getMinutes().toString().padStart(2, '0')}`;
}

const isBefore = (time1, time2) => {
    return time1 < time2;
}

export const getDayHours = month_day => {
    const day_hours = [];
    let current_hour;

    for (let h = 0; h < 24; h++) {
        let next_hour;

        current_hour = new Date(month_day.toDate());
        current_hour.setHours(h, 0, 0, 0);

        next_hour = new Date(month_day.toDate());
        next_hour.setHours(h + 1, 0, 0, 0);

        day_hours.push(new TimeSlot(current_hour, next_hour));
    }

    return day_hours;
}

const formatTimeToDayHour = time => {
    const hour = time.getHours().toString().padStart(2, '0');
    const minutes = time.getMinutes().toString().padStart(2, '0');

    return `${hour}:${minutes}`;
}

export class TimeSlot {
    #duration;
    #start;
    #end;
    constructor(start, end) {
        if (isBefore(end, start)) {
            start.setDate(start.getDate() - 1);
        }
        this.#start = start;
        this.#end = end;
        this.#duration = this.#calculateDuration();
    }

    get Duration() {
        return this.#duration;
    }

    get Start() {
        return this.#start;
    }

    get HourLocaleStart() {
        return formatTimeToDayHour(this.#start);
    }

    set Start(value) {
        if (isBefore(value, this.#end)) {
            this.#start = value;
            this.#duration = this.#calculateDuration();
        }
    }

    get End() {
        return this.#end;
    }

    get HourLocaleEnd() {
        return formatTimeToDayHour(this.#end);
    }

    set End(value) {
        if (isBefore(this.#start, value)) {
            this.#end = value;
            this.#duration = this.#calculateDuration();
        }
    }

    #calculateDuration() {
        return this.#end - this.#start;
    }

    #cloneWithMachingDates = (start, end) => {
        // this method receives the start and end of another TimeSlot and clones the dates of this TimeSlot to compare only the times
        const new_start = new Date(this.#start);
        const new_end = new Date(this.#end);

        new_start.setHours(start.getHours(), start.getMinutes(), 0, 0);
        new_end.setHours(end.getHours(), end.getMinutes(), 0, 0);

        return new TimeSlot(new_start, new_end);
    }

    inTimeframe = tsin => {
        tsin = this.#cloneWithMachingDates(tsin.Start, tsin.End);
        return !isBefore(tsin.Start, this.#start) && isBefore(tsin.End, this.#end);
    }

    overlaps = tsin => {
        tsin = this.#cloneWithMachingDates(tsin.Start, tsin.End);
        let overlaps = this.inTimeframe(tsin) || tsin.inTimeframe(this);

        overlaps = overlaps || !isBefore(this.#start, tsin.Start) && isBefore(this.#start, tsin.End);
        overlaps = overlaps || !isBefore(this.#end, tsin.Start) && isBefore(this.#end, tsin.End);

        return overlaps;
    }

    toString = () => {
        return `${this.#start.getHours()}:${this.#start.getMinutes()} -  ${this.#end.getHours()}:${this.#end.getMinutes()}>`;
    }
}