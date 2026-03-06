const fetchRefreshTokensByAccountIDUseCase = (fetchRefreshTokensByAccountIDRepository) => async (state) => {
  const projectID = state.route.params.project_id
  const accountID = state.route.params.account_id
  const sessionID = state.selectedSession?.id || ''
  const { page, limit } = state.tokens.filter
  const tokens = await fetchRefreshTokensByAccountIDRepository(projectID, accountID, {
    page,
    limit,
    session_id: sessionID,
  })
  state.tokens.items = tokens
}

export { fetchRefreshTokensByAccountIDUseCase }
