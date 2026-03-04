export const registerUseCase = (registerRepository) => async (state) => {
  try {
    if (state.registerForm) {
      const { valid } = await state.registerForm.validate()
      if (!valid) return
    }

    state.loading = true
    await registerRepository({
      name: state.name,
      email: state.email,
      password: state.password,
      client_id: state.clientId,
      redirect_uri: state.redirectUri,
    })
  } catch (error) {
    state.loading = false
    state.error = error.message || 'Registration failed'
  }
}
