const updateSignupByClientIDRepository = (axios) => async (projectId, clientId, signup) => {
    return await axios.put(`/projects/${projectId}/clients/${clientId}/signup`, signup)
}

export { updateSignupByClientIDRepository }
