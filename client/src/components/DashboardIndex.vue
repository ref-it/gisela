<template>
    <main class="p-8">
        <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
            <div class="relative p-6 focus-within:ring-2 focus-within:ring-inset focus-within:ring-[var(--accent-color)] border border-black/5 transition ease-in-out duration-300 rounded-lg shadow hover:bg-gray-100 bg-gray-50 dark:bg-zinc-700 dark:hover:bg-zinc-600">
                <div>
                    <span :class="['bg-[var(--accent-color)] text-[var(--accent-color-text)] inline-flex rounded-lg p-3 border border-black/10']">
                        <svg-icon type="mdi" :path="mdiListBox" class="h-6 w-6" aria-hidden="true"></svg-icon>
                    </span>
                </div>
                <div class="mt-8">
                    <h3 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">
                        <router-link to="/inventory/items" class="focus:outline-none">
                            <span class="absolute inset-0" aria-hidden="true"></span>
                            {{ t('dashboard.inventoryRegisterHeading') }}
                        </router-link>
                    </h3>
                    <p class="mt-2 text-sm text-gray-500 dark:text-gray-300">{{ t('dashboard.inventoryRegisterDescription') }}</p>
                </div>
                <span class="pointer-events-none absolute right-6 top-6 text-gray-300 group-hover:text-gray-400" aria-hidden="true">
                    <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M20 4h1a1 1 0 00-1-1v1zm-1 12a1 1 0 102 0h-2zM8 3a1 1 0 000 2V3zM3.293 19.293a1 1 0 101.414 1.414l-1.414-1.414zM19 4v12h2V4h-2zm1-1H8v2h12V3zm-.707.293l-16 16 1.414 1.414 16-16-1.414-1.414z" />
                    </svg>
                </span>
            </div>

            <div class="relative p-6 focus-within:ring-2 focus-within:ring-inset focus-within:ring-[var(--accent-color)] border border-black/5 transition ease-in-out duration-300 rounded-lg shadow hover:bg-gray-100 bg-gray-50 dark:bg-zinc-700 dark:hover:bg-zinc-600">
                <div>
                    <span :class="['bg-[var(--accent-color)] text-[var(--accent-color-text)] inline-flex rounded-lg p-3 border border-black/10']">
                        <svg-icon type="mdi" :path="mdiFolder" class="h-6 w-6" aria-hidden="true"></svg-icon>
                    </span>
                </div>
                <div class="mt-8">
                    <h3 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">
                        <router-link to="/lending/application/my" class="focus:outline-none">
                            <span class="absolute inset-0" aria-hidden="true"></span>
                            {{ t('nav.myLendings') }}
                        </router-link>
                    </h3>
                    <p class="mt-2 text-sm text-gray-500 dark:text-gray-300">{{ t('dashboard.inventoryRegisterDescription') }}</p>
                </div>
                <span class="pointer-events-none absolute right-6 top-6 text-gray-300 group-hover:text-gray-400" aria-hidden="true">
                    <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M20 4h1a1 1 0 00-1-1v1zm-1 12a1 1 0 102 0h-2zM8 3a1 1 0 000 2V3zM3.293 19.293a1 1 0 101.414 1.414l-1.414-1.414zM19 4v12h2V4h-2zm1-1H8v2h12V3zm-.707.293l-16 16 1.414 1.414 16-16-1.414-1.414z" />
                    </svg>
                </span>
            </div>
        
            <!--<div class="relative p-6 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 border border-black/5 transition ease-in-out duration-300 rounded-lg shadow hover:bg-gray-100 bg-gray-50 dark:bg-zinc-700 dark:hover:bg-zinc-600">
                <div>
                    <span :class="['bg-yellow-400 text-zinc-800 inline-flex rounded-lg p-3 border border-black/10']">
                        <svg-icon type="mdi" :path="action.icon" class="h-6 w-6" aria-hidden="true"></svg-icon>
                    </span>
                </div>
                <div class="mt-8">
                    <h3 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">
                        <router-link :to="action.href" class="focus:outline-none">
                            <span class="absolute inset-0" aria-hidden="true"></span>
                            {{ t('inventoryRegister') }}
                        </router-link>
                    </h3>
                    <p class="mt-2 text-sm text-gray-500 dark:text-gray-300">{{ action.description }}</p>
                </div>
                <span class="pointer-events-none absolute right-6 top-6 text-gray-300 group-hover:text-gray-400" aria-hidden="true">
                    <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M20 4h1a1 1 0 00-1-1v1zm-1 12a1 1 0 102 0h-2zM8 3a1 1 0 000 2V3zM3.293 19.293a1 1 0 101.414 1.414l-1.414-1.414zM19 4v12h2V4h-2zm1-1H8v2h12V3zm-.707.293l-16 16 1.414 1.414 16-16-1.414-1.414z" />
                    </svg>
                </span>
            </div>

            <div class="relative p-6 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 border border-black/5 transition ease-in-out duration-300 rounded-lg shadow hover:bg-gray-100 bg-gray-50 dark:bg-zinc-700 dark:hover:bg-zinc-600">
                <div>
                    <span :class="['bg-yellow-400 text-zinc-800 inline-flex rounded-lg p-3 border border-black/10']">
                        <svg-icon type="mdi" :path="action.icon" class="h-6 w-6" aria-hidden="true"></svg-icon>
                    </span>
                </div>
                <div class="mt-8">
                    <h3 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">
                        <router-link :to="action.href" class="focus:outline-none">
                            <span class="absolute inset-0" aria-hidden="true"></span>
                            {{ t('inventoryRegister') }}
                        </router-link>
                    </h3>
                    <p class="mt-2 text-sm text-gray-500 dark:text-gray-300">{{ action.description }}</p>
                </div>
                <span class="pointer-events-none absolute right-6 top-6 text-gray-300 group-hover:text-gray-400" aria-hidden="true">
                    <svg class="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M20 4h1a1 1 0 00-1-1v1zm-1 12a1 1 0 102 0h-2zM8 3a1 1 0 000 2V3zM3.293 19.293a1 1 0 101.414 1.414l-1.414-1.414zM19 4v12h2V4h-2zm1-1H8v2h12V3zm-.707.293l-16 16 1.414 1.414 16-16-1.414-1.414z" />
                    </svg>
                </span>
            </div>-->
        </div>
    </main>
</template>
  
<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import SvgIcon from '@jamescoyle/vue-icon';
import {
    mdiFolder,
    mdiListBox,
    mdiTooltipQuestion
} from '@mdi/js';

const { t } = useI18n()
  
/*const actions = [
    {
        title: 'Inventarverzeichnis',
        description: 'Sieh dir eine Auflistung des gesamten registrierten Inventars an.',
        href: '/inventory/items',
        icon: mdiListBox,
        iconForeground: 'text-yellow-400',
        iconBackground: 'bg-zinc-800',
    },
    {
        title: 'Meine Ausleihen',
        description: 'Lege neue Benutzer an und verwalte ihre Daten und Berechtigungen.',
        href: '/users',
        icon: mdiFolder,
        iconForeground: 'text-green-700',
        iconBackground: 'bg-green-100',
    },
    {
        title: 'Ausleihe-Antrag stellen',
        description: 'Verwalte Gruppen (Gremien, Rollen, etc.), lege neue an und lösche nicht mehr benötigte.',
        href: '/groups',
        icon: mdiTooltipQuestion,
        iconForeground: 'text-green-700',
        iconBackground: 'bg-green-100',
    }
]*/
</script>