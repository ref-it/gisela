<template>
    <router-link to="/" class="z-40 flex h-16 shrink-0 items-center gap-x-4 border-b px-4 shadow-sm shadow-black/20 sm:gap-x-6 sm:px-6 lg:px-6 text-center bg-zinc-800 text-[var(--theme-color)] border-zinc-900" title="Startseite">
        <!--<h1 class="mx-auto font-bold text-xl">GISELA</h1>-->
        <img class="h-8 w-auto" src="@/assets/img/logo.svg" alt="Logo">
    </router-link>
    <div class="flex grow flex-col gap-y-5 overflow-y-auto lg:border-r lg:border-gray-200 bg-zinc-50 px-6 py-4 dark:bg-zinc-800 lg:dark:border-zinc-900">
        <nav class="flex flex-1 flex-col">
            <ul role="list" class="flex flex-1 flex-col gap-y-7">
                <li>
                    <div class="text-xs font-semibold leading-6 text-zinc-500 dark:text-zinc-300 mb-2">{{ t('nav.headingInventory') }}</div>
                    <ul role="list" class="-mx-2 space-y-1">
                        <li>
                            <router-link to="/inventory/items" class="sidebar-nav-button">
                                <svg-icon type="mdi" :path="mdiListBox" aria-hidden="true"></svg-icon>
                                {{ t('nav.inventoryRegister') }}
                            </router-link>
                        </li>
                        <li>
                            <router-link to="/inventory/categories" class="sidebar-nav-button">
                                <svg-icon type="mdi" :path="mdiShape" aria-hidden="true"></svg-icon>
                                {{ t('nav.categories') }}
                            </router-link>
                        </li>
                        <li>
                            <router-link to="/inventory/places" class="sidebar-nav-button">
                                <svg-icon type="mdi" :path="mdiMapMarker" aria-hidden="true"></svg-icon>
                                {{ t('nav.places') }}
                            </router-link>
                        </li>
                        <li>
                            <router-link to="/inventory/groups" class="sidebar-nav-button">
                                <svg-icon type="mdi" :path="mdiFileTree" aria-hidden="true"></svg-icon>
                                {{ t('nav.groups') }}
                            </router-link>
                        </li>
                        <li>
                            <router-link to="/inventory/export" class="sidebar-nav-button">
                                <svg-icon type="mdi" :path="mdiTrayArrowUp" aria-hidden="true"></svg-icon>
                                {{ t('nav.exportData') }}
                            </router-link>
                        </li>
                    </ul>
                </li>
                <li>
                    <div class="text-xs font-semibold leading-6 text-zinc-500 dark:text-zinc-300 mb-2">{{ t('nav.headingLending') }}</div>
                    <ul role="list" class="-mx-2 space-y-1">
                        <li>
                            <router-link to="/lending/application/my" class="sidebar-nav-button">
                                <svg-icon type="mdi" :path="mdiFolder" aria-hidden="true"></svg-icon>
                                {{ t('nav.myLendings') }}
                            </router-link>
                        </li>
                        <li>
                            <router-link to="/lending/application/all" class="sidebar-nav-button">
                                <svg-icon type="mdi" :path="mdiFolderQuestion" aria-hidden="true"></svg-icon>
                                {{ t('nav.lendingRequests') }}
                            </router-link>
                        </li>
                    </ul>
                </li>
            </ul>
        </nav>
    </div>
</template>
  
<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import SvgIcon from '@jamescoyle/vue-icon';
import {
    mdiFileTree,
    mdiFolder,
    mdiFolderQuestion,
    mdiHome,
    mdiListBox,
    mdiMapMarker,
    mdiShape,
    mdiTrayArrowUp
} from '@mdi/js';

const { t } = useI18n()
</script>

<style>
nav .router-link-active,
nav .router-link-active:hover,
nav .router-link-active:focus,
nav .router-link-active:active {
    color: var(--accent-color-text);
    background-color: var(--accent-color);
}

nav .router-link-active svg,
nav .router-link-active:hover svg,
nav .router-link-active:focus svg,
nav .router-link-active:active svg {
    color: var(--accent-color-text);
}
</style>