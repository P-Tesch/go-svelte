<script context="module" lang="ts">
    declare const dialog: HTMLDialogElement
</script>

<script lang="ts">
    import type Modal from "$lib/types/Modal.js";

    export const modal: Modal = {
        showModal
    };

    function showModal(): void {
        dialog.showModal();
    }
</script>

<dialog id="dialog" class="modal">
    <div class="modal-box">
        <h3 class="text-lg font-bold">Hello!</h3>
        <p class="py-4">Press ESC key or click the button below to close</p>
        <div class="modal-action">
            <form method="dialog">
                <button class="btn">Close</button>
            </form>
        </div>
    </div>
</dialog>