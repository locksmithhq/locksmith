import { createPaginate } from 'js-query-pagination'
import { debounce } from '@/module/core/helpers/debounce'

const fetchProjectsUseCase = (fetchProjectsRepository) => async (state) => {
  const paginate = createPaginate()
    .page(state.filter.page)
    .limit(state.filter.limit)
    .likeOr('name', state.filter.projectName)
    .likeOr('description', state.filter.projectDescription)
    .sort('name')

  const debouncedFetch = debounce(
    (paginator) => fetchProjectsRepository(paginator),
    500,
  )

  state.projects = await debouncedFetch(paginate)
}

export { fetchProjectsUseCase }
