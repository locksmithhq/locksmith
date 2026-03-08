const revokeSessionUseCase = (revokeSessionRepository) => async (state, sessionID) => {
  const projectID = state.route.params.project_id
  await revokeSessionRepository(projectID, sessionID)
}

export { revokeSessionUseCase }
