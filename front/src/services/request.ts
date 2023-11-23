import type { Build } from "@/types/buildType";
import { assertEquals } from 'typia'

export default class Request {
  constructor() {
    
  }
  static URL: string = "http://localhost:5000/v1/api/build/weightRanking";
  static headers = {
    "Content-Type": "application/json",
  }

  static async post<Build>(url: string, body: string): Promise<Build> {
    const res = await fetch(url, {
      method: "POST",
      headers: this.headers,
      body: body,
    })
    if(!res.ok) throw new Error(res.statusText)
    return (await res.json() as Build)
    }

  static async get(body: string): Promise<Build> {
    const res = await fetch(this.URL, {
      method: "GET",
      headers: this.headers
    })
    if(!res.ok) throw new Error(res.statusText)
    return await res.json()
  }
}