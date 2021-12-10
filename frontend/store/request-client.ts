export class RequestClient {
  axios: any

  constructor(axios: any) {
    this.axios = axios
  }

  async get(uri: string, params: any = {}) {
    const queryString = Object.keys(params)
      .map((key) => key + '=' + params[key])
      .join('&')

    const query = queryString ? `${uri}?${queryString}` : uri
    return await this.axios.$get(query)
  }
}

export function createRequestClient(axios: any) {
  return new RequestClient(axios)
}
