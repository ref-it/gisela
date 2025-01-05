<template>
    <div class="p-6 sm:p-8 overflow-y-auto h-full">
        <div class="xl:flex xl:items-center">
            <div class="sm:flex-auto">
                <h1 class="text-base font-semibold leading-6 text-zinc-800 dark:text-white">{{ t('inventory.inventoryTitle') }}</h1>
                <p class="mt-2 tex sm:p-8t-sm text-zinc-800 dark:text-gray-300">{{ t('inventory.inventoryIntro') }}</p>
            </div>
            <div class="mt-4 xl:ml-16 xl:mt-0 flex whitespace-nowrap">
                <router-link to="/inventory/items/add" type="button" class="btn-primary">
                    <svg-icon type="mdi" :path="mdiFilePlus" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                    {{ t('inventory.addItem') }}
                </router-link>
                <router-link to="/inventory/items/retired" type="button" class="btn-neutral ml-2">
                    <svg-icon type="mdi" :path="mdiFileRemove" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                    {{ t('inventory.retiredItems') }}
                </router-link>
            </div>
        </div>
        <div class="mt-8 flow-root">
            <div class="mb-8 grid sm:grid-cols-2 xl:grid-cols-[repeat(4,minmax(0,1fr))_auto] gap-6">
                <div>
                    <label for="filter-title">{{ t('common.title') }}</label>
                    <input id="filter-title" type="text" v-model="showTitle">
                </div>
                <div>
                    <Listbox id="filter-category" as="div" v-model="showCategories" multiple>
                        <ListboxLabel>{{ t('common.category') }}</ListboxLabel>
                        <div class="relative">
                            <ListboxButton class="relative w-full cursor-default rounded-md bg-white dark:text-white dark:bg-white/5 dark:ring-white/10 py-2 pl-3 pr-10 h-10 text-left text-zinc-900 shadow-sm ring-1 ring-inset ring-zinc-300 focus:outline-none focus:ring-2 focus:ring-[var(--accent-color)] sm:text-sm sm:leading-6">
                                <span class="block truncate">{{ showCategories.map((cat: any) => cat.title).join(', ') }}</span>
                                <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                                    <svg-icon type="mdi" :path="mdiChevronDown" class="h-5 w-5 text-zinc-400" aria-hidden="true"></svg-icon>
                                </span>
                            </ListboxButton>

                            <transition leave-active-class="transition ease-in duration-100" leave-from-class="opacity-100" leave-to-class="opacity-0">
                                <ListboxOptions class="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white dark:bg-zinc-800 py-1 shadow-lg ring-1 ring-black dark:ring-white/10 ring-opacity-5 focus:outline-none text-sm">
                                    <ListboxOption as="template" v-for="item in categories" :key="item.id" :value="item" v-slot="{ active, selected, showCategories }">
                                        <li class="hover:bg-zinc-200 focus:bg-zinc-200 dark:hover:bg-zinc-700 dark:focus:bg-zinc-700 text-zinc-800 dark:text-white relative cursor-default select-none py-2 pl-3 pr-9">
                                            <span class="block truncate">{{ item.title }}</span>

                                            <span v-if="selected" class="text-[var(--accent-color)] absolute inset-y-0 right-0 flex items-center pr-4">
                                                <svg-icon type="mdi" :path="mdiCheck" class="h-5 w-5" aria-hidden="true"></svg-icon>
                                            </span>
                                        </li>
                                    </ListboxOption>
                                </ListboxOptions>
                            </transition>
                        </div>
                    </Listbox>
                </div>
                <div>
                    <Listbox id="filter-group" as="div" v-model="showGroups" multiple>
                        <ListboxLabel>{{ t('common.group') }}</ListboxLabel>
                        <div class="relative">
                            <ListboxButton class="relative w-full cursor-default rounded-md bg-white dark:text-white dark:bg-white/5 dark:ring-white/10 py-2 pl-3 pr-10 h-10 text-left text-zinc-900 shadow-sm ring-1 ring-inset ring-zinc-300 focus:outline-none focus:ring-2 focus:ring-[var(--accent-color)] sm:text-sm sm:leading-6">
                                <span class="block truncate">{{ showGroups.map((group: any) => group.title).join(', ') }}</span>
                                <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                                    <svg-icon type="mdi" :path="mdiChevronDown" class="h-5 w-5 text-zinc-400" aria-hidden="true"></svg-icon>
                                </span>
                            </ListboxButton>

                            <transition leave-active-class="transition ease-in duration-100" leave-from-class="opacity-100" leave-to-class="opacity-0">
                                <ListboxOptions class="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white dark:bg-zinc-800 py-1 shadow-lg ring-1 ring-black dark:ring-white/10 ring-opacity-5 focus:outline-none text-sm">
                                    <ListboxOption as="template" v-for="item in groups" :key="item.id" :value="item" v-slot="{ active, selected, showCategories }">
                                        <li class="hover:bg-zinc-200 focus:bg-zinc-200 dark:hover:bg-zinc-700 dark:focus:bg-zinc-700 text-zinc-800 dark:text-white relative cursor-default select-none py-2 pl-3 pr-9">
                                            <span class="block truncate">{{ item.title }}</span>

                                            <span v-if="selected" class="text-[var(--accent-color)] absolute inset-y-0 right-0 flex items-center pr-4">
                                                <svg-icon type="mdi" :path="mdiCheck" class="h-5 w-5" aria-hidden="true"></svg-icon>
                                            </span>
                                        </li>
                                    </ListboxOption>
                                </ListboxOptions>
                            </transition>
                        </div>
                    </Listbox>
                </div>
                <div>
                    <Listbox as="div" v-model="showBorrowable">
                        <ListboxLabel>{{ t('common.borrowable') }}</ListboxLabel>
                        <div class="relative">
                            <ListboxButton class="relative w-full cursor-default rounded-md bg-white dark:text-white dark:bg-white/5 dark:ring-white/10 py-2 pl-3 pr-10 text-left text-zinc-900 shadow-sm ring-1 ring-inset ring-zinc-300 focus:outline-none focus:ring-2 focus:ring-[var(--accent-color)] sm:text-sm sm:leading-6">
                            <span class="block truncate">{{ showBorrowable.title }}</span>
                            <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                                <svg-icon type="mdi" :path="mdiChevronDown" class="h-5 w-5 text-zinc-400" aria-hidden="true"></svg-icon>
                            </span>
                            </ListboxButton>

                            <transition leave-active-class="transition ease-in duration-100" leave-from-class="opacity-100" leave-to-class="opacity-0">
                            <ListboxOptions class="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white dark:bg-zinc-800 py-1 text-base shadow-lg ring-1 ring-black dark:ring-white/10 ring-opacity-5 focus:outline-none sm:text-sm">
                                <ListboxOption as="template" v-for="item in borrowableOptions" :key="item.id" :value="item" v-slot="{ active, borrowable }">
                                    <li class="hover:bg-zinc-200 focus:bg-zinc-200 dark:hover:bg-zinc-700 dark:focus:bg-zinc-700 text-zinc-800 dark:text-white relative cursor-default select-none py-2 pl-3 pr-9">
                                        <span class="block truncate">{{ item.title }}</span>

                                        <span v-if="borrowable" class="text-[var(--accent-color)] absolute inset-y-0 right-0 flex items-center pr-4">
                                            <CheckIcon class="h-5 w-5" aria-hidden="true" />
                                        </span>
                                    </li>
                                </ListboxOption>
                            </ListboxOptions>
                            </transition>
                        </div>
                    </Listbox>
                </div>
                <div class="flex">
                    <button type="button" class="btn-primary mt-auto" @click="start = 0; reloadItems()">
                        <svg-icon type="mdi" :path="mdiReload" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                        Neu laden
                    </button>
                </div>
            </div>
            <div class="-mx-6 -my-2 overflow-x-auto sm:-mx-8">
                <div class="inline-block min-w-full py-2 align-middle sm:px-8">
                    <div v-if="itemsTotal > 0" class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 sm:rounded-lg">
                        <table>
                            <thead>
                                <tr>
                                    <th scope="col">{{ t('common.inventoryNumber') }}</th>
                                    <th scope="col">{{ t('common.title') }}</th>
                                    <th scope="col">{{ t('common.category') }}</th>
                                    <th scope="col">{{ t('common.place') }}</th>
                                    <th scope="col">{{ t('common.group') }}</th>
                                    <th scope="col">{{ t('common.borrowable') }}</th>
                                    <th></th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="item in items" :key="item.id">
                                    <td>{{ item.id }}</td>
                                    <td>{{ item.title }}</td>
                                    <td>{{ item.category }}</td>
                                    <td>{{ item.place }}</td>
                                    <td>{{ item.group }}</td>
                                    <td>
                                        <span v-if="item.borrowable">
                                            <svg-icon type="mdi" :path="mdiCheck" class="text-green-600 dark:text-green-500" aria-hidden="true"></svg-icon>
                                            <span class="sr-only">{{ t('inventory.borrowable') }}</span>
                                        </span>
                                        <span v-else>
                                            <svg-icon type="mdi" :path="mdiClose" class="text-red-600 dark:text-red-300" aria-hidden="true"></svg-icon>
                                            <span class="sr-only">{{ t('inventory.notBorrowable') }}</span>
                                        </span>
                                    </td>
                                    <td class="whitespace-nowrap text-right">
                                        <Menu as="div" class="relative inline-block text-left">
                                            <div>
                                                <MenuButton class="rounded-md px-2 py-1 -my-1 hover:bg-zinc-200 dark:hover:bg-zinc-600 focus:bg-zinc-200 dark:focus:bg-zinc-600 active:bg-zinc-200 dark:active:bg-zinc-600">
                                                    <svg-icon type="mdi" :path="mdiDotsVertical" class="text-zinc-800 dark:text-white h-5 w-5" aria-hidden="true" />
                                                    <span class="sr-only">Optionen</span>
                                                </MenuButton>
                                            </div>

                                            <transition enter-active-class="transition ease-out duration-100" enter-from-class="transform opacity-0 scale-95" enter-to-class="transform opacity-100 scale-100" leave-active-class="transition ease-in duration-75" leave-from-class="transform opacity-100 scale-100" leave-to-class="transform opacity-0 scale-95">
                                                <MenuItems class="absolute right-0 z-10 mt-2 origin-top-right divide-y divide-zinc-200 dark:divide-zinc-700 rounded-md bg-white dark:bg-zinc-800 shadow-lg ring-1 ring-black/5 dark:ring-white/5">
                                                    <div class="py-1">
                                                        <MenuItem v-slot="{ active }">
                                                            <router-link :to="'/inventory/items/view/' + item.id" :class="[active ? 'bg-zinc-200 dark:bg-zinc-700' : '', 'text-zinc-800 dark:text-white group flex items-center px-4 py-2 text-sm']">
                                                                <svg-icon type="mdi" :path="mdiFileEye" class="mr-3 size-5 text-zinc-500 dark:text-zinc-300" aria-hidden="true" />
                                                                {{ t('common.view') }}
                                                            </router-link>
                                                        </MenuItem>
                                                        <MenuItem v-slot="{ active }">
                                                            <router-link :to="'/inventory/items/edit/' + item.id" :class="[active ? 'bg-zinc-200 dark:bg-zinc-700' : '', 'text-zinc-800 dark:text-white group flex items-center px-4 py-2 text-sm']">
                                                                <svg-icon type="mdi" :path="mdiFileEdit" class="mr-3 size-5 text-zinc-500 dark:text-zinc-300" aria-hidden="true" />
                                                                {{ t('common.edit') }}
                                                            </router-link>
                                                        </MenuItem>
                                                    </div>
                                                    <div v-if="item.borrowable" class="py-1">
                                                        <MenuItem v-slot="{ active }">
                                                            <button type="button"
                                                                :class="[active ? 'bg-zinc-200 dark:bg-zinc-700' : '', 'text-zinc-800 dark:text-white group flex items-center px-4 py-2 w-full text-sm']"
                                                                @click="addItemToCart(item.id)"
                                                            >
                                                                <svg-icon type="mdi" :path="mdiCartPlus" class="mr-3 size-5 text-zinc-500 dark:text-zinc-300" aria-hidden="true" />
                                                                {{ t('common.add') }}
                                                            </button>
                                                        </MenuItem>
                                                    </div>
                                                    <div class="py-1">
                                                        <MenuItem v-slot="{ active }">
                                                            <button type="button" :class="[active ? 'bg-zinc-200 dark:bg-zinc-700' : '', ' text-red-700 dark:text-red-300 group flex items-center px-4 py-2 text-sm']">
                                                                <svg-icon type="mdi" :path="mdiFileRemove" class="mr-3 size-5 text-red-700 dark:text-red-300" aria-hidden="true" />
                                                                {{ t('common.requestRetirement') }}
                                                            </button>
                                                        </MenuItem>
                                                    </div>
                                                </MenuItems>
                                            </transition>
                                        </Menu>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                    <div v-else>
                        <p class="text-zinc-800 dark:text-white">Die Suche ergab keine Ergebnisse.</p>
                    </div>
                </div>
            </div>
        </div>
        <div v-if="itemsTotal > 0" class="flex flex-col sm:flex-row mt-6">
            <div class="flex">
                <div class="flex rounded-l-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-[var(--accent-color)] w-20">
                    <input type="number" v-model="noOfItems" class="block !w-20 flex-1 border-0 bg-transparent py-2 text-zinc-800 dark:text-white dark:bg-white/5 placeholder:text-zinc-500 focus:ring-0 sm:text-sm sm:leading-6 !rounded-r-none" />
                </div>
                <div class="flex px-4 py-2 shadow-sm border-y border-gray-300 dark:text-white dark:border-white/10 w-full">
                    {{ t('common.itemsPerPage') }}
                </div>
                <button type="button" class="btn-neutral rounded-none rounded-r-md" @click="start = 0; reloadItems()">
                    <svg-icon type="mdi" :path="mdiReload" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                    <span class="sr-only">Neu laden</span>
                </button>
            </div>
            <div class="ml-auto mt-4 sm:mt-0">
                <div class="flex gap-1">
                    <button type="button" class="btn-neutral" @click="showFirst()" :title="t('common.first')" :disabled="!(start - noOfItems > -noOfItems)">
                        <svg-icon type="mdi" :path="mdiChevronDoubleLeft" class="h-5 w-5" aria-hidden="true"></svg-icon>
                    </button>
                    <button type="button" class="btn-neutral" @click="showPrevious()" :title="t('common.previous')" :disabled="!(start - noOfItems > -noOfItems)">
                        <svg-icon type="mdi" :path="mdiChevronLeft" class="h-5 w-5" aria-hidden="true"></svg-icon>
                    </button>
                    <!-- replace with select -->
                    <input type="number" min="1" :max="noOfPages" v-model="currentPage">
                    <button type="button" class="btn-neutral" @click="showNext()" :title="t('common.next')" :disabled="!(start + items.length < itemsTotal)">
                        <svg-icon type="mdi" :path="mdiChevronRight" class="h-5 w-5" aria-hidden="true"></svg-icon>
                    </button>
                    <button type="button" class="btn-neutral" @click="showLast()" :title="t('common.last')" :disabled="!(start + items.length < itemsTotal)">
                        <svg-icon type="mdi" :path="mdiChevronDoubleRight" class="h-5 w-5" aria-hidden="true"></svg-icon>
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script setup lang="ts">
import { type Ref, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useCartStore } from '@/stores/cart'
import axios from 'axios'
import {
    Menu,
    MenuButton,
    MenuItem,
    MenuItems,
    Listbox,
    ListboxButton,
    ListboxLabel,
    ListboxOption,
    ListboxOptions,
    Switch
} from '@headlessui/vue'
import SvgIcon from '@jamescoyle/vue-icon'
import {
    mdiCartPlus,
    mdiCheck,
    mdiChevronDoubleLeft,
    mdiChevronDoubleRight,
    mdiChevronDown,
    mdiChevronLeft,
    mdiChevronRight,
    mdiClose,
    mdiDotsVertical,
    mdiFileEye,
    mdiFileEdit,
    mdiFilePlus,
    mdiFileRemove,
    mdiReload
} from '@mdi/js';

