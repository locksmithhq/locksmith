import { createPaginate } from 'js-query-pagination'
import { debounce } from '@/module/core/helpers/debounce'

const fetchAclUseCase = (fetchAclRepository) => async (state) => {
  state.loading = true
  let paginate = createPaginate()
    .page(state.options.page)
    .limit(state.options.itemsPerPage || 100)

  for (const sort of state.options.sortBy) {
    paginate.sort((sort.order === 'asc' ? '' : '-') + sort.key)
  }

  for (const rule of state.filter.rules) {
    if (!rule.value) continue
    if (Array.isArray(rule.value) && rule.value.every((v) => !v)) continue

    const method = rule.condition

    if (typeof paginate[method] === 'function') {
      if (method === 'search') {
        paginate = paginate.search(rule.value, 'v0', 'v1', 'v2', 'v3')
      } else if (method === 'between') {
        paginate = paginate.between(rule.column, rule.value[0], rule.value[1])
      } else if (['whereIn', 'likeOr', 'likeAnd'].includes(method)) {
        const values =
          typeof rule.value === 'string'
            ? rule.value.split(',').map((s) => s.trim())
            : rule.value
        paginate = paginate[method](rule.column, ...values)
      } else {
        paginate = paginate[method](rule.column, rule.value)
      }
    }
  }

  const debouncedFetch = debounce(
    (paginator) => fetchAclRepository(paginator),
    500,
  )

  state.acls = await debouncedFetch(paginate)
  state.loading = false
}

export { fetchAclUseCase }
