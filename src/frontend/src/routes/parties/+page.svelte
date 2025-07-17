<script lang="ts">
    import Header from "$lib/components/Header.svelte";
    import BaseAddModal from "$lib/components/modals/add/BaseAddModal.svelte";
    import PartyAddModal from "$lib/components/modals/add/PartyAddModal.svelte";
    import Table from "$lib/components/Table.svelte";
    import type Party from "$lib/types/Party.js";
    import { onMount } from "svelte";

    let entities: Party[] = [];
    let modal: BaseAddModal;

    onMount(() => {
        fetch("/api/parties", {
            method: "GET",
        })
        .then(async (response: Response) => {
            if (response.status === 200) {
                entities = JSON.parse(await response.json())
                return;
            }
            console.log(response);
        })
        .catch((error) => console.log(error));
    });
</script>

<Header />
<Table entities="{entities}" modal="{modal}" label="Partes"/>
<PartyAddModal bind:modal="{modal}" />