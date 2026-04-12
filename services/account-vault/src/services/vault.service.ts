import { Account, CreateAccountDTO, AccountStatusDTO } from '../models/account'

export interface VaultService {
  addAccount(dto: CreateAccountDTO): Promise<Account>
  removeAccount(id: string, userId: string): Promise<void>
  getAccountsByUser(userId: string): Promise<AccountStatusDTO[]>
  findAvailable(userId: string, provider: string): Promise<{ id: string; apiKey: string } | null>
  incrementUsage(accountId: string): Promise<void>
}
