const countSessionsByAccountIDUseCase = (countSessionsByAccountIDRepository) => async (state) => {
  const projectID = state.route.params.project_id
  const accountID = state.route.params.account_id
  const count = await countSessionsByAccountIDRepository(projectID, accountID)
  const limit = state.sessions.filter.limit
  state.sessions.filter.totalPages = Math.max(1, Math.ceil(count / limit))
}

export { countSessionsByAccountIDUseCase }
