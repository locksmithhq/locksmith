const countRefreshTokensByAccountIDUseCase = (countRefreshTokensByAccountIDRepository) => async (state) => {
  const projectID = state.route.params.project_id
  const accountID = state.route.params.account_id
  const sessionID = state.selectedSession?.id || ''
  const count = await countRefreshTokensByAccountIDRepository(projectID, accountID, { session_id: sessionID })
  const limit = state.tokens.filter.limit
  state.tokens.filter.totalPages = Math.max(1, Math.ceil(count / limit))
}

export { countRefreshTokensByAccountIDUseCase }
