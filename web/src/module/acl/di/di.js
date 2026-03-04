import { aclController } from '../controller/aclController'
import { fetchAclUseCase } from '@/module/core/domain/usecase/fetchAclUseCase'
import { fetchAclRepository } from '@/module/core/data/repository/fetchAclRepository'
import { axiosInstance } from '@/plugins/axios'

const fetchAclRepositoryImpl = fetchAclRepository(axiosInstance)
const fetchAclUseCaseImpl = fetchAclUseCase(fetchAclRepositoryImpl)

const aclControllerImpl = aclController(fetchAclUseCaseImpl)

export { aclControllerImpl }
