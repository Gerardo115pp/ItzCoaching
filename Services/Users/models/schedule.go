package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

type TimeSlot struct {
	Start time.Time `json:"start_time"`
	End   time.Time `json:"end_time"`
}

type ExpertSchedule struct {
	Expert_id         int             `json:"owner"`
	IsExpertAvailable bool            `json:"-"` // for this field to be trusted it must be assigned from whats on the db not from the schedule.json file
	WeekAvailability  map[string]bool `json:"week_availability"`
	TimeAvailability  []*TimeSlot     `json:"time_availability"`
}

func parseSchedule(schedule_path string) (*ExpertSchedule, error) {
	file_content, err := ioutil.ReadFile(schedule_path)
	if err != nil {
		return nil, err
	}

	var schedule *ExpertSchedule = new(ExpertSchedule)

	err = json.Unmarshal(file_content, schedule)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (ts *TimeSlot) sanitize(raw_slot map[string]string) error {
	// Check if start_time and end_time are present and start_time is before end_time

	start_time_str, ok := raw_slot["start_time"]
	if !ok {
		return errors.New("start_time field not found")
	}

	end_time_str, ok := raw_slot["end_time"]
	if !ok {
		return errors.New("end_time field not found")
	}

	start_time, err := time.Parse("3:04pm", start_time_str)
	if err != nil {
		start_time, err = time.Parse("15:04", start_time_str)
		if err != nil {
			return fmt.Errorf("invalid start_time: %s", start_time_str)
		}
	}

	end_time, err := time.Parse("3:04pm", end_time_str)
	if err != nil {
		end_time, err = time.Parse("15:04", end_time_str)
		if err != nil {
			return fmt.Errorf("invalid end_time: %s", end_time_str)
		}
	}

	ts.Start = start_time
	ts.End = end_time

	// Check if start_time is before end_time
	if ts.Start.After(end_time) {
		return errors.New("start_time is after end_time")
	}

	ts.Start = start_time

	return nil
}

func (ts *TimeSlot) UnmarshalJSON(data []byte) error {
	var aux map[string]string = make(map[string]string)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	return ts.sanitize(aux)
}

func (ts *TimeSlot) MarshalJSON() ([]byte, error) {
	aux := map[string]string{
		"start_time": ts.Start.Format("15:04"),
		"end_time":   ts.End.Format("15:04"),
	}
	return json.Marshal(aux)
}
