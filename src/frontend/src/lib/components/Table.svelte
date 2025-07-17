<script lang="ts">
    import BaseAddModal from "$lib/components/modals/add/BaseAddModal.svelte";

    interface Props {
        entities: Array<any>,
        modal: BaseAddModal,
        label: string
    }

    let { entities, modal, label }: Props = $props();
    let headers: string[] = $state([]);

    if (entities?.length > 0) {
        headers = Object.getOwnPropertyNames(entities[0]);
    }

    function add(): void {
        modal.showModal();
    }
</script>

<div class="w-[92vw] h-[88vh] ml-[1.25vw] mr-[1.25vw] mt-[2.5vh] mb-[2.5vh] bg-base-200 rounded-box border-base-content/5">
    <div class="flex h-6 justify-between p-1">
        <div class="ml-2">
            <h2 class="text-xl">{label}</h2>
        </div>
        <div class="flex h-full gap-2 justify-end">
            <button onclick="{add}" aria-label="add" class="cursor-pointer z-10">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-10">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
                </svg>
            </button>
        </div>
    </div>
    <div>
        {#if entities === null || entities?.length === 0}
            <h1>Nenhum valor encontrado</h1>
        {:else}
            <table class="table table-zebra">
                <thead>
                    <tr>
                        <th></th>
                        {#each headers as header}
                            <th>{String(header).charAt(0).toUpperCase() + String(header).slice(1)}</th>
                        {/each}
                    </tr>
                </thead>
                <tbody>
                    {#each entities as entity, index}
                        <tr>
                            <td>{index + 1}</td>
                            {#each Object.values(entity) as info}
                                <td>{info}</td>
                            {/each}
                        </tr>
                    {/each}
                </tbody>
            </table>
        {/if}
    </div>
</div>