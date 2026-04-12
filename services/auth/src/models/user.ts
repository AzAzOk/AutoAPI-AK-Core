export interface User {
  id: string
  email: string
  passwordHash: string
  role: 'user' | 'admin'
  createdAt: Date
  updatedAt: Date
}

export interface RegisterDTO {
  email: string
  password: string
}

export interface LoginDTO {
  email: string
  password: string
}

export interface TokenPair {
  accessToken: string
  refreshToken: string
  expiresIn: number
}
