<style>
    .error {
        background: #c52f21;
        color: #f5a390;
        padding-left: 0.7rem;
        padding-right: 0.7rem;
        padding-top: 0.3rem;
        padding-bottom: 0.3rem;
        border-radius: 2rem;
    }
</style>

<script lang="ts">
    import Layout from "$lib/components/layout.svelte"
    import { action } from "$lib/scripts/core/action.svelte"
    import { href } from "$lib/scripts/core/href.svelte"
    import { keyof } from "$lib/scripts/core/keyof"
    import { type Form } from "$lib/types/server/main/lib/routes/index/form"
    let disabled = $state(false)
    let error = $state("")
</script>

<Layout title="index">
    <div class="contain">
        <form
            method="POST"
            {...action("/index", {
                onpending() {
                    disabled = true
                },
                ondone() {
                    disabled = false
                },
                onerror(errorLocal) {
                    disabled = false
                    error = errorLocal.message
                },
            })}
        >
            <fieldset role="group">
                <input {disabled} name={keyof<Form>("Address")} type="text" placeholder="Address" />
                <input {disabled} name={keyof<Form>("Depth")} type="number" placeholder="Depth" />
                <input {disabled} type="submit" value="Index" />
            </fieldset>
        </form>
        {#if error}
            <br />
            <span class="error">{error}</span>
            <br />
            <br />
        {/if}
        <p>Or go back to <a {...href("/search")}>search</a>.</p>
    </div>
</Layout>
