<script lang="ts">
    import Header from "$lib/components/Header.svelte";
    import BaseAddModal from "$lib/components/modals/add/BaseAddModal.svelte";
    import TransactionAddModal from "$lib/components/modals/add/TransactionAddModal.svelte";
    import Table from "$lib/components/Table.svelte";
    import type Transaction from "$lib/types/Transaction.js";
    import { onMount } from "svelte";

    let entities: Transaction[] = [];
    let modal: BaseAddModal;

    onMount(() => {
        fetch("/api/transactions", {
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
<Table entities="{entities}" modal="{modal}" label="Transações"/>
<TransactionAddModal bind:modal="{modal}" />