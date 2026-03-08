const revokeSessionUseCase = (repository) => async (state, sessionID) => {
  const projectID = state.route.params.id
  await repository(projectID, sessionID)
}

export { revokeSessionUseCase }
