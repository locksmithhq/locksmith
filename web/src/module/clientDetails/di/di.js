import { axiosInstance } from '@/plugins/axios'
import { clientDetailController } from '../controller/clientDetailController'
import { getClientByIDRepository } from '../data/repository/getClientByIDRepository'
import { getClientByIDUseCase } from '../domain/usecase/getClientByIDUseCase'
import { updateClientRepository } from '../data/repository/updateClientRepository'
import { updateClientUseCase } from '../domain/usecase/updateClientUseCase'
import { getLoginByClientIDUseCase } from '../domain/usecase/getLoginByClientIDUseCase'
import { getLoginByClientIDRepository } from '../data/repository/getLoginByClientIDRepository'
import { createLoginByClientIDRepository } from '../data/repository/createLoginByClientIDRepository'
import { createLoginByClientIDUseCase } from '../domain/usecase/createLoginByClientIDUseCase'
import { updateLoginByClientIDRepository } from '../data/repository/updateLoginByClientIDRepository'
import { updateLoginByClientIDUseCase } from '../domain/usecase/updateLoginByClientIDUseCase'
import { getSignupByClientIDRepository } from '../data/repository/getSignupByClientIDRepository'
import { getSignupByClientIDUseCase } from '../domain/usecase/getSignupByClientIDUseCase'
import { createSignupByClientIDRepository } from '../data/repository/createSignupByClientIDRepository'
import { createSignupByClientIDUseCase } from '../domain/usecase/createSignupByClientIDUseCase'
import { updateSignupByClientIDRepository } from '../data/repository/updateSignupByClientIDRepository'
import { updateSignupByClientIDUseCase } from '../domain/usecase/updateSignupByClientIDUseCase'

const clientDetailControllerImpl = clientDetailController(
    {
        getClientByIDUseCase: getClientByIDUseCase(getClientByIDRepository(axiosInstance)),
        updateClientUseCase: updateClientUseCase(updateClientRepository(axiosInstance)),
        getLoginByClientIDUseCase: getLoginByClientIDUseCase(getLoginByClientIDRepository(axiosInstance)),
        createLoginByClientIDUseCase: createLoginByClientIDUseCase(createLoginByClientIDRepository(axiosInstance)),
        updateLoginByClientIDUseCase: updateLoginByClientIDUseCase(updateLoginByClientIDRepository(axiosInstance)),
        getSignupByClientIDUseCase: getSignupByClientIDUseCase(getSignupByClientIDRepository(axiosInstance)),
        createSignupByClientIDUseCase: createSignupByClientIDUseCase(createSignupByClientIDRepository(axiosInstance)),
        updateSignupByClientIDUseCase: updateSignupByClientIDUseCase(updateSignupByClientIDRepository(axiosInstance)),
    }
)

export { clientDetailControllerImpl }
