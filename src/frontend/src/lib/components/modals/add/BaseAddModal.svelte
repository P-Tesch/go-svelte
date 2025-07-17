<script lang="ts">
    import type { Snippet } from "svelte";

    interface Props {
        children: Snippet,
        dialog: HTMLDialogElement,
        entity: string,
        submitPath: string,
    }

    let {children, dialog = $bindable(), entity = $bindable(), submitPath = $bindable()}: Props = $props();

    export function showModal(): void {
        dialog.showModal();
    }

    function submitForm(event: SubmitEvent): void {
        event.preventDefault();

        console.log(event.target);

        fetch(submitPath, {
            method: "POST", 
            body: new FormData(event.target as HTMLFormElement)
        })
        .then((response: Response) => dialog.close());
    }
</script>

<dialog bind:this={dialog} class="modal">
    <div class="modal-box">
        <div class="flex flex-row justify-between">
            <h3 class="text-lg font-bold">Adicionar {entity}</h3>
            <button aria-label="close" class="btn btn-ghost btn-circle" onclick="{(): void => dialog.close()}">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12" />
                </svg>
            </button>
        </div>
        <div class="modal-action">
            <form class="w-full" onsubmit="{submitForm}">
                {@render children()}
                <input type="submit" class="btn mt-5" value="Salvar">
            </form>
        </div>
    </div>
</dialog>