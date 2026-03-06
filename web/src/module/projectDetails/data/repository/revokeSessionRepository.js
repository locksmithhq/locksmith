const revokeSessionRepository = (axios) => async (projectID, sessionID) => {
  await axios.delete(`/projects/${projectID}/sessions/${sessionID}`)
}

export { revokeSessionRepository }
