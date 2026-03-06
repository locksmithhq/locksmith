import { createRouter, createWebHistory, RouterView } from 'vue-router'
import Login from '@/module/login/view/login.vue'
import Auth from '@/module/auth/view/auth.vue'
import Register from '@/module/register/view/register.vue'
import Skeleton from '@/module/skeleton/view/skeleton.vue'
import Project from '@/module/project/view/project.vue'
import ProjectDetails from '@/module/projectDetails/view/projectDetails.vue'
import ClientDetails from '@/module/clientDetails/view/clientDetails.vue'
import UserDetails from '@/module/userDetails/view/userDetails.vue'
import Dashboard from '@/module/dashborad/view/dashboard.vue'
import ACL from '@/module/acl/view/acl.vue'
import i18n from '@/plugins/i18n'
import { axiosInstance } from '@/plugins/axios'

const getCustomDomainClientId = () => {
  const hostname = window.location.hostname
  if (
    hostname === 'localhost' ||
    hostname === '127.0.0.1' ||
    hostname.includes('github.dev') ||
    hostname.includes('gitpod.io')
  )
    return null
  // Simulation: find client by domain in localStorage
  for (let i = 0; i < localStorage.length; i++) {
    const key = localStorage.key(i)
    if (key.startsWith('custom_domain_')) {
      const savedDomain = localStorage.getItem(key)
      if (savedDomain === hostname) {
        return key.replace('custom_domain_', '')
      }
    }
  }
  return null
}

const routes = [
  {
    path: '/',
    component: () => (getCustomDomainClientId() ? Auth : RouterView),
    beforeEnter: (to, from, next) => {
      const clientId = getCustomDomainClientId()
      if (clientId) {
        to.query.client_id = clientId
        // Fallback locale if missing
        if (!to.params.locale)
          to.params.locale = i18n.global.locale.value || 'en'
        return next()
      }
      return next('/en')
    },
  },
  {
    path: '/:locale',
    component: RouterView,
    beforeEnter: async (to, from, next) => {
      const locale = to.params.locale
      const supportedLocales = ['en', 'pt-br']

      if (!supportedLocales.includes(locale)) {
        return next('en')
      }

      if (i18n.global.locale.value !== locale) {
        i18n.global.locale.value = locale
      }

      const clientId = getCustomDomainClientId()
      if (clientId && !to.query.client_id) {
        to.query.client_id = clientId
      }

      return next()
    },
    children: [
      { path: 'auth', name: 'auth', component: Auth },
      { path: 'register', name: 'register', component: Register },
      { path: 'login', name: 'login', component: Login },
      {
        path: '',
        component: Skeleton,
        beforeEnter: async (to, _, next) => {
          try {
            await axiosInstance.get('/locksmith/status', {
              withCredentials: true,
            })
          } catch (error) {
            const locale = to.params.locale
            const supportedLocales = ['en', 'pt-br']

            if (!supportedLocales.includes(locale)) {
              return next('en/login')
            }

            if (i18n.global.locale.value !== locale) {
              i18n.global.locale.value = locale
            }
            return next(locale + '/login')
          }
          return next()
        },
        children: [
          {
            path: '',
            name: 'dashboard',
            component: Dashboard,
          },
          {
            path: 'projects',
            name: 'projects',
            component: Project,
          },
          {
            path: 'projects/:id',
            name: 'project-details',
            component: ProjectDetails,
          },
          {
            path: 'projects/:project_id/clients',
            name: 'clients',
            redirect: (to) => {
              return {
                name: 'project-details',
                params: {
                  ...to.params,
                  id: to.params.project_id,
                },
                query: {
                  tab: 'oauth',
                },
              }
            },
          },
          {
            path: 'projects/:project_id/clients/:client_id',
            name: 'client-details',
            component: ClientDetails,
          },
          {
            path: 'projects/:project_id/users/:account_id',
            name: 'user-details',
            component: UserDetails,
          },
          {
            path: 'acl',
            name: 'acl',
            component: ACL,
          },
        ],
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
