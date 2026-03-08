export const registerRepository =
  (axios) =>
  async ({ name, email, password, client_id, redirect_uri }) => {
    const response = await axios.post('/oauth2/register', {
      name,
      email,
      password,
      client_id,
      redirect_uri,
    })

    if (response.data.redirect_to) {
      window.location.href = response.data.redirect_to
    }

    return response.data
  }
