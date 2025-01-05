<template>
    <div class="p-8 overflow-y-auto">
        <div class="sm:flex sm:items-center">
            <div class="sm:flex-auto">
                <h1 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">Gremien</h1>
                <p class="mt-2 text-sm text-gray-700 dark:text-gray-300">Eine Auflistung aller angelegten Gremien.</p>
            </div>
            <div class="mt-4 sm:ml-16 sm:mt-0 sm:flex-none">
                <router-link to="/inventory/groups/add" type="button" class="btn-primary">Gremium hinzufügen</router-link>
            </div>
        </div>
        <div class="mt-8 flow-root">
            <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
                <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
                    <div class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 sm:rounded-lg">
                        <table>
                            <thead>
                                <tr>
                                    <th scope="col">{{ t('common.name') }}</th>
                                    <th scope="col">{{ t('common.description') }}</th>
                                    <th scope="col">{{ t('common.emailAddress') }}</th>
                                    <th scope="col">{{ t('common.roles') }}</th>
                                    <th></th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="group in groups" :key="group.id">
                                    <td>{{ group.title }}</td>
                                    <td>{{ group.description }}</td>
                                    <td>{{ group.email }}</td>
                                    <td>
                                        <span v-for="(role, index) in group.roles" :key="index" class="inline-flex items-center rounded-md bg-blue-50 dark:bg-blue-400/10 px-2 py-1 text-xs text-blue-700 dark:text-blue-400 ring-1 ring-inset ring-blue-700/10 dark:ring-blue-400/30 mr-2">{{ role }}</span>
                                    </td>
                                    <td class="whitespace-nowrap text-right">
                                        <router-link :to="'/inventory/groups/edit/' + group.id" type="button" class="btn-table-primary">
                                            <svg-icon type="mdi" :path="mdiFileEdit" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                                            {{ t('common.edit') }}
                                        </router-link>
                                        <button @click="openDeleteDialog(group.title)" type="button" class="btn-table-danger ml-2">
                                            <svg-icon type="mdi" :path="mdiDelete" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                                            {{ t('common.delete') }}
                                        </button>
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <TransitionRoot as="template" :show="open">
        <Dialog as="div" class="relative z-50" @close="open = false">
            <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
                <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
            </TransitionChild>

            <div class="fixed inset-0 z-60 w-screen overflow-y-auto">
                <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
                    <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0 tranzinc-y-4 sm:tranzinc-y-0 sm:scale-95" enter-to="opacity-100 tranzinc-y-0 sm:scale-100" leave="ease-in duration-200" leave-from="opacity-100 tranzinc-y-0 sm:scale-100" leave-to="opacity-0 tranzinc-y-4 sm:tranzinc-y-0 sm:scale-95">
                        <DialogPanel class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
                            <div class="bg-white px-4 pb-4 pt-5 sm:p-6 sm:pb-4">
                                <div class="sm:flex sm:items-start">
                                    <div class="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10">
                                        <svg-icon type="mdi" :path="mdiAlertCircle" class="h-6 w-6 text-red-600" aria-hidden="true"></svg-icon>
                                    </div>
                                    <div class="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
                                        <DialogTitle as="h3" class="text-base font-semibold leading-6 text-gray-900">Ort löschen</DialogTitle>
                                        <div class="mt-2">
                                            <p class="text-sm text-gray-500">Bist du dir sicher, dass du den Ort <b>{{ currentGroupName }}</b> löschen möchtest? Dieser wird noch in <b>{{ currentGroupUsageCount }}</b> Inventar-Gegenständen verwendet.</p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6 border-t border-gray-300">
                                <button type="button" class="inline-flex w-full justify-center rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 sm:ml-3 sm:w-auto" @click="open = false">Löschen</button>
                                <button type="button" class="mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:mt-0 sm:w-auto" @click="open = false" ref="cancelButtonRef">Abbrechen</button>
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
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import {
    Dialog,
    DialogPanel,
    DialogTitle,
    TransitionChild,
    TransitionRoot
} from '@headlessui/vue'
import SvgIcon from '@jamescoyle/vue-icon';
import {
    mdiFileEdit,
    mdiDelete,
    mdiAlertCircle
} from '@mdi/js';

const { t } = useI18n()

const groups: any = ref()
const currentGroupName = ref()
const currentGroupUsageCount = ref()
const open: Ref<boolean> = ref(false)

const getPlaces = () => {
    const url = '/api/inventory/groups/index'
    axios
        .get(url)
        .then((response: any) => {
            groups.value = response.data
        })
}

const openDeleteDialog = (groupName: string) => {
    currentGroupName.value = groupName
    open.value = true
}
  
const deleteItem = () => {
    axios.post('/api/inventory/places/delete')
        .then(() => {
            getPlaces()
        })
}

getPlaces()
</script>
