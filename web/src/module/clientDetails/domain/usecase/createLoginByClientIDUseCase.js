const createLoginByClientIDUseCase = (createLoginByClientIDRepository) => (state) => {
    const projectId = state.route.params.project_id
    const clientId = state.route.params.client_id
    const { loginConfig } = state
    return createLoginByClientIDRepository(projectId, clientId, loginConfig)
}

export { createLoginByClientIDUseCase }