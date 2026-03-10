import * as z from 'zod'

const accountSchema = z
  .object({
    id: z.string().nullable().optional(),
    name: z.string().trim().min(2, 'At least 2 chars.'),
    email: z.string().trim().email('Invalid email'),
    username: z.string().trim().min(2, 'At least 2 chars.'),
    password: z.string().trim().optional().or(z.literal('')),
    role_name: z.string().trim().min(2, 'At least 2 chars.'),
  })
  .superRefine((data, ctx) => {
    const isCreate = !data.id

    if (isCreate) {
      if (!data.password || data.password.trim() === '') {
        ctx.addIssue({
          code: 'custom',
          message: 'Password is required',
          path: ['password'],
        })
        return
      }
    }

  })

export { accountSchema }
