interface Message {
  role: 'user' | 'assistant'
  content: string
}

interface ChatRequest {
  model: string
  messages: Message[]
}

class ChatService {
  private baseURL = 'http://localhost:8080'

  async sendMessage(request: ChatRequest): Promise<string> {
    try {
      const response = await fetch(`${this.baseURL}/chat`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(request)
      })

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const result = await response.text()
      return result
    } catch (error) {
      console.error('Chat service error:', error)
      throw new Error('Failed to send message. Please check if the server is running.')
    }
  }
}

export const chatService = new ChatService()