const fetchSessionsByAccountIDUseCase = (repository) => async (state) => {
  const projectID = state.route.params.id
  const accountID = state.deviceDialog.accountID
  const params = {
    page: state.deviceDialog.filter.page,
    limit: state.deviceDialog.filter.limit,
  }
  const sessions = await repository(projectID, accountID, params)
  state.deviceDialog.sessions = sessions
}

export { fetchSessionsByAccountIDUseCase }
