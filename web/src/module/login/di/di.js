import { authorizeApplicationRepository } from '@/module/core/data/repository/authorizeApplicationRepository'
import { authorizeApplicationUseCase } from '@/module/core/domain/usecase/authorizeApplicationUseCase'
import { loginRepository } from '@/module/core/data/repository/loginRepository'
import { loginUseCase } from '@/module/core/domain/usecase/loginUseCase'
import { loginController } from '../controller/loginController'
import { axiosInstance } from '@/plugins/axios'
import { getConfigUseCase } from '@/module/core/domain/usecase/getConfigUseCase'
import { getConfigRepository } from '@/module/core/data/repository/getConfigRepository'

const authorizeApplicationRepositoryImpl =
  authorizeApplicationRepository(axiosInstance)

const authorizeApplicationUseCaseImpl = authorizeApplicationUseCase(
  authorizeApplicationRepositoryImpl,
)

const loginRepositoryImpl = loginRepository(axiosInstance)
const loginUseCaseImpl = loginUseCase(loginRepositoryImpl)

const getConfigUseCaseImpl = getConfigUseCase(
  getConfigRepository(axiosInstance),
)

const loginControllerImpl = loginController(
  authorizeApplicationUseCaseImpl,
  loginUseCaseImpl,
  getConfigUseCaseImpl,
)

export { loginControllerImpl }
