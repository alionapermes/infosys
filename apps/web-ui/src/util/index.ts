const baseURL = '//localhost:8081'

export interface Response {
  result: {
    is_connected: boolean
    error?: string | undefined
  }
}

export const get = (path: string) => fetch(`${baseURL}${path}`, {
  method: 'GET',
  headers: {
    'Content-Type': 'application/json',
  },
})

export const post = (path: string, params: object) => fetch(`${baseURL}${path}`, {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify(params),
})
