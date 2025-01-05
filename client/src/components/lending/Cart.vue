<template>
    <div class="p-6 sm:p-8 overflow-y-auto">
        <div class="sm:flex sm:items-center">
            <div class="sm:flex-auto">
                <h1 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">{{ t('lending.cartTitle') }}</h1>
                <p class="mt-2 text-sm text-gray-700 dark:text-gray-300">{{ t('lending.cartIntro') }}</p>
            </div>
        </div>
        <div class="mt-8 flow-root">
            <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
                <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
                    <div class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 sm:rounded-lg">
                        <ul class="divide-y divide-gray-200 dark:divide-zinc-800 bg-white dark:bg-zinc-700">
                            <li v-for="item in cartItems" :key="item.id" class="grid grid-cols-3">
                                <div>
                                    <img v-if="item.image != ''" class="w-full rounded-md border border-black/10 dark:border-white/10" :src="'data:image/jpg;base64,' + item.image" />
                                </div>
                                <div>
                                    <span>{{ item.title }}</span>
                                    <span>{{ item.id }}</span>
                                    <span>{{ item.number }}</span>
                                </div>
                                <div>
                                    <button type="button" class="btn-table-primary ml-2" :title="t('common.remove')" @click="removeItem(item.id)">
                                        <svg-icon type="mdi" :path="mdiCartRemove" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                                        {{ t('common.remove') }}
                                    </button>
                                </div>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useCartStore } from '@/stores/cart'
import axios from 'axios'
import SvgIcon from '@jamescoyle/vue-icon'
import {
  mdiCartRemove
} from '@mdi/js';

const { t } = useI18n()
const cartStore = useCartStore()

const cartItems: any = ref()
const groups: any = ref()
const categories: any = ref()
const places: any = ref()

console.log(cartStore.items)
const getItemData = () => {
    for (let i = 0; i < cartStore.items.length; i++) {
        console.log(i)
        const url = '/api/inventory/items/data/' + cartStore.items[i].id
        axios.get(url)
        .then((response: any) => {
            cartItems.value.push(response.data)
        })
    }
}

const removeItem = (itemID: number) => {
    for (let i = 0; i < cartStore.items.length; i++) {
        if (cartStore.items[i].id == itemID) {
            cartStore.items.splice(i, 1)
            break
        }
    }
}

getItemData()
</script>
