import { defineStore } from 'pinia'

interface Item {
    id: number,
    number: number
}

interface State {
    items: Item[]
}

export const useCartStore = defineStore('cart', {
    state: (): State => {
        return {
            items: []
        }
    }
})
