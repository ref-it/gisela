<template>
    <div class="p-8 overflow-y-auto">
        <div class="space-y-12">
          <div class="border-b border-gray-900/10 dark:border-white/10 pb-12">
            <h1 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">Ort bearbeiten</h1>

            <div class="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
              <div class="col-span-full">
                <label for="name" class="block text-sm font-medium leading-6 text-gray-900 dark:text-white">Name</label>
                <div class="mt-2">
                  <div class="flex rounded-md shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-yellow-400 w-full">
                    <input type="text" name="name" id="name" v-model="title" class="block flex-1 border-0 bg-transparent py-1.5 text-gray-900 dark:text-white dark:bg-white/5 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6" />
                  </div>
                </div>
              </div>

              <div class="col-span-full">
                <label for="description" class="block text-sm font-medium leading-6 text-gray-900 dark:text-white">Beschreibung</label>
                <div class="mt-2">
                  <textarea id="description" name="description" v-model="description" class="block w-full rounded-md border-0 py-1.5 text-gray-900 dark:text-white dark:bg-white/5 shadow-sm ring-1 ring-inset ring-gray-300 dark:ring-white/10 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-yellow-400 sm:text-sm sm:leading-6" />
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="mt-6 flex items-center justify-end gap-x-6">
          <router-link to="/inventory/categories" type="button" class="text-sm font-semibold leading-6 text-gray-900 dark:text-white">Abbrechen</router-link>
          <button @click="sendNewCategoryData" class="rounded-md bg-green-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600">Sichern</button>
        </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import axios from 'axios'

  const router = useRouter()
  const route = useRoute()

  const title = ref()
  const description = ref()

  const getCategoryData = () => {
     axios.post('/api/inventory/categories/data/' + route.params.id)
          .then((response: any) => {
            title.value = response.data.title
            description.value = response.data.description
          })
          .catch(() => {
            
          })
  }

  const sendNewCategoryData = () => {
     axios.post('/api/inventory/categories/edit/' + route.params.id, { title: title.value, description: description.value })
          .then(() => {
            router.push('/inventory/categories');
          })
          .catch(() => {
            
          })
  }

  getCategoryData()
</script>