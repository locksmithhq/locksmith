const createActionUseCase = (createActionRepository) => async (state) => {
  await createActionRepository(state.action.form)
}

export { createActionUseCase }
