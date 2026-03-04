const createModuleUseCase = (createModuleRepository) => async (state) => {
  await createModuleRepository(state.module.form)
}

export { createModuleUseCase }
