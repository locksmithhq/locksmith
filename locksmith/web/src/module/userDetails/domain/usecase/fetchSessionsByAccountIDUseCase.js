const fetchSessionsByAccountIDUseCase = (fetchSessionsByAccountIDRepository) => async (state) => {
  const projectID = state.route.params.project_id
  const accountID = state.route.params.account_id
  const { page, limit } = state.sessions.filter
  const sessions = await fetchSessionsByAccountIDRepository(projectID, accountID, { page, limit })
  state.sessions.items = sessions
}

export { fetchSessionsByAccountIDUseCase }
