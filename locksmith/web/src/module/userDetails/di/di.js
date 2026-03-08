import { axiosInstance } from '@/plugins/axios'
import { userDetailController } from '../controller/userDetailController'
import { getAccountByProjectIDRepository } from '../data/repository/getAccountByProjectIDRepository'
import { getAccountByProjectIDUseCase } from '../domain/usecase/getAccountByProjectIDUseCase'
import { fetchSessionsByAccountIDRepository } from '../data/repository/fetchSessionsByAccountIDRepository'
import { fetchSessionsByAccountIDUseCase } from '../domain/usecase/fetchSessionsByAccountIDUseCase'
import { countSessionsByAccountIDRepository } from '../data/repository/countSessionsByAccountIDRepository'
import { countSessionsByAccountIDUseCase } from '../domain/usecase/countSessionsByAccountIDUseCase'
import { fetchRefreshTokensByAccountIDRepository } from '../data/repository/fetchRefreshTokensByAccountIDRepository'
import { fetchRefreshTokensByAccountIDUseCase } from '../domain/usecase/fetchRefreshTokensByAccountIDUseCase'
import { countRefreshTokensByAccountIDRepository } from '../data/repository/countRefreshTokensByAccountIDRepository'
import { countRefreshTokensByAccountIDUseCase } from '../domain/usecase/countRefreshTokensByAccountIDUseCase'
import { revokeSessionRepository } from '../data/repository/revokeSessionRepository'
import { revokeSessionUseCase } from '../domain/usecase/revokeSessionUseCase'

const userDetailControllerImpl = userDetailController(
  getAccountByProjectIDUseCase(getAccountByProjectIDRepository(axiosInstance)),
  fetchSessionsByAccountIDUseCase(fetchSessionsByAccountIDRepository(axiosInstance)),
  countSessionsByAccountIDUseCase(countSessionsByAccountIDRepository(axiosInstance)),
  fetchRefreshTokensByAccountIDUseCase(fetchRefreshTokensByAccountIDRepository(axiosInstance)),
  countRefreshTokensByAccountIDUseCase(countRefreshTokensByAccountIDRepository(axiosInstance)),
  revokeSessionUseCase(revokeSessionRepository(axiosInstance)),
)

export { userDetailControllerImpl }
