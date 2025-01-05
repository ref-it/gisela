<template>
    <div class="p-8 overflow-y-auto">
        <div class="space-y-12">
          <div class="border-b border-gray-900/10 dark:border-white/10 pb-12">
            <h1 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">Daten exportieren</h1>

          <div class="mt-8 flow-root grid grid-cols-1 xl:grid-cols-2 3xl:grid-cols-3 gap-6">

            <div class="mt-2 flex rounded-md shadow-sm">
              <div class="inline-flex check-input flex h-6 items-center border border-r-0 border-zinc-300 px-3 py-2 rounded-l-md dark:bg-white/5 dark:border-white/10">
                <input type="checkbox" v-model="filterDateOfReceipt" class="h-4 w-4 rounded border-zinc-300 dark:border-zinc-500 dark:bg-zinc-600 text-[var(--accent-color)] focus:ring-[var(--accent-color)]" />
              </div>
              <span class="w-48 inline-flex items-center border border-r-0 px-3 py-2 text-zinc-700 sm:text-sm dark:bg-white/5 dark:border-white/10 dark:text-gray-300">Zugangsdatum</span>
              <div class="inset-y-0 flex items-center">
                <select v-model="relationDateOfReceipt" class="h-full w-16 border-0 ring-1 ring-r-0 ring-inset ring-gray-300 bg-transparent py-2 pl-3 pr-7 text-zinc-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-[var(--accent-color)] sm:text-sm dark:bg-white/5 dark:ring-white/10 dark:text-gray-300">
                  <option>=</option>
                  <option>&le;</option>
                  <option>&lt;</option>
                  <option>&ge;</option>
                  <option>&gt;</option>
                  <option>!=</option>
                </select>
              </div>
              <input type="date" v-model="dateOfReceipt" class="block w-full min-w-0 flex-1 rounded-none rounded-r-md border-0 py-2 text-gray-900 ml-[-1px] ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-[var(--accent-color)] sm:text-sm sm:leading-6 dark:bg-white/5 dark:ring-white/10 dark:text-gray-300 !rounded-l-none" />
            </div>

            <div class="mt-2 flex rounded-md shadow-sm">
              <div class="inline-flex check-input flex h-6 items-center border border-r-0 border-zinc-300 px-3 py-2 rounded-l-md dark:bg-white/5 dark:border-white/10">
                <input type="checkbox" v-model="filterDateOfRetirement" class="h-4 w-4 rounded border-zinc-300 dark:border-zinc-500 dark:bg-zinc-600 text-[var(--accent-color)] focus:ring-[var(--accent-color)]" />
              </div>
              <span class="w-48 inline-flex items-center border border-r-0 px-3 py-2 text-zinc-700 sm:text-sm dark:bg-white/5 dark:border-white/10 dark:text-gray-300">Abgangsdatum</span>
              <div class="inset-y-0 flex items-center">
                <select v-model="relationDateOfRetirement" class="h-full w-16 border-0 ring-1 ring-r-0 ring-inset ring-gray-300 bg-transparent py-2 pl-3 pr-7 text-zinc-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-[var(--accent-color)] sm:text-sm dark:bg-white/5 dark:ring-white/10 dark:text-gray-300">
                  <option>=</option>
                  <option>&le;</option>
                  <option>&lt;</option>
                  <option>&ge;</option>
                  <option>&gt;</option>
                  <option>!=</option>
                </select>
              </div>
              <input type="date" v-model="dateOfRetirement" class="block w-full min-w-0 flex-1 rounded-none rounded-r-md border-0 py-2 text-gray-900 ml-[-1px] ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-[var(--accent-color)] sm:text-sm sm:leading-6 dark:bg-white/5 dark:ring-white/10 dark:text-gray-300 !rounded-l-none" />
            </div>

            <div class="mt-2 flex rounded-md shadow-sm">
              <div class="inline-flex check-input flex h-6 items-center border border-r-0 border-zinc-300 px-3 py-2 rounded-l-md dark:bg-white/5 dark:border-white/10">
                <input type="checkbox" v-model="filterAcquisitionCost" class="h-4 w-4 rounded border-zinc-300 dark:border-zinc-500 dark:bg-zinc-600 text-[var(--accent-color)] focus:ring-[var(--accent-color)]" />
              </div>
              <span class="w-48 inline-flex items-center border border-r-0 px-3 py-2 text-zinc-700 sm:text-sm dark:bg-white/5 dark:border-white/10 dark:text-gray-300">Anschaffungskosten</span>
              <div class="inset-y-0 flex items-center">
                <select v-model="relationAcqusitionCost" class="h-full w-16 border-0 ring-1 ring-r-0 ring-inset ring-gray-300 bg-transparent py-2 pl-3 pr-7 text-zinc-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-[var(--accent-color)] sm:text-sm dark:bg-white/5 dark:ring-white/10 dark:text-gray-300">
                  <option>=</option>
                  <option>&le;</option>
                  <option>&lt;</option>
                  <option>&ge;</option>
                  <option>&gt;</option>
                  <option>!=</option>
                </select>
              </div>
              <input type="number" v-model="acquisitionCost" class="block w-full min-w-0 flex-1 rounded-none rounded-r-md border-0 py-2 text-gray-900 ml-[-1px] ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-[var(--accent-color)] sm:text-sm sm:leading-6 dark:bg-white/5 dark:ring-white/10 dark:text-gray-300 !rounded-l-none" />
            </div>

          </div>

          <div class="mt-6 flow-root grid grid-cols-1 xl:grid-cols-2 2xl:grid-cols-3 gap-6">

            <div>
              <div class="mt-2 grid grid-cols-[auto_1fr] rounded-md shadow-sm">
                <div class="flex row-span-2 h-full border border-r-0 border-zinc-300 px-3 pt-3.5 rounded-l-md dark:bg-white/5 dark:border-white/10">
                  <input type="checkbox" v-model="filterCategories" class="h-4 w-4 rounded border-zinc-300 dark:border-zinc-500 dark:bg-zinc-600 text-[var(--accent-color)] focus:ring-[var(--accent-color)]" />
                </div>
                <span class="w-full h-6 table inline-flex border rounded-tr-md px-3 py-2 text-zinc-700 sm:text-sm dark:bg-white/5 dark:border-white/10 dark:text-gray-300">Kategorien</span>
                <div class="w-full h-full border border-t-0 rounded-br-md px-3 py-2 text-zinc-700 sm:text-sm dark:bg-white/5 dark:border-white/10 dark:text-gray-300 grid gap-2">
                  <div v-for="(item, index) in categories" :key="index" class="grid grid-cols-[2rem_1fr]">
                    <div>
                      <input type="checkbox" v-model="item.selected" class="h-4 w-4 rounded border-zinc-300 dark:border-zinc-500 dark:bg-zinc-600 text-[var(--accent-color)] focus:ring-[var(--accent-color)]" />
                    </div>
                    <div>{{ item.title }}</div>
                  </div>
                </div>
              </div>
            </div>

            <div>
              <div class="mt-2 grid grid-cols-[auto_1fr] rounded-md shadow-sm">
                <div class="flex row-span-2 h-full border border-r-0 border-zinc-300 px-3 pt-3.5 rounded-l-md dark:bg-white/5 dark:border-white/10">
                  <input type="checkbox" v-model="filterGroups" class="h-4 w-4 rounded border-zinc-300 dark:border-zinc-500 dark:border-zinc-500 dark:bg-zinc-600 text-[var(--accent-color)] focus:ring-[var(--accent-color)]" />
                </div>
                <span class="w-full h-6 table inline-flex border rounded-tr-md px-3 py-2 text-zinc-700 sm:text-sm dark:bg-white/5 dark:border-white/10 dark:text-gray-300">Gremien</span>
                <div class="w-full h-full border border-t-0 rounded-br-md px-3 py-2 text-zinc-700 sm:text-sm dark:bg-white/5 dark:border-white/10 dark:text-gray-300 grid gap-2">
                  <div v-for="(item, index) in groups" :key="index" class="grid grid-cols-[2rem_1fr]">
                    <div>
                      <input type="checkbox" v-model="item.selected" class="h-4 w-4 rounded border-zinc-300 dark:border-zinc-500 dark:bg-zinc-600 text-[var(--accent-color)] focus:ring-[var(--accent-color)]" />
                    </div>
                    <div>{{ item.title }}</div>
                  </div>
                </div>
              </div>
            </div>

            <div>
              <div class="mt-2 grid grid-cols-[auto_1fr] rounded-md shadow-sm">
                <div class="flex row-span-2 h-full border border-r-0 border-zinc-300 px-3 pt-3.5 rounded-l-md dark:bg-white/5 dark:border-white/10">
                  <input type="checkbox" v-model="filterPlaces" class="h-4 w-4 rounded border-zinc-300 dark:border-zinc-500 dark:bg-zinc-600 text-[var(--accent-color)] focus:ring-[var(--accent-color)]" />
                </div>
                <span class="w-full h-6 table inline-flex border rounded-tr-md px-3 py-2 text-zinc-700 sm:text-sm dark:bg-white/5 dark:border-white/10 dark:text-gray-300">Lagerorte</span>
                <div class="w-full h-full border border-t-0 rounded-br-md px-3 py-2 text-zinc-700 sm:text-sm dark:bg-white/5 dark:border-white/10 dark:text-gray-300 grid gap-2">
                  <div v-for="(item, index) in places" :key="index" class="grid grid-cols-[2rem_1fr]">
                    <div>
                      <input type="checkbox" v-model="item.selected" class="h-4 w-4 rounded border-zinc-300 dark:border-zinc-500 dark:bg-zinc-600 text-[var(--accent-color)] focus:ring-[var(--accent-color)]" />
                    </div>
                    <div>{{ item.title }}</div>
                  </div>
                </div>
              </div>
            </div>

          </div>

          <div class="mt-6 flow-root grid grid-cols-1 xl:grid-cols-2 2xl:grid-cols-3 gap-6">

            <div class="mt-2 flex rounded-md shadow-sm">
              <div class="inline-flex check-input flex h-6 items-center border border-r-0 border-zinc-300 px-3 py-2 rounded-l-md dark:bg-white/5 dark:border-white/10">
                <input type="checkbox" v-model="filterRetired" class="h-4 w-4 rounded border-zinc-300 dark:border-zinc-500 dark:bg-zinc-600 text-[var(--accent-color)] focus:ring-[var(--accent-color)]" />
              </div>
              <span class="inline-flex w-full items-center border border-r-0 px-3 py-2 text-zinc-700 sm:text-sm dark:bg-white/5 dark:border-white/10 dark:text-gray-300">Ausgemustert</span>
              <div class="inline-flex items-center border border-l-0 border-zinc-300 px-3 py-2 rounded-r-md dark:bg-white/5 dark:border-white/10">
                <Switch v-model="retired" :class="[retired ? 'bg-[var(--accent-color)]' : 'bg-zinc-200 dark:bg-zinc-600', 'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-[var(--accent-color)] focus:ring-offset-2']">
                  <span aria-hidden="true" :class="[retired ? 'translate-x-5' : 'translate-x-0', 'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out']" />
                </Switch>
              </div>
            </div>

            <div class="mt-2 flex rounded-md shadow-sm">
              <div class="inline-flex check-input flex h-6 items-center border border-r-0 border-zinc-300 px-3 py-2 rounded-l-md dark:bg-white/5 dark:border-white/10">
                <input type="checkbox" v-model="filterBorrowable" class="h-4 w-4 rounded border-zinc-300 dark:border-zinc-500 dark:bg-zinc-600 text-[var(--accent-color)] focus:ring-[var(--accent-color)]" />
              </div>
              <span class="inline-flex w-full items-center border border-r-0 px-3 py-2 text-zinc-700 sm:text-sm dark:bg-white/5 dark:border-white/10 dark:text-gray-300">Ausleihbar</span>
              <div class="inline-flex items-center border border-l-0 border-zinc-300 px-3 py-2 rounded-r-md dark:bg-white/5 dark:border-white/10">
                <Switch v-model="borrowable" :class="[borrowable ? 'bg-[var(--accent-color)]' : 'bg-zinc-200 dark:bg-zinc-600', 'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-[var(--accent-color)] focus:ring-offset-2']">
                  <span aria-hidden="true" :class="[borrowable ? 'translate-x-5' : 'translate-x-0', 'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out']" />
                </Switch>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="mt-6 flex items-center justify-end gap-x-6">
        <button @click="exportCSV" :disabled="loadingCSV" class="btn-primary">
          <svg-icon v-if="loadingCSV" type="mdi" :path="mdiTimerSand" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
          <svg-icon v-else type="mdi" :path="mdiCommaBox" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
          <span v-if="loadingCSV">Lade &hellip;</span>
          <span v-else>Als CSV exportieren</span>
        </button>
        <button @click="exportPDF" :disabled="loadingPDF" class="btn-primary">
          <svg-icon v-if="loadingPDF" type="mdi" :path="mdiTimerSand" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
          <svg-icon v-else type="mdi" :path="mdiFilePdfBox" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
          <span v-if="loadingPDF">Lade &hellip;</span>
          <span v-else>Als PDF exportieren</span>
        </button>
      </div>
    </div>
  </template>
  
