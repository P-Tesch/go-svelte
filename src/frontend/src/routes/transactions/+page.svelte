<script lang="ts">
    import Header from "$lib/components/Header.svelte";
    import Table from "$lib/components/Table.svelte";

    const entities = [
        {name: "teste1", value: "teste1a"},
        {name: "teste2", value: "teste2a"},
        {name: "teste3", value: "teste3a"}
    ];

</script>

<Header />
<Table entities="{entities}" />