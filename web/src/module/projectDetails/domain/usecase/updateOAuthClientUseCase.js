const updateOAuthClientUseCase =
  (updateOAuthClientRepository) => async (state) => {
    const projectId = state.route.params.id
    return await updateOAuthClientRepository(
      projectId,
      state.oauthClient.form.id,
      state.oauthClient.form,
    )
  }

export { updateOAuthClientUseCase }
