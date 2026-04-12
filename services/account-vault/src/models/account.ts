export interface Account {
  id: string
  userId: string
  provider: 'openai' | 'claude' | 'gemini' | string
  encryptedKey: string     // AES-256-GCM зашифрованный API ключ
  limitPerDay: number      // лимит запросов в день (0 = без лимита)
  usedToday: number        // использовано сегодня
  resetAt: Date            // когда сбрасывается счётчик
  isActive: boolean
  createdAt: Date
  updatedAt: Date
}

export interface CreateAccountDTO {
  userId: string
  provider: string
  apiKey: string           // принимаем открытый ключ, шифруем внутри
  limitPerDay?: number
}

export interface AccountStatusDTO {
  id: string
  provider: string
  remaining: number
  resetsAt: string
  isActive: boolean
}
