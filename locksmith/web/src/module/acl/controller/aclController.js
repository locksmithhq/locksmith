import { onMounted, reactive } from 'vue'

const aclController = (fetchAclUseCase) => () => {
  const state = reactive({
    acls: [],
    loading: false,
    showFilter: false,
    totalItems: 0,
    filter: {
      page: 1,
      limit: 100,
      rules: [{ column: 'v1', condition: 'likeAnd', value: '' }],
      logic: 'AND',
      autoSort: true,
    },
    options: {
      page: 1,
      itemsPerPage: 100,
      sortBy: [],
    },
    loadItems: async (options) => {
      state.options = options
      await fetchAclUseCase(state)
    },
    applyFilter: async () => {
      state.filter.page = 1
      await fetchAclUseCase(state)
    },
    clearAllFilters: () => {
      state.filter.rules = [{ column: 'v1', condition: 'likeAnd', value: '' }]
      state.applyFilter()
    },
  })

  return state
}

export { aclController }
