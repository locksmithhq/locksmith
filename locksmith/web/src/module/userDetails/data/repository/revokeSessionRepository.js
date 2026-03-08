const revokeSessionRepository = (axios) => async (projectID, sessionID) => {
  const response = await axios.delete(`/projects/${projectID}/sessions/${sessionID}`)
  return response.data
}

export { revokeSessionRepository }