const { t } = useI18n()
const cartStore = useCartStore()

const items: any = ref()
const groups: any = ref()
const categories: any = ref()
const places: any = ref()

const start: Ref<number> = ref(0)
const noOfItems: Ref<number> = ref(10)
const itemsTotal: Ref<number> = ref(0)
const noOfPages: Ref<number> = ref(0)
const currentPage: Ref<number> = ref(1)

const showTitle: Ref<string> = ref('')
const showCategories: any = ref([])
const showGroups: any = ref([])

const borrowableOptions = [
    {
        id: 0,
        title: "egal"
    },
    {
        id: 1,
        title: "ja"
    },
    {
        id: 2,
        title: "nein"
    }
]
const showBorrowable = ref(borrowableOptions[0])

const getItems = () => {
    const formData = new FormData()
    let url = '/api/inventory/items/index?display=' + noOfItems.value + '&start=' + start.value
    axios.post(url, formData)
        .then((response: any) => {
            let itemsResponse = response.data.items
            itemsTotal.value = response.data.itemsTotal

            let url = '/api/inventory/groups/index'
            axios
                .get(url)
                .then((response: any) => {
                    groups.value = response.data
                    url = '/api/inventory/categories/index'
                    axios
                        .get(url)
                        .then((response: any) => {
                            categories.value = response.data
                            url = '/api/inventory/places/index'
                            axios
                                .get(url)
                                .then((response: any) => {
                                    places.value = response.data
                                    for (let i = 0; i < itemsResponse.length; i++) {
                                        for (let j = 0; j < groups.value.length; j++) {
                                            if (groups.value[j].id == itemsResponse[i].group) {
                                                itemsResponse[i].group = groups.value[j].title
                                            }
                                        }
                                        for (let j = 0; j < categories.value.length; j++) {
                                            if (categories.value[j].id == itemsResponse[i].category) {
                                                itemsResponse[i].category = categories.value[j].title
                                            }
                                        }
                                        for (let j = 0; j < places.value.length; j++) {
                                            if (places.value[j].id == itemsResponse[i].place) {
                                                itemsResponse[i].place= places.value[j].title
                                            }
                                        }
                                    }
                                    items.value = itemsResponse
                                    noOfPages.value = Math.trunc(itemsTotal.value / noOfItems.value)
                                })
                        })
                })
            })
}

