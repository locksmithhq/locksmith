import { createPaginate } from 'js-query-pagination'
import { debounce } from '@/module/core/helpers/debounce'

const fetchAccountsByProjectIDUseCase =
  (fetchAccountsByProjectIDRepository) => async (state) => {
    const paginate = createPaginate()
      .search(state.account.filter.search, 'name', 'email', 'username')
      .page(state.account.filter.page)
      .limit(state.account.filter.limit)

    const projectID = state.route.params.id

    const debouncedFetch = debounce(
      (projectID, paginate) =>
        fetchAccountsByProjectIDRepository(projectID, paginate),
      300,
    )

    state.accounts = await debouncedFetch(projectID, paginate)
  }

export { fetchAccountsByProjectIDUseCase }
