<template>
  <div class="p-8 overflow-y-auto">
      <div class="space-y-12">
        <div class="border-b border-zinc-900/10 dark:border-white/10 pb-12">
          <h1 class="text-base font-semibold leading-6 text-zinc-900 dark:text-white">Gegenstand hinzufügen</h1>

          <div class="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5">
            <div class="col-span-full">
              <label for="name">Name</label>
              <div class="mt-2">
                <input type="text" id="name" v-model="title" />
              </div>
            </div>

            <div class="col-span-full xl:col-span-2 2xl:col-span-3">
              <label for="description">Beschreibung</label>
              <div class="mt-2">
                <textarea id="description" v-model="description" />
              </div>
            </div>

            <div class="col-span-full xl:col-span-2 2xl:col-span-2">
              <label for="condition">Zustand</label>
              <div class="mt-2">
                <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-zinc-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-[var(--accent-color)] w-full">
                  <textarea id="condition" v-model="condition" />
                </div>
              </div>
            </div>

            <Listbox as="div" v-model="group">
              <ListboxLabel>Gremium</ListboxLabel>
              <div class="relative mt-2">
                <ListboxButton class="relative w-full cursor-default rounded-md bg-white dark:text-white dark:bg-white/5 dark:ring-white/10 py-2 pl-3 pr-10 text-left text-zinc-900 shadow-sm ring-1 ring-inset ring-zinc-300 focus:outline-none focus:ring-2 focus:ring-[var(--accent-color)] sm:text-sm sm:leading-6">
                  <span class="block truncate">{{ group.title }}</span>
                  <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                    <svg-icon type="mdi" :path="mdiChevronDown" class="h-5 w-5 text-zinc-400" aria-hidden="true"></svg-icon>
                  </span>
                </ListboxButton>

                <transition leave-active-class="transition ease-in duration-100" leave-from-class="opacity-100" leave-to-class="opacity-0">
                  <ListboxOptions class="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white dark:bg-zinc-800 py-1 text-base shadow-lg ring-1 ring-black dark:ring-white/10 ring-opacity-5 focus:outline-none sm:text-sm">
                    <ListboxOption as="template" v-for="item in groups" :key="item.id" :value="item" v-slot="{ active, group }">
                      <li :class="[active ? 'bg-[var(--accent-color)] text-zinc-800' : 'text-zinc-900 dark:text-white', 'relative cursor-default select-none py-2 pl-3 pr-9']">
                        <span :class="[group ? 'font-semibold' : 'font-normal', 'block truncate']">{{ item.title }}</span>

                        <span v-if="group" :class="[active ? 'text-zinc-800' : 'text-[var(--accent-color)]', 'absolute inset-y-0 right-0 flex items-center pr-4']">
                          <CheckIcon class="h-5 w-5" aria-hidden="true" />
                        </span>
                      </li>
                    </ListboxOption>
                  </ListboxOptions>
                </transition>
              </div>
            </Listbox>

            <Listbox as="div" v-model="category">
              <ListboxLabel>Kategorie</ListboxLabel>
              <div class="relative mt-2">
                <ListboxButton class="relative w-full cursor-default rounded-md bg-white dark:text-white dark:bg-white/5 dark:ring-white/10 py-2 pl-3 pr-10 text-left text-zinc-900 shadow-sm ring-1 ring-inset ring-zinc-300 focus:outline-none focus:ring-2 focus:ring-[var(--accent-color)] sm:text-sm sm:leading-6">
                  <span class="block truncate">{{ category.title }}</span>
                  <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                    <svg-icon type="mdi" :path="mdiChevronDown" class="h-5 w-5 text-zinc-400" aria-hidden="true"></svg-icon>
                  </span>
                </ListboxButton>

                <transition leave-active-class="transition ease-in duration-100" leave-from-class="opacity-100" leave-to-class="opacity-0">
                  <ListboxOptions class="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white dark:bg-zinc-800 py-1 text-base shadow-lg ring-1 ring-black dark:ring-white/10 ring-opacity-5 focus:outline-none sm:text-sm">
                    <ListboxOption as="template" v-for="item in categories" :key="item.id" :value="item" v-slot="{ active, category }">
                      <li :class="[active ? 'bg-[var(--accent-color)] text-zinc-800' : 'text-zinc-900 dark:text-white', 'relative cursor-default select-none py-2 pl-3 pr-9']">
                        <span :class="[category ? 'font-semibold' : 'font-normal', 'block truncate']">{{ item.title }}</span>

                        <span v-if="category" :class="[active ? 'text-zinc-800' : 'text-[var(--accent-color)]', 'absolute inset-y-0 right-0 flex items-center pr-4']">
                          <CheckIcon class="h-5 w-5" aria-hidden="true" />
                        </span>
                      </li>
                    </ListboxOption>
                  </ListboxOptions>
                </transition>
              </div>
            </Listbox>

            <Listbox as="div" v-model="place">
              <ListboxLabel>Lagerort</ListboxLabel>
              <div class="relative mt-2">
                <ListboxButton class="relative w-full cursor-default rounded-md bg-white dark:text-white dark:bg-white/5 dark:ring-white/10 py-2 pl-3 pr-10 text-left text-zinc-900 shadow-sm ring-1 ring-inset ring-zinc-300 focus:outline-none focus:ring-2 focus:ring-[var(--accent-color)] sm:text-sm sm:leading-6">
                  <span class="block truncate">{{ place.title }}</span>
                  <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                    <svg-icon type="mdi" :path="mdiChevronDown" class="h-5 w-5 text-zinc-400" aria-hidden="true"></svg-icon>
                  </span>
                </ListboxButton>

                <transition leave-active-class="transition ease-in duration-100" leave-from-class="opacity-100" leave-to-class="opacity-0">
                  <ListboxOptions class="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white dark:bg-zinc-800 py-1 text-base shadow-lg ring-1 ring-black dark:ring-white/10 ring-opacity-5 focus:outline-none sm:text-sm">
                    <ListboxOption as="template" v-for="item in places" :key="item.id" :value="item" v-slot="{ active, place }">
                      <li :class="[active ? 'bg-[var(--accent-color)] text-zinc-800' : 'text-zinc-900 dark:text-white', 'relative cursor-default select-none py-2 pl-3 pr-9']">
                        <span :class="[place ? 'font-semibold' : 'font-normal', 'block truncate']">{{ item.title }}</span>

                        <span v-if="place" :class="[active ? 'text-zinc-800' : 'text-[var(--accent-color)]', 'absolute inset-y-0 right-0 flex items-center pr-4']">
                          <CheckIcon class="h-5 w-5" aria-hidden="true" />
                        </span>
                      </li>
                    </ListboxOption>
                  </ListboxOptions>
                </transition>
              </div>
            </Listbox>

            <div class="">
              <label for="dateOfReceipt">Zugangsdatum</label>
              <div class="mt-2">
                <input type="date" id="dateOfReceipt" v-model="dateOfReceipt" />
              </div>
            </div>

            <div class="">
              <label for="dateOfRetirement">Abgangsdatum</label>
              <div class="mt-2">
                <input type="date" id="dateOfRetirement" v-model="dateOfRetirement" />
              </div>
            </div>

            <div class="">
              <label for="number">Anzahl</label>
              <div class="mt-2">
                <input type="number" id="number" v-model="number" />
              </div>
            </div>

            <div class="">
              <label for="partOf">Teil von</label>
              <div class="mt-2">
                  <input type="number" id="partOf" v-model="partOf" />
              </div>
            </div>

            <div class="">
              <label for="supplier">Lieferant</label>
              <div class="mt-2">
                  <input type="text" id="supplier" v-model="supplier" />
              </div>
            </div>

            <div class="">
                <label for="guaranteeUntil">Garantie bis</label>
                <div class="mt-2">
                  <input type="date" id="guaranteeUntil" v-model="guaranteeUntil" />
                </div>
            </div>

            <div class="">
              <label for="acquisitionCost">Anschaffungskosten</label>
              <div class="mt-2">
                <input type="number" id="acquisitionCost" v-model="acquisitionCost" />
              </div>
            </div>

            <div class="md:col-span-2">
                <label for="project">
                    Projekt <span class="text-sm font-normal text-zinc-500 dark:text-zinc-300">(mehrere durch Komma getrennt)</span>
                </label>
                <div class="mt-2">
                    <input type="text" id="project" v-model="project" />
                </div>
            </div>

            <div class="md:col-span-2">
                <label for="receipt">
                    Abrechnung <span class="text-sm font-normal text-zinc-500 dark:text-zinc-300">(mehrere durch Komma getrennt)</span>
                </label>
                <div class="mt-2">
                    <input type="text" id="receipt" v-model="receipt" />
                </div>
            </div>

            <div class="">
              <label for="dateOfAccounting">Buchungsdatum</label>
              <div class="mt-2">
                <input type="date" id="dateOfAccounting" v-model="dateOfAccounting" />
              </div>
            </div>

            <SwitchGroup as="div" class="flex items-center justify-between md:col-span-2">
                <span class="flex flex-grow flex-col">
                    <SwitchLabel as="span" class="text-sm font-bold leading-6 text-zinc-900 dark:text-white" passive>Ausleihbar</SwitchLabel>
                    <SwitchDescription as="span" class="text-sm text-zinc-500 dark:text-zinc-300">Dieser Gegenstand kann gemäß <a href="https://ordnungen.stura.eu/ordnung/ausleiheordnung.html" target="_blank">Ausleiheordnung</a> ausgeliehen werden.</SwitchDescription>
                </span>
                <Switch v-model="borrowable" :class="[borrowable ? 'bg-[var(--accent-color)]' : 'bg-zinc-200 dark:bg-white/5', 'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out ring-zinc-300 dark:ring-white/10 focus:outline-none focus:ring-2 focus:ring-[var(--accent-color)] focus:ring-offset-2 focus:ring-offset-white dark:focus:ring-offset-zinc-800']">
              <span aria-hidden="true" :class="[borrowable ? 'translate-x-5' : 'translate-x-0', 'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-zinc-400 shadow ring-0 transition duration-200 ease-in-out']" />
            </Switch>
            </SwitchGroup>

            <div class="col-span-full">
            <label for="name">Bild</label>
            <div class="mt-2 flex">
                <button class="btn-primary" @click="dialogCropImageOpen = true">
                    <svg-icon type="mdi" :path="mdiImagePlus" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                    Bild hinzufügen
                </button>
                <span v-if="imageAdded" class="ml-4 mt-2"><svg-icon type="mdi" :path="mdiCheck" class="text-green-600" aria-hidden="true"></svg-icon></span>
                </div>
            </div>

          </div>
        </div>
      </div>

      <div class="mt-6 flex items-center justify-end gap-x-6">
        <router-link to="/inventory/items" type="button" class="btn-neutral">
            <svg-icon type="mdi" :path="mdiCancel" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
            Abbrechen
        </router-link>
        <button @click="sendNewItemData" class="btn-primary">
            <svg-icon type="mdi" :path="mdiPlus" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
            Hinzufügen
        </button>
        </div>
    </div>

  <TransitionRoot as="template" :show="dialogCropImageOpen">
    <Dialog as="div" class="relative z-50" @close="dialogCropImageOpen = false">
      <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
        <div class="fixed inset-0 bg-zinc-500 bg-opacity-75 transition-opacity" />
      </TransitionChild>

      <div class="fixed inset-0 z-60 w-screen overflow-y-auto">
        <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
          <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0 tranzinc-y-4 sm:tranzinc-y-0 sm:scale-95" enter-to="opacity-100 tranzinc-y-0 sm:scale-100" leave="ease-in duration-200" leave-from="opacity-100 tranzinc-y-0 sm:scale-100" leave-to="opacity-0 tranzinc-y-4 sm:tranzinc-y-0 sm:scale-95">
            <DialogPanel class="relative transform overflow-hidden rounded-lg bg-white dark:bg-zinc-800 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-xl">
              <div class="bg-white dark:bg-zinc-800 px-4 pb-4 pt-5 sm:p-6 sm:pb-4 dark:text-white">
                <div class="sm:flex sm:items-start">
                  <div class="mt-3 text-center sm:mt-0 sm:text-left">
                    <input type="file" id="imageInput" class="hidden" accept="image/*" @change="readImageFile($event)">
                    <vue-cropper
                      ref="cropper"
                      :src="imageFile"
                      :aspect-ratio="1 / 1"
                      alt="Wähle ein Bild aus."
                    >
                    </vue-cropper>
                  </div>
                </div>
              </div>
              <div class="bg-zinc-50 dark:bg-zinc-700 px-3 py-3 sm:flex border-t border-zinc-300 dark:border-zinc-900">
                <div class="mr-auto">
                    <label class="inline-flex w-full justify-center rounded-md bg-blue-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-blue-500 sm:w-auto" for="imageInput">Bild laden</label>
                </div>
                <div class="sm:flex sm:flex-row-reverse ">
                    <button type="button" class="inline-flex w-full justify-center rounded-md bg-green-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-green-500 sm:ml-3 sm:w-auto" @click="cropImage">Hinzufügen</button>
                    <button type="button" class="mt-3 inline-flex w-full justify-center rounded-md bg-white dark:bg-zinc-600 px-3 py-2 text-sm font-semibold text-zinc-900 shadow-sm ring-1 ring-inset ring-zinc-300 dark:ring-zinc-600 hover:bg-zinc-50 dark:hover:bg-zinc-500 dark:text-white sm:mt-0 sm:w-auto" @click="dialogCropImageOpen = false" ref="cancelButtonRef">Abbrechen</button>
                </div>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>

