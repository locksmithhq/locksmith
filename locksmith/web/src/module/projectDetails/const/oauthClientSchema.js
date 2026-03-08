import * as z from 'zod'

const oauthClientSchema = z.object({
  name: z.string().trim().min(2, 'At least 2 chars.'),
  client_id: z.string().trim().min(2, 'At least 2 chars.').nullish(),
  redirect_uris: z.string().trim().min(2, 'At least 2 chars.'),
  grant_types: z
    .union([
      z.string().min(1, 'Select at least one grant type'),
      z.array(z.string()).min(1, 'Select at least one grant type'),
    ])
    .optional(),
})

export { oauthClientSchema }
