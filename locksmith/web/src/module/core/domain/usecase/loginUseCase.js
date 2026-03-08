export const loginUseCase = (loginRepository) => async (state) => {
  try {
    state.error = null
    if (state.loginForm) {
      const { valid } = await state.loginForm.validate()
      if (!valid) return
    }

    state.loading = true
    return await loginRepository({
      email: state.email,
      password: state.password,
      client_id: state.clientId,
      redirect_uri: state.redirectUri,
      state: state.stateParam,
      code_challenge: state.codeChallenge || '',
      code_challenge_method: state.codeChallengeMethod || '',
    })
  } catch (error) {
    state.loading = false
    state.error = 'Usuário ou senha inválidos'
  }
}
