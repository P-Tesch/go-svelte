<script lang="ts">
    import type { Snippet } from "svelte";

    interface Props {
        children: Snippet,
        dialog: HTMLDialogElement
    }

    let {children, dialog = $bindable()}: Props = $props();

    export function showModal(): void {
        dialog.showModal();
    }
</script>

<dialog bind:this={dialog} class="modal">
    <div class="modal-box">
        {@render children()}
        <div class="modal-action">
            <form method="dialog">
                <button class="btn">Close</button>
            </form>
        </div>
    </div>
</dialog>