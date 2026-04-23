import { onMounted, reactive } from 'vue'
import { useRoute } from 'vue-router'
import { generateCodeVerifier, generateCodeChallenge } from '@/module/core/utils/pkce'
import { useAppForm } from '@/module/core/composables/appForm'
import { loginSchema } from '../const/loginSchema'

const loginController =
  (authorizeApplicationUseCase, loginUseCase, getConfigUseCase) => () => {
    const route = useRoute()

    const { fields, errors, handleSubmit } = useAppForm(loginSchema, {
      email: '',
      password: '',
    })

    const state = reactive({
      clientId: null,
      redirectUri: null,
      responseType: route.query.response_type || 'code',
      stateParam: route.query.state || '',
      codeChallenge: '',
      codeChallengeMethod: 'S256',
      loading: false,
      error: null,
      form: fields,
      errors,
      // getters so loginUseCase can read state.email / state.password directly
      get email() { return state.form.email },
      get password() { return state.form.password },
      login: handleSubmit(() => loginUseCase(state)),
    })

    onMounted(async () => {
      const codeVerifier = generateCodeVerifier()
      const codeChallenge = await generateCodeChallenge(codeVerifier)

      document.cookie = `pkce_cv=${codeVerifier}; path=/; max-age=600; samesite=strict; secure`

      if (!document.cookie.split(';').some((c) => c.trim().startsWith('device_id='))) {
        let deviceId = localStorage.getItem('device_id')
        if (!deviceId) {
          deviceId = crypto.randomUUID()
          localStorage.setItem('device_id', deviceId)
        }
        document.cookie = `device_id=${deviceId}; path=/; max-age=31536000; samesite=lax`
      }

      state.codeChallenge = codeChallenge
      state.codeChallengeMethod = 'S256'

      const config = await getConfigUseCase()
      state.clientId = config.clientId
      state.redirectUri = config.baseUrl + '/api/locksmith/callback'
      await authorizeApplicationUseCase(state)
    })

    return state
  }

export { loginController }
