<template>
    <div class="p-8 overflow-y-auto">
        <div class="sm:flex sm:items-center mb-6">
            <div class="sm:flex-auto">
                <h1 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">Ausleihe-Anträge</h1>
                <p class="mt-2 text-sm text-gray-700 dark:text-gray-300">Eine Auflistung aller aktuellen Ausleihe-Anträge nach ihrem jeweiligen Bearbeitungsstatus.</p>
            </div>
            <div class="mt-4 sm:ml-16 sm:mt-0 sm:flex-none">
                <router-link to="/lending/application/new" type="button" class="btn-primary">Neue Anfrage</router-link>
                <button type="button" class="btn-neutral ml-2">Vergangene Ausleihen</button>
            </div>
        </div>
        <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4 overflow-x-auto">
            <div>
                <h2 class="text-sm font-bold text-gray-500 dark:text-gray-300">Offen</h2>
                <ul role="list" class="mt-3 grid grid-cols-1 gap-5">
                    <router-link to="" v-for="item in open" :key="item.name">
                        <li class="col-span-1 flex rounded-md shadow hover:bg-gray-100 bg-gray-50 transition ease-in-out duration-300 dark:bg-zinc-700 dark:hover:bg-zinc-600">
                            <div class="flex w-12 flex-shrink-0 items-top pt-3 justify-center rounded-l-md text-sm font-medium text-white bg-red-700 border border-black/5">
                                <svg-icon type="mdi" :path="icons.mdiAlertCircleOutline"></svg-icon>
                            </div>
                            <div class="flex flex-1 items-center justify-between truncate rounded-r-md dark:text-white border border-black/5">
                                <div class="flex-1 truncate px-4 py-2 text-sm">
                                    <p class="font-bold dark:text-white">{{ item.name }}</p>
                                    <p class="text-gray-500 dark:text-gray-300">[Ausleihe-Nr.]</p>
                                    <p class="text-gray-500 dark:text-gray-300">StuRa/Referat/FSR</p>
                                    <p class="text-gray-500 dark:text-gray-300">01.02.2024 &ndash; 31.03.2024</p>
                                </div>
                            </div>
                        </li>
                    </router-link>
                </ul>
            </div>

            <div>
                <h2 class="text-sm font-bold text-gray-500 dark:text-gray-300">In Bearbeitung</h2>
                <ul role="list" class="mt-3 grid grid-cols-1 gap-5">
                    <router-link to="" v-for="item in open" :key="item.name">
                        <li class="col-span-1 flex rounded-md shadow hover:bg-gray-100 bg-gray-50 transition ease-in-out duration-300 dark:bg-zinc-700 dark:hover:bg-zinc-600">
                            <div class="flex w-12 flex-shrink-0 items-top pt-3 justify-center rounded-l-md text-sm font-medium text-white bg-amber-600 border border-black/5">
                                <svg-icon type="mdi" :path="icons.mdiRecordCircleOutline"></svg-icon>
                            </div>
                            <div class="flex flex-1 items-center justify-between truncate rounded-r-md dark:text-white border border-black/5">
                                <div class="flex-1 truncate px-4 py-2 text-sm">
                                    <p class="font-bold dark:text-white">{{ item.name }}</p>
                                    <p class="text-gray-500 dark:text-gray-300">[Ausleihe-Nr.]</p>
                                    <p class="text-gray-500 dark:text-gray-300">StuRa/Referat/FSR</p>
                                    <p class="text-gray-500 dark:text-gray-300">01.02.2024 &ndash; 31.03.2024</p>
                                </div>
                            </div>
                        </li>
                    </router-link>
                </ul>
            </div>

            <div>
                <h2 class="text-sm font-bold text-gray-500 dark:text-gray-300">Bereit</h2>
                <ul role="list" class="mt-3 grid grid-cols-1 gap-5">
                    <router-link to="" v-for="item in open" :key="item.name">
                        <li class="col-span-1 flex rounded-md shadow hover:bg-gray-100 bg-gray-50 transition ease-in-out duration-300 dark:bg-zinc-700 dark:hover:bg-zinc-600">
                            <div class="flex w-12 flex-shrink-0 items-top pt-3 justify-center rounded-l-md text-sm font-medium text-white bg-green-600 border border-black/5">
                                <svg-icon type="mdi" :path="icons.mdiCheckCircleOutline"></svg-icon>
                            </div>
                            <div class="flex flex-1 items-center justify-between truncate rounded-r-md dark:text-white border border-black/5">
                                <div class="flex-1 truncate px-4 py-2 text-sm">
                                    <p class="font-bold dark:text-white">{{ item.name }}</p>
                                    <p class="text-gray-500 dark:text-gray-300">[Ausleihe-Nr.]</p>
                                    <p class="text-gray-500 dark:text-gray-300">StuRa/Referat/FSR</p>
                                    <p class="text-gray-500 dark:text-gray-300">01.02.2024 &ndash; 31.03.2024</p>
                                </div>
                            </div>
                        </li>
                    </router-link>
                </ul>
            </div>

            <div>
                <h2 class="text-sm font-bold text-gray-500 dark:text-gray-300">Ausgeliehen</h2>
                <ul role="list" class="mt-3 grid grid-cols-1 gap-5">
                    <router-link to="" v-for="item in open" :key="item.name">
                        <li class="col-span-1 flex rounded-md shadow hover:bg-gray-100 bg-gray-50 transition ease-in-out duration-300 dark:bg-zinc-700 dark:hover:bg-zinc-600">
                            <div class="flex w-12 flex-shrink-0 items-top pt-3 justify-center rounded-l-md text-sm font-medium text-white bg-sky-600 border border-black/5">
                                <svg-icon type="mdi" :path="icons.mdiArrowRightDropCircleOutline"></svg-icon>
                            </div>
                            <div class="flex flex-1 items-center justify-between truncate rounded-r-md dark:text-white border border-black/5">
                                <div class="flex-1 truncate px-4 py-2 text-sm">
                                    <p class="font-bold dark:text-white">{{ item.name }}</p>
                                    <p class="text-gray-500 dark:text-gray-300">[Ausleihe-Nr.]</p>
                                    <p class="text-gray-500 dark:text-gray-300">StuRa/Referat/FSR</p>
                                    <p class="text-gray-500 dark:text-gray-300">01.02.2024 &ndash; 31.03.2024</p>
                                </div>
                            </div>
                        </li>
                    </router-link>
                </ul>
            </div>
        </div>
    </div>
