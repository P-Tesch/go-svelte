<script lang="ts">
    import Header from "$lib/components/Header.svelte";
    import Table from "$lib/components/Table.svelte";
    import type Transaction from "$lib/types/Transaction.js";

    let entities: Transaction[] = [];

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

</script>

<Header />
<Table entities="{entities}" />