import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: { name: 'inventoryRegister' },
    },
    {
      path: '/inventory',
      children: [
        {
          path: 'items',
          children: [
            {
              path: '',
              name: 'inventoryRegister',
              component: () => import('../components/inventory/items/ItemsIndex.vue'),
            },
            {
              path: 'add',
              component: () => import('../components/inventory/items/ItemsAdd.vue')
            },
            {
              path: 'edit/:id',
              component: () => import('../components/inventory/items/ItemsEdit.vue')
            },
            {
              path: 'retire',
              component: () => import('../components/inventory/items/ItemsRetire.vue')
            },
            {
              path: 'retired',
              component: () => import('../components/inventory/items/ItemsIndexRetired.vue')
            },
            {
              path: 'view/:id',
              component: () => import('../components/inventory/items/ItemsView.vue')
            }
          ]
        },
        {
          path: 'categories',
          children: [
            {
              path: '',
              component: () => import('../components/inventory/categories/CategoriesIndex.vue'),
            },
            {
              path: 'add',
              component: () => import('../components/inventory/categories/CategoriesAdd.vue')
            },
            {
              path: 'edit/:id',
              component: () => import('../components/inventory/categories/CategoriesEdit.vue')
            }
          ]
        },
        {
          path: 'places',
          children: [
            {
              path: '',
              component: () => import('../components/inventory/places/PlacesIndex.vue'),
            },
            {
              path: 'add',
              component: () => import('../components/inventory/places/PlacesAdd.vue')
            },
            {
              path: 'edit/:id',
              component: () => import('../components/inventory/places/PlacesEdit.vue')
            }
          ]
        },
        {
          path: 'groups',
          children: [
            {
              path: '',
              component: () => import('../components/inventory/groups/GroupsIndex.vue'),
            },
            {
              path: 'add',
              component: () => import('../components/inventory/groups/GroupsAdd.vue')
            },
            {
              path: 'edit/:id',
              component: () => import('../components/inventory/groups/GroupsEdit.vue')
            }
          ]
        },
        {
          path: 'export',
          children: [
            {
              path: '',
              component: () => import('../components/inventory/export/ExportData.vue'),
            }
          ]
        },
      ]
    },
    {
      path: '/lending',
      children: [
        {
          path: 'application',
          children: [
            {
              path: 'all',
              component: () => import('../components/lending/Applications.vue')
            },
            {
              path: 'my',
              component: () => import('../components/lending/Applications.vue')
            },
            {
              path: 'new',
              component: () => import('../components/lending/Cart.vue')
            },
            {
              path: ':id',
              component: () => import('../components/lending/ApplicationView.vue')
            }
          ]
        },
      ]
    }
  ]
})

export default router
