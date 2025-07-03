<script lang="ts">
  import DateInput from "$lib/components/DateInput.svelte";
    import BaseAddModal from "$lib/components/modals/add/BaseAddModal.svelte";
  import TransactionStatus from "$lib/enums/TransactionStatus.js";

    let dialog: HTMLDialogElement;

    export let modal: BaseAddModal;
</script>

<BaseAddModal bind:this={modal} bind:dialog={dialog}>
    <h3 class="text-lg font-bold">Adicionar transação</h3>

    <legend class="fieldset-legend">Nome</legend>
    <label class="input">
        <svg class="h-[1.5em] opacity-50" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m2.25 0H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9Z" />
        </svg>
        <input type="text" class="grow" placeholder="Ex.: Conta de internet" />
    </label>

    <legend class="fieldset-legend">Valor</legend>
    <label class="input">
        <svg class="h-[1.5em] opacity-50" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 6v12m-3-2.818.879.659c1.171.879 3.07.879 4.242 0 1.172-.879 1.172-2.303 0-3.182C13.536 12.219 12.768 12 12 12c-.725 0-1.45-.22-2.003-.659-1.106-.879-1.106-2.303 0-3.182s2.9-.879 4.006 0l.415.33M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
        </svg>
        <input type="text" class="grow" placeholder="Ex.: R$ 89,90" />
    </label>
        
    <DateInput parent={dialog} label="Data de vencimento"/>

    <legend class="fieldset-legend">Status</legend>
    <select class="select">
        <option disabled selected>Selecione o status</option>
        {#each Object.entries(TransactionStatus) as value}
            <option value="{value[0]}">{value[1]}</option>
        {/each}
    </select>
</BaseAddModal>