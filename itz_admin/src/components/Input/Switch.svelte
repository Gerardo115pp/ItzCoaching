<script>
    export let checked = false;
    export let color = "#2196F3";
    
    export let onChange = () => {};

    let switch_style_sheet;
    let switch_color_rule;
    $: setSwitchColor(color);

    function getStyleSheet() {
        if (document.getElementById("switch-style")) {
            return document.getElementById("switch-style").sheet;
        }
        let style = document.createElement('style');
        style.id = "switch-style";
        document.head.appendChild(style);
        return style.sheet;
    }

    function setSwitchColor(new_color) {
        if (switch_style_sheet === undefined) {
            switch_style_sheet = getStyleSheet();
        }

        const new_color_rule = `.libery-switch > input:checked + .slider { background-color: ${new_color}; box-shadow: 0 0 1px ${new_color}; }`;
        if (new_color_rule !== switch_color_rule) {
            switch_color_rule = new_color_rule;
            if (switch_style_sheet.cssRules.length > 0) {
                switch_style_sheet.removeRule(0);
            }
            switch_style_sheet.insertRule(switch_color_rule, 0);
        }
    }
  </script>
  
<style>
.libery-switch {
    --libery-switch-width: 1.6vw;
    --libery-switch-height: calc(.56 * var(--libery-switch-width));

    position: relative;
    display: inline-block;
    /* width: 60px; */
    /* height: 34px; */
    width: var(--libery-switch-width);
    height: var(--libery-switch-height);
}

.libery-switch input {
    opacity: 0;
    width: 0;
    height: 0;
}

.slider {
    position: absolute;
    cursor: pointer;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: #ccc;
    -webkit-transition: 0.4s;
    transition: 0.4s;
    border-radius: 34px;
}

.slider:before {
    position: absolute;
    content: "";
    height: calc(.43 * var(--libery-switch-width));
    width: calc(.43 * var(--libery-switch-width));
    left: calc(.06 * var(--libery-switch-width));
    bottom: calc(.06 * var(--libery-switch-width));
    background-color: white;
    -webkit-transition: 0.4s;
    transition: 0.4s;
    border-radius: 50%;
}

input:checked + .slider:before {
    -webkit-transform: translateX(43%);
    -ms-transform: translateX(43%);
    transform: translateX(100%);
}
</style>

<label class="libery-switch">
    <input on:change={onChange}  type="checkbox" bind:checked />
    <span class="slider" />
</label>