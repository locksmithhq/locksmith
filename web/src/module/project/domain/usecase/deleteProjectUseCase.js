const deleteProjectUseCase = (deleteProjectRepository) => async (id) => {
  await deleteProjectRepository(id)
}

export { deleteProjectUseCase }
