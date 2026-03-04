import { createPaginate } from 'js-query-pagination'
import { debounce } from '@/module/core/helpers/debounce'

const countAccountsByProjectIDUseCase =
  (countAccountsByProjectIDRepository) => async (state) => {
    const paginate = createPaginate()
      .search(state.account.filter.search, 'name', 'email', 'username')
      .limit(state.account.filter.limit)

    const projectID = state.route.params.id

    const debouncedCount = debounce(
      (projectID, paginate) => countAccountsByProjectIDRepository(projectID, paginate),
      300,
    )

    const result = await debouncedCount(projectID, paginate)
    if (result) {
      const total = result.total ?? 0
      state.account.filter.totalPages = Math.ceil(total / state.account.filter.limit) || 1
    }
  }

export { countAccountsByProjectIDUseCase }
