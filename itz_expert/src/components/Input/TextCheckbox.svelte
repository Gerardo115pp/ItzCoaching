<script>
    import { onMount } from 'svelte';

    // Checkbox functionality
    export let checked = false;
    export let disabled = false;
    export let text_label = "Label";
    let checkbox_element;


    // Styles
    export let isSquared = false;
    export let background = "red";
    export let background_unchecked = "black";
    export let border_width = "1px";
    export let border_radius = "50%";
    export let color = "white";
    export let font = "sans-serif";
    export let font_size = "inherit";
    export let font_weight = "500";
    export let line_height = ".5";
    export let letter_spacing = "2px";
    export let padding = "1vw 1vw";
    export let width = "";
    
    onMount(() => {
        setStyles();
    });

    const checkBoxClicked = () => {
        if (!disabled) {
            checked = !checked;
        }
    }

    const getWidth = () => {
        if (width === "") {
            return `calc(${text_label.length}.3 * 2.2em)`
        } else {
            return width;
        }
    }

    function setStyles() {
        checkbox_element.style.setProperty('--libery-text-checkbox-background', background);
        checkbox_element.style.setProperty('--libery-text-checkbox-background-unchecked', background_unchecked);
        checkbox_element.style.setProperty('--libery-text-checkbox-border-width', border_width);
        checkbox_element.style.setProperty('--libery-text-checkbox-color', color);
        checkbox_element.style.setProperty('--libery-text-checkbox-font', font);
        checkbox_element.style.setProperty('--libery-text-checkbox-font-size', font_size);
        checkbox_element.style.setProperty('--libery-text-checkbox-font-weight', font_weight);
        checkbox_element.style.setProperty('--libery-text-checkbox-line-height', line_height);
        checkbox_element.style.setProperty('--libery-text-checkbox-letter-spacing', letter_spacing);
        checkbox_element.style.setProperty('--libery-text-checkbox-padding', padding);
        checkbox_element.style.setProperty('--libery-text-checkbox-width', getWidth());
        checkbox_element.style.setProperty('--libery-text-checkbox-border-radius', isSquared ? "0" : border_radius);
    }

</script>

<div bind:this={checkbox_element}  role="checkbox" aria-checked={checked} aria-disabled={disabled} on:click={checkBoxClicked} class="libery-text-checkbox {checked ? 'ltc-checked' : ''}">
    {text_label}
</div>

<style>
    .libery-text-checkbox {
        --libery-text-checkbox-background: red;
        --libery-text-checkbox-background-unchecked: black;
        --libery-text-checkbox-border-width: 1px;
        --libery-text-checkbox-color: white;
        --libery-text-checkbox-font: sans-serif;
        --libery-text-checkbox-font-size: inherit;
        --libery-text-checkbox-font-weight: inherit;
        --libery-text-checkbox-line-height: .5;
        --libery-text-checkbox-letter-spacing: 2px;
        --libery-text-checkbox-padding: 0.5rem;
        --libery-text-checkbox-width: max-content;
        --libery-text-checkbox-border-radius: 0;

        box-sizing: border-box;
        display: table-cell;
        vertical-align: middle;
        width: var(--libery-text-checkbox-width);
        background: var(--libery-text-checkbox-background-unchecked);
        border: var(--libery-text-checkbox-border-width) solid var(--libery-text-checkbox-background);
        color: var(--libery-text-checkbox-background);
        font-family: var(--libery-text-checkbox-font);
        font-size: var(--libery-text-checkbox-font-size);
        font-weight: var(--libery-text-checkbox-font-weight);
        letter-spacing: var(--libery-text-checkbox-letter-spacing);
        padding: var(--libery-text-checkbox-padding);
        border-radius: var(--libery-text-checkbox-border-radius);
        line-height: var(--libery-text-checkbox-line-height);
        user-select: none;
        text-align: center;
        text-transform: full-width;
    }

    .libery-text-checkbox.ltc-checked {
        background: var(--libery-text-checkbox-background);
        color: var(--libery-text-checkbox-color);
    }

    .libery-text-checkbox[aria-disabled="true"] {
        opacity: 0.5;
        cursor: not-allowed;
    }
</style>