<script>
    import Input from '../../../components/Input/Input.svelte';
    import TextCheckbox from '../../../components/Input/TextCheckbox.svelte';
    import FieldData from '../../../libs/FieldData';

    let is_schedule_active = false;
    const avaliable_hours = [];
    const avaliable_days = {
        monday: false,
        tuesday: false,
        wednesday: false,
        thursday: false,
        friday: false,
        saturday: false,
        sunday: false,
    };
    
    let time_from = new FieldData('time-from', /.+/, 'time-from', 'time');
    let time_to = new FieldData('time-to', /.+/, 'time-to', 'time');


</script>

<figure id="itz-expert-schedule-manager">
    <header id="schedule-manager-header">
        <h3>
            Elige tu horario
        </h3>
    </header>
    <div id="schedule-manager-wrapper">
        <div class="smw-schdule-definer" id="time-definer">
            <h4 id="time-definer-label" class="itz-subtitle">
                Tus horas disponibles
            </h4>    
            <div class="inputs-wrapper">
                <div class="text-input-wrapper">
                    <Input 
                        field_data={time_from}
                        id="time-from" 
                        type="time"
                        input_label="Desde"
                        input_padding="0"
                        isClear
                    />
                </div>
                <div class="text-input-wrapper">
                    <Input 
                        field_data={time_to}
                        id="time-to" 
                        type="time"
                        input_label="Hasta"
                        input_padding="0"
                        isClear
                    />
                </div>
                <button class="full-btn-thin material-symbols-outlined">
                    add
                </button>
            </div>
            <ul id="expert-available-hours"></ul>
        </div>
        <div class="smw-schdule-definer" id="weekdays-definer">
            <h4 id="date-definer-label" class="itz-subtitle">
                Tus días disponibles
            </h4>
            <div class="inputs-wrapper">
                <TextCheckbox 
                    id="smw-sd-monday"
                    text_label="L"
                    border_width="2px"
                    background="var(--theme-red)"
                    background_unchecked="var(--clear-color)"
                    font_size="var(--font-size-3)"
                    font_weight="bolder"
                    font="Helvetica"
                    bind:checked={avaliable_days.monday}
                />
                <TextCheckbox 
                    id="smw-sd-tuesday"
                    text_label="M"
                    border_width="2px"
                    background="var(--theme-red)"
                    background_unchecked="var(--clear-color)"
                    font_size="var(--font-size-3)"
                    font_weight="bolder"
                    font="Helvetica"
                    bind:checked={avaliable_days.tuesday}
                />
                <TextCheckbox 
                    id="smw-sd-wednesday"
                    text_label="M"
                    border_width="2px"
                    background="var(--theme-red)"
                    background_unchecked="var(--clear-color)"
                    font_size="var(--font-size-3)"
                    font_weight="bolder"
                    font="Helvetica"
                    bind:checked={avaliable_days.wednesday}
                />
                <TextCheckbox 
                    id="smw-sd-thursday"
                    text_label="J"
                    border_width="2px"
                    background="var(--theme-red)"
                    background_unchecked="var(--clear-color)"
                    font_size="var(--font-size-3)"
                    font_weight="bolder"
                    font="Helvetica"
                    bind:checked={avaliable_days.thursday}
                />
                <TextCheckbox 
                    id="smw-sd-friday"
                    text_label="V"
                    border_width="2px"
                    background="var(--theme-red)"
                    background_unchecked="var(--clear-color)"
                    font_size="var(--font-size-3)"
                    font_weight="bolder"
                    font="Helvetica"
                    bind:checked={avaliable_days.friday}
                />
                <TextCheckbox 
                    id="smw-sd-saturday"
                    text_label="S"
                    border_width="2px"
                    background="var(--theme-red)"
                    background_unchecked="var(--clear-color)"
                    font_size="var(--font-size-3)"
                    font_weight="bolder"
                    font="Helvetica"
                    bind:checked={avaliable_days.saturday}
                />
                <TextCheckbox 
                    id="smw-sd-sunday"
                    text_label="D"
                    border_width="2px"
                    background="var(--theme-red)"
                    background_unchecked="var(--clear-color)"
                    font_size="var(--font-size-3)"
                    font_weight="bolder"
                    font="Helvetica"
                    bind:checked={avaliable_days.sunday}
                />
            </div>
        </div>
        <div class="smw-schdule-definer" id="availability-definer">
            <h4 class="itz-subtitle">
                Estas disponible?
            </h4>
            <div class="inputs-wrapper">
                <TextCheckbox 
                    id="smw-sd-availability"
                    text_label="activar"
                    border_width="2px"
                    background="var(--theme-red)"
                    background_unchecked="var(--clear-color)"
                    font_size="var(--font-size-2)"
                    border_radius="10px"
                    padding="calc(var(--font-size-3) * .6)"
                    font_weight="bolder"
                    font="var(--font-text)"
                    bind:checked={is_schedule_active}
                />
            </div>
        </div>
        <div class="smw-schdule-definer" id="save-btn-wrapper">
            <button class="full-btn">guardar</button>
        </div>
    </div>
</figure>

<style>
    /* DEBUG: BOXES */
    /* #schedule-manager-wrapper * {
        border: 1px solid blue;
    } */

    #itz-expert-schedule-manager {
        border: 1px solid var(--dark-light-color);
        border-radius: var(--boxes-roundness);
        padding: var(--spacing-3);
    }

    #schedule-manager-header {
        font-size: var(--font-size-h3);
        margin-bottom: var(--spacing-4);
    }

    #schedule-manager-wrapper {
        display: grid;
        grid-template-columns: repeat(7, 1fr);
        gap: var(--spacing-2);
    }

    #time-definer {
        grid-column: 1 / 3;
    }

    #time-definer .inputs-wrapper {
        display: flex;
        gap: var(--spacing-4);
    }

    .text-input-wrapper {
        width: min(30%, 10vw);
    }

    #expert-available-hours {
        display: flex;
        flex-wrap: wrap;
        gap: var(--spacing-2);
        height: calc(2 * var(--spacing-4));
    }

    #weekdays-definer {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        grid-column: 3 / 6;
    }

    #date-definer-label {
        margin-bottom: auto;
    }

    #weekdays-definer .inputs-wrapper {
        display: flex;
        width: 100%;
        justify-content: space-around;
        margin-bottom: auto;
    }

    #availability-definer {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        grid-column: 6 / 8;
    }

    #availability-definer h4 {
        margin-bottom: auto;
    }

    #availability-definer .inputs-wrapper {
        display: flex;
        width: 100%;
        justify-content: center;
        margin-bottom: auto;
    }

    #save-btn-wrapper {
        grid-column: 1 / 8;
    }

</style>