const updateProjectUseCase = (updateProjectRepository) => async (state) => {
  return await updateProjectRepository(state.form.id, state.form)
}

export { updateProjectUseCase }
