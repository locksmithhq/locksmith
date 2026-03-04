import { projectController } from '../controller/projectController'
import { fetchProjectsUseCase } from '../domain/usecase/fetchProjectsUseCase'
import { createProjectUseCase } from '../domain/usecase/createProjectUseCase'
import { updateProjectUseCase } from '../domain/usecase/updateProjectUseCase'
import { deleteProjectUseCase } from '../domain/usecase/deleteProjectUseCase'
import { fetchProjectsRepository } from '../data/repository/fetchProjectsRepository'
import { createProjectRepository } from '../data/repository/createProjectRepository'
import { updateProjectRepository } from '../data/repository/updateProjectRepository'
import { deleteProjectRepository } from '../data/repository/deleteProjectRepository'
import { axiosInstance } from '@/plugins/axios'
import { fetchAclRepository } from '@/module/core/data/repository/fetchAclRepository'
import { fetchAclUseCase } from '@/module/core/domain/usecase/fetchAclUseCase'

const fetchProjectsRepositoryImpl = fetchProjectsRepository(axiosInstance)
const createProjectRepositoryImpl = createProjectRepository(axiosInstance)
const updateProjectRepositoryImpl = updateProjectRepository(axiosInstance)
const deleteProjectRepositoryImpl = deleteProjectRepository(axiosInstance)

const fetchProjectsUseCaseImpl = fetchProjectsUseCase(
  fetchProjectsRepositoryImpl,
)
const createProjectUseCaseImpl = createProjectUseCase(
  createProjectRepositoryImpl,
)
const updateProjectUseCaseImpl = updateProjectUseCase(
  updateProjectRepositoryImpl,
)
const deleteProjectUseCaseImpl = deleteProjectUseCase(
  deleteProjectRepositoryImpl,
)

const projectControllerImpl = projectController(
  fetchProjectsUseCaseImpl,
  createProjectUseCaseImpl,
  updateProjectUseCaseImpl,
  deleteProjectUseCaseImpl,
  fetchAclUseCase(fetchAclRepository(axiosInstance)),
)

export { projectControllerImpl }
