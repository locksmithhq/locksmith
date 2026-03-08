const countSessionsByAccountIDUseCase = (repository) => async (state) => {
  const projectID = state.route.params.id
  const accountID = state.deviceDialog.accountID
  const total = await repository(projectID, accountID)
  const limit = state.deviceDialog.filter.limit
  state.deviceDialog.filter.totalPages = Math.ceil(total / limit) || 1
}

export { countSessionsByAccountIDUseCase }
