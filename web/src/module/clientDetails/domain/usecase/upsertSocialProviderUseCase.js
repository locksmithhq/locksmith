const upsertSocialProviderUseCase = (upsertSocialProviderRepository) => async (state, provider) => {
    const projectId = state.route.params.project_id
    const clientId = state.route.params.client_id
    const data = state.socialProviders[provider]
    await upsertSocialProviderRepository(projectId, clientId, provider, data)
}

export { upsertSocialProviderUseCase }
