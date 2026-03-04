const getConfigUseCase = (getConfigRepository) => async () => {
  return await getConfigRepository()
}

export { getConfigUseCase }
