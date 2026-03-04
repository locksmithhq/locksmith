const fetchModulesUseCase = (fetchModulesRepository) => async (state) => {
  state.module.modules = await fetchModulesRepository()
}

export { fetchModulesUseCase }
