import { registerController } from '../controller/registerController'
import { authorizeApplicationUseCase } from '@/module/core/domain/usecase/authorizeApplicationUseCase'
import { authorizeApplicationRepository } from '@/module/core/data/repository/authorizeApplicationRepository'
import { registerUseCase } from '@/module/core/domain/usecase/registerUseCase'
import { registerRepository } from '@/module/core/data/repository/registerRepository'
import { axiosInstance } from '@/plugins/axios'

const authorizeApplicationRepositoryImpl =
    authorizeApplicationRepository(axiosInstance)
const authorizeApplicationUseCaseImpl = authorizeApplicationUseCase(
    authorizeApplicationRepositoryImpl,
)

const registerRepositoryImpl = registerRepository(axiosInstance)
const registerUseCaseImpl = registerUseCase(registerRepositoryImpl)

const registerControllerImpl = registerController(
    authorizeApplicationUseCaseImpl,
    registerUseCaseImpl,
)

export { registerControllerImpl }
