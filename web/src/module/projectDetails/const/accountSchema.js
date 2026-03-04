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

    if (data.password && data.password.trim() !== '') {
      const pwd = data.password

      if (pwd.length < 8) {
        ctx.addIssue({
          code: 'custom',
          message: 'Password must be at least 8 characters long',
          path: ['password'],
        })
      }

      if (!/[a-z]/.test(pwd)) {
        ctx.addIssue({
          code: 'custom',
          message: 'Password must contain at least one lowercase letter',
          path: ['password'],
        })
      }

      if (!/[A-Z]/.test(pwd)) {
        ctx.addIssue({
          code: 'custom',
          message: 'Password must contain at least one uppercase letter',
          path: ['password'],
        })
      }

      if (!/\d/.test(pwd)) {
        ctx.addIssue({
          code: 'custom',
          message: 'Password must contain at least one number',
          path: ['password'],
        })
      }

      if (!/[@$!%*?&]/.test(pwd)) {
        ctx.addIssue({
          code: 'custom',
          message: 'Password must contain at least one special character (@$!%*?&)',
          path: ['password'],
        })
      }
    }
  })

export { accountSchema }
