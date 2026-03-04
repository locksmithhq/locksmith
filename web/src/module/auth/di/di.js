import { authController } from '../controller/authController'
import { authorizeApplicationUseCase } from '@/module/core/domain/usecase/authorizeApplicationUseCase'
import { authorizeApplicationRepository } from '@/module/core/data/repository/authorizeApplicationRepository'
import { loginUseCase } from '@/module/core/domain/usecase/loginUseCase'
import { loginRepository } from '@/module/core/data/repository/loginRepository'
import { changePasswordUseCase } from '@/module/core/domain/usecase/changePasswordUseCase'
import { changePasswordRepository } from '@/module/core/data/repository/changePasswordRepository'
import { axiosInstance } from '@/plugins/axios'

const authorizeApplicationRepositoryImpl =
  authorizeApplicationRepository(axiosInstance)
const authorizeApplicationUseCaseImpl = authorizeApplicationUseCase(
  authorizeApplicationRepositoryImpl,
)

const loginRepositoryImpl = loginRepository(axiosInstance)
const loginUseCaseImpl = loginUseCase(loginRepositoryImpl)

const changePasswordRepositoryImpl = changePasswordRepository(axiosInstance)
const changePasswordUseCaseImpl = changePasswordUseCase(
  changePasswordRepositoryImpl,
)

const authControllerImpl = authController(
  authorizeApplicationUseCaseImpl,
  loginUseCaseImpl,
  changePasswordUseCaseImpl,
)

export { authControllerImpl }
