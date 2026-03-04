export const loginRepository =
  (axios) =>
  async ({ email, password, client_id, redirect_uri, state, code_challenge, code_challenge_method }) => {
    const response = await axios.post('/oauth2/login', {
      email,
      password,
      client_id,
      redirect_uri,
      state: state || '',
      code_challenge: code_challenge || '',
      code_challenge_method: code_challenge_method || '',
    })

    if (response.data.redirect_to && !response.data.must_change_password) {
      window.location.href = response.data.redirect_to
    }

    return response.data
  }
