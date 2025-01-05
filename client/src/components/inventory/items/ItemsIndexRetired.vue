<template>
    <div class="p-6 overflow-y-auto">
      <div class="sm:flex sm:items-center">
        <div class="sm:flex-auto">
          <h1 class="text-base font-semibold leading-6 text-zinc-800 dark:text-white">{{ t('inventory.inventoryTitle') }}</h1>
          <p class="mt-2 text-sm text-gray-700 dark:text-gray-300">{{ t('inventory.inventoryIntro') }}</p>
        </div>
        <div class="mt-4 sm:ml-16 sm:mt-0 flex whitespace-nowrap">
            <router-link to="/inventory/items" type="button" class="btn-neutral">
                <svg-icon type="mdi" :path="mdiFile" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                Vorhandene Gegenstände
            </router-link>
        </div>
      </div>
      <div class="mt-8 flow-root">
        <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
            <div class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 sm:rounded-lg">
              <table class="min-w-full divide-y divide-gray-300 dark:divide-zinc-900">
                <thead class="bg-zinc-100 dark:bg-zinc-600">
                  <tr>
                    <th scope="col" class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-zinc-800 dark:text-white sm:pl-6">{{ t('common.inventoryNumber') }}</th>
                    <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-zinc-800 dark:text-white">{{ t('common.title') }}</th>
                    <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-zinc-800 dark:text-white">{{ t('common.category') }}</th>
                    <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-zinc-800 dark:text-white">{{ t('common.place') }}</th>
                    <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-zinc-800 dark:text-white">{{ t('common.group') }}</th>
                    <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-zinc-800 dark:text-white">{{ t('common.borrowable') }}</th>
                    <th scope="col" class="relative py-3.5 pl-3 pr-4 sm:pr-6">
                      <span class="sr-only"></span>
                    </th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-zinc-300 dark:divide-zinc-800 bg-zinc-50 dark:bg-zinc-700">
                  <tr v-for="item in items" :key="item.id">
                    <td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-zinc-800 dark:text-white sm:pl-6">{{ item.id }}</td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-zinc-800 dark:text-white">{{ item.title }}</td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-zinc-800 dark:text-white">{{ item.category }}</td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-zinc-800 dark:text-white">{{ item.place }}</td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-zinc-800 dark:text-white">{{ item.group }}</td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-zinc-800 dark:text-white">
                      <span v-if="item.borrowable == true"><svg-icon type="mdi" :path="mdiCheck" class="text-green-600" aria-hidden="true" :aria-label="t('inventory.borrowable')"></svg-icon></span>
                      <span v-else><svg-icon type="mdi" :path="mdiClose" class="text-red-600" aria-hidden="true" :aria-label="t('inventory.notBorrowable')"></svg-icon></span>
                    </td>
                    <td class="relative whitespace-nowrap py-3 pl-3 pr-4 text-right text-sm font-medium sm:pr-4">
                      <router-link :to="'/inventory/items/view/' + item.id" type="button" class="btn-table-primary">
                        <svg-icon type="mdi" :path="mdiFileEye" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                        {{ t('common.view') }}
                      </router-link>
                      <router-link :to="'/inventory/items/edit/' + item.id" type="button" class="btn-table-primary ml-2">
                        <svg-icon type="mdi" :path="mdiFileEdit" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                        {{ t('common.edit') }}
                      </router-link>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
      <div class="flex flex-col sm:flex-row mt-6">
        <div class="flex">
          <div class="flex rounded-l-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-24">
            <input type="number" v-model="noOfItems" class="block w-24 flex-1 border-0 bg-transparent py-2 text-zinc-800 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
          </div>
          <div class="flex px-3 py-2 shadow-sm border-y border-gray-300 dark:text-white dark:border-white/10 w-full">
            {{ t('common.itemsPerPage') }}
          </div>
          <button type="button" class="rounded-r-md bg-white px-3 py-2 text-sm font-semibold text-zinc-800 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 dark:bg-white/10 dark:text-white dark:hover:bg-white/20 dark:ring-white/10" @click="reloadItems()"><svg-icon type="mdi" :path="mdiReload" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon></button>
        </div>
        <div class="ml-auto mt-4 sm:mt-0">
          <div>
            <button type="button" v-if="start - noOfItems > -noOfItems" class="inline-flex items-center gap-x-1.5 rounded-md bg-white px-3 py-2 text-sm font-semibold text-zinc-800 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 dark:bg-white/10 dark:text-white dark:hover:bg-white/20 dark:ring-white/10" @click="showPrevious()">
              <svg-icon type="mdi" :path="mdiChevronLeft" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
              {{ t('common.previous') }}
            </button>
            <button type="button" v-if="start + items.length < itemsTotal" class="ml-2 inline-flex items-center gap-x-1.5 rounded-md bg-white px-3 py-2 text-sm font-semibold text-zinc-800 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 dark:bg-white/10 dark:text-white dark:hover:bg-white/20 dark:ring-white/10" @click="showNext()">
              <svg-icon type="mdi" :path="mdiChevronRight" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
              {{ t('common.next') }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import SvgIcon from '@jamescoyle/vue-icon'
import {
    mdiCheck,
    mdiChevronLeft,
    mdiChevronRight,
    mdiClose,
    mdiFile,
    mdiFileEye,
    mdiFileEdit,
    mdiReload
} from '@mdi/js';

const { t } = useI18n()

const items: any = ref()
const groups: any = ref()
const categories: any = ref()
const places: any = ref()

const start = ref(0)
const noOfItems = ref(10)
const itemsTotal = ref(0)

const getItems = () => {
  const formData = new FormData()
    const url = '/api/inventory/items/index?display=' + noOfItems.value + '&start=' + start.value + '&retired=true'
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
                                })
                        })
                })
            })
}

const reloadItems = () => {
    const formData = new FormData()
    const url = '/api/inventory/items/index?display=' + noOfItems.value + '&start=' + start.value + '&retired=true'
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


getItems()
</script>