<script setup lang="ts">
import { type Ref, ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import SvgIcon from '@jamescoyle/vue-icon';
import {
    mdiCancel,
    mdiCheck,
    mdiChevronDown,
    mdiImagePlus,
    mdiPlus
} from '@mdi/js';
import {
    Dialog,
    DialogPanel,
    Listbox,
    ListboxButton,
    ListboxLabel,
    ListboxOption,
    ListboxOptions,
    Switch,
    SwitchDescription,
    SwitchGroup,
    SwitchLabel,
    TransitionChild,
    TransitionRoot
} from '@headlessui/vue'
import { CheckIcon } from '@heroicons/vue/20/solid'

import VueCropper from 'vue-cropperjs';
import 'cropperjs/dist/cropper.css';

const groups: any = ref()
const categories: any = ref()
const places: any = ref()

const router = useRouter()

const title = ref()
const description = ref()
const condition = ref()
const guaranteeUntil = ref()
const number = ref()
const partOf = ref()
const group = ref()
const category = ref()
const place = ref()
const supplier = ref()
const dateOfReceipt = ref()
const dateOfRetirement = ref()
const dateOfAccounting = ref()
const acquisitionCost = ref()
const project = ref()
const receipt = ref()
const borrowable: Ref<boolean> = ref(false)

const dialogCropImageOpen: Ref<boolean> = ref(false)
const imageFile = ref()
const cropper = ref()
const imageCropped = ref()
const imageAdded: Ref<boolean> = ref(false)

const getData = () => {
    let url = '/api/inventory/groups/index'
    axios
        .get(url)
        .then((response: any) => {
            groups.value = response.data
            group.value = groups.value[0]
        })
    
    url = '/api/inventory/categories/index'
    axios
        .get(url)
        .then((response: any) => {
            categories.value = response.data
            category.value = categories.value[0]
        })

    url = '/api/inventory/places/index'
    axios
        .get(url)
        .then((response: any) => {
            places.value = response.data
            place.value = places.value[0]
        })
}

const sendNewItemData = () => {
    const formData = new FormData()
    formData.append("title", title.value)
    formData.append("description", description.value)
    formData.append("condition", condition.value)
    formData.append("category", category.value.id)
    formData.append("place", place.value.id)
    formData.append("group", group.value.id)
    formData.append("number", number.value)
    formData.append("image", imageCropped.value)
    formData.append("dateOfReceipt", dateOfReceipt.value)
    formData.append("dateOfRetirement", dateOfRetirement.value)
    formData.append("dateOfAccounting", dateOfAccounting.value)
    formData.append("acquisitionCost", acquisitionCost.value)
    formData.append("project", project.value)
    formData.append("receipt", receipt.value)
    formData.append("supplier", supplier.value)
    formData.append("guaranteeUntil", guaranteeUntil.value)
    formData.append("borrowable", borrowable.value.toString())
    formData.append("parentItem", partOf.value)

    axios.post('/api/inventory/items/add', formData)
    .then(() => {
        router.push('/inventory/items');
    })
    .catch(() => {
        
    })
}

const readImageFile = (event: Event) => {
    console.log(1)
    const target = event.target as HTMLInputElement
    let file = null

    if (target && target.files) {
        file = target.files[0]
        console.log(2)

        if (file.type.indexOf('image/') === -1) {
            alert('Please select an image file')
            return
        }

        if (typeof FileReader === 'function') {
            const reader = new FileReader()

            reader.onload = (event) => {
                imageFile.value = event.target?.result
                // rebuild cropperjs with the updated source
                cropper.value.replace(event.target?.result)
            };

            reader.readAsDataURL(file);
        } else {
            alert('Sorry, FileReader API not supported')
        }
    }
}

const cropImage = () => {
    dialogCropImageOpen.value = false
    imageCropped.value = cropper.value.getCroppedCanvas().toDataURL()
    imageAdded.value = true
}

getData()
</script>

<style scoped>
.cropper-container {
    width: 100% !important;
    max-height: 40rem;
}
</style>