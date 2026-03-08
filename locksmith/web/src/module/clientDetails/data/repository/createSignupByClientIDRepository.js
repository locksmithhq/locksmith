const createSignupByClientIDRepository = (axios) => async (projectId, clientId, signup) => {
    return await axios.post(`/projects/${projectId}/clients/${clientId}/signup`, signup)
}

export { createSignupByClientIDRepository }
