import type { Build } from "@/types/buildType";
export default class Request {
  constructor() {
    
  }
  static URL: string = "http://localhost:5000/v1/api/build/weightRanking";

  static async getBuild(body: string): Promise<Build> {
    return await fetch(this.URL, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: body,
    }).then(response => {
      if (!response.ok) {
        throw new Error(response.statusText)
      }
      return response.json() as Promise<{ data: Build }>
    }).then(data => {
      return data.data
  })
  }
}