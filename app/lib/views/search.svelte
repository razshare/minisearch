<style lang="scss">
    .results {
        position: absolute;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
        display: grid;
        grid-template-areas:
            "searchbar"
            "separator"
            "items";
        grid-template-rows: auto auto 1fr;
        .searchbar {
            padding: 2rem;
            grid-area: searchbar;
        }
        & > hr {
            grid-area: separator;
            margin: 0;
        }
        .items {
            grid-area: items;
            overflow-y: auto;
            text-align: start;
            & > .item {
                padding: 2rem;
                & > .hint {
                    opacity: 0.5;
                }
            }
        }
    }
</style>

<script lang="ts">
    import Layout from "$lib/components/layout.svelte"
    import { action } from "$lib/scripts/core/action.svelte"
    import { href } from "$lib/scripts/core/href.svelte"
    import { keyof } from "$lib/scripts/core/keyof"
    import { type Form } from "$lib/types/server/main/lib/routes/search/form"
    import type { Props } from "$lib/types/server/main/lib/routes/search/props"
    let { Query, Items }: Props = $props()
</script>

<Layout title="Search">
    <div class="results">
        <div class="searchbar">
            <form method="GET" {...action("/search")}>
                <fieldset role="group">
                    <input name={keyof<Form>("Query")} type="text" placeholder="Search" value={Query} />
                    <input type="submit" value="Search" />
                </fieldset>
            </form>
            <span>Or <a {...href("/index")}>index</a> a new page.</span>
        </div>
        <hr />
        <div class="items">
            {#each Items as item (item)}
                <div class="item">
                    <a href={item.address} target="_blank">{item.description}</a>
                    <span class="hint"> - {item.address}</span>
                </div>
                <hr />
            {/each}
        </div>
    </div>
</Layout>
