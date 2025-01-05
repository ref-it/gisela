<template>
  <div class="p-6 sm:p-8 overflow-y-auto">
      <div class="space-y-12">
        <div class="border-b border-gray-900/10 dark:border-white/10 pb-12">
          <h1 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">Gegenstand bearbeiten</h1>

          <div class="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5">
            <div class="col-span-full">
              <label for="name" class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Name</label>
              <div class="mt-2">
                <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-full">
                  <input type="text" name="name" id="name" v-model="title" class="block flex-1 border-0 bg-transparent py-2 text-gray-900 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
                </div>
              </div>
            </div>

            <div class="col-span-full xl:col-span-2 2xl:col-span-3">
              <label for="description" class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Beschreibung</label>
              <div class="mt-2">
                <textarea id="description" name="description" v-model="description" class="block w-full rounded-md border-0 py-2 text-gray-900 dark:text-white dark:bg-white/5 shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-yellow-400 sm:text-sm sm:leading-6" />
              </div>
            </div>

            <div class="col-span-full xl:col-span-2 2xl:col-span-2">
              <label for="condition" class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Zustand</label>
              <div class="mt-2">
                <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-full">
                  <textarea id="condition" v-model="condition" class="block flex-1 border-0 bg-transparent py-2 text-gray-900 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
                </div>
              </div>
            </div>

            <Listbox as="div" v-model="group">
              <ListboxLabel class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Gremium</ListboxLabel>
              <div class="relative mt-2">
                <ListboxButton class="relative w-full cursor-default rounded-md bg-white dark:text-white dark:bg-white/5 dark:ring-white/10 py-2 pl-3 pr-10 text-left text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:outline-none focus:ring-2 focus:ring-yellow-400 sm:text-sm sm:leading-6">
                  <span class="block truncate">{{ group.title }}</span>
                  <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                    <ChevronUpDownIcon class="h-5 w-5 text-gray-400" aria-hidden="true" />
                  </span>
                </ListboxButton>

                <transition leave-active-class="transition ease-in duration-100" leave-from-class="opacity-100" leave-to-class="opacity-0">
                  <ListboxOptions class="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white dark:bg-zinc-800 py-1 text-base shadow-lg ring-1 ring-black dark:ring-white/10 ring-opacity-5 focus:outline-none sm:text-sm">
                    <ListboxOption as="template" v-for="item in groups" :key="item.id" :value="item" v-slot="{ active, group }">
                      <li :class="[active ? 'bg-yellow-400 text-zinc-800' : 'text-gray-900 dark:text-white', 'relative cursor-default select-none py-2 pl-3 pr-9']">
                        <span :class="[group ? 'font-semibold' : 'font-normal', 'block truncate']">{{ item.title }}</span>

                        <span v-if="group" :class="[active ? 'text-zinc-800' : 'text-yellow-400', 'absolute inset-y-0 right-0 flex items-center pr-4']">
                          <CheckIcon class="h-5 w-5" aria-hidden="true" />
                        </span>
                      </li>
                    </ListboxOption>
                  </ListboxOptions>
                </transition>
              </div>
            </Listbox>

            <Listbox as="div" v-model="category">
              <ListboxLabel class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Kategorie</ListboxLabel>
              <div class="relative mt-2">
                <ListboxButton class="relative w-full cursor-default rounded-md bg-white dark:text-white dark:bg-white/5 dark:ring-white/10 py-2 pl-3 pr-10 text-left text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:outline-none focus:ring-2 focus:ring-yellow-400 sm:text-sm sm:leading-6">
                  <span class="block truncate">{{ category.title }}</span>
                  <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                    <ChevronUpDownIcon class="h-5 w-5 text-gray-400" aria-hidden="true" />
                  </span>
                </ListboxButton>

                <transition leave-active-class="transition ease-in duration-100" leave-from-class="opacity-100" leave-to-class="opacity-0">
                  <ListboxOptions class="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white dark:bg-zinc-800 py-1 text-base shadow-lg ring-1 ring-black dark:ring-white/10 ring-opacity-5 focus:outline-none sm:text-sm">
                    <ListboxOption as="template" v-for="item in categories" :key="item.id" :value="item" v-slot="{ active, category }">
                      <li :class="[active ? 'bg-yellow-400 text-zinc-800' : 'text-gray-900 dark:text-white', 'relative cursor-default select-none py-2 pl-3 pr-9']">
                        <span :class="[category ? 'font-semibold' : 'font-normal', 'block truncate']">{{ item.title }}</span>

                        <span v-if="category" :class="[active ? 'text-zinc-800' : 'text-yellow-400', 'absolute inset-y-0 right-0 flex items-center pr-4']">
                          <CheckIcon class="h-5 w-5" aria-hidden="true" />
                        </span>
                      </li>
                    </ListboxOption>
                  </ListboxOptions>
                </transition>
              </div>
            </Listbox>

            <Listbox as="div" v-model="place">
              <ListboxLabel class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Lagerort</ListboxLabel>
              <div class="relative mt-2">
                <ListboxButton class="relative w-full cursor-default rounded-md bg-white dark:text-white dark:bg-white/5 dark:ring-white/10 py-2 pl-3 pr-10 text-left text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:outline-none focus:ring-2 focus:ring-yellow-400 sm:text-sm sm:leading-6">
                  <span class="block truncate">{{ place.title }}</span>
                  <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                    <ChevronUpDownIcon class="h-5 w-5 text-gray-400" aria-hidden="true" />
                  </span>
                </ListboxButton>

                <transition leave-active-class="transition ease-in duration-100" leave-from-class="opacity-100" leave-to-class="opacity-0">
                  <ListboxOptions class="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white dark:bg-zinc-800 py-1 text-base shadow-lg ring-1 ring-black dark:ring-white/10 ring-opacity-5 focus:outline-none sm:text-sm">
                    <ListboxOption as="template" v-for="item in places" :key="item.id" :value="item" v-slot="{ active, place }">
                      <li :class="[active ? 'bg-yellow-400 text-zinc-800' : 'text-gray-900 dark:text-white', 'relative cursor-default select-none py-2 pl-3 pr-9']">
                        <span :class="[place ? 'font-semibold' : 'font-normal', 'block truncate']">{{ item.title }}</span>

                        <span v-if="place" :class="[active ? 'text-zinc-800' : 'text-yellow-400', 'absolute inset-y-0 right-0 flex items-center pr-4']">
                          <CheckIcon class="h-5 w-5" aria-hidden="true" />
                        </span>
                      </li>
                    </ListboxOption>
                  </ListboxOptions>
                </transition>
              </div>
            </Listbox>

            <div class="">
              <label for="dateOfReceipt" class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Zugangsdatum</label>
              <div class="mt-2">
                <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-full">
                  <input type="date" id="dateOfReceipt" v-model="dateOfReceipt" class="block flex-1 border-0 bg-transparent py-2 text-gray-900 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
                </div>
              </div>
            </div>

            <div class="">
              <label for="dateOfRetirement" class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Abgangsdatum</label>
              <div class="mt-2">
                <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-full">
                  <input type="date" id="dateOfRetirement" v-model="dateOfRetirement" class="block flex-1 border-0 bg-transparent py-2 text-gray-900 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
                </div>
              </div>
            </div>

            <div class="">
              <label for="number" class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Anzahl</label>
              <div class="mt-2">
                <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-full">
                  <input type="number" id="number" v-model="number" class="block flex-1 border-0 bg-transparent py-2 text-gray-900 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
                </div>
              </div>
            </div>

            <div class="">
              <label for="partOf" class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Teil von</label>
              <div class="mt-2">
                <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-full">
                  <input type="number" id="partOf" v-model="partOf" class="block flex-1 border-0 bg-transparent py-2 text-gray-900 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
                </div>
              </div>
            </div>

            <div class="">
            <label for="supplier" class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Lieferant</label>
            <div class="mt-2">
                <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-full">
                <input type="text" id="supplier" v-model="supplier" class="block flex-1 border-0 bg-transparent py-2 text-gray-900 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
                </div>
            </div>
            </div>

            <div class="">
                <label for="guaranteeUntil" class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Garantie bis</label>
                <div class="mt-2">
                    <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-full">
                        <input type="date" id="guaranteeUntil" v-model="guaranteeUntil" class="block flex-1 border-0 bg-transparent py-2 text-gray-900 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
                    </div>
                </div>
            </div>

            <div class="">
              <label for="acquisitionCost" class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Anschaffungskosten</label>
              <div class="mt-2">
                <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-full">
                  <input type="number" id="acquisitionCost" v-model="acquisitionCost" class="block flex-1 border-0 bg-transparent py-2 text-gray-900 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
                </div>
              </div>
            </div>

            <div class="md:col-span-2">
                <label for="project">
                    <span class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Projekt <span class="text-sm font-normal text-gray-500 dark:text-gray-300">(mehrere durch Komma getrennt)</span></span>
                </label>
                <div class="mt-2">
                    <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-full">
                        <input type="text" id="project" v-model="project" class="block flex-1 border-0 bg-transparent py-2 text-gray-900 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
                    </div>
                </div>
            </div>

            <div class="md:col-span-2">
                <label for="receipt">
                    <span class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Abrechnung <span class="text-sm font-normal text-gray-500 dark:text-gray-300">(mehrere durch Komma getrennt)</span></span>
                </label>
                <div class="mt-2">
                    <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-full">
                        <input type="text" id="receipt" v-model="receipt" class="block flex-1 border-0 bg-transparent py-2 text-gray-900 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
                    </div>
                </div>
            </div>

            <div class="">
              <label for="dateOfAccounting" class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Buchungsdatum</label>
              <div class="mt-2">
                <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-full">
                  <input type="date" id="dateOfAccounting" v-model="dateOfAccounting" class="block flex-1 border-0 bg-transparent py-2 text-gray-900 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
                </div>
              </div>
            </div>

            <SwitchGroup as="div" class="flex items-center justify-between md:col-span-2">
                <span class="flex flex-grow flex-col">
                    <SwitchLabel as="span" class="text-sm font-bold leading-6 text-gray-900 dark:text-white" passive>Ausleihbar</SwitchLabel>
                    <SwitchDescription as="span" class="text-sm text-gray-500 dark:text-gray-300">Dieser Gegenstand kann gemäß <a href="https://ordnungen.stura.eu/ordnung/ausleiheordnung.html" target="_blank">Ausleiheordnung</a> ausgeliehen werden.</SwitchDescription>
                </span>
                <Switch v-model="borrowable" :class="[borrowable ? 'bg-yellow-400' : 'bg-gray-200', 'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-yellow-400 focus:ring-offset-2']">
              <span aria-hidden="true" :class="[borrowable ? 'translate-x-5' : 'translate-x-0', 'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white dark:bg-zinc-800 shadow ring-0 transition duration-200 ease-in-out']" />
            </Switch>
            </SwitchGroup>

            <div class="col-span-full">
                <label for="name" class="block text-sm font-bold leading-6 text-gray-900 dark:text-white">Bild</label>
                <div v-if="imagePreview == null" class="mt-2 flex">
                    <button class="inline-flex items-center gap-x-1.5 rounded-md bg-yellow-400 px-3 py-2 text-zinc-800 text-sm font-semibold shadow-sm hover:bg-yellow-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-yellow-600" @click="dialogCropImageOpen = true">
                        <svg-icon type="mdi" :path="mdiImagePlus" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                        Bild hinzufügen
                    </button>
                    <span v-if="imageAdded" class="ml-4 mt-2"><svg-icon type="mdi" :path="mdiCheck" class="text-green-600" aria-hidden="true"></svg-icon></span>
                </div>
                <div v-else class="mt-2">
                    <img class="w-60 rounded-t-md border-2 border-b-0 border-yellow-400" :src="'data:image/jpeg;base64,' + imagePreview" />
                    <button class="inline-flex w-60 items-center gap-x-1.5 rounded-b-md bg-yellow-400 px-3 py-2 text-zinc-800 text-sm font-semibold shadow-sm hover:bg-yellow-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-yellow-600" @click="removeImage">
                        <svg-icon type="mdi" :path="mdiImageRemove" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                        Bild entfernen
                    </button>
                </div>
            </div>

          </div>
        </div>
      </div>

      <div class="mt-6 flex items-center justify-end gap-x-6">
        <router-link to="/inventory/items" type="button" class="inline-flex items-center gap-x-1.5 rounded-md bg-blue-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600">
            <svg-icon type="mdi" :path="mdiCancel" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
            Abbrechen
        </router-link>
        <button @click="sendNewItemData" class="inline-flex items-center gap-x-1.5 rounded-md bg-green-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-green-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-green-600">
            <svg-icon type="mdi" :path="mdiContentSave" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
            Speichern
        </button>
        </div>
    </div>

  <TransitionRoot as="template" :show="dialogCropImageOpen">
    <Dialog as="div" class="relative z-50" @close="dialogCropImageOpen = false">
      <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
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
              <div class="bg-gray-50 dark:bg-zinc-700 px-3 py-3 sm:flex border-t border-gray-300 dark:border-zinc-900">
                <div class="mr-auto">
                    <label class="inline-flex w-full justify-center rounded-md bg-blue-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-blue-500 sm:w-auto" for="imageInput">Bild laden</label>
                </div>
                <div class="sm:flex sm:flex-row-reverse ">
                    <button type="button" class="inline-flex w-full justify-center rounded-md bg-green-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-green-500 sm:ml-3 sm:w-auto" @click="cropImage">Hinzufügen</button>
                    <button type="button" class="mt-3 inline-flex w-full justify-center rounded-md bg-white dark:bg-zinc-600 px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-zinc-600 hover:bg-gray-50 dark:hover:bg-zinc-500 dark:text-white sm:mt-0 sm:w-auto" @click="dialogCropImageOpen = false" ref="cancelButtonRef">Abbrechen</button>
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
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'
import SvgIcon from '@jamescoyle/vue-icon';
import {
    mdiCancel,
    mdiCheck,
    mdiContentSave,
    mdiImagePlus,
    mdiImageRemove
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
import { CheckIcon, ChevronUpDownIcon } from '@heroicons/vue/20/solid'

import VueCropper from 'vue-cropperjs';
import 'cropperjs/dist/cropper.css';

const groups: any = ref()
const categories: any = ref()
const places: any = ref()

const route = useRoute()
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
const borrowable = ref(false)
const imagePreview = ref()

const dialogCropImageOpen = ref(false)
const imageFile = ref()
const cropper = ref()
const imageCropped = ref()
const imageAdded = ref(false)

const getItem = () => {
    const url = '/api/inventory/items/data/' + route.params.id
    axios.get(url)
    .then((response: any) => {
        let itemResponse = response.data

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

                                title.value = itemResponse.title
                                description.value = itemResponse.description
                                condition.value = itemResponse.condition
                                guaranteeUntil.value = itemResponse.guaranteeUntil

                                if (itemResponse.number != 0) {
                                  number.value = itemResponse.number
                                }

                                if (itemResponse.parentItem != 0) {
                                  partOf.value = itemResponse.parentItem
                                }

                                for (let i = 0; i < groups.value.length; i++) {
                                  if (groups.value[i].id == itemResponse.group) {
                                    group.value = groups.value[i]
                                  }
                                }

                                for (let i = 0; i < categories.value.length; i++) {
                                  if (categories.value[i].id == itemResponse.category) {
                                    category.value = categories.value[i]
                                  }
                                }

                                for (let i = 0; i < places.value.length; i++) {
                                  if (places.value[i].id == itemResponse.place) {
                                    place.value = places.value[i]
                                  }
                                }

                                supplier.value = itemResponse.supplier
                                dateOfReceipt.value = itemResponse.dateOfReceipt
                                dateOfRetirement.value = itemResponse.dateOfRetirement
                                dateOfAccounting.value = itemResponse.dateOfAccounting

                                if (itemResponse.acquisitionCost != 0) {
                                  acquisitionCost.value = itemResponse.acquisitionCost
                                }

                                project.value = itemResponse.project
                                receipt.value = itemResponse.receipt
                                borrowable.value = itemResponse.borrowable

                                if (itemResponse.image != '') {
                                  imageCropped.value = itemResponse.image
                                  imagePreview.value = itemResponse.imagePreview
                                }
                            })
                    })
            })
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

    axios.post('/api/inventory/items/edit/' + route.params.id, formData)
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

const removeImage = () => {
    imageCropped.value = undefined
    imagePreview.value = undefined
}

getItem()
</script>

<style scoped>
.cropper-container {
    width: 100% !important;
    max-height: 40rem;
}
</style>