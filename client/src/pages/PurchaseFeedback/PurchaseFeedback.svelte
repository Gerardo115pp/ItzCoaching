<script>
    import { PostPaymentConfirmationRequest } from "../../libs/HttpRequests";
    import { getUrlPARAM } from "../../libs/itz-utils";
    import itz_logo from "../../svg/MainLogo.svg";
    import { onMount } from "svelte";

    let session_id = getUrlPARAM("session_id");

    onMount(() => {
        confirmPayment();
    });

    const confirmPayment = () => {
        if (session_id === undefined && session_id === "") {
            console.error("No session id provided");
            return;
        }

        const request = new PostPaymentConfirmationRequest(session_id);

        const on_success = response => {
            console.log(response);
        };

        const on_error = error => {
            console.log(error);
        };

        request.do(on_success, on_error);
    }
</script>

<main id="successful-appointment">
    <div id="sa-center-content">
        <figure id="itz-logo-wrapper">
            {@html itz_logo}
        </figure>
        <h1>¡Tu cita se a reservado!</h1>
    </div>
</main>

<style>
    #successful-appointment {
        box-sizing: border-box;
        width: 100%;
        height: 100vh;
        display: grid;
        justify-content: center;
        align-items: center;
        background-color: var(--ready);
    }

    #sa-center-content {
        display: flex;
        flex-direction: column;
        align-items: center;
        color: var(--theme-pearl);
    }

    #itz-logo-wrapper {
        width: 70%;
    }

    :global(#itz-logo-wrapper svg) {
        width: 100%;
        height: 100%;
        fill: var(--theme-pearl);
    }

    #successful-appointment h1 {
        font-size: var(--font-size-h3);
        font-family: var(--font-titles);
    }
</style>