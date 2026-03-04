export const oauthClient = (data = {}) => ({
  id: data.id || '',
  name: data.name || '',
  client_id: data.client_id || '',
  client_secret: data.client_secret || '',
  grant_types: data.grant_types || '',
  redirect_uris: data.redirect_uris || '',
  active: data.active !== undefined ? data.active : true,
  project_id: data.project_id || '',
  created_at: data.created_at || '',
  updated_at: data.updated_at || '',
})
