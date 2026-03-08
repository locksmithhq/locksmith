const createOAuthClientUseCase =
  (createOAuthClientRepository) => async (state) => {
    const projectId = state.route.params.id
    return await createOAuthClientRepository(projectId, state.oauthClient.form)
  }

export { createOAuthClientUseCase }
