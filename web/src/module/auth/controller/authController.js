import { onMounted, reactive, useTemplateRef } from 'vue'
import { useRoute } from 'vue-router'

const authController =
  (authorizeApplicationUseCase, loginUseCase, changePasswordUseCase) => () => {
    const route = useRoute()

    const state = reactive({
      clientId: route.query.client_id || '',
      redirectUri: route.query.redirect_uri || '',
      responseType: route.query.response_type || 'code',
      stateParam: route.query.state || '',
      codeChallenge: route.query.code_challenge || '',
      codeChallengeMethod: route.query.code_challenge_method || '',
      client: null,
      loading: true, // Initial loading for authorize
      error: null,
      email: '',
      password: '',
      rememberMe: false,
      emailRules: [
        (v) => !!v || 'Email is required',
        (v) => /.+@.+\..+/.test(v) || 'Email must be valid',
      ],
      passwordRules: [(v) => !!v || 'Password is required'],
      mustChangePassword: false,
      changePasswordJWT: '',
      newPassword: '',
      confirmPassword: '',
      newPasswordRules: [
        (v) => !!v || 'New password is required',
      ],
      confirmPasswordRules: [
        (v) => !!v || 'Confirmation is required',
        (v) => v === state.newPassword || 'Passwords do not match',
      ],
      loginForm: useTemplateRef('authLoginForm'),
      countDown: 60,
      login: async () => {
        state.syncCustomInputs()
        const response = await loginUseCase(state)
        if (response && response.must_change_password) {
          state.mustChangePassword = true
          state.changePasswordJWT = response.change_password_jwt
          state.loading = false
          const interval = setInterval(() => {
            state.countDown--
            if (state.countDown <= 0) {
              clearInterval(interval)
              state.mustChangePassword = false
              state.countDown = 60
              state.changePasswordJWT = ''
              state.newPassword = ''
              state.confirmPassword = ''
              state.loginForm.value = {
                email: '',
                password: '',
              }
            }
          }, 1000)
        }
      },
      changePassword: async () => {
        await changePasswordUseCase(state)
      },
      syncCustomInputs: () => {
        // Helper to find in document or shadow roots
        const findField = (selectors) => {
          for (const selector of selectors) {
            const field = document.querySelector(selector)
            if (field) return field

            // Also check inside any shadow roots if needed
            const shadowHosts = document.querySelectorAll('*')
            for (const host of shadowHosts) {
              if (host.shadowRoot) {
                const shadowField = host.shadowRoot.querySelector(selector)
                if (shadowField) return shadowField
              }
            }
          }
          return null
        }

        const emailField = findField([
          'input[name="email"]',
          '#email',
          'input[type="email"]',
        ])
        const passwordField = findField([
          'input[name="password"]',
          '#password',
          'input[type="password"]',
        ])
        const rememberMeField = findField([
          'input[name="rememberMe"]',
          '#rememberMe',
          'input[type="checkbox"]',
        ])

        if (emailField) state.email = emailField.value
        if (passwordField) state.password = passwordField.value
        if (rememberMeField) state.rememberMe = rememberMeField.checked
      },
      getClientName: () => {
        return state.client?.name || 'OAuth Application'
      },
    })

    onMounted(async () => {
      // Ensure a stable device_id cookie exists for session fingerprinting
      if (!document.cookie.split(';').some((c) => c.trim().startsWith('device_id='))) {
        let deviceId = localStorage.getItem('device_id')
        if (!deviceId) {
          deviceId = (crypto.randomUUID?.() ?? ([1e7]+-1e3+-4e3+-8e3+-1e11).replace(/[018]/g, c => (c ^ (Math.random() * 16 >> c / 4)).toString(16)))
          localStorage.setItem('device_id', deviceId)
        }
        document.cookie = `device_id=${deviceId}; path=/; max-age=31536000; samesite=lax`
      }

      try {
        state.loading = true
        state.error = null

        if (!state.clientId) {
          throw new Error('client_id is required')
        }

        if (!state.redirectUri) {
          throw new Error('redirect_uri is required')
        }

        // Authorize returns the client info
        state.client = await authorizeApplicationUseCase(state)

        if (!state.client) {
          throw new Error('Could not fetch client information')
        }

        const clientName = state.client.name
        if (clientName) {
          document.title = clientName
          const metaTitle = document.querySelector('meta[name="apple-mobile-web-app-title"]')
          if (metaTitle) metaTitle.setAttribute('content', clientName)
        }

        const faviconURL = state.client.login?.favicon_url
        if (faviconURL) {
          let link = document.querySelector("link[rel~='icon']")
          if (!link) {
            link = document.createElement('link')
            link.rel = 'icon'
            document.head.appendChild(link)
          }
          link.href = `/api/oauth2/favicon?client_id=${state.clientId}`
        }

        const existingManifest = document.querySelector("link[rel='manifest']")
        if (existingManifest) existingManifest.remove()
        const manifestLink = document.createElement('link')
        manifestLink.rel = 'manifest'
        const locale = route.params.locale || 'en'
        manifestLink.href = `/api/oauth2/manifest?client_id=${state.clientId}&redirect_uri=${encodeURIComponent(state.redirectUri)}&locale=${locale}`
        document.head.appendChild(manifestLink)
      } catch (err) {
        state.error = err.message
        console.error('Failed to initialize auth:', err)
      } finally {
        state.loading = false
      }
    })

    return state
  }

export { authController }
