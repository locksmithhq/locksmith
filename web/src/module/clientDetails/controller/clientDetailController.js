import { onMounted, reactive } from 'vue'
import { useRoute } from 'vue-router'

const clientDetailController = (usecase) => () => {
  const route = useRoute()
  const state = reactive({
    route,
    activeTab: 'config',
    defaultClient: {},
    client: {},
    isEdit: false,
    showLoginPreview: false,
    customDomain: '',
    loginConfig: { enabled: true },
    registerConfig: { enabled: true },
    editClient: () => {
      state.defaultClient = { ...state.client }
      state.isEdit = true
    },
    resetClient: () => {
      state.client = { ...state.defaultClient }
      state.isEdit = false
    },
    saveClient: async () => {
      await usecase.updateClientUseCase(state)
      state.isEdit = false
      state.saveCustomDomain()
    },
    saveRegisterConfig: async () => {
      state.saving = true
      try {
        if (state.registerConfig.id) {
          await usecase.updateSignupByClientIDUseCase(state)
        } else {
          await usecase.createSignupByClientIDUseCase(state)
        }
      } finally {
        state.saving = false
      }
    },
    saveCustomDomain: () => {
      if (state.client.client_id) {
        localStorage.setItem(
          `custom_domain_${state.client.client_id}`,
          state.customDomain,
        )
      }
    },
    loadCustomDomain: () => {
      if (state.client.client_id) {
        const saved = localStorage.getItem(
          `custom_domain_${state.client.client_id}`,
        )
        if (saved) {
          state.customDomain = saved
        }
      }
    },
    saving: false,
    save: async () => {
      state.saving = true
      try {
        if (state.loginConfig.id) {
          await usecase.updateLoginByClientIDUseCase(state)
        } else {
          await usecase.createLoginByClientIDUseCase(state)
        }
      } finally {
        state.saving = false
      }
    },
  })

  onMounted(async () => {
    await usecase.getClientByIDUseCase(state)
    await usecase.getLoginByClientIDUseCase(state)
    try {
      await usecase.getSignupByClientIDUseCase(state)
    } catch {
      // signup config may not exist yet
    }
  })

  return state
}

export { clientDetailController }
