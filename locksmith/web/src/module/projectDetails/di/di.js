import { axiosInstance } from '@/plugins/axios'
import { projectDetailController } from '../controller/projectDetailController'
import { getProjectByIDRepository } from '../data/repository/getProjectByIDRepository'
import { getProjectByIDUseCase } from '../domain/usecase/getProjectByIDUseCase'
import { getClientsByProjectIDRepository } from '../data/repository/getClientsByProjectIDRepository'
import { getClientsByProjectIDUseCase } from '../domain/usecase/getClientsByProjectIDUseCase'
import { createOAuthClientRepository } from '../data/repository/createOAuthClientRepository'
import { createOAuthClientUseCase } from '../domain/usecase/createOAuthClientUseCase'
import { updateOAuthClientRepository } from '../data/repository/updateOAuthClientRepository'
import { updateOAuthClientUseCase } from '../domain/usecase/updateOAuthClientUseCase'
import { createAccountUseCase } from '../domain/usecase/createAccountUseCase.js'
import { createAccountRepository } from '../data/repository/createAccountRepository'
import { updateAccountUseCase } from '../domain/usecase/updateAccountUseCase'
import { updateAccountRepository } from '../data/repository/updateAccountRepository'
import { fetchAccountsByProjectIDUseCase } from '../domain/usecase/fetchAccountsByProjectIDUseCase'
import { fetchAccountsByProjectIDRepository } from '../data/repository/fetchAccountsByProjectIDRepository'
import { createRoleRepository } from '../data/repository/createRoleRepository'
import { createRoleUseCase } from '../domain/usecase/createRoleUseCase'

import { createModuleRepository } from '../data/repository/createModuleRepository'
import { createModuleUseCase } from '../domain/usecase/createModuleUseCase'

import { createActionRepository } from '../data/repository/createActionRepository'
import { createActionUseCase } from '../domain/usecase/createActionUseCase'

import { fetchRolesRepository } from '../data/repository/fetchRolesRepository'
import { fetchRolesUseCase } from '../domain/usecase/fetchRolesUseCase'

import { fetchModulesRepository } from '../data/repository/fetchModulesRepository'
import { fetchModulesUseCase } from '../domain/usecase/fetchModulesUseCase'

import { fetchActionsRepository } from '../data/repository/fetchActionsRepository'
import { fetchActionsUseCase } from '../domain/usecase/fetchActionsUseCase'
import { fetchProjectAclUseCase } from '../domain/usecase/fetchProjectAclUseCase'
import { fetchProjectAclRepository } from '../data/repository/fetchProjectAclRepository'
import { createProjectAclUseCase } from '../domain/usecase/createProjectAclUseCase'
import { createProjectAclRepository } from '../data/repository/createProjectAclRepository'
import { fetchSessionsByProjectIDUseCase } from '../domain/usecase/fetchSessionsByProjectIDUseCase'
import { fetchSessionsByProjectIDRepository } from '../data/repository/fetchSessionsByProjectIDRepository'
import { countSessionsByProjectIDUseCase } from '../domain/usecase/countSessionsByProjectIDUseCase'
import { countSessionsByProjectIDRepository } from '../data/repository/countSessionsByProjectIDRepository'
import { countAccountsByProjectIDUseCase } from '../domain/usecase/countAccountsByProjectIDUseCase'
import { countAccountsByProjectIDRepository } from '../data/repository/countAccountsByProjectIDRepository'
import { fetchSessionsByAccountIDUseCase } from '../domain/usecase/fetchSessionsByAccountIDUseCase'
import { fetchSessionsByAccountIDRepository } from '../data/repository/fetchSessionsByAccountIDRepository'
import { countSessionsByAccountIDUseCase } from '../domain/usecase/countSessionsByAccountIDUseCase'
import { countSessionsByAccountIDRepository } from '../data/repository/countSessionsByAccountIDRepository'
import { revokeSessionUseCase } from '../domain/usecase/revokeSessionUseCase'
import { revokeSessionRepository } from '../data/repository/revokeSessionRepository'

const projectDetailControllerImpl = projectDetailController(
  getProjectByIDUseCase(getProjectByIDRepository(axiosInstance)),
  getClientsByProjectIDUseCase(getClientsByProjectIDRepository(axiosInstance)),
  createOAuthClientUseCase(createOAuthClientRepository(axiosInstance)),
  updateOAuthClientUseCase(updateOAuthClientRepository(axiosInstance)),
  createAccountUseCase(createAccountRepository(axiosInstance)),
  updateAccountUseCase(updateAccountRepository(axiosInstance)),
  fetchAccountsByProjectIDUseCase(
    fetchAccountsByProjectIDRepository(axiosInstance),
  ),
  countAccountsByProjectIDUseCase(
    countAccountsByProjectIDRepository(axiosInstance),
  ),
  createRoleUseCase(createRoleRepository(axiosInstance)),
  createModuleUseCase(createModuleRepository(axiosInstance)),
  createActionUseCase(createActionRepository(axiosInstance)),
  fetchRolesUseCase(fetchRolesRepository(axiosInstance)),
  fetchModulesUseCase(fetchModulesRepository(axiosInstance)),
  fetchActionsUseCase(fetchActionsRepository(axiosInstance)),
  fetchProjectAclUseCase(fetchProjectAclRepository(axiosInstance)),
  createProjectAclUseCase(createProjectAclRepository(axiosInstance)),
  fetchSessionsByProjectIDUseCase(fetchSessionsByProjectIDRepository(axiosInstance)),
  countSessionsByProjectIDUseCase(countSessionsByProjectIDRepository(axiosInstance)),
  fetchSessionsByAccountIDUseCase(fetchSessionsByAccountIDRepository(axiosInstance)),
  countSessionsByAccountIDUseCase(countSessionsByAccountIDRepository(axiosInstance)),
  revokeSessionUseCase(revokeSessionRepository(axiosInstance)),
)

export { projectDetailControllerImpl }
