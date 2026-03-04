const createSignupByClientIDUseCase = (createSignupByClientIDRepository) => (state) => {
    const projectId = state.route.params.project_id
    const clientId = state.route.params.client_id
    const { registerConfig } = state
    return createSignupByClientIDRepository(projectId, clientId, registerConfig)
}

export { createSignupByClientIDUseCase }
