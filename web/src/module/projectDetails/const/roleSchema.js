import * as z from 'zod'

const roleSchema = z.object({
  title: z
    .string()
    .startsWith('role:', 'Role must start with "role:"')
    .trim()
    .min(2, 'At least 2 chars.'),
})

export { roleSchema }
