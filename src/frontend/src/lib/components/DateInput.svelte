<script lang="ts">
    import Pikaday from "pikaday";

    interface Props {
        parent: HTMLElement,
        label: string
    }

    const {parent, label}: Props = $props();

    let datePicker: HTMLInputElement;
    $effect(() => {
        if (datePicker !== undefined) {
            const picker = new Pikaday({
                field: datePicker,
                container: parent,
                reposition: false,
                position: "top",
                i18n: {
                    previousMonth: "Mês anterior",
                    nextMonth: "Próximo mês",
                    months: ["Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"],
                    weekdays: ["Domingo", "Segunda-feira", "Terça-feira", "Quarta-feira", "Quinta-Feira", "Sexta-feira", "Sábado"],
                    weekdaysShort: ["Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sab"],
                },
                toString(date, _format) {
                  return date.toLocaleDateString();
                },
            });
            return () => picker.destroy();
        }
    });
</script>


<legend class="fieldset-legend">{label}</legend>
<label class="input">
    <svg class="h-[1.5em] opacity-50" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" d="M12 6v12m-3-2.818.879.659c1.171.879 3.07.879 4.242 0 1.172-.879 1.172-2.303 0-3.182C13.536 12.219 12.768 12 12 12c-.725 0-1.45-.22-2.003-.659-1.106-.879-1.106-2.303 0-3.182s2.9-.879 4.006 0l.415.33M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
    </svg>
    <input type="text" class="grow pika-single"  bind:this={datePicker} placeholder="Ex.: 03/07/2025" />
</label>

<style>
    :global(.pika-single.is-bound) {
        margin-left: -9rem;
        margin-bottom: 13rem;
    }
</style>