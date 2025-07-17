<script lang="ts">
    import Pikaday from "pikaday";

    interface Props {
        parent: HTMLElement,
        label: string,
        name: string
    }

    const {parent, label, name}: Props = $props();

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
<label class="input w-full">
    <svg class="h-[1.5em] opacity-50" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" d="M6.75 3v2.25M17.25 3v2.25M3 18.75V7.5a2.25 2.25 0 0 1 2.25-2.25h13.5A2.25 2.25 0 0 1 21 7.5v11.25m-18 0A2.25 2.25 0 0 0 5.25 21h13.5A2.25 2.25 0 0 0 21 18.75m-18 0v-7.5A2.25 2.25 0 0 1 5.25 9h13.5A2.25 2.25 0 0 1 21 11.25v7.5" />
    </svg>
    <input type="text" name="{name}" class="grow pika-single"  bind:this={datePicker} placeholder="Ex.: 03/07/2025" />
</label>

<style>
    :global(.pika-single.is-bound) {
        margin-left: -9rem;
        margin-bottom: 13rem;
    }
</style>