const reloadItems = () => {
    const formData = new FormData()
    formData.append("title", showTitle.value)

    let categoriesRequest: string = ""
    for (let i = 0; i < showCategories.value.length; i++) {
        if (categoriesRequest != "") {
            categoriesRequest += ","
        }
        categoriesRequest += showCategories.value[i].id
    }
    formData.append("categories", categoriesRequest)

    let groupsRequest: string = ""
    for (let i = 0; i < showGroups.value.length; i++) {
        if (groupsRequest != "") {
            groupsRequest += ","
        }
        groupsRequest += showGroups.value[i].id
    }
    formData.append("groups", groupsRequest)

    let url = '/api/inventory/items/index?display=' + noOfItems.value + '&start=' + start.value
    if (showBorrowable.value.id == 1) {
        url += '&borrowable=true'
    } else if (showBorrowable.value.id == 2) {
        url += '&borrowable=false'
    }

    axios.post(url, formData)
    .then((response: any) => {
        let itemsResponse = response.data.items
        itemsTotal.value = response.data.itemsTotal
        for (let i = 0; i < itemsResponse.length; i++) {
            for (let j = 0; j < groups.value.length; j++) {
                if (groups.value[j].id == itemsResponse[i].group) {
                    itemsResponse[i].group = groups.value[j].title
                }
            }
            for (let j = 0; j < categories.value.length; j++) {
                if (categories.value[j].id == itemsResponse[i].category) {
                    itemsResponse[i].category = categories.value[j].title
                }
            }
            for (let j = 0; j < places.value.length; j++) {
                if (places.value[j].id == itemsResponse[i].place) {
                    itemsResponse[i].place= places.value[j].title
                }
            }
        }
        items.value = itemsResponse
    });
}

const showPrevious = () => {
    if (start.value - noOfItems.value < 0) {
        start.value = 0
    } else {
        start.value -= noOfItems.value
    }
    reloadItems()
}

const showNext = () => {
    start.value += noOfItems.value
    reloadItems()
}

const showFirst = () => {
    start.value = 0
    reloadItems()
}

const showLast = () => {
    const remainder = itemsTotal.value % noOfItems.value
    if (remainder > 0) {
        start.value = itemsTotal.value - remainder
    } else {
        start.value = itemsTotal.value - noOfItems.value
    }
    reloadItems()
}

const addItemToCart = (itemID: number) => {
    if (cartStore.items.find(o => o.id === itemID)) {
        cartStore.items.find((o, i) => {
            if (o.id === itemID) {
                cartStore.items[i].number++
            }
            return
        })
    } else {
        cartStore.items.push({ id: itemID, number: 1 })
    }
}

getItems()
</script>
