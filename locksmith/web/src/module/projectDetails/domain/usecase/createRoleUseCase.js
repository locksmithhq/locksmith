const createRoleUseCase = (createRoleRepository) => async (state) => {
  await createRoleRepository(state.role.form)
}

export { createRoleUseCase }
