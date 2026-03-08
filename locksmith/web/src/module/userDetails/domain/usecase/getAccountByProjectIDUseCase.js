const getAccountByProjectIDUseCase = (getAccountByProjectIDRepository) => async (state) => {
  const projectID = state.route.params.project_id
  const accountID = state.route.params.account_id
  const account = await getAccountByProjectIDRepository(projectID, accountID)
  state.account = account
}

export { getAccountByProjectIDUseCase }
