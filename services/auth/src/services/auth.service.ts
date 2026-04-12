import { RegisterDTO, LoginDTO, TokenPair } from '../models/user'

export interface AuthService {
  register(dto: RegisterDTO): Promise<{ id: string; email: string }>
  login(dto: LoginDTO): Promise<TokenPair>
  refresh(refreshToken: string): Promise<TokenPair>
  logout(refreshToken: string): Promise<void>
  validateToken(token: string): Promise<{ userId: string; role: string }>
}
