export type Props = search.Props

export declare namespace search {
    export type Props = {
        Query: string
        Items: null|(schema.Result[])
    }
}

export declare namespace schema {
    export type Result = {
        id: string
        address: string
        description: string
    }
}