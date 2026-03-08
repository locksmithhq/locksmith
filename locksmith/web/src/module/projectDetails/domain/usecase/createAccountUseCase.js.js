const createAccountUseCase = (createAccountRepository) => async (state) => {
  const projectId = state.route.params.id
  return await createAccountRepository(projectId, state.account.form)
}

export { createAccountUseCase }
