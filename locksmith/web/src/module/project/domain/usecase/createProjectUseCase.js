const createProjectUseCase = (createProjectRepository) => async (state) => {
  return await createProjectRepository(state.form)
}

export { createProjectUseCase }
