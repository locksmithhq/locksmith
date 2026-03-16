import { createPaginate } from 'js-query-pagination'
import { debounce, isDebounceCancel } from '@/module/core/helpers/debounce'

const fetchAccountsByProjectIDUseCase =
  (fetchAccountsByProjectIDRepository) => {
    const debouncedFetch = debounce(
      (projectID, paginate) =>
        fetchAccountsByProjectIDRepository(projectID, paginate),
      300,
    )

    return async (state) => {
      const paginate = createPaginate()
        .search(state.account.filter.search, 'name', 'email', 'username')
        .page(state.account.filter.page)
        .limit(state.account.filter.limit)

      const projectID = state.route.params.id

      try {
        state.accounts = await debouncedFetch(projectID, paginate)
      } catch (e) {
        if (!isDebounceCancel(e)) throw e
      }
    }
  }

export { fetchAccountsByProjectIDUseCase }
