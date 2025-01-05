<template>
  <div>
    <TransitionRoot as="template" :show="sidebarOpen">
      <Dialog as="div" class="relative z-50 lg:hidden" @close="sidebarOpen = false">
        <TransitionChild as="template" enter="transition-opacity ease-linear duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="transition-opacity ease-linear duration-300" leave-from="opacity-100" leave-to="opacity-0">
          <div class="fixed inset-0 bg-zinc-900/80"></div>
        </TransitionChild>

        <div class="fixed inset-0 flex">
          <TransitionChild as="template" enter="transition ease-in-out duration-300 transform" enter-from="-tranzinc-x-full" enter-to="tranzinc-x-0" leave="transition ease-in-out duration-300 transform" leave-from="tranzinc-x-0" leave-to="-tranzinc-x-full">
            <DialogPanel class="relative mr-16 flex w-full max-w-72 flex-1">
              <TransitionChild as="template" enter="ease-in-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in-out duration-300" leave-from="opacity-100" leave-to="opacity-0">
                <div class="absolute left-full top-0 flex w-16 justify-center pt-5">
                  <button type="button" class="-m-2.5 p-2.5" @click="sidebarOpen = false">
                    <span class="sr-only">Close sidebar</span>
                    <svg-icon type="mdi" :path="mdiClose" class="h-6 w-6 text-white" aria-hidden="true"></svg-icon>
                  </button>
                </div>
              </TransitionChild>
              <!-- Sidebar component, swap this element with another sidebar if you like -->
              <div class="inset-y-0 flex flex-col w-72">
                <SidebarNav />
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </Dialog>
    </TransitionRoot>

    <!-- Static sidebar for desktop -->
    <div class="hidden lg:fixed lg:inset-y-0 lg:flex lg:w-72 lg:flex-col lg:z-40">
      <SidebarNav />
    </div>

    <div class="lg:pl-72 lg:fixed lg:w-full flex inset-y-0 flex-col">
      <div class="sticky top-0 z-40 flex h-16 shrink-0 items-center gap-x-4 border-b px-4 shadow-sm shadow-black/20 sm:gap-x-6 px-8 bg-zinc-800 text-white border-zinc-900">
        <button type="button" class="-m-2.5 p-2.5 text-gray-400 hover:text-gray-500 lg:hidden" @click="sidebarOpen = true">
          <span class="sr-only">Open sidebar</span>
          <svg-icon type="mdi" :path="mdiMenu" class="h-6 w-6" aria-hidden="true"></svg-icon>
        </button>

        <div class="flex flex-1 gap-x-4 self-stretch lg:gap-x-6">
          <div class="flex items-center gap-x-4 lg:gap-x-6 ml-auto">
            <router-link to="/lending/application/new" class="-m-2.5 p-2.5 text-zinc-300 hover:text-zinc-400" :title="t('lending.cartTitle')">
              <svg-icon type="mdi" :path="mdiCartVariant" class="h-6 w-6" aria-hidden="true"></svg-icon>
            </router-link>
            <button type="button" class="-m-2.5 p-2.5 text-zinc-300 hover:text-zinc-400" :title="t('common.aboutGISELA') + ' &hellip;'" @click="openAbout = true">
              <svg-icon type="mdi" :path="mdiInformation" class="h-6 w-6" aria-hidden="true"></svg-icon>
            </button>
            <a class="-m-2.5 p-2.5 text-zinc-300 hover:text-zinc-400" href="mailto:it@stura-ilmenau.de" title="Schreib ne Mail an Ref IT">
              <span class="sr-only">Schreib ne Mail an Ref IT</span>
              <svg-icon type="mdi" :path="mdiHelpCircle" class="h-6 w-6" aria-hidden="true"></svg-icon>
            </a>

            <!-- Separator -->
            <div class="h-6 w-px bg-zinc-700" aria-hidden="true"></div>

            <!-- Profile dropdown -->
            <Menu as="div" class="relative">
              <MenuButton class="-m-1.5 flex items-center p-1.5">
                <span class="sr-only">Open user menu</span>
                <span class="hidden lg:flex lg:items-center">
                  <span class="text-sm font-semibold leading-6 text-white" aria-hidden="true">Tom Cook</span>
                  <svg-icon type="mdi" :path="mdiChevronDown" class="ml-2 h-5 w-5 text-zinc-300" aria-hidden="true"></svg-icon>
                </span>
              </MenuButton>
              
              <transition enter-active-class="transition ease-out duration-100" enter-from-class="transform opacity-0 scale-95" enter-to-class="transform opacity-100 scale-100" leave-active-class="transition ease-in duration-75" leave-from-class="transform opacity-100 scale-100" leave-to-class="transform opacity-0 scale-95">
                <MenuItems class="absolute right-0 z-10 mt-2 w-56 origin-top-right divide-y divide-gray-100 dark:divide-zinc-600 rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none dark:bg-zinc-700 dark:shadow-black/20">
                  <div class="py-1">
                    <MenuItem v-slot="{ active }">
                      <router-link to="/logout" :class="[active ? 'bg-gray-100 text-gray-900 dark:text-gray-50  dark:bg-zinc-800' : 'text-gray-700 dark:text-gray-50 dark:hover:bg-zinc-600', 'group flex items-center px-4 py-2 text-sm']">
                        <svg-icon type="mdi" :path="mdiLogout" class="mr-3 h-5 w-5 text-gray-400 group-hover:text-gray-500 dark:text-gray-50 dark:group-hover:text-gray-50" aria-hidden="true"></svg-icon>
                        Abmelden
                      </router-link>
                    </MenuItem>
                  </div>
                </MenuItems>
              </transition>
            </Menu>
          </div>
        </div>
      </div>
      <RouterView />
    </div>
  </div>

  <TransitionRoot as="template" :show="openAbout">
    <Dialog class="relative z-50" @close="openAbout = false">
      <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0" enter-to="opacity-100" leave="ease-in duration-200" leave-from="opacity-100" leave-to="opacity-0">
        <div class="fixed inset-0 bg-gray-500/75 transition-opacity" />
      </TransitionChild>

      <div class="fixed inset-0 z-50 w-screen overflow-y-auto">
        <div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
          <TransitionChild as="template" enter="ease-out duration-300" enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95" enter-to="opacity-100 translate-y-0 sm:scale-100" leave="ease-in duration-200" leave-from="opacity-100 translate-y-0 sm:scale-100" leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95">
            <DialogPanel class="relative transform overflow-hidden rounded-lg bg-white dark:bg-zinc-800 px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-sm sm:p-6">
              <div>
                <div class="text-center">
                  <DialogTitle as="h3" class="text-base font-semibold text-zinc-800 dark:text-white">GISELA</DialogTitle>
                  <div class="mt-2">
                    <p class="text-zinc-800 dark:text-white"><b>G</b>emeinsames <b>I</b>nventarverwaltungs-<b>S</b>ystem für <b>e</b>ffiziente und <b>l</b>eichte <b>A</b>usleihe</p>
                  </div>
                  <div class="mt-6">
                    <p class="text-zinc-800 dark:text-white">&copy; 2022&ndash;{{ new Date().getFullYear() }} Marc Schlagenhauf</p>
                    <p class="text-zinc-800 dark:text-white">
                      <a class="underline underline-offset-[6px] decoration-[var(--accent-color)] hover:decoration-2" href="https://www.gnu.org/licenses/agpl-3.0.html" target="_blank" rel="noopener noreferrer">AGPLv3</a>
                    </p>
                  </div>
                </div>
              </div>
              <div class="mt-8">
                <button type="button" class="inline-flex w-full justify-center rounded-md bg-[var(--accent-color)] px-3 py-2 font-semibold text-[var(--accent-color-text)] shadow-sm hover:bg-[var(--accent-color-1)] focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-[var(--accent-color)]" @click="openAbout = false">{{ t('common.close') }}</button>
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
import { useI18n } from 'vue-i18n'
import { useCartStore } from '@/stores/cart'
import {
  Dialog,
  DialogPanel,
  DialogTitle,
  Menu,
  MenuButton,
  MenuItem,
  MenuItems,
  TransitionChild,
  TransitionRoot,
} from '@headlessui/vue'
import SvgIcon from '@jamescoyle/vue-icon'
import {
  mdiAccount,
  mdiCartVariant,
  mdiChevronDown,
  mdiClose,
  mdiFormTextboxPassword,
  mdiHelpCircle,
  mdiInformation,
  mdiLogout,
  mdiMenu
} from '@mdi/js';
import SidebarNav from '@/components/SidebarNav.vue'

const userNavigation = [
  {
    name: 'Persönliches Profil',
    href: '/profile/personal-data',
    icon: mdiAccount
  },
  {
    name: 'Passwort ändern',
    href: '/profile/change-password',
    icon: mdiFormTextboxPassword
  },
]

const { t } = useI18n()
const cartStore = useCartStore()

const sidebarOpen = ref(false)
const openAbout = ref(false)

cartStore.$subscribe((mutation, state) => {
    localStorage.setItem('cart', JSON.stringify(state))
})
</script>
