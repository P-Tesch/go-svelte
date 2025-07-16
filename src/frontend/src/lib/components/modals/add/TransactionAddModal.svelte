<script lang="ts">
    import DateInput from "$lib/components/DateInput.svelte";
    import BaseAddModal from "$lib/components/modals/add/BaseAddModal.svelte";
    import TransactionStatus from "$lib/enums/TransactionStatus.js";
    import TransactionType from "$lib/enums/TransactionTypes.js";
    import { floatField, intField } from "$lib/scripts/Utils.svelte.js";

    let dialog: HTMLDialogElement;
    let statusValue: keyof typeof TransactionStatus | "null";
    let typeValue: keyof typeof TransactionType | "null";
    let entity: string = "transação";
    let quotesValue: number | null;

    export let modal: BaseAddModal;
</script>

<BaseAddModal bind:this={modal} bind:dialog bind:entity>
    <legend class="fieldset-legend">Nome</legend>
    <label class="input w-full">
        <svg class="h-[1.5em] opacity-50" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" d="M9.879 7.519c1.171-1.025 3.071-1.025 4.242 0 1.172 1.025 1.172 2.687 0 3.712-.203.179-.43.326-.67.442-.745.361-1.45.999-1.45 1.827v.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 5.25h.008v.008H12v-.008Z" />
        </svg>
        <input type="text" class="grow" placeholder="Ex.: Conta de internet" />
    </label>

    <legend class="fieldset-legend">Nº de parcelas</legend>
    <label class="input w-full">
        <svg class="h-[1.5em] opacity-50" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" d="M8.625 12a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Zm0 0H8.25m4.125 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Zm0 0H12m4.125 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Zm0 0h-.375M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
        </svg>
        <input bind:value={quotesValue} type="text" class="grow" placeholder="Ex.: 3" onbeforeinput="{intField}" />
    </label>

    <div class="flex flex-row justify-between w-full gap-2">
        <div class="w-full">
            <legend class="fieldset-legend">Valor {(quotesValue ?? 1) > 1 ? " da parcela" : ""}</legend>
            <label class="input w-full">
                <svg class="h-[1.5em] opacity-50" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 6v12m-3-2.818.879.659c1.171.879 3.07.879 4.242 0 1.172-.879 1.172-2.303 0-3.182C13.536 12.219 12.768 12 12 12c-.725 0-1.45-.22-2.003-.659-1.106-.879-1.106-2.303 0-3.182s2.9-.879 4.006 0l.415.33M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
                </svg>
                <input type="text" class="grow" placeholder="Ex.: R$ 89,90" onbeforeinput="{floatField}" />
            </label>
        </div>

        {#if (quotesValue ?? 1) > 1}
            <div class="w-full">
                <legend class="fieldset-legend">Valor total</legend>
                <label class="input w-full">
                    <svg class="h-[1.5em] opacity-50" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M12 6v12m-3-2.818.879.659c1.171.879 3.07.879 4.242 0 1.172-.879 1.172-2.303 0-3.182C13.536 12.219 12.768 12 12 12c-.725 0-1.45-.22-2.003-.659-1.106-.879-1.106-2.303 0-3.182s2.9-.879 4.006 0l.415.33M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
                    </svg>
                    <input type="text" class="grow" placeholder="Ex.: R$ 269,70" onbeforeinput="{floatField}" />
                </label>
            </div>
        {/if}
    </div>
        
    <DateInput parent={dialog} label="Data de vencimento"/>

    <legend class="fieldset-legend">Status</legend>
    <label class="input w-full">
        <svg class="h-[1.5em] opacity-50" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" d="m11.25 11.25.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9-3.75h.008v.008H12V8.25Z" />
        </svg>
        <select bind:value={statusValue} class="grow select select-ghost -ml-4 pr-0 focus:outline-0 focus-within:outline-0 focus:shadow-transparent" style="{statusValue === 'null' ? 'color: #697987;' : 'color: --color-base-content;'}">
            <option disabled selected value="null">Ex.: Pendente</option>
            {#each Object.entries(TransactionStatus) as value}
                <option value="{value[0]}">{value[1]}</option>
            {/each}
        </select>
    </label>

    <legend class="fieldset-legend">Tipo</legend>
    <label class="input w-full">
        {#if typeValue === "IN"}
            <svg class="h-[1.5em] opacity-50" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" d="m11.25 9-3 3m0 0 3 3m-3-3h7.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
            </svg>
        {:else}
            <svg class="h-[1.5em] opacity-50" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" d="m12.75 15 3-3m0 0-3-3m3 3h-7.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
            </svg>
        {/if}

        <select bind:value={typeValue} class="grow select select-ghost -ml-4 pr-0 focus:outline-0 focus-within:outline-0 focus:shadow-transparent" style="{statusValue === 'null' ? 'color: #697987;' : 'color: --color-base-content;'}">
            <option disabled selected value="null">Ex.: Saída</option>
            {#each Object.entries(TransactionType) as value}
                <option value="{value[0]}">{value[1]}</option>
            {/each}
        </select>
    </label>
</BaseAddModal>
