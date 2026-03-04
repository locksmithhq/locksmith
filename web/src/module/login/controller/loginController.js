import { onMounted, reactive, useTemplateRef } from 'vue'
import { useRoute } from 'vue-router'
import { generateCodeVerifier, generateCodeChallenge } from '@/module/core/utils/pkce'

const loginController =
  (authorizeApplicationUseCase, loginUseCase, getConfigUseCase) => () => {
    const route = useRoute()

    const state = reactive({
      clientId: null,
      redirectUri: null,
      responseType: route.query.response_type || 'code',
      stateParam: route.query.state || '',
      codeChallenge: '',
      codeChallengeMethod: 'S256',
      loginOption: 0,
      loading: false,
      email: '',
      password: '',
      emailRules: [
        (v) => !!v || 'E-mail is required.',
        (v) => /.+@.+/.test(v) || 'E-mail must be valid.',
      ],
      passwordRules: [
        (v) => !!v || 'Password is required.',
        (v) => v.length >= 2 || 'Password must be at least 5 characters long.',
      ],
      login: async () => {
        await loginUseCase(state)
      },
      loginForm: useTemplateRef('meuPingulin'),
    })

    onMounted(async () => {
      const codeVerifier = generateCodeVerifier()
      const codeChallenge = await generateCodeChallenge(codeVerifier)

      // Store code_verifier in a short-lived cookie so the server-side callback can use it
      document.cookie = `pkce_cv=${codeVerifier}; path=/; max-age=600; samesite=lax`

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
