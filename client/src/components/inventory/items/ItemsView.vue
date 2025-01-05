<template>
    <div class="p-6 sm:p-8 overflow-y-auto dark:text-white">
        <p>{{ item.id }}</p>
        <h1 v-if="item.title != ''" class="font-bold text-3xl">{{ item.title }}</h1>
        <div class="mt-4">
            <button v-if="item.borrowable" type="button" class="btn-table-primary mr-2 mb-2">
                <svg-icon type="mdi" :path="mdiCartPlus" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                {{ t('common.add') }}
            </button>
            <router-link :to="'/inventory/items/edit/' + item.id" type="button" class="btn-table-primary mr-2 mb-2">
                <svg-icon type="mdi" :path="mdiFileEdit" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                {{ t('common.edit') }}
            </router-link>
            <button type="button" class="btn-table-danger mb-2">
                <svg-icon type="mdi" :path="mdiFileRemove" class="-ml-0.5 h-5 w-5" aria-hidden="true"></svg-icon>
                {{ t('common.requestRetirement') }}
            </button>
        </div>
        <div class="grid md:grid-cols-[1fr_10rem] xl:grid-cols-[1fr_20rem] 2xl:grid-cols-[1fr_30rem] gap-6 mt-6">
            <div class="detailsList">
                <div v-if="item.description != ''" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('common.description') }}</div>
                    <div class="allow-line-breaks py-1.5">{{ item.description }}</div>
                </div>
                <div v-if="item.condition != ''" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('inventory.condition') }}</div>
                    <div class="allow-line-breaks py-1.5">{{ item.condition }}</div>
                </div>
                <div v-if="item.category != 0" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('common.category') }}</div>
                    <div class="py-1.5">{{ item.category }}</div>
                </div>
                <div v-if="item.place != 0" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('common.place') }}</div>
                    <div class="py-1.5">{{ item.place }}</div>
                </div>
                <div v-if="item.group != ''" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('common.group') }}</div>
                    <div class="py-1.5">{{ item.group }}</div>
                </div>
                <div v-if="item.number != 0" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('inventory.number') }}</div>
                    <div class="py-1.5">{{ item.number }}</div>
                </div>
                <div v-if="item.dateOfReceipt != ''" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('inventory.dateOfReceipt') }}</div>
                    <div class="py-1.5">{{ item.dateOfReceipt }}</div>
                </div>
                <div v-if="item.dateOfRetirement != ''" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('inventory.dateOfRetirement') }}</div>
                    <div class="py-1.5">{{ item.dateOfRetirement }}</div>
                </div>
                <div v-if="item.dateOfAccounting != ''" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('inventory.dateOfAccounting') }}</div>
                    <div class="py-1.5">{{ item.dateOfAccounting }}</div>
                </div>
                <div v-if="item.acquisitionCost != 0" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('inventory.acqusitionCost') }}</div>
                    <div class="py-1.5">{{ item.acquisitionCost }} €</div>
                </div>
                <div v-if="item.project[0] != null" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('inventory.project') }}</div>
                    <div class="py-1.5"><a v-for="(p, index) in item.project" :key="index" :href="'https://finanzen.stura-ilmenau.de/projekt/' + p" target="_blank" class="inline-block border border-black/10 dark:border-white/10 rounded-md mr-2 px-3 py-1.5 hover:bg-black/10 dark:hover:bg-white/10">{{ p }}</a></div>
                </div>
                <div v-if="item.receipt[0] != null" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">Abrechnung</div>
                    <div class="py-1.5"><a v-for="(r, index) in item.receipt" :key="index" :href="'https://finanzen.stura-ilmenau.de/auslagen/' + r" target="_blank" class="inline-block border border-black/10 dark:border-white/10 rounded-md mr-2 px-3 py-1.5 hover:bg-black/10 dark:hover:bg-white/10">{{ r }}</a></div>
                </div>
                <div v-if="item.supplier != ''" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('inventory.supplier') }}</div>
                    <div class="py-1.5">{{ item.supplier }}</div>
                </div>
                <div v-if="item.guaranteeUntil != ''" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('inventory.guaranteeUntil') }}</div>
                    <div class="py-1.5">{{ item.guaranteeUntil }}</div>
                </div>
                <div class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('inventory.borrowable') }}</div>
                    <div class="py-1.5">
                        <span v-if="item.borrowable"><svg-icon type="mdi" :path="mdiCheck" class="text-green-600" aria-hidden="true" :aria-label="t('inventory.borrowable')"></svg-icon></span>
                        <span v-else><svg-icon type="mdi" :path="mdiClose" class="text-red-600" aria-hidden="true" :aria-label="t('inventory.notBorrowable')"></svg-icon></span>
                    </div>
                </div>
                <div v-if="item.parentItem != 0" class="grid grid-cols-[10rem_1fr] border-t border-black/10 dark:border-white/10">
                    <div class="pl-1.5 pr-3 py-1.5 font-bold">{{ t('inventory.partOf') }}</div>
                    <div class="py-1.5"><a :href="'/inventory/items/view/' + item.parentItem">{{ item.parentItem }}</a></div>
                </div>
            </div>
            <div>
                <img v-if="item.image != ''" class="w-full rounded-md border border-black/10 dark:border-white/10" :src="'data:image/jpg;base64,' + item.image" />
            </div>
        </div>
    </div>
</template>
  
<script setup lang="ts">
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import SvgIcon from '@jamescoyle/vue-icon';
import {
    mdiCartPlus,
    mdiCheck,
    mdiClose,
    mdiFileEdit,
    mdiFileRemove
} from '@mdi/js';

const route = useRoute()
const { t } = useI18n()

const item: any = ref()
const groups: any = ref()
const categories: any = ref()
const places: any = ref()

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

                                for (let i = 0; i < groups.value.length; i++) {
                                    if (groups.value[i].id == itemResponse.group) {
                                        itemResponse.group = groups.value[i].title
                                    }
                                }
                                for (let i = 0; i < categories.value.length; i++) {
                                    if (categories.value[i].id == itemResponse.category) {
                                        itemResponse.category = categories.value[i].title
                                    }
                                }
                                for (let i = 0; i < places.value.length; i++) {
                                    if (places.value[i].id == itemResponse.place) {
                                        itemResponse.place= places.value[i].title
                                    }
                                }
                                item.value = itemResponse
                            })
                    })
            })
    })
}

getItem()
</script>

<style scoped>
.detailsList > div:last-of-type {
    border-bottom-width: 1px;
}

.allow-line-breaks {
  white-space: pre-wrap;
}
</style>
