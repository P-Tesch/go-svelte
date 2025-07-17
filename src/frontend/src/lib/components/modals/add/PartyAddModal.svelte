<script lang="ts">
    import BaseAddModal from "$lib/components/modals/add/BaseAddModal.svelte";
    import PartyTypes from "$lib/enums/PartyTypes.js";

    let dialog: HTMLDialogElement;
    let typeValue: keyof typeof PartyTypes | "null";
    let entity: string = "parte";
    let submitPath: string = "/api/parties";

    export let modal: BaseAddModal;
</script>

<BaseAddModal bind:this={modal} bind:dialog bind:entity bind:submitPath>
    <legend class="fieldset-legend">Nome</legend>
    <label class="input w-full">
        <svg class="h-[1.5em] opacity-50" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" d="M9.879 7.519c1.171-1.025 3.071-1.025 4.242 0 1.172 1.025 1.172 2.687 0 3.712-.203.179-.43.326-.67.442-.745.361-1.45.999-1.45 1.827v.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 5.25h.008v.008H12v-.008Z" />
        </svg>
        <input type="text" class="grow" placeholder="Ex.: Oi Telecom" name="name" />
    </label>

    <legend class="fieldset-legend">Status</legend>
    <label class="input w-full">
        <svg class="h-[1.5em] opacity-50" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" d="m11.25 11.25.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9-3.75h.008v.008H12V8.25Z" />
        </svg>
        <select bind:value={typeValue} name="type" class="grow select select-ghost -ml-4 pr-0 focus:outline-0 focus-within:outline-0 focus:shadow-transparent" style="{typeValue === 'null' ? 'color: #697987;' : 'color: --color-base-content;'}">
            <option disabled selected value="null">Ex.: Pessoa jurídica</option>
            {#each Object.entries(PartyTypes) as value}
                <option value="{value[0]}">{value[1]}</option>
            {/each}
        </select>
    </label>
</BaseAddModal>
