import { onMounted, reactive } from 'vue'
import { useRoute } from 'vue-router'

const userDetailController =
  (
    getAccountByProjectIDUseCase,
    fetchSessionsByAccountIDUseCase,
    countSessionsByAccountIDUseCase,
    fetchRefreshTokensByAccountIDUseCase,
    countRefreshTokensByAccountIDUseCase,
    revokeSessionUseCase,
  ) =>
  () => {
    const route = useRoute()

    const state = reactive({
      route,
      account: null,
      selectedSession: null,
      sessions: {
        items: [],
        loading: false,
        filter: {
          page: 1,
          limit: 50,
          totalPages: 1,
        },
        fetch: async () => {
          state.sessions.loading = true
          try {
            await Promise.all([
              fetchSessionsByAccountIDUseCase(state),
              countSessionsByAccountIDUseCase(state),
            ])
          } finally {
            state.sessions.loading = false
          }
        },
        select: async (session) => {
          state.selectedSession = session
          state.tokens.filter.page = 1
          state.tokens.items = []
          await state.tokens.fetch()
        },
        revoke: async (session) => {
          session._revoking = true
          try {
            await revokeSessionUseCase(state, session.id)
            await state.sessions.fetch()
            if (state.selectedSession?.id === session.id) {
              const updated = state.sessions.items.find((s) => s.id === session.id)
              if (updated) state.selectedSession = updated
            }
          } finally {
            session._revoking = false
          }
        },
      },
      tokens: {
        items: [],
        loading: false,
        filter: {
          page: 1,
          limit: 50,
          totalPages: 1,
        },
        fetch: async () => {
          if (!state.selectedSession) return
          state.tokens.loading = true
          try {
            await Promise.all([
              fetchRefreshTokensByAccountIDUseCase(state),
              countRefreshTokensByAccountIDUseCase(state),
            ])
          } finally {
            state.tokens.loading = false
          }
        },
      },
    })

    onMounted(async () => {
      await getAccountByProjectIDUseCase(state)
      await state.sessions.fetch()
      if (state.sessions.items.length > 0) {
        await state.sessions.select(state.sessions.items[0])
      }
    })

    return state
  }

export { userDetailController }
