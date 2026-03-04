import { debounce } from '@/module/core/helpers/debounce'

const countSessionsByProjectIDUseCase =
  (countSessionsByProjectIDRepository) => async (state) => {
    const projectID = state.route.params.id
    const search = state.session.filter.search

    const debouncedCount = debounce(
      (projectID, search) => countSessionsByProjectIDRepository(projectID, search),
      300,
    )

    const result = await debouncedCount(projectID, search)
    if (result) {
      const total = result.total ?? 0
      state.session.filter.totalPages = Math.ceil(total / state.session.filter.limit) || 1
    }
  }

export { countSessionsByProjectIDUseCase }
