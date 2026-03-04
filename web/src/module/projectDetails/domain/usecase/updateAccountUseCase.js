const updateAccountUseCase = (updateAccountRepository) => async (state) => {
  const projectId = state.route.params.id
  const accountId = state.account.form.id
  return await updateAccountRepository(projectId, accountId, state.account.form)
}

export { updateAccountUseCase }
