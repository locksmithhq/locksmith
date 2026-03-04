import { onMounted, reactive, useTemplateRef } from 'vue'
import { useRoute } from 'vue-router'

const registerController =
  (authorizeApplicationUseCase, registerUseCase) => () => {
    const route = useRoute()

    const state = reactive({
      clientId: route.query.client_id || '',
      redirectUri: route.query.redirect_uri || '',
      responseType: route.query.response_type || 'code',
      stateParam: route.query.state || '',
      client: null,
      loading: true,
      error: null,
      name: '',
      email: '',
      password: '',
      nameRules: [(v) => !!v || 'Name is required'],
      emailRules: [
        (v) => !!v || 'Email is required',
        (v) => /.+@.+\..+/.test(v) || 'Email must be valid',
      ],
      passwordRules: [(v) => !!v || 'Password is required'],
      registerForm: useTemplateRef('registerForm'),
      registerConfig: {
        layout: 'split-right',
        input_variant: 'outlined',
        show_social: true,
        use_custom_html: false,
        custom_html: '',
        custom_css: '',
      },
      register: async () => {
        if (state.registerConfig.use_custom_html) {
          state.syncCustomInputs()
        }
        await registerUseCase(state)
      },
      syncCustomInputs: () => {
        const nameField =
          document.querySelector('input[name="name"]') ||
          document.querySelector('#name')
        const emailField =
          document.querySelector('input[name="email"]') ||
          document.querySelector('#email')
        const passwordField =
          document.querySelector('input[name="password"]') ||
          document.querySelector('#password')

        if (nameField) state.name = nameField.value
        if (emailField) state.email = emailField.value
        if (passwordField) state.password = passwordField.value
      },
    })

    onMounted(async () => {
      try {
        state.loading = true
        state.error = null

        if (!state.clientId) {
          throw new Error('client_id is required')
        }

        state.client = await authorizeApplicationUseCase(state)

        if (!state.client) {
          throw new Error('Could not fetch client information')
        }

        if (state.client.signup) {
          state.registerConfig = {
            ...state.registerConfig,
            ...state.client.signup,
          }
        }
      } catch (err) {
        state.error = err.message
      } finally {
        state.loading = false
      }
    })

    return state
  }

export { registerController }
