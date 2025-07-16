<script lang="ts">
     import { onMount } from "svelte";

    let navOpen = $state(false);

    onMount(() => {
        navOpen = sessionStorage.getItem("navOpen") === "true"
    });

    function toggleNav(): void {
        navOpen = !navOpen;
        sessionStorage.setItem("navOpen", "" + navOpen);
    }
</script>

<div class="w-[100vw] h-16 bg-base-300 left-0 top-0 flex flex-row z-40">
    <a href="/" class="flex flex-row h-full items-center ml-3 text-2xl">
        <img src="favicon.png" class="h-full pt-2 pb-2 pr-2" alt="logo"/>
        <strong>Velt</strong>
    </a>
</div>
<div class="{navOpen ? "w-64" : "w-20"} h-[100vh] bg-base-300 right-0 left-auto top-2 absolute z-50" style="transition: width 0.3s ease;">
    <button onclick="{toggleNav}" class="cursor-pointer w-full mt-2 flex items-center justify-end pr-5 mb-8">
        <h2 class="text-2xl mr-2 {navOpen ? "" : "hidden"}">Ocultar</h2>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-10">
            <path stroke-linecap="round" stroke-linejoin="round" d="{navOpen ? "m5.25 4.5 7.5 7.5-7.5 7.5m6-15 7.5 7.5-7.5 7.5" : "m18.75 4.5-7.5 7.5 7.5 7.5m-6-15L5.25 12l7.5 7.5"}" />
        </svg>
    </button>
    <a href="/transactions" class="cursor-pointer w-full mt-2 flex items-center justify-end pr-5 mb-8">
        <h2 class="text-2xl mr-2 {navOpen ? "" : "hidden"}">Transações</h2>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-10">
            <path stroke-linecap="round" stroke-linejoin="round" d="M3 7.5 7.5 3m0 0L12 7.5M7.5 3v13.5m13.5 0L16.5 21m0 0L12 16.5m4.5 4.5V7.5" />
        </svg>
    </a>
</div>