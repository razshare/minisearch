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
    import { source } from "$lib/scripts/core/source"
    import { type Form } from "$lib/types/server/main/lib/routes/index/form"
    import type { Progress } from "$lib/types/server/main/lib/routes/index/progress"
    let pending = $state(false)
    let error = $state("")
    const progress = source("/events/index-progress").selectJson<Progress>()
</script>

<Layout title="index">
    <div class="contain">
        <form
            method="POST"
            {...action("/index", {
                onpending() {
                    pending = true
                },
                ondone() {
                    pending = false
                },
                onerror(errorLocal) {
                    pending = false
                    error = errorLocal.message
                },
            })}
        >
            <fieldset role="group">
                <input disabled={pending} name={keyof<Form>("Address")} type="text" placeholder="Address" />
                <input disabled={pending} name={keyof<Form>("Depth")} type="number" placeholder="Depth" />
                <input disabled={pending} type="submit" value="Index" />
            </fieldset>
        </form>
        {#if $progress && $progress.Current > 0 && $progress.Current !== $progress.Maximum}
            <progress value={$progress.Current} max={$progress.Maximum} />
        {/if}
        {#if error}
            <br />
            <span class="error">{error}</span>
            <br />
            <br />
        {/if}
        <p>Or go back to <a {...href("/search")}>search</a>.</p>
    </div>
</Layout>