</template>
  
<script setup lang="ts">
import SvgIcon from '@jamescoyle/vue-icon';
import {
  mdiAlertCircleOutline,
  mdiArrowRightDropCircleOutline,
  mdiCheckCircleOutline,
  mdiMenu,
  mdiRecordCircleOutline
} from '@mdi/js';
import { EllipsisVerticalIcon } from '@heroicons/vue/20/solid'

const icons = {
  mdiAlertCircleOutline: mdiAlertCircleOutline,
  mdiRecordCircleOutline: mdiRecordCircleOutline,
  mdiCheckCircleOutline: mdiCheckCircleOutline,
  mdiArrowRightDropCircleOutline: mdiArrowRightDropCircleOutline
}
  
  const open = [
    { name: 'Graph API', initials: 'GA', href: '#', members: 16, bgColor: 'bg-pink-600' },
    { name: 'Component Design', initials: 'CD', href: '#', members: 12, bgColor: 'bg-purple-600' },
    { name: 'Templates', initials: 'T', href: '#', members: 16, bgColor: 'bg-yellow-500' },
    { name: 'React Components', initials: 'RC', href: '#', members: 8, bgColor: 'bg-green-500' },
  ]

  const inProcessing = [
    { name: 'Graph API', initials: 'GA', href: '#', members: 16, bgColor: 'bg-pink-600' },
    { name: 'Component Design', initials: 'CD', href: '#', members: 12, bgColor: 'bg-purple-600' },
    { name: 'Templates', initials: 'T', href: '#', members: 16, bgColor: 'bg-yellow-500' },
    { name: 'React Components', initials: 'RC', href: '#', members: 8, bgColor: 'bg-green-500' },
  ]

  const ready = [
    { name: 'Graph API', initials: 'GA', href: '#', members: 16, bgColor: 'bg-pink-600' },
    { name: 'Component Design', initials: 'CD', href: '#', members: 12, bgColor: 'bg-purple-600' },
    { name: 'Templates', initials: 'T', href: '#', members: 16, bgColor: 'bg-yellow-500' },
    { name: 'React Components', initials: 'RC', href: '#', members: 8, bgColor: 'bg-green-500' },
  ]

  const lend = [
    { name: 'Graph API', initials: 'GA', href: '#', members: 16, bgColor: 'bg-pink-600' },
    { name: 'Component Design', initials: 'CD', href: '#', members: 12, bgColor: 'bg-purple-600' },
    { name: 'Templates', initials: 'T', href: '#', members: 16, bgColor: 'bg-yellow-500' },
    { name: 'React Components', initials: 'RC', href: '#', members: 8, bgColor: 'bg-green-500' },
  ]
  </script>