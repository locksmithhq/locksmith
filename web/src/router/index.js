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

let _cachedClientId = null
let _resolved = false

const resolveCustomDomain = async () => {
  if (_resolved) return _cachedClientId
  const hostname = window.location.hostname
  const skip = ['localhost', '127.0.0.1']
  const skipIncludes = ['github.dev', 'gitpod.io']
  if (skip.includes(hostname) || skipIncludes.some((s) => hostname.includes(s))) {
    _resolved = true
    return null
  }
  try {
    const { data } = await axiosInstance.get('/oauth2/resolve-domain', { params: { hostname } })
    _cachedClientId = data.client_id || null
  } catch {
    _cachedClientId = null
  }
  _resolved = true
  return _cachedClientId
}

const routes = [
  {
    path: '/',
    component: RouterView,
    beforeEnter: async (to, from, next) => {
      const clientId = await resolveCustomDomain()
      if (clientId) {
        return next({ name: 'auth', params: { locale: i18n.global.locale.value || 'en' }, query: { client_id: clientId } })
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

      const clientId = _cachedClientId
      if (clientId && !to.query.client_id) {
        to.query.client_id = clientId
      }

      return next()
    },
    children: [
      { path: 'auth', name: 'auth', component: Auth },
      { path: 'register', name: 'register', component: Register },
      {
        path: 'login',
        name: 'login',
        component: Login,
        beforeEnter: async (to, from, next) => {
          const clientId = await resolveCustomDomain()
          if (clientId) {
            return next({ name: 'auth', params: { locale: to.params.locale || 'en' }, query: { client_id: clientId } })
          }
          return next()
        },
      },
      {
        path: '',
        component: Skeleton,
        beforeEnter: async (to, _, next) => {
          const locale = to.params.locale
          const supportedLocales = ['en', 'pt-br']
          const redirectToLogin = () => {
            if (!supportedLocales.includes(locale)) return next('en/login')
            if (i18n.global.locale.value !== locale) i18n.global.locale.value = locale
            return next(locale + '/login')
          }

          try {
            await axiosInstance.get('/locksmith/status', { withCredentials: true })
          } catch (statusError) {
            if (statusError.response?.status !== 401) return redirectToLogin()
            try {
              await axiosInstance.post('/locksmith/r', null, { withCredentials: true })
            } catch {
              return redirectToLogin()
            }
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