<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import { Switch } from '@headlessui/vue'

import SvgIcon from '@jamescoyle/vue-icon';
import {
  mdiCommaBox,
  mdiFilePdfBox,
  mdiTimerSand
} from '@mdi/js'

  const date = new Date()
  const dd = String(date.getDate()).padStart(2, '0')
  const mm = String(date.getMonth() + 1).padStart(2, '0')
  const yyyy = date.getFullYear()
  const today = yyyy + '-' + mm + '-' + dd

  const filterDateOfReceipt = ref(false)
  const filterDateOfRetirement = ref(false)
  const filterAcquisitionCost = ref(false)
  const filterCategories = ref(false)
  const filterGroups = ref(false)
  const filterPlaces = ref(false)
  const filterRetired = ref(false)
  const filterBorrowable = ref(false)

  const relationDateOfReceipt = ref('≥')
  const relationDateOfRetirement = ref('≥')
  const relationAcqusitionCost = ref('≥')

  const dateOfReceipt = ref(today)
  const dateOfRetirement = ref(today)
  const acquisitionCost = ref(410)
  const retired = ref(false)
  const borrowable = ref(true)

  const categories: any = ref([])
  const groups:any = ref([])
  const places: any = ref([])

  const loadingCSV = ref(false)
  const loadingPDF = ref(false)

    const getData = () => {
        let url = '/api/inventory/categories/index'
        axios
            .get(url)
            .then((response: any) => {
              for (let i = 0; i < response.data.length; i++) {
                const item = { id: response.data[i].id, title: response.data[i].title, selected: false}
                categories.value.push(item)
              }
            })
        url = '/api/inventory/groups/index'
        axios
            .get(url)
            .then((response: any) => {for (let i = 0; i < response.data.length; i++) {
                const item = { id: response.data[i].id, title: response.data[i].title, selected: false}
                groups.value.push(item)
              }
            })
        url = '/api/inventory/places/index'
        axios
            .get(url)
            .then((response: any) => {
              for (let i = 0; i < response.data.length; i++) {
                const item = { id: response.data[i].id, title: response.data[i].title, selected: false}
                places.value.push(item)
              }
            })
    }

    const exportCSV = () => {
      loadingCSV.value = true
      let categoriesList: number[] = []
      for (let i = 0; i < categories.value.length; i++) {
        if (categories.value[i].selected) {
          categoriesList.push(categories.value[i].id)
        }
      }

      let groupsList: number[] = []
      for (let i = 0; i < groups.value.length; i++) {
        if (groups.value[i].selected) {
          groupsList.push(groups.value[i].id)
        }
      }

      let placesList: number[] = []
      for (let i = 0; i < places.value.length; i++) {
        if (places.value[i].selected) {
          placesList.push(places.value[i].id)
        }
      }

      const url = '/api/inventory/export/csv'
      axios.post(url,
        {
          filterDateOfReceipt: filterDateOfReceipt.value,
          filterDateOfRetirement: filterDateOfRetirement.value,
          filterAcquisitionCost: filterAcquisitionCost.value,
          filterCategories: filterCategories.value,
          filterGroups: filterGroups.value,
          filterPlaces: filterPlaces.value,
          filterRetired: filterRetired.value,
          filterBorrowable: filterBorrowable.value,

          relationDateOfReceipt: relationDateOfReceipt.value,
          relationDateOfRetirement: relationDateOfRetirement.value,
          relationAcqusitionCost: relationAcqusitionCost.value,

          dateOfReceipt: dateOfReceipt.value,
          dateOfRetirement: dateOfRetirement.value,
          acquisitionCost: acquisitionCost.value,
          categories: categoriesList,
          groups: groupsList,
          places: placesList,
          retired: retired.value,
          borrowable: borrowable.value
        },
        { 'responseType': 'blob' }
      )
      .then((response: any) => {
          var fileURL = window.URL.createObjectURL(new Blob([response.data], {type: 'text/csv'}))
          var fileLink = document.createElement('a')
          fileLink.href = fileURL
          fileLink.setAttribute('download', 'gisela-export.csv')
          document.body.appendChild(fileLink)
          fileLink.click()
          loadingCSV.value = false
      })
    }

    const exportPDF = () => {
      loadingPDF.value = true
      let categoriesList: number[] = []
      for (let i = 0; i < categories.value.length; i++) {
        if (categories.value[i].selected) {
          categoriesList.push(categories.value[i].id)
        }
      }

      let groupsList: number[] = []
      for (let i = 0; i < groups.value.length; i++) {
        if (groups.value[i].selected) {
          groupsList.push(groups.value[i].id)
        }
      }

      let placesList: number[] = []
      for (let i = 0; i < places.value.length; i++) {
        if (places.value[i].selected) {
          placesList.push(places.value[i].id)
        }
      }

      const url = '/api/inventory/export/pdf'
      axios.post(url,
        {
          filterDateOfReceipt: filterDateOfReceipt.value,
          filterDateOfRetirement: filterDateOfRetirement.value,
          filterAcquisitionCost: filterAcquisitionCost.value,
          filterCategories: filterCategories.value,
          filterGroups: filterGroups.value,
          filterPlaces: filterPlaces.value,
          filterRetired: filterRetired.value,
          filterBorrowable: filterBorrowable.value,

          relationDateOfReceipt: relationDateOfReceipt.value,
          relationDateOfRetirement: relationDateOfRetirement.value,
          relationAcqusitionCost: relationAcqusitionCost.value,

          dateOfReceipt: dateOfReceipt.value,
          dateOfRetirement: dateOfRetirement.value,
          acquisitionCost: acquisitionCost.value,
          categories: categoriesList,
          groups: groupsList,
          places: placesList,
          retired: retired.value,
          borrowable: borrowable.value
        },
        { 'responseType': 'blob' }
      )
      .then((response: any) => {
          var fileURL = window.URL.createObjectURL(new Blob([response.data], {type: 'application/pdf'}))
          var fileLink = document.createElement('a')
          fileLink.href = fileURL
          fileLink.setAttribute('download', 'gisela-export.pdf')
          document.body.appendChild(fileLink)
          fileLink.click()
          loadingPDF.value = false
      })
    }

  getData()
</script>

<style scoped>
.description {
  white-space: pre;
}

.check-input {
  padding-top: 1.25rem;
  padding-bottom: 1.25rem;
}
</style>