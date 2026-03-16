import { createPaginate } from 'js-query-pagination'
import { debounce, isDebounceCancel } from '@/module/core/helpers/debounce'

const countAccountsByProjectIDUseCase =
  (countAccountsByProjectIDRepository) => {
    const debouncedCount = debounce(
      (projectID, paginate) => countAccountsByProjectIDRepository(projectID, paginate),
      300,
    )

    return async (state) => {
      const paginate = createPaginate()
        .search(state.account.filter.search, 'name', 'email', 'username')
        .limit(state.account.filter.limit)

      const projectID = state.route.params.id

      try {
        const result = await debouncedCount(projectID, paginate)
        if (result) {
          state.account.filter.totalItems = result.total ?? 0
        }
      } catch (e) {
        if (!isDebounceCancel(e)) throw e
      }
    }
  }

export { countAccountsByProjectIDUseCase }
