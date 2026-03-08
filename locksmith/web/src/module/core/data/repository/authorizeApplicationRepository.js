export const authorizeApplicationRepository =
  (axios) =>
  async ({ client_id, redirect_uri, response_type, state, code_challenge, code_challenge_method }) => {
    const response = await axios.post('/oauth2/authorize', null, {
      params: {
        client_id,
        redirect_uri,
        response_type: response_type || 'code',
        state: state || '',
        code_challenge: code_challenge || '',
        code_challenge_method: code_challenge_method || '',
      },
      withCredentials: true,
    })

    return response.data
  }
