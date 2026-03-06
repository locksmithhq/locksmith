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
    roles: [],
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
      state.saveError = null
      state.saveSuccess = false
      try {
        if (state.registerConfig.id) {
          await usecase.updateSignupByClientIDUseCase(state)
        } else {
          await usecase.createSignupByClientIDUseCase(state)
          await usecase.getSignupByClientIDUseCase(state)
        }
        state.saveSuccess = true
      } catch (err) {
        state.saveError = err?.response?.data || 'Error saving register config'
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
    saveError: null,
    saveSuccess: false,
    save: async () => {
      state.saving = true
      state.saveError = null
      state.saveSuccess = false
      try {
        if (state.loginConfig.id) {
          await usecase.updateLoginByClientIDUseCase(state)
        } else {
          await usecase.createLoginByClientIDUseCase(state)
          await usecase.getLoginByClientIDUseCase(state)
        }
        state.saveSuccess = true
      } catch (err) {
        state.saveError = err?.response?.data || 'Error saving login config'
      } finally {
        state.saving = false
      }
    },
  })

  onMounted(async () => {
    await usecase.getClientByIDUseCase(state)
    try {
      await usecase.getLoginByClientIDUseCase(state)
    } catch {
      // login config may not exist yet
    }
    try {
      await usecase.fetchRolesUseCase(state)
    } catch {
      // roles may not be accessible
    }
    try {
      await usecase.getSignupByClientIDUseCase(state)
    } catch {
      // signup config may not exist yet
    }
  })

  return state
}

export { clientDetailController }
