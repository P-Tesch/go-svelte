<script lang="ts">
    let { entities }: {entities: Array<any>} = $props();
    let headers: string[] = $state([]);

    if (entities.length > 0) {
        headers = Object.getOwnPropertyNames(entities[0]);
    }
</script>

<div class="w-[92vw] h-[88vh] ml-[1.25vw] mr-[1.25vw] mt-[2.5vh] mb-[2.5vh] bg-base-200 rounded-box border-base-content/5">
    {#if entities.length === 0}
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